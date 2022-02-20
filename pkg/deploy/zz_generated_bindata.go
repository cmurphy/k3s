// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/ccm.yaml
// manifests/coredns.yaml
// manifests/local-storage.yaml
// manifests/metrics-server/aggregated-metrics-reader.yaml
// manifests/metrics-server/auth-delegator.yaml
// manifests/metrics-server/auth-reader.yaml
// manifests/metrics-server/metrics-apiservice.yaml
// manifests/metrics-server/metrics-server-deployment.yaml
// manifests/metrics-server/metrics-server-service.yaml
// manifests/metrics-server/resource-reader.yaml
// manifests/rolebindings.yaml
// manifests/traefik.yaml
// +build !no_stage

package deploy

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ccmYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xc1\xab\xdb\x30\x0c\xc6\xef\xfe\x2b\x4c\x2f\x85\x81\x5b\xc6\x2e\x23\xc7\xed\xb0\x7b\x61\xbb\x2b\xb6\xd6\x7a\x75\x2c\x23\xc9\x29\xdb\x5f\x3f\xd2\xf4\x3d\x4a\xf2\x5a\x92\xde\x44\x90\x7e\xdf\xa7\xc8\x1f\x94\xf8\x0b\x59\x22\xe5\xc6\x72\x0b\x7e\x07\x55\x4f\xc4\xf1\x1f\x68\xa4\xbc\x3b\x7f\x95\x5d\xa4\x7d\xff\xd9\x9c\x63\x0e\x8d\xfd\x9e\xaa\x28\xf2\x81\x12\x9a\x0e\x15\x02\x28\x34\xc6\xda\x0c\x1d\x36\xf6\xfc\x45\x9c\x4f\x54\x83\xf3\x94\x95\x29\x25\x64\xd7\x41\x86\x23\xb2\xe1\x9a\x50\x1a\xe3\x2c\x94\xf8\x83\xa9\x16\x19\x06\x9d\xf5\x44\x1c\x62\xbe\xd7\x33\xd6\x32\x0a\x55\xf6\x78\x6b\x4a\x08\x82\x62\xac\xed\x91\xdb\xdb\xb7\x23\xea\x08\x60\x04\xc5\x6b\x59\x4b\x18\xca\x99\xc6\x66\x33\x47\x62\x8f\x59\x27\xc8\x3b\x54\x01\xf5\xa7\xd5\xd0\x4c\x61\x6a\x73\xfb\x69\xbb\x62\x76\x2f\x0a\x5a\x27\x88\xd1\xcb\x22\x88\x20\xf7\xd1\x4f\x3d\xa4\x28\xfa\xf1\x56\x43\x79\x59\x8d\x07\xef\xa9\x3e\xfa\x7b\x8b\x40\x65\x78\x74\xa2\x98\xb5\xa7\x54\xbb\x47\xb7\x7d\x37\xfe\x9a\x5d\xcc\xa1\x50\x7c\x76\xe6\x99\xd0\x65\x76\x77\xe7\xcc\xeb\x29\xf9\x16\x73\x88\xf9\xb8\x3a\x2c\x94\xf0\x80\xbf\x87\xee\xb7\x35\x9f\x28\x1b\x6b\xe7\xf1\x5c\xa4\x23\xb5\xfd\x83\x5e\xaf\xb9\x1c\x11\x3f\x05\x79\xd9\xec\xd8\x24\x05\xfc\xd0\x59\x5b\x74\xf2\x57\x14\x3b\xf3\x3f\x00\x00\xff\xff\x37\xa3\x5f\x88\x54\x04\x00\x00")

func ccmYamlBytes() ([]byte, error) {
	return bindataRead(
		_ccmYaml,
		"ccm.yaml",
	)
}

