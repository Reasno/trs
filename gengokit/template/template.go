// Code generated by go-bindata. DO NOT EDIT.
// sources:
// NAME-service/handlers/handlers.gotemplate (62B)
// NAME-service/svc/client/grpc/client.gotemplate (2.205kB)
// NAME-service/svc/client/http/client.gotemplate (105B)
// NAME-service/svc/endpoints.gotemplate (4.609kB)
// NAME-service/svc/transport_grpc.gotemplate (2.962kB)
// NAME-service/svc/transport_http.gotemplate (106B)

package template

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _handlersHandlersGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\x8c\x03\xef\x57\x73\xc0\x37\x51\x3a\xd5\xe8\xd0\x02\x4a\x3c\xc6\x32\x03\x00\x00\xff\xff\xd6\x21\xab\x2e\x3e\x00\x00\x00")

func handlersHandlersGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersHandlersGotemplate,
		"handlers/handlers.gotemplate",
	)
}

func handlersHandlersGotemplate() (*asset, error) {
	bytes, err := handlersHandlersGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "handlers/handlers.gotemplate", size: 62, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xcb, 0xd5, 0x72, 0x80, 0xc6, 0xf9, 0x82, 0x4b, 0xe0, 0x8b, 0x90, 0xb8, 0x9b, 0xbc, 0x5d, 0x8d, 0x12, 0xd4, 0x8e, 0x54, 0xf6, 0x72, 0xcb, 0xef, 0xf5, 0x12, 0xd0, 0xe1, 0xb8, 0x41, 0xc8}}
	return a, nil
}

