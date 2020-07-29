// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sqls/1.sql.down.tmpl (306B)
// sqls/1.sql.up.tmpl (5.222kB)
// sqls/2.sql.down.tmpl (1B)
// sqls/2.sql.up.tmpl (6.086kB)

package postgres

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _sqls1SqlDownTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x41\xaa\x42\x31\x0c\x45\xe7\x7f\x15\x1d\x7e\xd7\xf0\x16\x53\x6a\xbd\x96\x40\xfb\x12\x6e\x52\xdd\xbe\x20\x88\x22\x7d\x4e\xcf\xc9\x49\x72\xa1\x5a\xba\x09\xee\xc9\xb4\x4b\x15\x78\x36\x70\x88\xbb\xe8\xee\xdb\xdf\xc2\x53\x3b\x5e\x26\xca\xb9\x23\xa9\x95\x3c\x1d\x5c\xd3\x83\xe0\x98\x2e\x3e\x78\x0f\xfc\x72\xa5\xc6\x92\x13\xae\x93\xf5\xeb\x5c\x53\x5a\x1e\xd2\x58\x3e\xab\xeb\xdc\x9f\x5b\x52\x50\x5a\x03\xb3\x23\x72\xc8\x80\x47\x19\xf6\x7f\xda\x1e\x01\x00\x00\xff\xff\xfd\xeb\x29\xcd\x32\x01\x00\x00")

func sqls1SqlDownTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls1SqlDownTmpl,
		"sqls/1.sql.down.tmpl",
	)
}

func sqls1SqlDownTmpl() (*asset, error) {
	bytes, err := sqls1SqlDownTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/1.sql.down.tmpl", size: 306, mode: os.FileMode(0644), modTime: time.Unix(1595836250, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3f, 0x40, 0x4, 0xc0, 0xd9, 0x9a, 0x73, 0xb0, 0x49, 0xeb, 0x46, 0x2f, 0x53, 0xd1, 0xa1, 0xfa, 0x4, 0x82, 0x7c, 0x55, 0xec, 0x58, 0x5f, 0x7d, 0x33, 0x4f, 0xa0, 0x74, 0x4, 0xbe, 0x4, 0x13}}
	return a, nil
}