func ccmYaml() (*asset, error) {
	bytes, err := ccmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ccm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdf\x6f\xdb\x38\x12\x7e\xf7\x5f\x41\xe8\x90\xb7\x93\x13\x5f\xd0\x5e\x8e\x6f\xa9\xed\xb6\x01\x1a\xd7\xb0\x9d\x02\xc5\x61\x11\xd0\xe4\xd8\xe2\x86\xe2\x70\x49\xca\x8d\xb7\x9b\xff\x7d\x41\xfd\xb2\xe8\x28\xd9\x34\xdb\xf5\x8b\x45\x0e\xe7\x1b\xea\xe3\xf0\x9b\x11\x33\xf2\x0b\x58\x27\x51\x53\xb2\x1b\x0d\xee\xa4\x16\x94\x2c\xc1\xee\x24\x87\x4b\xce\xb1\xd0\x7e\x90\x83\x67\x82\x79\x46\x07\x84\x68\x96\x03\x25\x1c\x2d\x08\xed\xea\xb1\x33\x8c\x03\x25\x77\xc5\x1a\x52\xb7\x77\x1e\xf2\x41\x9a\xa6\x83\x2e\xb4\x5d\x33\x3e\x64\x85\xcf\xd0\xca\xdf\x99\x97\xa8\x87\x77\x17\x6e\x28\xf1\x74\x37\x5a\x83\x67\x4d\xe4\xb1\x2a\x9c\x07\xbb\x40\x05\x51\x58\xc5\xd6\xa0\x5c\x78\x22\x65\x1c\xab\xc1\x43\xe9\xbf\x46\xf4\xce\x5b\x66\x8c\xd4\xdb\x2a\x50\x2a\x60\xc3\x0a\xe5\x5d\xbb\xdf\x6a\x57\xb4\xd9\xb6\x2d\x14\x38\x3a\x48\x09\x33\xf2\x83\xc5\xc2\x94\xc8\x29\x49\x92\x01\x21\x16\x1c\x16\x96\x43\x3d\x07\x5a\x18\x94\xba\x04\x4b\x89\xab\x98\xa9\x06\x06\x45\xf5\xd0\x92\x10\x86\x3b\xb0\xeb\xda\x57\x49\xe7\xcb\x87\x6f\xcc\xf3\xec\x71\x3c\x21\x1d\xc7\x1d\xd8\x7d\x4d\xc6\x33\xd1\x95\xfc\x4b\xf4\xbf\x4f\xf9\x3b\xa9\x85\xd4\xdb\x88\x79\xa6\x35\xfa\xd2\xbd\xa6\xbf\x0f\x37\x3a\x11\x56\x78\x2c\x8c\x60\x1e\x28\x49\xbc\x2d\x20\xf9\xf9\x07\x88\x0a\x16\xb0\x29\xf7\x57\x53\xfa\xcc\x0b\x0f\x08\x79\x9c\x5d\x4f\x20\xbb\x62\xfd\x2b\x70\x5f\x66\x47\xef\x65\x78\xf5\x15\x68\x6f\xd7\x18\xf5\x46\x6e\xaf\x99\x79\xcd\xc5\x6a\x96\x8f\xd1\xc2\x46\x2a\xa0\xe4\x8f\x92\xd3\x21\x7d\x73\x4e\xbe\x97\x8f\xe1\x07\xd6\xa2\x75\xed\x30\x03\xa6\x7c\xd6\x0e\x2d\x30\xb1\x6f\x47\x87\xe3\x20\x27\xdf\xc7\x9f\x6e\x96\xab\xe9\xe2\x76\xf2\xf9\xfa\xf2\x6a\xf6\x70\x42\xa4\x4e\x99\x10\x76\xc8\xac\x61\x44\x9a\xb7\xd5\xc3\x21\x12\x29\xaf\x01\x91\xda\x01\x2f\x2c\x74\xe6\x37\x4c\x29\x9f\x59\x2c\xb6\x59\x3f\x4a\xbb\xf6\xe1\xb0\x51\x74\xde\x91\x53\xf0\xfc\xb4\xa6\xe2\x74\x86\x02\x3e\x96\xd3\xdd\xa0\xde\x2b\xf2\xf6\xac\x33\x61\x41\x21\x13\x64\xf4\xc6\xf5\x6f\xa1\x27\x98\xb1\x98\x83\xcf\xa0\x70\x84\xfe\x6f\xf4\xe6\xbc\x35\x6c\xd0\x7e\x63\x56\x90\x61\xb5\x93\x70\x27\xd5\x6e\xc8\x51\x6f\xda\x25\x9c\xf1\x0c\xc8\xf9\x61\x07\x0a\xd1\x0c\xe2\xcd\x74\x6c\x4c\xac\x99\x62\x9a\x57\xfc\x3c\x3c\x4a\x0e\x66\x8c\x3b\x6d\x33\x64\x02\x46\xe1\x3e\x87\xd7\x69\xef\xd1\x65\xbb\x70\x29\x33\xa6\x5e\x52\x39\x1e\x5f\xc1\x0a\x38\x09\x39\x35\x99\x2d\x93\x81\x33\xc0\x83\xf7\xbf\x2c\x18\x25\x39\x73\x94\x8c\x06\x84\x84\x5b\xea\x61\xbb\xaf\x80\xfd\xde\x00\x25\x0b\x54\x4a\xea\xed\x4d\x79\xdf\x2b\x7d\xe8\xce\xd0\x9a\x83\x9c\xdd\xdf\x68\xb6\x63\x52\xb1\x75\x48\xda\x12\x0e\x14\x70\x8f\xb6\x5a\x93\x07\x11\xfb\xd4\xd9\x78\xff\xd6\x3d\xe4\x46\xb5\xc0\x5d\x76\x4a\xa2\x23\xff\xa7\x5e\xbe\x79\xbd\x2a\x07\x24\x5a\xe9\xf7\x63\xc5\x9c\x9b\x55\x3c\x54\x3c\xa6\xbc\x52\x8b\x94\x5b\xe9\x25\x67\x2a\xa9\x5d\x5c\x24\x08\xb3\xa3\x43\x29\xa9\x41\x05\xb6\xab\x99\xe1\x97\x92\x3b\xd8\x07\x96\x6b\xb8\x4b\x21\x50\xbb\xcf\x5a\xed\x93\x4e\xc6\xa2\x09\x9e\x68\x29\x49\xa6\xf7\xd2\x79\x97\x3c\x02\xd0\x28\x20\x0d\x0a\x78\xa4\xbb\x1c\xb5\xb7\xa8\x52\xa3\x98\x86\x17\x62\x12\x02\x9b\x0d\x70\x4f\x49\x32\xc3\x25\xcf\x40\x14\x0a\x5e\x1e\x32\x67\x81\xa1\x9f\x11\x2b\x44\x58\x46\x09\x11\x7e\xa1\x4e\x1d\x85\x44\x47\x89\x92\xba\xb8\xaf\x17\x85\xb7\x66\x52\x83\x6d\xa9\x4e\x1f\x5d\x94\xea\x27\x73\xb6\x05\x4a\x4e\xbe\x2f\xbf\x2e\x57\xd3\xeb\xdb\xc9\xf4\xfd\xe5\xcd\xa7\xd5\xed\x62\xfa\xe1\x6a\xb9\x5a\x7c\x7d\x38\xb1\x4c\xf3\x0c\xec\x69\x2e\x83\x7c\x82\x48\x6b\x88\xe6\x9f\x8e\x86\x17\xc3\xb7\x31\xe0\xbc\x50\x6a\x8e\x4a\xf2\x3d\x25\x57\x9b\x19\xfa\xb9\x05\x07\x65\xa1\x68\xb4\xa0\x53\xd1\x5b\x45\x90\xb9\xf4\xd1\x4c\x48\xe6\x1c\xed\x9e\x92\xd1\x7f\xcf\xae\x65\xa4\x6c\xbf\x15\xe0\x8e\x57\x73\x53\x50\x32\x3a\x3b\xcb\x7b\x31\x22\x08\x66\xb7\x8e\x92\xff\x93\x24\x0d\x12\x96\xfc\x9b\x24\x91\xc0\x36\xa5\x24\x21\xbf\xb4\x2e\x3b\x54\x45\x0e\xd7\x21\xc1\xa3\x14\x6e\x98\x0d\x15\x2c\xad\x16\x75\xe2\xe7\x61\xfd\x9c\xf9\x8c\x46\x12\x1e\xbd\x0b\x13\x21\xe5\x29\x09\x8d\xc1\x41\x89\xd1\xc6\x71\xda\x53\x9d\xa3\xf5\x94\x74\xb4\xb9\x91\xc1\x18\xd7\x58\xf4\xc8\x51\x51\x72\x33\x99\xff\x28\x4e\xea\xb9\xe9\xc5\x5a\x8d\x9f\xc1\x8a\x2a\x46\x83\x96\x83\xb7\x92\xf7\xef\xac\x8b\x56\x16\xcb\x20\x3b\xa8\x3d\xdc\xfb\xee\xd1\x32\xa5\xf0\xdb\xdc\xca\x9d\x54\xb0\x85\xa9\xe3\x4c\x95\x52\x42\x43\x35\x73\x5d\xba\x39\x33\x6c\x2d\x95\xf4\x12\x8e\x92\x83\x09\x11\x4f\xa4\x64\x36\x5d\xdd\xbe\xbb\x9a\x4d\x6e\x97\xd3\xc5\x97\xab\xf1\x34\x32\x0b\x8b\xe6\xd8\x81\x29\xd5\x73\x70\x0b\x44\xff\x5e\x2a\xa8\xdb\xa6\xf8\x18\x95\xdc\x81\x06\xe7\xe6\x16\xd7\xd0\xc5\xcb\xbc\x37\x1f\xc0\xc7\x21\x4c\x95\x28\x47\xbd\x09\xa9\xd3\x81\x92\x8b\xb3\x8b\xb3\x68\xda\xf1\x0c\x02\xc9\x1f\x57\xab\x79\xc7\x20\xb5\xf4\x92\xa9\x09\x28\xb6\x5f\x02\x47\x2d\x1c\x8d\x7b\x03\x03\x56\xa2\x68\x6d\xa3\xae\xcd\xcb\x1c\xb0\xf0\x07\x63\xc7\xe6\x0a\xce\xc1\xb9\x55\x66\xc1\x65\xa8\x44\x6c\xdd\x30\xa9\x0a\x0b\x1d\xeb\x79\xd4\x61\xc9\x1f\xa6\x22\xee\xcb\x3a\x4c\x8c\x2e\x46\xaf\x66\xe2\x19\x22\xfe\xf3\x0f\xf3\x20\xb4\x6b\xa4\x71\x52\x75\xf4\xb5\xa1\x52\x8e\x1f\x50\x16\xde\xf4\xcc\x31\x6f\xfd\x42\x5f\x52\xe1\x21\x77\xc7\x29\x5d\x16\xb3\x46\xee\x22\x5b\x73\x04\xbd\xc6\xda\xb1\x6d\x44\x7b\x3d\x0f\xd6\x27\x1b\xff\xfa\x4b\xa2\xa7\xa7\xeb\xb4\x27\x4f\x36\x75\x8f\x3e\xc4\x0e\xed\x6b\xa8\x8b\x55\xa6\x24\x41\x95\x92\x1e\xb3\xe3\x96\x99\x27\x3f\xc8\x5e\xd0\x23\x36\xdd\x50\xdd\xfd\x74\x90\x5e\xda\x4d\xc6\xfd\x5e\x5f\xcc\x3a\xc6\xd5\x9c\x76\xbf\x44\x66\xcb\x87\x93\x41\xa7\x46\xa4\x47\x15\xc0\x74\xa5\xfd\xb8\x10\xa4\x3d\x32\xff\x84\x43\xa5\xcf\x69\x8f\x92\x9b\x58\xf0\x63\x97\x3f\x03\x00\x00\xff\xff\x13\xea\x61\x9b\x42\x11\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5b\x6f\xdb\xb6\x17\x7f\xd7\xa7\x38\x7f\xfd\x9b\x87\x0d\xa5\x1c\x6f\x03\x32\xb0\xd8\x83\x9b\x38\x59\x80\xc4\x36\x6c\x6f\x43\x51\x14\x06\x4d\x1d\xc7\x6c\x28\x92\x20\x29\xb7\x6e\x96\xef\x3e\x50\x94\x1d\xc9\x71\x2e\xc6\xb6\xb7\xe9\x45\xe0\xe1\xf9\x9d\xfb\x85\xcc\x88\xdf\xd1\x3a\xa1\x15\x85\x55\x37\xb9\x15\x2a\xa7\x30\x41\xbb\x12\x1c\x7b\x9c\xeb\x52\xf9\xa4\x40\xcf\x72\xe6\x19\x4d\x00\x14\x2b\x90\x82\xd4\x9c\x49\x62\x98\x5f\x12\x63\xf5\x4a\x04\x3c\x5a\xe2\x22\x8e\xb0\x1a\x18\xd9\x9d\x61\x1c\x29\xdc\x96\x73\x24\x6e\xed\x3c\x16\x09\x21\x24\x69\x6a\xb6\x73\xc6\x33\x56\xfa\xa5\xb6\xe2\x1b\xf3\x42\xab\xec\xf6\x67\x97\x09\xdd\x59\x75\xe7\xe8\xd9\xc6\xb0\x53\x59\x3a\x8f\x76\xac\x25\xbe\xde\x2a\x1b\xb8\x6d\x29\xd1\xd1\x84\x00\x33\xe2\xc2\xea\xd2\x38\x0a\x1f\xd3\xf4\x53\x02\x60\xd1\xe9\xd2\x72\xac\x28\x4a\xe7\xe8\xd2\xb7\x90\x9a\x60\x9b\xf3\xa8\xfc\x4a\xcb\xb2\x40\x2e\x99\x28\xaa\x1b\xae\xd5\x42\xdc\x14\xcc\xb8\x0a\xbe\x42\x3b\xaf\xa0\x37\xe8\xc3\xb5\x14\xae\xfa\x7f\x61\x9e\x2f\xd3\x4f\x2f\xab\x44\x95\x1b\x2d\x94\xdf\xab\x36\x12\x75\xbe\xa3\xeb\xfb\x57\x09\x5e\x61\x90\xda\x02\x72\x8b\xcc\x63\x25\x74\xbf\x7d\xce\x6b\xcb\x6e\xb0\x8e\xff\x63\xa1\xf5\x3d\x97\xcc\x39\x7c\x65\x04\xfe\x7e\xb6\xdf\x0b\x95\x0b\x75\xf3\xfa\xa4\xcf\x85\xca\x93\x90\xf9\x31\x2e\x02\xf3\xc6\xc7\x67\xb4\x27\x00\x8f\xab\xec\x35\xb5\xe5\xca\xf9\x67\xe4\xbe\x2a\xaf\xbd\x0d\xf4\x6f\xb5\x0d\x33\xc6\x75\xb6\x5d\x7b\x86\x46\xea\x75\x81\x07\x74\xec\xd3\xaa\x9c\x41\x4e\xab\xdc\x1b\x29\x38\x73\x14\xba\x09\x80\x43\x89\xdc\x6b\x1b\x6e\x00\x8a\x90\xdf\x2b\x36\x47\xe9\x22\x21\x84\xd9\x3c\xa3\xcb\x63\x61\x24\xf3\x58\xc3\x1b\x46\x86\x4f\xb6\x24\xbd\x24\x0b\x60\x63\x62\xf8\x8c\x15\xda\x0a\xbf\x3e\x0d\x65\x39\xa8\x3c\x4e\xa3\x27\x24\x74\x34\xe1\x56\x78\xc1\x99\x4c\x6b\x7e\xd7\x4a\xd0\xe0\xb0\xec\x84\xcf\x6b\x89\xb6\xaa\x9e\x86\xc5\x00\x04\x6e\x71\x4d\x21\x3d\xad\xf5\xf5\xf2\x5c\x2b\x37\x54\x72\x9d\x36\xb8\x00\xb4\x09\x68\x6d\x29\xa4\xfd\xaf\xc2\x79\x97\xee\x11\x52\x59\x1e\x2a\x2c\x0b\x99\xb1\x0a\x3d\x56\x5d\xc2\xb5\xf2\x56\x4b\x62\x24\x53\x78\x80\x5c\x00\x5c\x2c\x90\x7b\x0a\xe9\x40\x4f\xf8\x12\xf3\x52\xe2\x21\x8a\x0b\x16\xfa\xe2\x9f\xd2\x18\xdc\x60\x42\xa1\xdd\x46\x90\xbc\x54\xac\xf1\x13\x05\xbb\x41\x0a\x47\x77\x93\x0f\x93\x69\xff\x7a\x76\xd6\x3f\xef\xfd\x76\x35\x9d\x8d\xfb\x17\x97\x93\xe9\xf8\xc3\xfd\x91\x65\x8a\x2f\xd1\x76\xf6\x0b\xa2\xab\xe3\xec\x38\xfb\xa1\xdb\x16\x38\x2a\xa5\x1c\x69\x29\xf8\x9a\xc2\xe5\x62\xa0\xfd\xc8\xa2\xc3\x6d\xc2\x83\xbd\x45\xc1\x54\xfe\x90\x6e\xf2\x92\xa1\x04\x9c\x67\xd6\x37\xce\x84\xc4\xed\xd1\x20\x75\xd0\xf3\x4e\xa4\xd6\xbf\xec\xb3\xd3\x6a\xcb\x11\xf7\xc0\x75\xa8\x3d\xd7\xd4\x1d\x43\x15\x11\x24\x32\x35\x22\x5f\x04\xfe\x11\xf3\x4b\xda\x52\xb0\xe5\x40\xb5\x7a\x2c\x6c\x34\x3c\x9b\x0d\x7a\xd7\xfd\xc9\xa8\x77\xda\x6f\x08\x5b\x31\x59\xe2\xb9\xd5\x05\x6d\xe5\x76\x21\x50\xe6\xf5\x7c\x7d\x44\x8f\xba\x37\x3d\x9e\x6d\xc7\x4c\xd2\xf4\xea\x00\x87\x22\xfd\x9a\x99\xb6\xb6\x47\x05\x53\xc7\x77\x77\x54\xb6\xd7\xda\xc3\xd0\x9c\x44\x7a\x35\x37\x9e\x1d\x9b\x61\x87\x28\xa5\x7d\xb3\xe7\x9b\xbb\x70\xa7\x55\x84\x23\x39\x2e\x58\x29\x3d\xa9\xae\x29\xa4\xde\x96\x98\x26\xcd\x3a\x84\xba\x4e\x03\xa0\xa1\x29\xfa\x5e\xaf\xbc\x6b\x9d\x23\x85\x3f\x98\xf0\xe7\xda\x9e\x0b\xeb\xfc\xa9\x56\xae\x2c\xd0\x26\x36\x3e\x4a\x36\x45\x7b\x86\x12\x3d\x56\x9e\xd7\x7b\x6c\x13\xb2\x64\xe7\x95\xf7\xec\x7a\xd8\x16\xe8\x13\x9b\x61\x03\x6c\xd4\x2a\x85\x3f\x49\x15\x90\xbb\x3a\x37\xd5\x04\x09\x15\x70\xcd\x4c\x4a\x3f\xd6\xd4\xbb\x6d\xe6\xaa\xfb\x94\xa6\x9b\xce\x1d\xf5\xa6\xbf\xce\xce\x87\xe3\xd9\x60\x38\x98\x5d\x5d\x4e\xa6\xfd\xb3\xd9\x60\x78\xd6\x9f\xa4\x6f\x1f\x30\xc1\x3a\x97\xd2\x8f\xe9\xd1\xdd\x06\x77\x35\x3c\xed\x5d\xcd\x26\xd3\xe1\xb8\x77\xd1\xaf\xa4\xdc\x1f\x55\x4f\x92\xf0\xdd\xd7\xff\x78\xbe\xaf\xd6\x97\x0f\x2f\x80\xda\xd8\xff\xff\xaf\x33\x17\xaa\xe3\x96\xd5\xe9\xcb\x52\x48\x84\x1b\xf4\xda\x78\x07\x69\x41\x1d\x35\x34\x05\x6d\x62\xfb\xe6\xfa\x61\x0e\x30\x87\xf0\x46\x1b\x0f\x42\xb5\x6a\xd1\x7c\xd7\x3a\xb2\xb9\xd3\xb2\xf4\x55\x1c\x7e\x79\x33\x1c\x4d\x7b\xe3\x8b\x16\xc3\xbb\x77\xad\xa3\x6b\xc3\x9d\xf8\x86\x97\xea\xfd\xda\xa3\x7b\x0d\xba\x68\xa3\x57\x5a\x86\xca\x79\x09\x89\x8e\xf1\xda\x3f\x15\xbb\xad\xb8\xcd\x85\x05\x52\xc0\xf1\xc9\xc9\x09\x10\x03\x6f\xee\x9a\x8e\xc4\xa0\xf2\x65\xa1\x73\x38\x39\xee\xee\xde\x76\xb2\xac\xda\xf3\xcc\xe6\xfa\x8b\xfa\x2f\xd4\xcf\x86\xda\x16\x40\xec\x62\x4f\x80\x97\x28\x0d\xda\x91\xce\xb3\x35\x2b\xe4\x36\x8a\x3b\x5d\x1c\x48\xb1\xd1\x47\x3a\xdf\xfb\xa2\x8a\xbd\x1d\xa5\x11\x53\x33\x35\x9f\x4d\x4f\xaf\xe0\x1d\x10\x1c\xb4\x76\x0b\x61\xad\xb6\x98\x13\x29\xe6\x96\xd9\x35\x99\x97\x6e\x3d\xd7\x5f\x69\x37\xfb\xf1\xa7\xac\x9b\xfc\x15\x00\x00\xff\xff\x95\xeb\x4b\xd2\x75\x0e\x00\x00")

func localStorageYamlBytes() ([]byte, error) {
	return bindataRead(
		_localStorageYaml,
		"local-storage.yaml",
	)
}

func localStorageYaml() (*asset, error) {
	bytes, err := localStorageYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local-storage.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAggregatedMetricsReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x31\x6b\xf4\x30\x0c\xc6\xf1\xdd\x9f\x42\x78\x7e\x93\x97\x6e\xc5\x6b\x87\xee\x1d\xba\x94\x1b\x94\xf8\x21\x27\xce\xb1\x83\x24\xe7\x68\x3f\x7d\xb9\x70\xdc\x58\x68\x27\x0d\x7f\x7e\x0f\xe8\x22\x35\x27\x7a\x29\xdd\x1c\xfa\xd6\x0a\x02\x6f\xf2\x0e\x35\x69\x35\x91\x4e\x3c\x8f\xdc\xfd\xdc\x54\xbe\xd8\xa5\xd5\xf1\xf2\x6c\xa3\xb4\xff\xfb\x53\x58\xe1\x9c\xd9\x39\x05\xa2\xca\x2b\x12\xd9\xa7\x39\xd6\xc4\xcb\xa2\x58\xd8\x91\x87\x15\xae\x32\xdb\xa0\xe0\x0c\x0d\x44\x85\x27\x14\xbb\x11\xfa\x61\xfd\xb1\x30\x78\x1b\x76\xc1\x35\x51\x74\xed\x88\xbf\x71\xc8\xe2\x7f\x71\x9c\x57\xa9\x0f\xa8\xbd\xc0\x52\x18\x88\x37\x79\xd5\xd6\x37\x4b\xf4\x11\xef\x7f\xdd\x7d\x3c\x05\x22\x85\xb5\xae\x33\x8e\xbe\xb5\x6c\xf1\x1f\xc5\xda\x32\xec\xc8\x3b\x74\x3a\xd2\x02\xbf\x95\x22\x76\xdc\x2b\xfb\x7c\x8e\xa7\xf0\x1d\x00\x00\xff\xff\xe5\x1d\x7a\x17\x89\x01\x00\x00")

func metricsServerAggregatedMetricsReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAggregatedMetricsReaderYaml,
		"metrics-server/aggregated-metrics-reader.yaml",
	)
}

