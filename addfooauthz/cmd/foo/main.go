package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	addtransports "github.com/cage1016/gokit-istio-security/internal/app/add/transports"
	"github.com/cage1016/gokit-istio-security/internal/app/foo/endpoints"
	"github.com/cage1016/gokit-istio-security/internal/app/foo/service"
	"github.com/cage1016/gokit-istio-security/internal/app/foo/transports"
	pb "github.com/cage1016/gokit-istio-security/pb/foo"
)

const (
	defZipkinV2URL string = ""
	defServiceName string = "foo"
	defLogLevel    string = "error"
	defServiceHost string = "localhost"
	defHTTPPort    string = "8180"
	defGRPCPort    string = "8181"
	defAddURL      string = ""
	envZipkinV2URL string = "QS_ZIPKIN_V2_URL"
	envServiceName string = "QS_SERVICE_NAME"
	envLogLevel    string = "QS_LOG_LEVEL"
	envServiceHost string = "QS_SERVICE_HOST"
	envHTTPPort    string = "QS_HTTP_PORT"
	envGRPCPort    string = "QS_GRPC_PORT"
	envAddURL      string = "QS_ADD_URL"
)

type config struct {
	serviceName string
	logLevel    string
	serviceHost string
	httpPort    string
	grpcPort    string
	zipkinV2URL string
	addURL      string
}

// Env reads specified environment variable. If no value has been found,
// fallback is returned.
func env(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowInfo())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	cfg := loadConfig(logger)
	logger = log.With(logger, "service", cfg.serviceName)
	level.Info(logger).Log("version", service.Version, "commitHash", service.CommitHash, "buildTimeStamp", service.BuildTimeStamp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// addsvc grpc connection
	var conn *grpc.ClientConn
	{
		var err error
		if cfg.addURL != "" {
			conn, err = grpc.Dial(cfg.addURL, grpc.WithInsecure())
			if err != nil {
				level.Error(logger).Log("serviceName", cfg.addURL, "error", err)
				os.Exit(1)
			}
		}
	}

	tracer := initOpentracing()
	zipkinTracer := initZipkin(cfg.serviceName, cfg.httpPort, cfg.zipkinV2URL, logger)
	service := NewServer(conn, tracer, zipkinTracer, logger)
	endpoints := endpoints.New(service, logger, tracer, zipkinTracer)

	hs := health.NewServer()
	hs.SetServingStatus(cfg.serviceName, healthgrpc.HealthCheckResponse_SERVING)

	wg := &sync.WaitGroup{}

	go startHTTPServer(ctx, wg, endpoints, tracer, zipkinTracer, cfg.httpPort, logger)
	go startGRPCServer(ctx, wg, endpoints, tracer, zipkinTracer, cfg.grpcPort, hs, logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	cancel()
	wg.Wait()

	fmt.Println("main: all goroutines have told us they've finished")
}

func loadConfig(logger log.Logger) (cfg config) {
	cfg.serviceName = env(envServiceName, defServiceName)
	cfg.logLevel = env(envLogLevel, defLogLevel)
	cfg.serviceHost = env(envServiceHost, defServiceHost)
	cfg.httpPort = env(envHTTPPort, defHTTPPort)
	cfg.grpcPort = env(envGRPCPort, defGRPCPort)
	cfg.zipkinV2URL = env(envZipkinV2URL, defZipkinV2URL)
	cfg.addURL = env(envAddURL, defAddURL)
	return cfg
}

func NewServer(conn *grpc.ClientConn, tracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) service.FooService {
	addservice := addtransports.NewGRPCClient(conn, tracer, zipkinTracer, logger)
	service := service.New(addservice, logger)
	return service
}

func initOpentracing() stdopentracing.Tracer {
	return stdopentracing.GlobalTracer()
}

func initZipkin(serviceName, httpPort, zipkinV2URL string, logger log.Logger) (zipkinTracer *zipkin.Tracer) {
	var (
		err           error
		hostPort      = fmt.Sprintf("localhost:%s", httpPort)
		useNoopTracer = (zipkinV2URL == "")
		reporter      = zipkinhttp.NewReporter(zipkinV2URL)
	)
	zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
	zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer))
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	if !useNoopTracer {
		logger.Log("tracer", "Zipkin", "type", "Native", "URL", zipkinV2URL)
	}

	return
}

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup, endpoints endpoints.Endpoints, tracer stdopentracing.Tracer, zipkinTracer *zipkin.Tracer, port string, logger log.Logger) {
	wg.Add(1)
	defer wg.Done()

	if port == "" {
		level.Error(logger).Log("protocol", "HTTP", "exposed", port, "err", "port is not assigned exist")
		return
	}

	p := fmt.Sprintf(":%s", port)
	// create a server
	srv := &http.Server{Addr: p, Handler: transports.NewHTTPHandler(endpoints, tracer, zipkinTracer, logger)}
	level.Info(logger).Log("protocol", "HTTP", "exposed", port)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			level.Info(logger).Log("Listen", err)
		}
	}()

	<-ctx.Done()

	// shut down gracefully, but wait no longer than 5 seconds before halting
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ignore error since it will be "Err shutting down server : context canceled"
	srv.Shutdown(shutdownCtx)

	level.Info(logger).Log("protocol", "HTTP", "Shutdown", "http server gracefully stopped")
}

func startGRPCServer(ctx context.Context, wg *sync.WaitGroup, endpoints endpoints.Endpoints, tracer stdopentracing.Tracer, zipkinTracer *zipkin.Tracer, port string, hs *health.Server, logger log.Logger) {
	wg.Add(1)
	defer wg.Done()

	p := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", p)
	if err != nil {
		level.Error(logger).Log("protocol", "GRPC", "listen", port, "err", err)
		os.Exit(1)
	}

	var server *grpc.Server
	level.Info(logger).Log("protocol", "GRPC", "exposed", port)
	server = grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterFooServer(server, transports.MakeGRPCServer(endpoints, tracer, zipkinTracer, logger))
	healthgrpc.RegisterHealthServer(server, hs)
	reflection.Register(server)

	go func() {
		// service connections
		err = server.Serve(listener)
		if err != nil {
			fmt.Printf("grpc serve : %s\n", err)
		}
	}()

	<-ctx.Done()

	// ignore error since it will be "Err shutting down server : context canceled"
	server.GracefulStop()

	fmt.Println("grpc server gracefully stopped")
}