var _sqls1SqlUpTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x97\xdf\x6f\xe2\x38\x10\xc7\xdf\xf3\x57\xcc\x43\xa5\x26\x12\x42\x77\x27\xed\x4b\xd1\x3e\xa4\x60\x38\x4e\xbd\x84\x0b\xa0\xee\x3d\x45\xde\xc4\x65\xad\x86\x24\x6b\x1b\x5a\x6e\xb5\xff\xfb\xc9\xce\x2f\x27\x24\xa1\x05\x7a\xba\x3c\x81\x33\xf3\x9d\xf1\xf8\x33\xb6\x33\xf6\x90\xbd\x42\xe0\x7a\xe0\xa1\xc5\x83\x3d\x46\x30\x5d\x3b\xe3\xd5\xdc\x75\x40\x30\xba\xd9\x10\xe6\x73\x22\x7c\x41\xb7\x84\x0b\xbc\x4d\x4d\xcb\x00\x00\xf0\xd0\x6a\xed\x39\x4b\x58\x79\xf3\xd9\x0c\x79\x60\x2f\x8d\x9b\x1b\xe3\x1e\xcd\xe6\x8e\x7a\xef\xa0\xc7\xe1\x2e\x0d\xb1\x20\xa1\x8f\x05\x7c\x06\xc7\x7d\x34\xad\x91\xe6\x2b\x4d\x46\x06\x72\x26\x23\xe3\xe6\x06\x1e\x6c\x67\xb6\xb6\x67\x08\xd2\x28\xdd\xf0\xef\xd1\xc8\x30\x02\x46\xb0\x20\x20\xf0\xd7\x88\x00\x7d\x82\x38\x11\x40\x5e\x29\x17\x1c\x36\x09\x4b\xfd\x2d\xdd\x30\x2c\x68\x12\x73\xc3\x54\xba\x34\x84\xe2\x11\xe4\x55\x28\x87\x78\x17\x45\x46\x31\x1a\x24\x31\x17\x0c\xd3\x58\x34\x15\xfc\xf4\x99\x1c\x4a\x3b\xf9\xa4\x8c\x6e\x31\x3b\xc0\x33\x39\x0c\xd4\x0b\x9c\xa6\x11\xcd\x66\x53\x16\x03\x5e\xa8\xf8\xa6\xfe\xc2\x3f\x49\x4c\x0c\x6b\x64\x18\x38\x12\x84\xe5\x59\x37\xf3\x94\x3a\xc9\x4b\x2c\xdf\x27\xf0\xe3\xc7\x50\xfd\xfe\xf9\xb3\x7f\xb2\x49\x8a\x7d\x1c\x74\x4c\x14\xf6\x98\x05\xdf\x30\x33\x7f\xfb\xd5\x82\xde\x19\x6b\x32\x6f\x98\x6d\x8c\xb7\xe4\x28\xc4\xa7\x4f\x56\x19\x22\x33\x0b\x09\x0f\x18\x4d\xa5\x6a\xbb\x19\x84\xe4\x09\xef\x22\x01\xb7\xb7\x99\x47\x36\x51\x55\xc6\xaa\x8e\xe6\x2f\x2d\x1e\x71\xf2\x62\x5a\x99\x93\x46\xd2\x5b\x9c\x8e\xd6\x41\x2f\x61\xf7\x1a\x4c\x3c\x77\x51\x02\x3d\x9f\x02\xfa\x32\x5f\xae\x96\x50\xc3\x5f\xb9\xbb\x8e\xae\x38\x32\xf2\x1e\x2a\x5c\x8f\x1d\xee\xd1\xd4\xf5\x10\xac\x17\x13\x7b\x85\x5a\x24\xd4\xd0\xd4\xf5\x00\xd9\xe3\xdf\xc1\x73\x1f\x0d\xf4\x05\x8d\xd7\x2b\x04\x0b\xcf\x1d\xa3\xc9\xda\x43\x5d\xdd\x78\x9a\x9d\x94\xb0\x2d\xe5\xfc\x1a\xfc\x68\x52\x6f\xe9\x18\x35\x39\x5f\x45\x6b\x0d\x93\x99\x31\xc2\x93\x1d\x0b\x88\x34\xec\x31\xfb\x2f\xc0\xc9\x9c\x9e\xc9\xe1\xb8\x3c\x12\xea\xf2\xa9\xa0\xbe\xbb\x93\xaf\x71\x20\x59\xdb\x63\x76\xa0\xf1\xa6\xaa\x62\x1b\x86\xfa\x6a\x5c\x0f\x45\x4d\xf5\x12\x1c\x9b\xc9\x7d\x14\x92\xc5\x92\xb7\x00\xf9\x2e\x1e\x4b\x9d\xf7\xed\x68\xad\x3b\xd5\xe9\x18\x52\xc1\x6f\xc6\xd9\xc5\xf4\xfb\x8e\x1c\x21\x7a\x0e\xa1\xe7\xee\x6c\x55\x35\xaf\x07\x54\xa9\x79\x09\x4e\xf5\xc4\x3e\x0c\xa6\x24\x22\x2d\x9b\x9c\x1a\x2e\xc1\xd2\xb0\xca\xea\x5e\x39\x34\xf6\x9d\xf6\x12\x37\x63\x74\x57\xfa\x0d\xd9\x5e\x8c\xbd\xd4\x38\x1f\xf9\x7e\xd9\xff\x2d\xe5\xaa\x70\x57\x24\x5c\xea\x5d\x44\x77\x99\xd0\x47\x91\xbd\xe3\x84\xd5\x80\x51\x03\x1d\xd4\x34\x0e\xd5\x1a\xfd\x7d\x96\x1a\x04\x71\xe2\x87\xbb\x34\xa2\x01\x16\xc4\xcf\x63\xf9\x85\x52\x86\x00\x98\xf9\xf8\xa0\x08\xd1\xbe\x5a\x5a\xea\x67\xb7\x8a\xd4\x68\xbb\xb2\x9c\xee\x97\xde\xeb\x6c\xb3\x1d\x32\x0b\xb2\xc5\x34\xea\xb0\x68\x54\xac\x73\x15\x4a\x31\xca\xd5\xcd\x6e\x5f\xc4\xfc\x9a\x24\x95\xf6\x11\xf3\x4f\x38\xe2\x79\x6f\xa5\x2c\xd9\xd3\x90\xb0\xde\x2c\x8e\xee\xd4\x78\x8f\x05\x66\x27\x72\xef\xbd\x89\xcb\xe7\xac\x4b\xd5\xbb\x1c\xbb\x50\xa3\xa1\x1f\x26\x5b\x4c\x63\x5f\xad\x82\x5f\x56\x61\xed\xcc\xff\x5a\x23\x30\xd5\xf0\xa0\xac\x4e\x37\x72\xd7\xdc\x20\x94\xde\x25\x1b\x44\x95\xd0\xa5\x1b\x44\xc2\x80\x91\x34\xc2\x01\x81\x3d\x25\x2f\x90\x26\x11\x0d\xa8\x3c\x03\xaa\x23\xc9\xcc\xaf\xdb\x12\xf5\x41\x75\xa9\xce\xff\xca\x5e\x95\x3f\x2d\xc0\xdc\x58\xa2\x07\x34\x5e\xe9\x1f\x20\xc3\xbc\x41\xec\x25\xe8\x32\x45\x43\xd5\x0e\xf3\xcc\xd6\x5e\x36\x62\xd4\x6c\x65\xe3\x0f\xcb\xa6\x93\xb6\x45\x02\xc6\xd4\x73\xff\x04\xd3\x34\xcd\xce\x73\xf5\x0f\x77\xae\xed\xb0\xb2\x9a\x95\x31\x1f\xd2\x10\x3e\xb7\x9e\xc9\xc3\x62\x43\xb2\xac\xba\x8c\x66\x53\x89\xe9\x8e\xdd\x92\xb5\x5b\xc2\x91\x70\x5e\xbc\x4a\xb4\xa8\x66\x29\xa8\x6b\x95\x9f\x43\x85\x0e\xd4\x27\x5b\xd4\x57\x9b\x70\x59\xf2\x56\x41\xed\xc3\x49\x4a\xba\xde\x04\x79\x70\xff\x77\x63\x05\x1a\x9d\xd2\x86\xce\xe9\x2d\xba\x13\x40\x15\xa7\x71\x26\xc8\xa8\x5c\x07\xad\x66\x3d\x2c\x8c\x8b\x1a\x70\xc1\x68\xbc\xf1\xf1\x66\x63\x9a\x75\xcb\x8a\xda\xbb\x3b\x41\x5e\xc5\x00\x6e\x07\xb7\xd9\x4f\xab\x06\x15\xcf\xa9\xd2\xc0\x56\xcd\x37\x54\xbb\xb4\xbd\x84\x66\xcc\x56\x54\x6b\x98\x66\x16\x39\xac\x2d\x27\x5a\x63\xf5\x54\xb8\x6c\xe5\x4c\x3d\x7c\x9e\x39\x7c\x86\x86\x48\x51\x86\xdc\xa0\xc6\x44\x93\x8c\xaa\x0d\x6a\x7d\xd0\xad\x5d\x74\x42\x53\xbb\x86\x48\x99\x62\x7d\x7d\x8c\x99\xe7\xae\x17\xd2\xa8\x75\xd9\xba\x70\xea\x3f\xeb\xff\x0d\x00\x00\xff\xff\x2d\xa2\x7e\xa8\x66\x14\x00\x00")