func metricsServerAggregatedMetricsReaderYaml() (*asset, error) {
	bytes, err := metricsServerAggregatedMetricsReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/aggregated-metrics-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthDelegatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xbb\x4e\xc4\x40\x0c\x45\xfb\xf9\x8a\xf9\x81\x09\xda\x0e\xb9\x03\x0a\xfa\x45\xa2\x77\x26\x97\xc5\x24\x19\x47\xb6\x27\x12\x7c\x3d\x5a\x11\xd1\xf0\x68\xaf\x74\xcf\x39\xa5\x94\xc4\x9b\x3c\xc3\x5c\xb4\x51\xb6\x91\xeb\xc0\x3d\x5e\xd5\xe4\x83\x43\xb4\x0d\xf3\xad\x0f\xa2\x37\xfb\x69\x44\xf0\x29\xcd\xd2\x26\xca\x0f\x4b\xf7\x80\x9d\x75\xc1\xbd\xb4\x49\xda\x25\xad\x08\x9e\x38\x98\x52\xce\x8d\x57\x50\x5e\x11\x26\xd5\x8b\xc3\x76\x18\xf9\xbb\x07\x56\xba\xd2\xcb\x84\x05\x17\x0e\xb5\x64\xba\xe0\x8c\x97\xeb\x8b\x37\x79\x34\xed\xdb\x3f\x19\x29\xe7\x1f\x01\xdf\xbe\xdf\x05\xde\xc7\x37\xd4\x70\x4a\xe5\xf8\x3e\xc1\x76\xa9\xb8\xab\x55\x7b\x8b\x3f\x72\x8f\xd9\x37\xae\xa0\x3c\xf7\x11\xe5\x8b\x9f\x3e\x03\x00\x00\xff\xff\x69\xfc\x98\x93\x34\x01\x00\x00")

func metricsServerAuthDelegatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthDelegatorYaml,
		"metrics-server/auth-delegator.yaml",
	)
}