var _svcClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x41\x6f\xea\x38\x10\x3e\xc7\xbf\xe2\x13\xea\x01\x9e\xa8\x73\xaf\xd4\xcb\xd2\xee\x53\xa5\xdd\x16\xb1\x68\xaf\x2b\x93\x0c\xc1\x6a\xb0\x53\xdb\xc0\x22\x2b\xff\xfd\xc9\x76\x42\x43\x4b\x79\x3d\xb4\x38\x33\xdf\x8c\x67\xbe\x6f\xc6\x79\x8e\x99\x2e\x09\x15\x29\x32\xc2\x51\x89\xd5\x11\xce\xec\xac\xe5\x78\x78\xc1\xf3\xcb\x12\x8f\x0f\x4f\x4b\xce\xf2\x1c\x0b\x32\x3b\xa5\xa4\xaa\x12\x00\x07\x59\xd7\xd0\x7b\x32\x07\x23\x1d\xc1\x6d\xa4\xc5\x5a\xd6\x14\xc1\xff\x92\xb1\x52\xab\x3b\x78\xcf\xbb\x73\xdb\x0e\x1c\x78\x10\x8e\x86\xde\xf0\xdd\xb6\x2c\x40\xe6\xa2\x78\x15\x15\xa1\x32\x4d\x81\xc6\xe8\xbd\x2c\xc9\x42\xa0\x5a\xcc\x67\x28\x6a\x49\xca\x61\xad\x0d\xdc\x86\x42\x82\x7f\xc8\xec\x65\x41\xfc\x59\x6c\xa9\x6d\x61\xbb\x4f\xd6\x0c\xd2\x30\x26\xb7\x8d\x36\x0e\x63\x96\x8d\x0a\xad\x1c\xfd\xef\x46\x2c\x1b\x55\x5a\x57\x35\xf1\x4a\xd7\x42\x55\x5c\x9b\x2a\x0f\xe8\x11\x0b\x2e\xe9\x36\xbb\x15\x2f\xf4\x36\xaf\xf4\xed\xab\x74\x79\xf8\x23\x55\x36\x5a\xaa\x10\x1c\x90\xce\x08\x65\x63\xe2\x2f\xf0\x27\xc0\x29\x71\x9e\x63\x19\xa8\xea\xca\x66\xd9\xc8\x7b\xfe\x14\xab\x9b\x0b\xb7\xc1\x6d\xdb\x22\xb7\xfb\x62\xc4\xb2\x66\x85\xe0\x9c\xff\x71\xee\x1e\xb1\x49\xa0\x29\x7a\x3a\xae\x52\xef\xc1\xfa\x4c\x07\x18\x72\x3b\xa3\x2c\x84\xea\xd9\xc0\x4a\x14\xaf\x49\xdd\x73\x1e\x0b\xad\x14\x15\x4e\x6a\xc5\xf1\xe4\x20\x6d\x60\x35\xe4\x31\x64\x1b\xad\xac\x5c\xc9\x5a\xba\x23\xf4\x3a\xd2\x5d\x88\xba\x26\x03\xa7\x51\x4a\x51\x4f\x21\x54\x89\x5a\x38\x32\x28\x6a\x6d\x69\x9a\x40\xef\x39\xd9\x7a\xa7\x8a\x50\xd3\x38\x18\xf1\x23\x90\xc0\x67\xf1\xea\x99\x56\x6a\x0a\xdd\x04\x9c\x05\xe7\xfc\x8c\xcf\x0e\xf4\x12\xdd\x13\x8c\xed\xbe\xe0\x8f\x1d\xf7\x76\x0a\x32\x46\x9b\x09\x3c\x63\x99\xf7\xb7\x38\x48\xb7\xc1\x8d\x23\xdc\xdd\x83\xb7\x2d\xcb\x06\xd6\xe6\xb5\x0a\xec\x04\xd7\x8d\xa3\x4f\x94\xb1\x2c\x81\x8d\x50\x15\xe1\x46\xf6\xb8\x7e\xaa\xfe\x26\xb7\xd1\xa5\x8d\x49\xb3\x6c\x2f\x0c\xbc\x5f\xea\xbf\xf4\x81\x0c\x6e\x64\x37\x74\x7d\x65\xe8\xc7\xe3\x54\x6b\x8c\xf2\xf1\x7f\x76\x2d\xf0\x1e\xe7\xed\x3f\xd3\x21\x31\x30\x4e\xb1\x59\xe0\x6f\xda\x9d\x47\xde\xf7\x6d\xb5\x2d\xf7\x7e\x58\x6f\x32\x8e\x86\x50\xf9\xd1\xf8\xa8\x0a\x5d\xd2\xcf\xc5\x7c\x36\xf0\x2e\xe8\x6d\x47\xd6\xf5\x98\x07\xba\x88\x89\x43\x41\x3d\xa8\x59\x71\xef\x7f\xea\xc8\xef\x8d\xe4\xbd\x7b\x79\x6c\xfa\x42\x7c\xdb\x63\x3b\xa5\x39\xe7\x9d\x65\x72\x22\x69\x3c\x89\x96\x36\x69\x41\xaa\xec\x24\xec\x4e\xfd\x81\x65\x69\xb0\x71\x36\x0c\x3e\x00\x86\xea\x7d\x94\x2e\xac\x4c\x4c\xf7\x89\xf5\x3b\x00\xd7\xe4\x9c\xbe\xdf\x9d\xb5\x53\x28\x59\xb3\xf4\x42\x05\x5a\x90\xf4\x41\x22\x8a\x5d\xaf\x21\xed\xe6\x55\x4e\xc3\xee\x09\x9c\xbf\x17\x3c\x45\xf4\x90\x3f\xc3\x32\xb9\x8d\x88\x6b\xbb\x27\xe3\x2c\x44\xc8\x1b\x17\xfa\x42\x1f\x30\xd4\xd4\xc7\xb0\xad\x02\x3b\x4b\xe6\xb6\xd4\x5b\x21\xd5\x17\xd0\x74\x07\xc7\xdc\xc8\xad\x30\xb2\x3e\x86\x98\xf5\xae\x86\x54\x10\xdd\x7b\xd1\x2d\xf4\xd5\x46\xc6\xff\xa1\x7b\x61\xf9\x2c\xfd\x4e\xe3\x70\x2f\x62\x31\x52\x39\x32\x6b\x51\x90\x6f\x27\x18\x0f\xbe\x06\x4b\x9d\xa5\xba\xef\xee\xdf\xe3\xf8\xf8\xc7\xef\x87\x6d\x72\x9a\x90\x98\xa0\x57\xec\x34\x3f\x1f\x94\x4b\x6b\xf0\x2d\xe5\xae\x6d\xcc\x45\xe1\x52\x40\x87\xf8\x4a\xb7\xdf\x6b\x92\x2e\x88\x02\x5e\x11\x39\xa2\xbe\x25\xdc\xb5\x3e\x2e\xe9\xd6\x57\xf0\x4d\xd5\xde\x02\x83\x7d\x3d\x17\x14\x8b\x8e\x2f\x04\x7b\xfb\x24\xd7\xaf\x00\x00\x00\xff\xff\xbd\x85\xb8\x4c\x9d\x08\x00\x00")

func svcClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcClientGrpcClientGotemplate,
		"svc/client/grpc/client.gotemplate",
	)
}

func svcClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := svcClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/client/grpc/client.gotemplate", size: 2205, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x25, 0xad, 0xc2, 0x90, 0x30, 0x6b, 0x57, 0xe0, 0xce, 0xad, 0x6b, 0x3b, 0x24, 0x4d, 0x91, 0x56, 0x25, 0x72, 0xd3, 0xcb, 0x64, 0xab, 0x95, 0x34, 0x88, 0x71, 0x50, 0x27, 0x48, 0xf6, 0x80, 0xbe}}
	return a, nil
}