func sqls1SqlUpTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls1SqlUpTmpl,
		"sqls/1.sql.up.tmpl",
	)
}

func sqls1SqlUpTmpl() (*asset, error) {
	bytes, err := sqls1SqlUpTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/1.sql.up.tmpl", size: 5222, mode: os.FileMode(0644), modTime: time.Unix(1595836129, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x34, 0xfd, 0xe1, 0xcd, 0x6d, 0x56, 0x62, 0x51, 0xb2, 0xcc, 0x48, 0x52, 0xe2, 0xfe, 0xb1, 0xd0, 0x48, 0x7, 0x50, 0xfb, 0x1, 0xbc, 0x2c, 0x66, 0x43, 0xfa, 0x31, 0xde, 0x19, 0x4d, 0xf5, 0x9}}
	return a, nil
}

var _sqls2SqlDownTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x06\x04\x00\x00\xff\xff\xa9\x06\x09\x63\x01\x00\x00\x00")

func sqls2SqlDownTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls2SqlDownTmpl,
		"sqls/2.sql.down.tmpl",
	)
}

func sqls2SqlDownTmpl() (*asset, error) {
	bytes, err := sqls2SqlDownTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/2.sql.down.tmpl", size: 1, mode: os.FileMode(0644), modTime: time.Unix(1594966792, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0xb8, 0x5, 0xea, 0x7a, 0xc0, 0x14, 0xe2, 0x35, 0x56, 0xe9, 0x8b, 0xb3, 0x74, 0x70, 0x2a, 0x8, 0x34, 0x42, 0x68, 0xf9, 0x24, 0x89, 0xa0, 0x2f, 0x8, 0x80, 0x84, 0x93, 0x94, 0xa1, 0xe4}}
	return a, nil
}