func metricsServerAuthDelegatorYaml() (*asset, error) {
	bytes, err := metricsServerAuthDelegatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-delegator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x3d\x4e\xc4\x40\x0c\x46\xfb\x39\xc5\x5c\xc0\x41\xdb\xa1\xe9\xa0\xa1\x5f\x24\x7a\x67\xf2\x01\x26\x1b\x4f\x64\x7b\x22\xe0\xf4\x28\x68\xf9\x69\x76\x7b\xfb\x7d\xef\x11\x51\xe2\x55\x9e\x60\x2e\x4d\x4b\xb6\x91\xeb\xc0\x3d\x5e\x9b\xc9\x27\x87\x34\x1d\xe6\x5b\x1f\xa4\xdd\x6c\x87\x11\xc1\x87\x34\x8b\x4e\x25\x1f\xdb\x09\xf7\xa2\x93\xe8\x4b\x5a\x10\x3c\x71\x70\x49\x39\x2b\x2f\x28\x79\x41\x98\x54\x27\x87\x6d\x30\xda\x79\x64\xe0\x09\x76\x3e\xf1\x95\x2b\x4a\x9e\xfb\x08\xf2\x0f\x0f\x2c\xc9\xda\x09\x47\x3c\xef\x10\x5e\xe5\xc1\x5a\x5f\xaf\xe8\xa4\x9c\xff\x44\x7e\x77\xf1\x1e\xd0\x3d\x84\x78\x95\x7f\xe3\xd0\x90\xfa\xfd\xfe\xa3\xe1\x7d\x7c\x43\x0d\x2f\x89\xce\xa0\x47\xd8\x26\x15\x77\xb5\xb6\xae\x71\x21\xe5\xb2\xfe\x57\x00\x00\x00\xff\xff\xc7\x9e\x8d\xd1\x49\x01\x00\x00")

func metricsServerAuthReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthReaderYaml,
		"metrics-server/auth-reader.yaml",
	)
}