var _svcClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func svcClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcClientHttpClientGotemplate,
		"svc/client/http/client.gotemplate",
	)
}

func svcClientHttpClientGotemplate() (*asset, error) {
	bytes, err := svcClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/client/http/client.gotemplate", size: 105, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa1, 0xf0, 0x36, 0xf9, 0x16, 0xea, 0x9d, 0x4e, 0x73, 0x64, 0xc5, 0xad, 0xb3, 0x1b, 0x4, 0xe, 0xd8, 0xc8, 0x1e, 0xf7, 0x7a, 0x39, 0x40, 0x4c, 0xb2, 0x12, 0x83, 0x35, 0xca, 0x82, 0x6f, 0xd0}}
	return a, nil
}

var _svcEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcf\x6f\xdc\xba\x11\x3e\x8b\x7f\xc5\x64\xe1\xc2\xbb\x81\x2c\xdf\x1d\xec\xa1\x4d\xdc\xd6\x40\xe2\x04\xb1\xdb\x1e\x82\x20\xe0\x4a\xb3\x2b\xc2\x14\xc9\x90\xd4\xfe\xa8\xa0\xff\xbd\x18\x52\xd2\x6a\xb3\xb2\xeb\xf7\x8e\x0f\xef\x60\xd8\x26\x87\xc3\x6f\xbe\x6f\x66\x38\xba\xbe\x86\xf7\xba\x40\xd8\xa0\x42\xcb\x3d\x16\xb0\x3a\x80\xb7\xb5\x73\x19\x7c\xf8\x0c\xf7\x9f\x1f\xe1\xf6\xc3\xdd\x63\xc6\xae\xaf\xe1\x2b\xda\x5a\x29\xa1\x36\xd1\x00\x76\x42\x4a\xd0\x5b\xb4\x3b\x2b\x3c\x82\x2f\x85\x83\xb5\x90\x18\x8c\xff\x8d\xd6\x09\xad\x6e\xa0\x69\xb2\xee\xef\xb6\x1d\x6d\xc0\x07\xee\x71\xbc\x4b\xff\xb7\x2d\x63\x86\xe7\x4f\x7c\x83\xe0\xb6\x39\x23\xfb\xc7\xde\x2d\xe4\x5a\x79\x2e\x94\x83\x0a\x7d\xa9\x0b\x07\x5e\x43\xc5\x9f\x10\x84\x2a\xc4\x56\x14\x35\x97\x80\xaa\x30\x5a\x28\xef\x60\x6d\x75\x05\x0e\xed\x56\xe4\xe8\x52\xf2\x64\xf1\x67\x8d\xce\x03\x57\x05\x58\x74\x46\x2b\x87\xe0\x0f\x06\x83\x27\x32\xa5\x20\xb4\xc3\xa3\x97\x14\xb8\x83\x1d\x4a\x49\xbf\x51\xe5\xba\x40\xeb\xc8\x01\xf9\x2b\xb0\xfb\x7f\xad\x6d\x77\x30\x78\x4b\xc3\x02\x27\x72\xd6\xa0\x6b\x0b\xae\x36\x46\x5b\x22\xd7\x5b\xae\x1c\xfd\x4d\xd7\x09\x2e\xc5\x7f\xb9\x17\x5a\x91\xb7\xb5\xb6\x15\xf7\x2e\x63\x4c\x54\xc1\x62\xce\x92\xd9\xba\xf2\x33\x96\xcc\x28\x72\xdc\xfb\x19\x63\xc9\x6c\x23\x7c\x59\xaf\xb2\x5c\x57\xd7\x1b\x7d\xf5\x24\xfc\x35\xfd\xf4\x88\xc9\xc4\xac\x60\xd6\x34\xd9\x97\xbf\xdd\x05\x47\x5f\xb8\x2f\xe1\xaa\x6d\x67\x6c\x11\x08\xbd\x1d\x28\xca\xb5\x94\x98\x7b\xd7\x63\xf5\xe5\x28\x74\xf0\x25\xf7\x90\xeb\xca\x50\x60\x5c\x01\x2f\x8a\x9e\xcf\x0c\xee\xfc\xa5\x23\x67\x15\x72\xe5\x89\xbe\x15\x42\xed\xb0\x20\x9e\x38\x94\x28\x0d\x5a\x70\xde\xd6\xb9\x4f\x69\xbb\xbb\x6a\xfa\x26\xa1\xbc\x06\x4e\xee\x9c\x50\x1b\x89\x60\xb8\xe5\x15\x7a\xb4\x94\x4a\xb4\x7e\xa7\x80\x47\x85\x6c\x0a\xc2\x5f\x3a\xba\x6c\x5d\xcb\xc0\xf4\xba\x56\x39\xb1\xd8\x41\x56\x48\x44\x6b\xd0\x26\x64\x34\x68\x3a\x6b\xd0\x5e\xf5\x17\x92\xc3\x15\x77\xc2\x65\xf0\x77\x6d\x01\xf7\xbc\x32\x12\x53\x38\xe8\x1a\x2a\xb1\x29\x3d\x18\xee\x48\xe5\x11\x55\x04\x70\xb8\x28\xde\x63\xac\x2e\xea\x1c\x03\x0d\x5c\x41\xe9\xbd\xc9\xfe\xc9\x55\x21\x09\xe3\x4e\xf8\x12\x90\xe7\x65\x97\xac\x30\xef\x6f\x5f\xc0\x4e\x58\x2c\xa0\x36\xd1\xa9\x33\x98\x8b\xb5\xc8\xc1\x70\x5f\x66\x30\xbf\x0b\xf8\x84\x23\xff\x2b\xbe\x92\x07\xe0\x50\x09\xe7\x63\xa2\x43\x81\x4e\x6c\x14\x1d\x15\x6a\xab\x9f\x30\x50\xf9\x10\x65\x19\x0a\x23\x40\xc4\x53\xb1\xa3\x18\xe4\xa2\x67\x32\x5b\x8c\xd9\xcd\xa5\x40\xe5\x4f\xd9\x1d\x09\x77\xac\x31\x79\xa0\x4a\x8c\xee\xb0\x78\x49\x46\xaa\x86\xc8\x95\x20\x86\x2b\x8c\x69\x75\xc4\x2b\x94\x47\xbb\xe6\x94\x50\xd3\x4a\x90\xb3\xe1\xb2\xe9\x3a\xaf\x5d\xec\x48\x5d\x61\x5d\x07\x1d\xee\x71\xf7\xbe\x8b\x27\xd7\xd5\x4a\xa8\xc0\x53\xd5\x41\x1c\x09\x9b\x76\xdd\xc0\xd7\x56\x81\x08\x99\x4c\x00\x73\x2e\x25\xda\x98\xcc\x1d\xd8\x8c\x85\x70\xce\x08\x6d\x58\xd3\x58\xae\x36\x08\x17\x02\x6e\x96\x90\xf5\xf6\x9f\xa2\x18\x6d\xcb\x92\xa6\xb9\x10\xd9\x3d\xaf\xb0\x6d\xfb\xf3\x00\x30\x04\x91\xf5\x8b\xac\x69\xae\x68\xb5\x6d\x59\xcb\x18\xa5\x1b\xdc\xe3\x6e\xb8\x72\xde\x95\x1f\x98\x55\xd6\x34\xc3\x45\xd1\xf1\x43\xd0\x74\x31\x02\xd8\x30\x96\x8c\x52\x00\x0a\x5d\x71\xa1\x32\x96\x6c\xb9\xa5\xee\xf2\x32\x6e\x6a\x19\x2c\x49\x9a\xe6\x51\x7f\xd4\x3b\xb4\x70\x1e\xc2\x12\x3e\xf1\x27\x9c\x08\xae\x07\xba\xa0\x4b\x62\x3c\xc9\x82\xb1\xe4\xa8\xda\xcd\xf2\x08\xb4\x79\x35\x94\xb3\x8b\x6e\x88\xc6\x17\x20\xa6\x23\x00\x2d\x63\x49\x27\xf4\x80\x83\x68\x1e\x57\xc9\x2b\xb4\x0c\xaa\xcc\x47\x89\xb0\x80\x11\xb0\x79\xee\xf7\xd0\xb5\xeb\xec\x7d\xfc\x9d\x52\xd1\xbd\x0d\x9a\xfd\x43\x93\x19\xe1\xfc\x1a\x1f\xa3\xc7\x83\xe9\x05\x5c\xc0\xfc\xdc\x28\xbe\x52\x23\xab\x14\xd0\x5a\x6d\x17\xd0\xb0\x24\xe9\x5f\xb1\xb0\x48\x80\x31\x9b\x52\x23\xf7\x7b\xc2\xb0\x60\x49\x22\xd6\xc1\xf4\xcd\x12\x94\x90\xc1\x47\xcf\x89\x12\x32\xb8\x61\x49\xd2\xb2\x61\xb5\xbf\x21\x7b\x0d\xb6\x45\x4a\x5e\x88\xe9\x9e\x74\x22\x97\x72\xe4\x84\xe1\xd0\x1b\x2f\x3c\x06\x86\x63\x79\x8c\x49\xbf\xf0\x38\xc5\x7b\x24\xfe\xd9\x84\x8b\x35\x31\x3e\x7b\x5a\x16\x67\xb5\x76\x12\x3c\xf9\x9e\x96\xae\x1f\x1a\x86\x56\xd5\x90\x50\xc3\xf8\x30\x5a\x8e\x22\x8c\xd4\x21\xef\x3f\x29\xa2\xce\xc7\x14\x87\x67\x49\x10\xce\x6d\x07\x41\x5d\xf6\x4b\x72\x05\x44\xd1\x6a\x42\xcb\x29\x35\xa3\x9e\xc3\xce\xb6\x13\x29\x2e\xb7\xc7\x02\x19\x6b\xf6\x1f\xcb\xcd\x5f\xa5\xbc\xdd\xe7\x68\x3c\xec\x2c\x37\x2e\xbe\x66\x03\x7b\x6b\x81\xb2\xa0\xa7\xbc\x6b\x83\xc7\xb6\x13\xe4\x0d\xcf\xc0\xc4\x7c\x92\x7d\x12\x45\x21\x71\xc7\x6d\x1c\x13\xff\xe5\xfa\xc1\x91\x46\x26\x63\xe4\x81\xba\x39\xbd\x50\x9e\x9c\x57\x83\x75\x78\x82\x71\x8b\xf6\x30\x48\x49\x65\x45\xcd\xba\x1f\x4a\xc8\xdf\x67\x43\x0f\x34\x3d\x52\xe9\xe8\x8d\xc8\xb9\xa2\x01\x85\x9e\x75\x2c\xe8\xd8\xea\x00\x8a\x34\x88\x83\x0b\xee\x73\x59\x17\x58\xc4\x99\x71\x85\x04\x81\x62\x36\x58\x64\x67\x6c\xcc\x8f\x98\x52\x98\x3d\x78\xee\x6b\x37\x4b\x61\xf6\x45\xa8\xcd\x6c\xc1\xfa\xf6\xf0\x76\xd4\x1f\x9e\x3b\x0f\x13\xac\xa4\x47\x34\x59\x96\x39\x6f\x85\xda\x84\x74\x12\xaa\x5b\xbe\x59\x42\xc5\xcd\xb7\xb8\xf5\x3d\xd2\xdf\xb4\x4d\x68\x90\x57\xf0\xff\xda\x57\x92\xcc\x46\x19\x35\xbb\x81\xa6\x4d\xbb\xa3\xa3\x3e\x49\x6a\xfc\x20\x28\x21\x7d\x83\xcb\x01\x56\x13\xdb\xc8\x8f\x14\xf4\x13\x6d\xf7\xc0\xbe\xe1\xfe\xfb\x3b\x78\xa3\x9f\x62\x2a\x1a\xae\x44\x3e\x5f\x57\x3e\x7b\x30\x56\x28\xbf\x9e\xcf\x6e\x7b\x17\x83\x82\x97\x7f\x71\x97\x50\x68\x74\xa0\xb4\x07\xdc\x0b\xe7\xdf\x81\x43\x1c\x0b\x3f\xe4\x8e\xcb\x36\x7a\x46\xa0\x16\x8b\xae\x49\x15\x28\xd1\xe3\xbc\x47\x10\xf6\x8e\x01\x08\x95\x1f\xe1\x0f\xf4\xbd\x9e\x28\xb1\x0e\x2e\x96\x4b\x38\xa1\xac\xab\xb4\xc9\x56\x0b\xcb\x11\xf2\xf9\xa4\xc9\xa2\x2f\xbd\x13\xca\x63\xd9\x7d\xe4\x2b\x94\x58\x1c\xb3\x21\x7e\x63\x6d\xd0\xf7\xb9\x3b\x1e\x9c\x63\x0a\xef\x4a\x54\xc3\xae\x1e\xa5\x6b\xe7\x2c\x66\x5d\x1a\xab\xac\x2b\x84\x3a\x1a\x43\xfc\x70\xe3\xf1\xeb\x4f\xe4\x34\x3f\x5a\x91\xc7\xc1\x7e\x84\xa1\x14\x79\x19\x8e\x3a\x54\x53\x10\xba\xa1\xa9\x3b\xdd\x8f\x8c\xda\x76\x23\xd3\x79\x54\xa1\xdd\xc6\x04\x4e\xcf\x3b\xf3\x44\xb3\x66\xcf\xc5\xf5\xbb\x7b\xd3\x19\xa8\xb4\x8b\x33\x30\x6e\x31\x47\xb1\x8d\xc3\x75\x08\xf1\x97\x6f\x96\x0c\x1e\x10\x07\x37\x23\x2f\x61\xa3\x9f\xf9\x87\xba\x27\xa0\x94\x91\x05\x7a\x2e\x64\x98\xcf\xfb\x72\x0a\x9f\x7e\xdd\x77\x05\x97\xc2\x1f\xb2\x97\x5a\xc8\x49\xec\xe3\x4e\xf2\x9b\x19\xfd\xb3\xcf\xfc\x71\xfa\xcc\xc9\xb1\x74\x7a\x08\x7c\xae\xed\xfc\x2f\x00\x00\xff\xff\xad\x43\xd8\xc0\x01\x12\x00\x00")

func svcEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcEndpointsGotemplate,
		"svc/endpoints.gotemplate",
	)
}

func svcEndpointsGotemplate() (*asset, error) {
	bytes, err := svcEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/endpoints.gotemplate", size: 4609, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0xf9, 0x28, 0x6, 0x55, 0x9b, 0xe2, 0x56, 0x63, 0xa0, 0x59, 0xb, 0x8c, 0x89, 0x39, 0xaa, 0x3, 0x39, 0xd1, 0x27, 0xdf, 0x37, 0x9b, 0xe2, 0xf0, 0x89, 0xc6, 0x79, 0x2d, 0x85, 0x2b, 0xf1}}
	return a, nil
}

var _svcTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x28\x16\x56\xe1\xd0\x7b\x0e\x90\x4b\x93\x6e\x5b\xec\xe6\x03\x59\xa3\x7b\x28\x8a\x82\x96\xc6\x12\x61\x89\x54\x48\xca\x89\x97\xd0\x7f\x5f\x0c\xf5\x61\x39\x76\x1c\x1f\x0c\x58\xe4\xe3\xcc\x9b\xf7\x86\x23\xcd\xe7\x70\xad\x53\x84\x0c\x15\x1a\xe1\x30\x85\xe5\x16\x9c\xa9\xad\xe5\x70\x73\x0f\x77\xf7\x0b\xf8\x7c\xf3\x6d\xc1\xd9\x7c\x0e\x8f\x68\x6a\xa5\xa4\xca\x5a\x00\x3c\xcb\xa2\x00\xbd\x41\xf3\x6c\xa4\x43\x70\xb9\xb4\xb0\x92\x05\x06\xf0\x77\x34\x56\x6a\x75\x09\xde\xf3\xee\x7f\xd3\x8c\x36\xe0\x46\x38\x1c\xef\xd2\x73\xd3\x30\x56\x89\x64\x2d\x32\x04\xbb\x49\x18\xe1\x17\x7d\x58\xa8\x8c\xde\xc8\x14\x2d\x58\x34\x1b\x34\x17\x56\xa6\x08\x4b\xa9\x52\xa9\x32\x0b\x2b\x6d\xc0\xe5\x08\xd9\xe3\xc3\x35\x38\x23\x94\xad\xb4\x71\x81\xcb\x37\x07\xb5\x93\x85\xfc\x0f\x6d\x80\x0c\xbb\xf3\xcc\x54\x09\xff\x27\x84\xe3\x8c\xc9\x92\x16\x61\xca\xa2\x89\x42\x37\xcf\x9d\xab\x26\x2c\x9a\x24\x5a\x39\x7c\x71\x13\xc6\xa2\x49\xa6\x75\x56\x20\xcf\x74\x21\x54\xc6\xb5\xc9\x42\x88\x79\x89\x4e\xa4\xc2\x09\xc2\xd0\xc2\x90\x01\x26\x99\x74\x79\xbd\xe4\x89\x2e\xe7\x99\xbe\x58\x4b\x37\xa7\xdf\x3e\x05\x3a\xd6\x97\x4a\x6c\x64\x82\x2c\xaa\x96\x30\xf1\x9e\x3f\x7c\xfa\x16\x68\x3d\x08\x97\xc3\x45\xd3\x4c\x58\x1c\x74\xb9\x15\x6b\xfc\xf2\xf8\x70\xdd\xb2\x87\x52\xac\xd1\x82\x00\x8b\x0e\xf4\x0a\x50\xa5\x95\x96\xca\x59\x10\x1b\x21\x0b\xb1\x2c\x10\x04\xed\x07\x79\xbc\xe7\x5d\x1a\x7e\x27\x4a\x6c\x9a\x5e\x82\x55\xad\x92\x57\x91\xa7\xbb\x50\x9f\xfb\x7f\x33\xd0\x95\x93\x5a\x59\xe0\x9c\xef\xd5\xdb\x89\x79\x1f\xb6\x63\xa8\x96\xfc\x8d\x5c\xe0\x59\x64\x47\x58\x0b\x97\x57\xf0\xe3\xe7\xdb\xc1\x3c\x8b\xa2\x63\xbb\x9f\x70\xa5\x0d\x4e\x7b\x07\x16\xfa\xba\xb5\x2b\x9e\xb1\xa8\x79\x9d\xe3\x0a\x44\x55\xa1\x4a\xa7\x7b\xcb\x43\x39\x9c\xf3\x98\x45\x06\x5d\x6d\x14\xfc\x4e\xd9\xda\x1c\x3e\xd8\xe3\x3d\x2c\xf4\xdf\xfa\x19\x0d\xec\x95\x04\x4d\xc3\x22\xef\x8d\x50\x19\xc2\x07\x49\x85\x0c\xfb\xb7\xe8\x72\x9d\x5a\x42\x44\xde\xf7\xc7\x3f\xc8\x4e\x8b\x4b\xd8\x2f\xe9\x0e\x9f\x3b\xd5\x59\x14\x45\x83\xf2\xdc\xfb\xe1\x48\x6f\xc2\x8c\x10\x37\x98\xe8\x34\x98\x35\x42\x3c\xe2\x53\x8d\xb6\x05\x7c\x56\x47\x01\xb6\xd2\xca\x62\x40\xec\x29\xc1\x39\xa7\x45\xd2\xce\xfb\x0b\xea\x22\x62\xde\xb0\x26\xb4\xdc\x4e\x10\x90\x65\x55\x60\x89\xd4\x15\x74\xa3\xbc\xff\xa2\x83\x14\xc7\xbd\x96\xca\xa1\x59\x89\x04\x99\xdb\x56\x38\x8e\x63\x9d\xa9\x13\x07\x9e\xbd\xaf\xdf\x11\xf9\x00\x5e\xe9\xf7\x55\xa8\xb4\x40\xc3\x76\xe4\x5b\xe6\x5d\x98\x30\x24\x46\xd9\x9d\xde\x15\x72\x7e\x0d\xef\x52\x0d\xb7\x68\x6a\xe1\xe3\x2e\x55\xbc\x0b\x3f\xb0\x9f\x26\xee\x05\xba\xe1\xc2\xbb\xae\x9d\x81\xc1\x27\xf8\x18\xee\xcd\x0e\xdf\x39\xba\xd8\x56\x3d\xa9\x18\xa6\x87\xa0\xd6\xd5\x11\x6a\x06\x68\x8c\xa6\xe4\x2c\xfa\x45\xa1\xab\xb0\x42\xb4\xa9\xa7\x0e\xf4\x6c\xaf\x14\x75\x0b\x71\x0b\x5c\x62\x16\xc9\x55\x38\xf4\xdb\x15\x28\x59\x50\xa8\xfe\x86\x28\x59\x84\x78\xe1\xa2\x75\x6b\x06\x2b\x7e\x0e\xb5\x78\x46\xc7\x59\xc3\xbc\x6f\x8d\x22\x9b\x3a\xa9\xdb\xae\x7e\x5f\xe7\xf9\x1c\x4e\x5d\x00\x90\x34\xf0\x5e\x0d\xfb\xf6\x40\x87\xf8\x93\x8c\x72\xb9\x70\x64\xc3\x06\x0d\x8d\xcb\xd0\xe8\xed\x90\x3c\xec\x37\xd3\x45\x76\x1a\x04\xd4\x16\xcd\x45\xaa\x4b\x21\xd5\x29\x30\x87\x07\x23\x4b\x61\x64\xb1\xa5\x23\xab\xba\x00\xa9\xc2\xa4\x1e\xcd\xdc\x53\x75\x4c\x7f\x1d\x76\x09\xd5\xf2\x88\x4f\xbb\xae\xf4\xd4\x12\xa3\xa7\xb1\xf5\xd4\x52\x97\x57\xfd\x99\x63\xf6\x1c\xb4\xd7\xc8\xcf\xa7\x13\x4e\xb5\xe3\xe5\x2c\xa7\x4e\x4e\xa2\xa3\x56\xb5\x27\x7a\xc8\x5b\x5e\xbd\xef\x42\x97\x22\x78\x76\xc2\xd9\xaa\xd8\x9e\x65\xd5\xc9\x42\x8e\x79\x35\x30\x38\xd3\x2c\x5b\x91\x8a\xfd\xa9\xf3\x6e\xd3\xc8\x2f\x5b\x1d\x33\xec\x2b\x16\x15\x1a\xcb\xda\x1a\x0e\xde\x96\xc7\x67\x51\x99\x0e\x48\x7e\x7b\x13\xbf\x06\x10\x5d\x9a\xa8\xeb\x19\x6c\x02\xe5\xd0\x04\x65\x1a\x66\x84\x5c\xc1\x66\x3c\x33\xda\x0f\x1c\x84\x35\x6e\x83\xdb\x69\x4a\x1f\x9b\xda\xe5\x24\x71\x9f\x85\x06\x74\x29\x1c\x4c\xd7\x31\x3c\xe7\x32\xc9\x03\xb4\x28\xa0\x20\xbb\xba\x28\x42\xa5\xe1\xa5\x43\xdf\x67\xfc\x5a\x28\xad\x64\x22\x8a\xaf\x28\x52\x34\x7f\xe1\x96\x3e\x7f\x5c\x97\xc8\xea\xb6\x65\xa4\x83\x44\x28\x58\x62\x1f\x22\x49\xd0\x5a\x4c\x29\x37\x4a\x97\xa3\xe9\x32\xd3\x3e\x49\x71\x35\xd4\xfa\xaf\x74\xf9\x77\x51\xd4\xd8\x8e\x44\xaa\xf5\xc7\x1f\x3f\xe3\x77\x81\x6f\xb0\x9b\xae\xe3\x5d\x84\xf0\x6e\x1d\xac\x4b\xdc\x0b\x6b\xd8\xff\x01\x00\x00\xff\xff\x71\x92\xdd\x9a\x92\x0b\x00\x00")

func svcTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_grpcGotemplate,
		"svc/transport_grpc.gotemplate",
	)
}

func svcTransport_grpcGotemplate() (*asset, error) {
	bytes, err := svcTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_grpc.gotemplate", size: 2962, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xc7, 0xff, 0xf, 0x16, 0xfd, 0xd8, 0x43, 0xfc, 0x25, 0x2b, 0x9c, 0xfe, 0x3c, 0xf, 0x3a, 0x3e, 0xc7, 0x21, 0xdf, 0x32, 0x41, 0x25, 0x39, 0x7d, 0x4b, 0xb6, 0xc2, 0xc, 0x27, 0x81, 0x3c}}
	return a, nil
}

var _svcTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x18\xa9\x9a\xdc\xcd\xe5\xb0\x1b\x64\xd5\x88\x65\xce\xa3\xb4\x26\x69\xcb\xf9\xba\xa1\x75\x78\xda\xe1\x5f\x78\xfe\x1b\x49\x11\xcb\x2f\x00\x00\xff\xff\xdd\x3a\x4a\x8f\x6a\x00\x00\x00")

func svcTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_httpGotemplate,
		"svc/transport_http.gotemplate",
	)
}

func svcTransport_httpGotemplate() (*asset, error) {
	bytes, err := svcTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_http.gotemplate", size: 106, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x57, 0x56, 0xc6, 0xb4, 0xe5, 0x1f, 0xf4, 0x1d, 0xa5, 0xda, 0x23, 0xea, 0x8f, 0xfb, 0xff, 0xae, 0x4b, 0x12, 0xe4, 0xf6, 0xbf, 0x11, 0xa6, 0x4, 0x83, 0x53, 0xfd, 0xbf, 0xce, 0x4a, 0x47}}
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
	"handlers/handlers.gotemplate":      handlersHandlersGotemplate,
	"svc/client/grpc/client.gotemplate": svcClientGrpcClientGotemplate,
	"svc/client/http/client.gotemplate": svcClientHttpClientGotemplate,
	"svc/endpoints.gotemplate":          svcEndpointsGotemplate,
	"svc/transport_grpc.gotemplate":     svcTransport_grpcGotemplate,
	"svc/transport_http.gotemplate":     svcTransport_httpGotemplate,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"handlers": &bintree{nil, map[string]*bintree{
		"handlers.gotemplate": &bintree{handlersHandlersGotemplate, map[string]*bintree{}},
	}},
	"svc": &bintree{nil, map[string]*bintree{
		"client": &bintree{nil, map[string]*bintree{
			"grpc": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{svcClientGrpcClientGotemplate, map[string]*bintree{}},
			}},
			"http": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{svcClientHttpClientGotemplate, map[string]*bintree{}},
			}},
		}},
		"endpoints.gotemplate":      &bintree{svcEndpointsGotemplate, map[string]*bintree{}},
		"transport_grpc.gotemplate": &bintree{svcTransport_grpcGotemplate, map[string]*bintree{}},
		"transport_http.gotemplate": &bintree{svcTransport_httpGotemplate, map[string]*bintree{}},
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