var _sqls2SqlUpTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x98\x5b\x77\xa2\x48\x17\x86\xef\xbf\x5f\xc1\x9d\xc9\xfa\x48\x5b\x1c\x84\xc0\xdc\x8c\x1a\x12\x89\xc7\x46\x10\x75\x7a\x86\x55\x42\xa9\x28\xa7\x70\xd2\x78\x91\xdf\x3e\x0b\x4c\xd2\x9d\x0c\x65\x3a\x59\x86\x78\x11\x52\x96\xae\xb5\x9f\xfd\xb2\x77\xbd\x1b\xe4\xde\x50\x52\x54\x42\xee\xa9\x7d\x22\x48\x66\x8e\x6d\x7e\xf3\x03\x68\x40\x33\xb6\x7d\x2f\x22\xce\x6c\x8b\x24\x3c\xe8\x22\x92\xb0\x50\x64\x86\x76\x90\xed\x93\x84\x19\x22\x18\x23\xcb\x80\x31\x49\x24\x81\xf5\xb8\x3e\x27\x46\xf5\x8e\x26\x0d\x89\xb3\xca\x62\xc7\xaa\xd3\x3e\x17\xcc\x6f\xd2\x71\x6b\xbd\x35\xa2\x66\x8b\x19\x57\x48\xa2\x72\x23\xa9\x15\x72\x7f\x25\x2a\x34\xa0\x84\x0b\xc0\x5f\x50\x80\xa0\x38\x11\x70\x62\xed\x12\xb7\x7d\xfe\xc7\xff\x3e\x8b\x55\xb5\x7c\xb5\xae\xdf\x2d\xae\x55\xf5\x1a\xba\xab\xc4\x63\xda\x4a\x9a\x71\x0c\xb4\x8c\x35\xbf\x9e\x0a\xab\xc9\xde\x7a\x46\x70\xbb\x9d\x04\x8d\xfe\x83\x52\x37\x3b\x6d\x2e\x5d\xe6\xac\xfd\x61\x0e\x9b\xff\x3b\x15\xda\x25\xd5\x9c\xde\xf6\xbf\x8f\x06\x42\xdb\x1f\xc5\xf5\x86\x49\x6b\x2c\x9f\x71\x5c\x49\x1d\x49\x95\x5e\xae\xde\x87\x8c\x63\x0e\x51\xe4\x27\xa1\x89\x5e\x50\xbf\x5d\xae\x51\x9b\x09\x36\xcd\x91\x63\x59\xeb\xb1\x26\x5d\x36\x0c\xca\x0d\x2b\x64\xa5\x0a\x03\xbb\x7a\xf6\xd7\x3f\x3f\x7e\x54\xff\xfe\xff\x79\x15\x26\xf1\x72\x57\x0d\x7d\x07\x45\xaf\xd1\x6a\x62\xad\x26\xb2\xcc\xf3\x36\x45\x5f\xd0\x1c\x01\x18\x91\xe1\x44\x86\x39\x20\xf2\xc7\x80\x91\x9f\x08\xc3\x5e\xca\x77\x15\xb4\x1a\x1a\x33\x3d\x69\x76\xe7\xe3\x02\x60\xcb\xaa\x46\x89\xfb\xc5\xb0\xb4\xe3\x4e\xd2\x60\x69\x6d\x9d\x5e\xf7\x41\xa7\xec\x8e\x31\xe4\x67\xc5\xb0\xa6\xef\x99\x30\xfe\x62\x5e\x79\xc4\x84\xa1\x22\xb3\xe1\xad\xb0\x02\x93\x9a\xdc\x74\xbe\xb3\xbb\xff\xf2\xce\x7d\x3f\xfb\x7b\x3f\x2c\x8e\x36\x40\xa1\x6b\x47\xd1\xcf\x9e\xdb\x37\xa0\x91\x2d\x9f\x32\xc9\x3f\x14\x67\x40\x12\x6b\x74\xff\x4b\x1a\xb6\xd5\x99\xf5\xe2\x44\x68\xea\x1b\x63\xc4\x02\xfd\x5a\x5b\x4d\xb7\x19\x16\xf6\x70\xc6\xb4\xc1\xeb\x26\xa4\x58\x91\xe6\x5f\x6f\xf3\x22\x7d\x29\x52\xdc\xe3\x19\x2f\x56\x0f\xb4\x0e\xfe\x7e\x1d\x59\x81\x86\x30\xed\xf6\x7b\x7a\x3c\x5d\x84\x03\xb8\xe8\xab\x94\x2e\x5d\x6d\x32\x42\xec\x91\x7f\x24\x05\x06\xda\x89\x28\xd0\x52\x2f\xbd\x40\xe1\xec\x35\xda\x6a\xb0\xa5\xa1\xda\x2a\x88\xf4\x8c\x10\x6b\x24\xc7\x52\xa0\x3f\x3c\x11\x09\xe8\x29\x9c\x50\x12\x93\x76\x26\x11\x10\xee\x5b\x0b\x6d\xe1\xa2\x56\x86\x88\x75\xa7\x23\x49\xb0\xb7\xb6\xd3\x10\x41\x4a\xbc\xb6\x3c\xeb\x02\xc5\x68\xaf\x76\x3e\x77\xa7\x51\x10\xec\x0e\xd6\x01\xc6\x61\x3e\x56\x07\xc5\xae\x54\x5a\xf6\xd4\xac\xdf\x50\x86\x0d\xc6\x92\x25\x19\xd4\x1a\xc2\x2c\x79\xb0\xa8\x83\xd9\x63\x2c\xeb\x58\xd9\x3f\xda\x5c\x69\x02\x30\x3d\x75\x78\x3d\x46\xf4\xca\x83\xc1\x12\xf1\x7c\xdd\x72\xd6\xf3\x83\x02\x60\x3c\xf0\x18\x02\x3c\xf9\xe6\xa1\x29\x2e\xeb\x8e\x77\x79\xf6\x0d\xd8\x79\x53\x5d\xad\xc1\x95\xac\xca\xaa\xc0\xa7\x9c\x3e\xb6\x33\x02\x7f\xe3\xa1\x82\xe6\x65\x29\x91\xe5\x0a\x3c\x9a\x17\x19\xee\xd0\x40\xf1\x6e\x30\x6d\xb0\x05\x92\x24\xdd\x2a\x4d\xa6\xbb\x03\x4e\xbc\x09\xae\xd8\x45\xde\x60\x96\x1d\xfb\x5f\x49\x66\x28\x83\x6d\x7b\x3c\xe5\xa4\x88\xe1\x1e\x8c\xe9\x20\x59\x3b\x49\xfe\xf0\x90\xda\x68\xf3\x11\xcd\x0e\xa1\xbd\xac\xe8\x7c\x27\x43\xfd\xb9\x6b\xd8\xd6\xdb\xb7\x13\x33\xd3\x1c\x56\xe5\x48\xa1\x8b\x87\x89\x52\x42\x17\xbb\x78\x29\xa1\x8b\xdd\xb3\x94\xd0\xc5\x9e\x55\x4a\xe8\x62\xc3\x28\x25\x74\xf1\x51\x7d\xdc\xee\xc2\x9c\x49\x65\x74\x17\x2e\x74\x09\xdd\x85\x0b\x5d\x42\x77\xe1\x42\x97\x50\xe2\xb8\xd0\x25\x94\x38\x2e\x74\x19\x25\x8e\x31\x37\x7c\x89\xe3\x62\x27\x11\x0a\x9f\xcc\x35\x5f\xe7\x23\xd8\x3e\xfc\xaf\x2f\x64\x1a\xc8\x9b\xb7\x93\xe6\xae\x6e\xd7\x9b\xc6\x6c\x7e\x27\xa9\x9b\x49\x85\xc4\xb4\x39\x5e\xe6\xdf\x0c\x07\xe2\xfb\x8d\xcd\x2d\x39\xc1\xb5\x06\xed\x5e\xcf\xe8\xad\x65\xbe\x51\x21\x31\x69\xbf\x91\xde\x8b\xb1\x01\xb9\xd0\x76\x48\xe2\x39\xb6\x1d\xe5\xef\x01\x53\x44\x12\x41\xe8\xa7\xb6\x85\x42\x92\x80\x29\x8c\x61\xf8\x3b\xef\xa9\x8a\x65\x21\x2a\xcf\xa3\xd9\x9f\x68\x0b\xdd\xc0\x41\xdf\x4c\xdf\x7d\xfc\x22\x0e\x13\x94\x3d\x86\xf9\xfe\xc2\x41\x4f\x3f\xa6\x01\x0d\x2e\x00\x93\x8d\x22\x80\x17\x41\x4d\x64\xa8\x82\x6d\x41\x64\xd9\x37\xd4\xfd\xc4\x6c\x31\x77\x65\x9f\xc0\x7e\xaa\xfa\x94\x74\xff\x0d\x00\x00\xff\xff\x09\x7b\x95\x0f\xc6\x17\x00\x00")