func metricsServerAuthReaderYaml() (*asset, error) {
	bytes, err := metricsServerAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsApiserviceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\x4d\x6a\xc4\x30\x0c\x85\xf7\x3e\x85\x2e\xe0\x74\xb2\x2b\xde\x75\x59\x68\x61\x20\x65\xf6\x1a\x8f\x3a\x08\xe3\x1f\x24\x39\x90\xdb\x97\x30\x71\xba\x13\x7a\xef\xfb\x90\xbc\xf7\x0e\x1b\xdf\x48\x94\x6b\x09\x80\x8d\x85\x9e\xac\x26\x68\x5c\xcb\x94\xde\x75\xe2\xfa\xb6\xce\x77\x32\x9c\x5d\xe2\xf2\x08\xf0\x71\xfd\x5c\x48\x56\x8e\xe4\x32\x19\x3e\xd0\x30\x38\x80\x82\x99\x02\x1c\xd5\x29\x93\x09\x47\x3d\x0c\x4e\x1b\xc5\xbd\xa4\x2f\x70\x1f\x07\x71\x34\xfd\x1e\x91\x9c\x81\x36\x8c\x14\x20\xf5\x3b\x79\xdd\xd4\x28\x3b\x80\xa7\xd4\xde\x4e\x64\xc8\x01\xd6\xf1\xc0\xb8\x14\x80\x8b\x52\xec\x42\x4b\xe2\xf6\xf3\xb5\xdc\x48\xf8\x77\x0b\x60\xd2\x69\x88\xae\xc2\x55\xd8\xb6\x6f\x2e\x9c\x7b\x0e\x30\x5f\x2e\xff\xb2\x91\xbe\xd6\x7f\x01\x00\x00\xff\xff\x25\x03\x92\xf5\x2a\x01\x00\x00")

func metricsServerMetricsApiserviceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsApiserviceYaml,
		"metrics-server/metrics-apiservice.yaml",
	)
}

func metricsServerMetricsApiserviceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xdf\x6f\x1b\x37\x0c\x7e\xf7\x5f\x41\x78\xe8\xdb\x2e\x76\xba\x76\x1b\x04\xe4\x21\x48\xdc\x36\x40\x93\x19\xb1\x3b\xa0\x4f\x85\xa2\xa3\x63\x21\x92\xa8\x91\x3c\x37\xb7\xa2\xff\xfb\x20\x5f\x7a\xb9\x4b\x93\xa2\xc3\xd6\x7b\x3a\xf0\xe3\x8f\x8f\x1f\x25\xaa\xaa\xaa\x89\xcd\xfe\x4f\x64\xf1\x94\x0c\xec\x0e\x27\x37\x3e\xd5\x06\x56\xc8\x3b\xef\xf0\xd8\x39\x6a\x92\x4e\x22\xaa\xad\xad\x5a\x33\x01\x48\x36\xa2\x81\x88\xca\xde\x49\x25\xc8\x3b\xe4\x3b\xb3\x64\xeb\xd0\xc0\x4d\x73\x85\x95\xb4\xa2\x18\x27\x0f\x2b\xd8\x9c\x65\xd6\x97\x39\xc5\x1c\xa8\x8d\xf8\x9f\x4a\x00\x04\x7b\x85\x41\x4a\x24\xc0\xcd\xef\x52\xd9\x9c\xbf\x0a\x97\x8c\xae\x78\x08\x06\x74\x4a\xdc\x79\x47\xab\x6e\xfb\x76\x10\xfe\x74\x02\x00\xc5\x98\x83\x55\xbc\x0b\x1d\x10\x2e\xdf\x13\xa4\xcb\x17\x46\x05\xbe\x55\x02\xe0\x0b\xcf\xf2\x65\xf6\xc4\x5e\xdb\x93\x60\x45\x2e\xf6\xf9\xa7\x5d\xd3\x55\xa2\x1a\x2b\xc7\x5e\xbd\xb3\x61\x7a\xe7\x2f\xa3\xa9\x5d\x3c\x4d\x48\x29\x20\x5b\xf5\x94\x06\xac\x2a\xb8\xc1\xd6\xc0\xf4\xe4\x2e\xeb\x71\x5d\x53\x92\x3f\x52\x68\xa7\xbd\x0f\x00\xe5\x12\x49\x6c\x60\xba\xb8\xf5\xa2\x32\xfd\x2a\xc1\x9e\x1b\x53\xc0\x83\x32\x26\x4e\xa8\x28\x07\x9e\x66\x8e\x92\x32\x85\x2a\x07\x9b\xf0\x3b\x73\x02\xe0\x66\x83\x4e\x0d\x4c\x2f\x68\xe5\xb6\x58\x37\x01\xbf\xbf\x64\xb4\xa2\xc8\xff\x47\xad\x1d\x85\x26\x62\x2f\xd7\x4f\x10\x8b\xc6\xe0\x13\x68\xcc\x20\x04\x1f\x11\x9c\x4d\x20\x76\x83\xa1\x85\x46\x10\x36\x4c\xb1\x12\xc7\xe5\x8c\x81\x8f\xf6\x1a\x05\x6c\xaa\x67\xc4\xc0\x68\xeb\x8a\x52\x68\xa1\x88\x62\x7d\x42\x96\xc9\x97\x96\xba\x93\xa4\x31\x57\xb5\xe7\x9e\x1d\xc6\xac\xed\xa9\x67\x03\x9f\x3e\xdf\x19\xef\x63\xcd\x83\xe0\x47\xa7\x0e\x1d\x09\x03\xcf\x3e\xad\xde\xaf\xd6\x8b\xf3\x0f\xa7\x8b\x57\xc7\xef\xde\xae\x3f\x5c\x2e\x5e\x9f\xad\xd6\x97\xef\x3f\x3f\x63\x9b\xdc\x16\x79\x16\x3d\x33\x31\xd6\xd5\x38\x93\xd9\xcd\x0f\x5e\x1e\x3c\xef\x13\x5a\xbe\x1e\x9d\xa0\xaa\x72\xc8\x5a\x78\x1f\xcd\x34\xe6\x11\x22\xe8\x1a\xc6\x2a\x13\xeb\xd1\x8b\x17\x2f\x7e\x19\x81\x65\x6c\x01\xb5\xca\x8c\x1b\xe4\x52\xd8\xd6\x35\xa3\x48\xa5\x6d\x46\x39\x3a\x4b\x8a\x9c\x6c\x38\x5b\xfe\xbc\xb8\xed\x7f\xdf\x90\x68\x69\xf8\xd1\x54\x8d\x60\x77\x4d\x44\xad\x36\xb2\x2f\x3c\x72\xec\x5a\xab\x18\x85\x42\x53\x2e\xc3\xd1\xe1\x4b\xe9\x3d\x8a\xb9\x61\x87\x83\xfe\x8a\xf1\xaf\x06\x45\x47\x36\x00\x97\x1b\x03\x87\xf3\x79\x1c\x59\x23\x46\xe2\xd6\xc0\x6f\xf3\x73\xdf\x03\x85\xc4\x48\xb1\x6e\x5e\x5b\xd5\x2c\x83\xe8\x7e\xb2\x4b\x62\x35\x30\x92\xab\xec\x05\x52\x72\x14\x0c\xac\x4f\x96\x03\xc2\xb6\xf6\x09\x45\x96\x4c\x57\x38\x64\x58\xb2\xbf\x46\x1d\x93\xce\x56\xb7\x06\x66\x25\xaa\xfd\x7b\x8c\xec\x6b\x3e\xa4\x04\x20\x6e\x8b\x85\xec\x9b\xf5\x7a\xb9\x1a\x20\x3e\x79\xf5\x36\x9c\x62\xb0\xed\x0a\x1d\xa5\x5a\x0c\xcc\x87\x7c\x91\x3d\xd5\x3d\xf4\x7c\x00\xa9\x8f\x48\x8d\xf6\xd8\xe1\x00\x93\xc6\x39\x14\x59\x6f\x19\x65\x4b\xa1\x1e\xa3\x1b\xeb\x43\xc3\x38\x40\xef\x25\x0a\x7e\x87\xff\x5a\x89\x12\xf4\x03\x84\xf8\xf5\x1b\x4a\x1c\xce\x7f\xb8\x14\xfb\x5b\x57\xde\x10\x4a\x8a\xb7\x3a\x3e\xcc\xb6\x2e\xeb\xfd\x92\x48\x5f\xf9\x80\xdd\xd3\x62\x40\xb9\xc1\xa1\x5b\x93\x8e\xe5\x82\x52\x71\x7b\x1c\x7c\x27\xc8\xfb\x0b\x30\x6c\xc7\x86\x40\x1f\x97\xec\x77\x3e\xe0\x35\x2e\xc4\xd9\xb0\x7f\x71\x0c\x6c\x6c\x90\xfb\x1c\xdd\x62\x3d\x2f\xdb\xf4\x91\x8b\xf1\x70\x0b\x42\xb7\x77\x97\xdd\xc8\xca\x8a\xf9\x27\x00\x00\xff\xff\xa9\x1e\x89\xbe\xc4\x08\x00\x00")

func metricsServerMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerDeploymentYaml,
		"metrics-server/metrics-server-deployment.yaml",
	)
}

func metricsServerMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3f\x4b\x04\x31\x10\xc5\xfb\x7c\x8a\x61\xfb\x28\xe2\x15\x92\xd6\x5a\x38\x50\xec\xe7\x72\x0f\x0d\x97\x4d\xc2\xcc\xec\x82\xdf\x5e\x76\xf6\x9a\x83\xed\x92\x37\xef\xcf\x2f\xc6\x18\x78\x94\x6f\x88\x96\xde\x12\xad\x2f\xe1\x56\xda\x35\xd1\x27\x64\x2d\x19\x61\x86\xf1\x95\x8d\x53\x20\x6a\x3c\x23\xd1\x0c\x93\x92\x35\x2a\x64\x85\xdc\x65\x1d\x9c\x91\xe8\xb6\x5c\x10\xf5\x4f\x0d\x73\x20\xaa\x7c\x41\xd5\x2d\x49\x7e\x91\x06\x83\x3e\x95\xfe\xbc\x37\x4d\x1f\x0f\x55\xd3\x81\x31\xd7\x45\x0d\xe2\x8e\xb2\x2d\x4c\x26\x0b\xa6\xa0\x03\x79\x2b\x56\x54\x64\xeb\x72\x1f\x79\xd3\xc8\x63\x1c\x30\x8e\x2e\xe6\x24\xd1\x9f\x89\x4e\xa7\x57\x8f\xec\x24\xbf\x66\x43\xfd\x3f\xa4\x5b\xcf\xbd\x26\xfa\x7a\x3f\xbb\x62\x2c\x3f\xb0\xb3\xa7\x76\xdf\x7f\x00\x00\x00\xff\xff\x7e\x3b\x1f\x83\x35\x01\x00\x00")

func metricsServerMetricsServerServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerServiceYaml,
		"metrics-server/metrics-server-service.yaml",
	)
}

func metricsServerMetricsServerServiceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerResourceReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x3f\x4f\xc4\x30\x0c\xc5\xf7\x7c\x0a\xeb\xf6\xf4\xc4\x86\xb2\x01\x03\xfb\x21\xb1\xbb\xa9\xb9\x33\x6d\xe3\xca\x76\x8a\xe0\xd3\xa3\x6b\xcb\x1f\x81\x74\x42\x62\xca\x4b\xe2\x9f\x9f\xde\x8b\x31\x06\x9c\xf8\x91\xd4\x58\x4a\x02\x6d\x31\x37\x58\xfd\x24\xca\x6f\xe8\x2c\xa5\xe9\xaf\xad\x61\xd9\xcf\x57\xa1\xe7\xd2\x25\xb8\x1b\xaa\x39\xe9\x41\x06\x0a\x23\x39\x76\xe8\x98\x02\x40\xc1\x91\x12\xd8\xab\x39\x8d\x69\x24\x57\xce\x16\x8d\x74\x26\x0d\x5a\x07\xb2\x14\x22\xe0\xc4\xf7\x2a\x75\xb2\x33\x11\x61\xb7\x0b\x00\x4a\x26\x55\x33\x6d\x6f\x93\x74\xb6\x88\x22\x1d\x7d\x53\x7b\x73\xf4\xed\x8e\x23\xd9\x84\x79\xf9\x9e\x49\xdb\x0d\x3d\x92\x2f\xe7\xc0\xb6\x8a\x17\xf4\x7c\x0a\xff\x0b\x79\xcb\xa5\xe3\x72\xfc\x7b\x56\x19\xe8\x40\x4f\xe7\xb1\x8f\xb4\x17\x2c\x03\xc0\xef\x5a\x2f\x1b\x58\x6d\x9f\x29\xfb\xd2\xe7\xca\x3e\x90\xce\x9c\xe9\x26\x67\xa9\xc5\x3f\xf1\x1f\x1c\x7c\xf5\x96\xa0\xaf\x2d\xc5\x75\x7f\x78\x0f\x00\x00\xff\xff\x39\x82\xcc\x46\x05\x02\x00\x00")

func metricsServerResourceReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerResourceReaderYaml,
		"metrics-server/resource-reader.yaml",
	)
}

func metricsServerResourceReaderYaml() (*asset, error) {
	bytes, err := metricsServerResourceReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/resource-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rolebindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x21\xbb\x72\x38\xdc\x72\xf0\xd8\x0e\xdd\x03\xb4\x3b\x6d\xb3\x09\x6b\x59\x14\x48\x2a\x41\xfb\xeb\x0b\xa7\x6e\x82\xa4\x76\x90\xb4\xdd\x24\x41\x7c\x1f\x1f\xf9\x20\xd3\x13\x8a\x12\xa7\xca\x4b\x0d\xcd\x12\x8a\x6d\x58\xe8\x0d\x8c\x38\x2d\xbb\xff\xba\x24\xfe\xb3\xfd\xeb\x3a\x4a\x6d\xe5\xef\x63\x51\x43\x59\x71\xc4\x3b\x4a\x2d\xa5\xb5\xeb\xd1\xa0\x05\x83\xca\x79\x9f\xa0\xc7\xca\x77\xa5\xc6\x00\x99\x14\x65\x8b\x12\x86\x6b\x44\x0b\xd0\xf6\x94\x9c\x70\xc4\x15\x3e\x0f\xbf\x21\xd3\x83\x70\xc9\x17\xc8\xce\xfb\x2f\xe0\x03\x47\x5f\xd5\xb0\xaf\x0e\xfa\x99\x46\x86\x96\xfa\x05\x1b\xd3\xca\x85\x9b\x20\x8f\x8a\x32\xe3\xc2\xb9\x10\x82\xfb\xfe\xb4\x26\xc6\xf4\xd9\xfe\x3f\x0d\x0d\x27\x13\x8e\x11\xc5\x49\x89\x78\xd2\xb8\x0e\x15\xc1\x2f\x16\xce\x7b\x41\xe5\x22\x0d\x8e\x6f\x89\x5b\x54\xe7\xfd\x16\xa5\x1e\x9f\xd6\x68\x57\xd6\x42\x8f\x9a\xa1\x39\x17\x88\xa4\xb6\x3f\xec\xc0\x9a\xcd\x84\x56\x42\xdb\xb1\x74\x94\xd6\xa3\xdf\x29\xf1\x8f\x3f\x99\x23\x35\x74\x33\x61\x42\x10\x53\x9b\x99\x92\xe9\xfe\x96\xb9\x9d\xd3\x1c\xfc\x1f\xb5\x7f\xb8\xb4\xf9\x88\xcf\xec\xee\xf7\xb3\x7d\x0a\x38\x06\x7b\xf0\x78\x1d\xe3\x2c\xdc\x97\x01\xef\x01\x00\x00\xff\xff\x46\xd3\x6d\x9d\x0f\x04\x00\x00")

func rolebindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_rolebindingsYaml,
		"rolebindings.yaml",
	)
}

func rolebindingsYaml() (*asset, error) {
	bytes, err := rolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x41\x6f\xd3\x40\x10\x85\xef\xfe\x15\x23\x4b\x39\xae\x9d\x54\x08\x55\xbe\x85\xd4\x40\x05\x14\x94\xa4\xa0\x9e\xa2\xf5\x7a\x12\xaf\xb2\xde\xb5\x66\xc6\x81\x50\xfa\xdf\xd1\xc6\x4d\x9b\x4a\x95\x8a\x10\x1c\x3d\x9e\xf9\xe6\xcd\x7b\xab\x3b\xfb\x15\x89\x6d\xf0\x05\x34\xe8\xda\xcc\x68\x11\x87\x99\x0d\xf9\x6e\x92\x6c\xad\xaf\x0b\x78\x8f\xae\x9d\x35\x9a\x24\x69\x51\x74\xad\x45\x17\x09\x80\xd7\x2d\x16\x20\xa4\x71\x6d\xb7\xca\x50\x7d\x5f\xe3\x4e\x1b\x2c\x60\xdb\x57\xa8\x78\xcf\x82\x6d\xc2\x1d\x9a\x38\x62\x22\xa4\x80\x46\xa4\xe3\x22\xcf\x47\xb7\x1f\xae\xdf\x94\xf3\xab\x72\x59\x2e\x56\xd3\x2f\x97\x77\xa3\x9c\x45\x8b\x35\xf9\xa1\x91\xf3\x13\xb8\x9a\x8c\xb3\xc9\xab\x6c\x32\x1e\x67\xb2\xf9\x99\x28\xa5\x92\x7f\x24\xfd\xff\xc9\x7e\x2a\x19\x80\x51\x22\x0e\x60\xe3\x42\xa5\x5d\x36\xac\xb9\xc0\xb5\xee\x9d\xcc\x71\x63\x59\x68\x5f\x40\x3a\xba\x5d\xdc\x2c\x96\xe5\xa7\xd5\x45\xf9\x76\x7a\xfd\x71\xb9\x9a\x97\xef\x2e\x17\xcb\xf9\xcd\x6a\x3e\xfd\x76\x37\x4a\x13\x80\x9d\x76\x3d\xf2\x2c\x78\x41\x2f\x05\xfc\x52\x07\x2e\x55\xda\x0c\x1b\x00\xd0\xeb\xca\x61\x1d\xcf\xec\xf1\x50\xeb\x02\x09\x1f\x7f\x7f\xc7\x8a\xd1\xf4\x84\xc7\x02\x80\x38\x7e\xfc\x78\x1e\x50\x4f\xbd\x0f\xf1\xd6\xe0\x1f\x7a\x3b\x0a\x2d\x4a\x83\x3d\x47\xe7\xe3\x92\x02\xd2\xf3\xf1\xf9\x59\xfa\x6c\x03\x1b\xd2\x1d\x16\x90\x46\xec\xd0\xd2\x51\xd8\xd9\x1a\xe9\x01\x19\x43\x20\x8f\x82\x7c\xe9\x37\x84\x7c\xa2\xab\xeb\x2b\x67\xb9\xc1\x7a\x81\xb4\xb3\x06\x5f\x50\x4c\x36\x90\x95\xfd\xcc\x69\xe6\xab\x43\xea\xe9\xe0\xba\x32\xae\x67\x41\x52\x86\xac\x58\xa3\xdd\x20\xc5\xb6\x7a\xf3\xc0\x1c\x9e\x49\x4a\xda\x9b\x06\x29\x6f\x2d\x51\x20\xac\x95\xb3\x15\x69\xda\xab\xfb\x9c\x8f\x77\x8a\xde\x14\x90\x9e\x65\xaf\xb3\xc9\x50\x92\xe0\x90\x4e\xcd\x52\xb0\xc5\x18\xf0\xec\x7e\xe7\xb4\xae\x83\xe7\xcf\xde\xed\x8f\x8c\xd0\xc5\x89\x40\x05\xa4\xe5\x0f\xcb\xc2\xe9\x93\x41\x1f\x6a\x54\x14\x1c\x66\x8f\x16\x45\x53\x4d\xf0\x42\xc1\xa9\xce\x69\x8f\x2f\xb0\x00\x70\xbd\x46\x13\x53\xba\x0a\x0b\xd3\x60\xdd\x3b\xfc\xb3\x35\xad\x8e\x96\xfd\x1d\xff\x77\x00\x00\x00\xff\xff\x68\x8d\x6e\xbc\x69\x04\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"ccm.yaml":           ccmYaml,
	"coredns.yaml":       corednsYaml,
	"local-storage.yaml": localStorageYaml,
	"metrics-server/aggregated-metrics-reader.yaml": metricsServerAggregatedMetricsReaderYaml,
	"metrics-server/auth-delegator.yaml":            metricsServerAuthDelegatorYaml,
	"metrics-server/auth-reader.yaml":               metricsServerAuthReaderYaml,
	"metrics-server/metrics-apiservice.yaml":        metricsServerMetricsApiserviceYaml,
	"metrics-server/metrics-server-deployment.yaml": metricsServerMetricsServerDeploymentYaml,
	"metrics-server/metrics-server-service.yaml":    metricsServerMetricsServerServiceYaml,
	"metrics-server/resource-reader.yaml":           metricsServerResourceReaderYaml,
	"rolebindings.yaml":                             rolebindingsYaml,
	"traefik.yaml":                                  traefikYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"ccm.yaml":           &bintree{ccmYaml, map[string]*bintree{}},
	"coredns.yaml":       &bintree{corednsYaml, map[string]*bintree{}},
	"local-storage.yaml": &bintree{localStorageYaml, map[string]*bintree{}},
	"metrics-server": &bintree{nil, map[string]*bintree{
		"aggregated-metrics-reader.yaml": &bintree{metricsServerAggregatedMetricsReaderYaml, map[string]*bintree{}},
		"auth-delegator.yaml":            &bintree{metricsServerAuthDelegatorYaml, map[string]*bintree{}},
		"auth-reader.yaml":               &bintree{metricsServerAuthReaderYaml, map[string]*bintree{}},
		"metrics-apiservice.yaml":        &bintree{metricsServerMetricsApiserviceYaml, map[string]*bintree{}},
		"metrics-server-deployment.yaml": &bintree{metricsServerMetricsServerDeploymentYaml, map[string]*bintree{}},
		"metrics-server-service.yaml":    &bintree{metricsServerMetricsServerServiceYaml, map[string]*bintree{}},
		"resource-reader.yaml":           &bintree{metricsServerResourceReaderYaml, map[string]*bintree{}},
	}},
	"rolebindings.yaml": &bintree{rolebindingsYaml, map[string]*bintree{}},
	"traefik.yaml":      &bintree{traefikYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