func sqls2SqlUpTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls2SqlUpTmpl,
		"sqls/2.sql.up.tmpl",
	)
}

func sqls2SqlUpTmpl() (*asset, error) {
	bytes, err := sqls2SqlUpTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/2.sql.up.tmpl", size: 6086, mode: os.FileMode(0644), modTime: time.Unix(1595838726, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x8a, 0x83, 0xa6, 0xc2, 0x2a, 0xa8, 0xfa, 0x5, 0xe9, 0xbb, 0x7f, 0x7, 0x65, 0x50, 0x39, 0x4a, 0x3b, 0xe8, 0xa7, 0x85, 0x2e, 0x4d, 0xde, 0x67, 0xcd, 0x6f, 0xdf, 0xc6, 0x6f, 0x7b, 0x19}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"sqls/1.sql.down.tmpl": sqls1SqlDownTmpl,
	"sqls/1.sql.up.tmpl":   sqls1SqlUpTmpl,
	"sqls/2.sql.down.tmpl": sqls2SqlDownTmpl,
	"sqls/2.sql.up.tmpl":   sqls2SqlUpTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"sqls": &bintree{nil, map[string]*bintree{
		"1.sql.down.tmpl": &bintree{sqls1SqlDownTmpl, map[string]*bintree{}},
		"1.sql.up.tmpl":   &bintree{sqls1SqlUpTmpl, map[string]*bintree{}},
		"2.sql.down.tmpl": &bintree{sqls2SqlDownTmpl, map[string]*bintree{}},
		"2.sql.up.tmpl":   &bintree{sqls2SqlUpTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
