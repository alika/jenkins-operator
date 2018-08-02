// Code generated by go-bindata. DO NOT EDIT.
// sources:
// pkg/bindata/.DS_Store
// pkg/bindata/environment/jenkins-jvm-environment
// pkg/bindata/environment/required-plugins
// pkg/bindata/init-groovy/0-jenkins-config.groovy
// pkg/bindata/job-scripts/install-dsl-job.sh
// pkg/bindata/job-scripts/install-xml-job.sh
// pkg/bindata/jobdsl/seed-job-dsl
// pkg/bindata/plugin-scripts/install-plugin.sh
package bindata

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

var _Ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6c\x14\xd5\x1b\xff\xbe\xa5\xd0\x73\x52\x68\x4f\xb9\xb5\xdd\xe5\xdf\xff\x4a\xa1\xb4\x61\x4b\xb6\x17\xda\x72\xb5\x37\x2e\xb5\x18\x08\x6d\x69\x69\xc1\x32\xbb\x67\x68\x47\x67\x67\xd6\x9d\xd9\x6d\x69\xa9\xd6\x28\x3e\x29\x20\x6f\x80\x26\xc4\xf8\xa2\xf0\x64\x82\x26\xc6\x07\x7d\x51\x04\x13\x43\x52\x05\x79\x20\x46\x5f\x78\xf4\x12\x8d\x3c\x69\x76\xe6\x94\xee\x2e\xdd\xb6\xc4\x10\x1a\x72\x7e\xc9\xe6\x37\x73\xbe\xef\x7c\x67\xe6\x9c\xdf\x99\x33\xdf\xd9\x01\x00\x6c\x89\xf3\x6a\x80\x0a\x00\x20\xe0\x32\x5b\x0a\x00\x6f\x02\xf8\xf3\x20\x0d\x44\xfc\x1e\xc2\xa2\x64\x1c\x00\x28\x75\x8f\x19\x40\x08\xaa\xc0\x82\x30\xc4\x40\x9b\xa9\x42\x06\x3c\xa2\x7e\x0e\x00\xe4\xc1\x8b\x60\xa6\xd5\x8f\x82\x0d\x56\xc4\x6c\x6b\xe3\x71\x3b\x0c\x70\xeb\x54\xdd\xa5\x4a\x0e\x2a\x18\x90\x00\x0d\x62\xf3\x88\x2f\x21\x21\x21\x21\x21\x21\xf1\x30\xd0\x25\x92\xf7\xa4\x2f\x44\x42\x42\x62\xc1\x21\xf9\x7c\xf0\x0b\x6e\x12\x3c\xe1\x32\x0a\xbb\x47\x70\x4e\x4a\x1d\x26\xd8\x2f\xb8\x49\xf0\x84\xcb\x28\xfc\x3c\x82\x73\x04\x13\xc1\x4c\xb0\x5f\x70\x93\xe0\x09\x97\xc5\x43\x0b\x3d\x82\x45\xcb\x28\x32\x14\x64\x82\xfd\x82\x9b\x1e\x4f\xdf\x48\x48\x3c\x0d\x28\x98\x2d\xf7\xe6\xe9\xb9\x77\x56\xcf\xe8\x50\x75\x67\xd8\x8c\x44\xdd\x90\xce\xcc\x5b\x22\x3c\x39\x58\xa0\xeb\x83\xa9\x76\xe4\x99\xf6\xb4\x2c\x7f\xc3\xf7\xf7\x1f\xb6\xf3\x59\xed\xe9\xed\x3b\x4f\x80\x7c\x88\x82\x0e\x71\x18\x04\x0d\x8c\xcc\xeb\x0d\x0d\x5b\xd1\x90\x6e\x86\x00\xe0\x9b\x50\x54\xd7\x2c\x3b\x18\xbc\x85\x9e\x45\x39\x8b\x97\xe4\x12\x42\xf2\xc8\x32\x72\xb4\x73\xc8\x1c\xee\xb4\x15\x3b\x6e\xb5\x28\xb1\xfe\xe4\xd9\x01\xc5\x1e\x0a\x89\xe3\x2e\xd3\xd4\x1f\x1c\x2b\xa1\x43\x9a\x3a\x3c\xc0\x56\xb6\x9a\x86\xad\x68\x86\x1a\x73\x2a\x6b\x5c\x0d\x29\xb1\x23\x3d\x9a\xc1\xcd\xe1\x16\x33\x6e\x70\xab\x3f\xc5\x40\x29\x25\x74\x80\x15\x8f\x8d\xd5\xd6\x34\x04\xfc\x35\xf5\xb5\xe3\x01\xff\x58\x43\x43\x30\xe0\xaf\xab\xad\x1f\x1f\xa7\xa4\x68\x7d\xf5\x8e\xf6\x63\x91\x13\xa3\x63\x27\xc7\x5f\x39\x2f\x3a\x0f\x53\x06\x2e\x15\x17\xe6\xba\x69\x6e\x85\xf5\x90\x69\xea\x38\x97\x63\xfa\x68\x2d\x9e\x33\xb0\x6e\x25\x5a\xdd\xde\xf4\x9c\x9b\xea\xcd\xdb\x53\xbd\x49\x29\x5d\xba\xb7\xfd\xb9\x8e\x01\xc6\xe2\x96\x7a\x50\xd5\x15\x5b\x4b\xa8\x6d\x8a\xad\x5a\x03\xac\xc0\x1a\x32\x87\xdb\xc3\xa6\x71\x20\xa6\x26\x9c\x2e\x2c\x0c\x2b\x7a\x38\xae\x2b\xb6\xda\xac\xeb\x9d\xda\xa8\x6a\xf5\x84\x4d\x3d\x1e\x31\xac\x5e\x5b\x1d\xb1\x93\x25\x7d\x96\x19\xb3\x5b\x9d\xc2\x5e\x2d\x6c\x1a\xc9\xb2\x01\xb6\x3c\x59\x7f\x7f\xd4\xd6\x4c\xc3\x3a\xa4\xc6\x2c\xcd\x34\x28\xa5\x97\x97\xad\x2e\x2d\xdb\x10\xa8\xae\xdf\xf6\x6c\xdb\x64\x7e\x01\x2b\xa4\x2b\xe8\xaa\x9e\x84\x66\x69\x21\x5d\xed\x1e\xd6\xb8\x3d\x74\x58\xb1\xc2\xaa\xc1\x35\x63\xb0\x4f\xe3\xaa\x61\x6b\xc7\x35\x35\x46\x0b\x31\x40\xbb\x0c\x25\xa2\x4e\x16\x15\x97\x14\x7a\x7d\xde\xff\x65\xab\x45\x58\x19\xe9\x8d\x87\xb4\x97\xe3\x9a\x7d\xc2\x6d\xc3\xef\x5d\x4b\xd9\x55\x72\x84\x2b\xb6\xfa\xbc\xc9\x93\x11\xb9\x63\xf1\xfa\xbd\xe5\x84\xf4\x27\xcb\x5b\x63\xaa\x62\xab\x7c\xb2\xb0\x80\xe5\x57\x54\x7a\x69\x97\xa5\x8d\xaa\x4c\x21\xd4\x29\xa9\xda\x44\x69\xd7\x4b\x9a\xc1\x99\x45\xdd\x92\x9a\x5a\xea\xed\xd6\x95\x90\xaa\x33\x4e\x89\x53\xd4\xd0\x48\xbd\x3d\x09\xf7\x5e\x59\x87\x28\xdc\xbe\x82\x7a\x7b\xc3\x66\x24\xa2\x1a\xb6\x45\x89\xdb\x6c\xb3\xb7\x95\xb0\x6b\xe4\x85\x64\xc3\xfb\x14\xcb\xde\x1f\x55\x8d\xe9\x4b\xda\x43\xc8\xe1\xa4\xa5\x99\x73\x95\x97\x35\x55\xb8\xe3\xee\xdc\x7d\x59\x53\x50\xcc\xa9\xe4\xfa\x5c\x02\x01\xd8\x09\x07\xa1\x0f\xc2\x60\xc0\x30\xbc\x05\x6f\xc3\x69\x38\x03\x17\xe1\x03\xb8\x02\x1f\xc3\xe7\xf0\x35\x5c\x83\x1b\xf0\x2d\xdc\x84\x3b\xf0\x33\xdc\x83\xdf\xe1\x0f\xf8\x13\xfe\x82\x7f\x90\x20\xc5\x3c\x5c\x8a\x25\xb8\x16\xcb\x70\x1d\x06\x71\x0b\xee\xc4\x26\x6c\xc6\x16\xec\xc0\x03\x78\x10\x3b\xb1\x0b\x8f\x62\x18\x55\x3c\x8e\x83\x18\xc5\x11\x1c\xc5\x31\x3c\x89\xa7\xf0\x34\x9e\xc1\xb3\x78\x1e\x2f\xe0\xbb\xf8\x1e\x5e\xc1\xab\xf8\x09\x7e\x8a\x5f\xe2\x75\xfc\x0e\x6f\xbb\xd7\xe7\x99\x9a\x16\xfb\xd2\x67\x05\xde\x99\x87\x7a\xc5\xb3\xc0\x73\x78\x46\xf5\xee\xda\xbd\xa7\xe5\xc9\xa8\xf7\xc7\x65\xc9\x11\x5a\xbe\x62\xe5\xaa\xd5\xff\x5f\x57\xb9\xa9\xb6\x71\x47\xf3\x83\xa1\xcd\x18\xcb\x54\x4d\x39\x52\x72\xb5\xe2\x68\x68\x4a\x23\xce\x88\xa6\x8b\xb2\xa8\xb8\xc4\xeb\x5b\x43\x4b\xb3\x89\xbb\x5b\x33\xb8\x3a\x42\x92\xb3\x81\xe5\x4e\x7a\x8b\x4b\x8a\xfc\xcf\xf8\x7c\x2c\x29\x26\xe2\x9c\xae\x2f\xf7\xf9\x98\x87\x5d\x25\xc4\x89\x45\x37\xfa\xaa\x28\x53\x08\x5b\xe4\x86\x0e\xd2\x1a\xc2\x38\x65\x8b\x5d\xeb\x66\xda\x40\x99\x45\x59\x8e\x6b\xdd\x4a\xb7\x13\xd6\x41\xd9\x12\x27\xd4\xce\x35\x94\x32\xa0\xd4\x39\x69\x29\xf7\x51\x86\x84\x66\x91\xe3\x2c\x62\x7c\x1f\x2e\xc3\x67\xf0\x15\xdc\x80\x9b\xf0\x03\xdc\x85\x5f\xe0\x57\xb8\x8f\x88\xb9\x58\x88\x45\x58\x8c\x3e\x5c\x83\xa5\xb8\x01\x2b\x71\x23\x06\xb0\x0a\xeb\xb1\x11\xb7\xe2\x36\xdc\x8e\xbb\x70\x37\xee\xc5\x76\xec\xc0\x2e\xec\xc6\x1e\xec\xc5\x3e\x0c\x23\x77\xe4\xa8\x61\x0c\x2d\x8c\x63\x02\x47\xf0\x35\x7c\x1d\xdf\x70\x64\x79\x16\xdf\xc1\x73\x78\x11\x2f\x65\x88\x70\x6f\x86\x08\x3f\x9c\x4b\x84\x29\x0b\x60\xe4\xd5\xf2\xbb\x39\x73\xfb\xf3\x47\xf2\x7f\xd4\x05\x32\xd1\x19\x33\x74\xd3\x18\x9c\xde\x24\x90\x90\x90\x90\x90\x90\x90\x78\x7a\x90\xef\x64\xf5\xd3\xff\xaa\x9b\x60\x40\xc4\x39\xb7\x33\xb2\xf6\xfe\xec\x9e\x29\xaf\x2f\x63\xf7\x6e\x7d\xf4\xdb\x6c\x9e\x7c\x9e\x9e\x33\xec\x29\xe4\x39\xaf\x2a\x1a\xd8\x50\x05\x83\x8e\xb7\x09\x09\x38\xb1\xc0\xb2\xf9\xc6\x60\x5d\xc0\xbf\xb9\x6e\xcb\x7f\xce\xe6\xb3\xdc\x6e\x6a\x1e\x9f\xc5\x25\x7d\xe4\x72\xcf\xce\xe2\x29\x73\x77\x99\xbb\x2f\xe4\xdc\x3d\xbb\x6e\x65\xd6\x2e\xb3\xf6\xc7\x9e\xb5\x67\x91\x5f\xea\x82\xf7\xf7\xf5\x89\x8a\xd9\x3c\xf9\x3c\x3d\x67\xc8\xd1\xb3\x78\x66\x64\xe7\x59\x36\xe5\x33\xb6\x71\x7f\xca\x88\x9d\x2c\x73\xa9\x22\x0f\x00\x76\x65\xff\xfe\x4f\x42\x42\xe2\x29\x06\xe6\xb4\x75\xb6\xb5\xcc\xb2\xd7\xe7\x11\x6f\xe0\xc7\x84\xcf\x17\x53\x15\xb3\x7c\x08\x80\xe2\x83\x61\x04\x08\xc2\xb4\xef\xc2\xf8\x10\x40\xae\xff\x72\xfd\x97\xeb\x3f\xfc\x1b\x00\x00\xff\xff\x7b\xc3\x97\x69\x04\x30\x00\x00")

func Ds_storeBytes() ([]byte, error) {
	return bindataRead(
		_Ds_store,
		".DS_Store",
	)
}

func Ds_store() (*asset, error) {
	bytes, err := Ds_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".DS_Store", size: 12292, mode: os.FileMode(420), modTime: time.Unix(1532986524, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _environmentJenkinsJvmEnvironment = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\x72\x0c\x73\x8c\xf7\x0f\x08\x09\xb6\xd2\x8d\xc8\x2d\x36\x34\x30\x32\xc9\x55\xd0\x8d\xc8\xad\x80\xb2\x5c\xb2\x52\xf3\xb2\x33\xf3\x8a\xf5\x32\xf3\x8a\x4b\x12\x73\x72\xf4\x8a\x4a\xf3\x82\x53\x4b\x4a\x0b\xc2\x33\xab\x12\x8b\x52\x6c\xd3\x12\x73\x8a\x53\x91\x94\x39\xfb\x78\xea\xa5\x64\x16\x27\x26\xe5\xa4\xa6\xd8\x96\x14\x95\xa6\x72\x79\xb9\xfa\x79\x7b\xfa\x05\x43\xed\xd0\x4d\x2c\x4a\x2f\xcd\x4d\xcd\x2b\x29\x0e\x4a\x4d\xcc\xc9\xd5\x2b\x48\x2c\x2e\x2e\x4f\xd1\x53\xd1\x70\x74\xf1\xf5\xf4\x8b\x0f\x0d\x76\x0d\xd2\xb4\x85\xf1\x02\x1c\x83\x83\xc3\xfd\x83\x5c\x34\x15\x30\xf4\x15\xe5\xe7\xa4\x16\xe3\xd0\x06\xe6\x01\x02\x00\x00\xff\xff\xca\x81\xb8\xf8\xd8\x00\x00\x00")

func environmentJenkinsJvmEnvironmentBytes() ([]byte, error) {
	return bindataRead(
		_environmentJenkinsJvmEnvironment,
		"environment/jenkins-jvm-environment",
	)
}

func environmentJenkinsJvmEnvironment() (*asset, error) {
	bytes, err := environmentJenkinsJvmEnvironmentBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "environment/jenkins-jvm-environment", size: 216, mode: os.FileMode(420), modTime: time.Unix(1532036458, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _environmentRequiredPlugins = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\xc1\x0d\x03\x21\x0c\x44\xd1\x7b\x7a\x71\x03\xe9\xc6\xc0\x84\xb0\x58\x78\x35\x18\xad\x94\xea\x73\xf2\xf1\xeb\x1d\x7e\x25\x1a\x56\x0c\xb5\xfd\x36\x0d\xec\x78\x95\x33\xac\x49\x31\xaf\x13\x94\xdb\x4e\x1f\x2b\xed\xf2\x22\x6d\x5b\xa6\x9e\xf8\x3a\xc7\x0f\x72\xd3\x2f\xd4\x48\x78\x9c\xf3\x63\xfe\x88\xf6\x4e\x74\x0d\x67\xd2\x3c\x05\x5c\x08\xe4\xf0\x1f\x00\x00\xff\xff\xc2\xc9\xda\x4d\x83\x00\x00\x00")

func environmentRequiredPluginsBytes() ([]byte, error) {
	return bindataRead(
		_environmentRequiredPlugins,
		"environment/required-plugins",
	)
}

func environmentRequiredPlugins() (*asset, error) {
	bytes, err := environmentRequiredPluginsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "environment/required-plugins", size: 131, mode: os.FileMode(420), modTime: time.Unix(1532377175, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initGroovy0JenkinsConfigGroovy = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\xb0\xee\xc1\x52\x11\x33\x40\xd1\x53\xdb\x2c\x60\x64\xb7\x6d\x52\xef\xd6\x8d\xbb\xe8\x61\xb1\x28\x68\x6a\x2c\xd3\xa5\x38\xc6\xcc\xd0\xd9\x34\xf0\x7f\x2f\x44\xc9\xb6\x94\x38\x41\x7b\x48\x00\x52\x8f\x6f\x3e\xde\xbc\xf1\xd7\x5f\x55\x84\xb8\x7b\xc8\x32\x57\x6f\x91\x44\x6d\x20\xfc\xed\x02\xeb\x1a\x4b\xf0\xfa\x9b\xc3\xf5\x3a\x96\x8c\x41\x33\xd8\x48\x4e\x1e\x4e\x1f\x86\xf8\xdb\xf6\x34\x43\x6b\xc4\x61\xb8\xc6\xb0\x72\x55\xa4\x74\x78\x89\xca\x32\xad\xf4\x5b\x58\x99\xe8\xe5\x9a\x62\xbd\xbc\x61\x8e\x40\x4f\x03\x1c\xf1\xfc\x6d\x7d\x0a\x8f\x54\xe9\x0e\x61\x9d\xde\xfa\x58\x35\x58\x13\x65\x8d\xe4\xfe\x81\x2d\xe1\x06\xac\xfc\x6f\x3c\x0b\x19\x81\xea\x4c\x9d\xc7\x34\x7e\x8f\x10\xe1\x46\xa0\x9e\x46\x59\x43\x10\x67\x8d\x20\x9d\xad\x78\x63\x76\x66\x8b\xcc\xa0\x37\xb8\x2c\xd9\xeb\xe6\xef\x2d\xfb\x85\x25\xb7\x95\x19\x9a\xb2\x57\xee\x53\x6c\x9b\xe3\xa1\xb3\xb7\xb8\x7c\x6f\x82\xa9\xa0\x86\x20\x59\x56\xc2\x4a\xb9\xc0\x62\x82\x05\x75\xa5\x3a\x90\xae\x40\x6e\xba\xdb\xbc\xc8\xb6\xe4\x82\xf8\xa0\x46\x93\xc9\x1b\x65\x09\x8c\xb8\x50\x29\x8f\xd6\x78\x15\x19\x48\x8d\x1f\x1f\xf5\x47\x06\xda\xef\xc7\xa3\x96\xb3\xd5\xe8\x0e\x8c\xaf\xd5\x95\x0a\x70\xaf\x7e\x49\x37\x73\x72\x3b\x23\xb0\xe8\x9a\x90\x00\xf9\xca\x78\x86\x22\xeb\xbd\xd1\x29\x0a\x4c\xad\xc5\x18\x24\xef\xf1\x5f\xa4\x60\x73\xc3\x7c\x8f\x54\xee\xf7\xe3\x22\x3b\xe4\xaf\x19\x64\x48\xdc\x63\x2c\xda\xbc\x0e\xba\x74\x49\xfd\x14\xbd\xbf\xc6\x20\x84\xfe\xb7\x60\x61\x86\x55\x05\xe5\x4d\x98\x76\x72\x26\x11\x16\xdd\x93\xbc\xc8\x8e\xaa\x32\xc8\xd4\x7b\xbc\x9f\x06\x0c\x0f\x35\x46\xbe\x03\x53\x1e\xea\xe8\xe7\x73\x9e\xe9\xc0\x53\x64\xc3\xde\x32\x48\x6a\x6d\x37\x2b\x2a\x92\x6f\x7b\x4b\xbe\x6d\xed\x66\x68\x8f\x93\x62\x67\x0d\xd3\xc8\x98\x17\x4f\x1f\x35\x69\x7d\x24\x9f\x9f\x88\xcf\x42\xa6\x65\xed\xc2\xb4\x2c\x09\x98\x13\x36\x5d\xbc\xab\x8d\x7b\xe1\x89\xd9\x35\xb3\x72\xbe\xa0\xb8\x55\x65\x6b\x50\x65\x1b\x87\xb2\x72\xc9\xa3\xa3\x41\xb3\x7a\xe6\xcd\x1b\x75\x9e\x7b\x3a\x17\x8a\x50\xbc\x12\xe5\xf6\xc3\x6c\x3e\xca\x16\x20\x3f\x2e\x84\x5c\xa8\xde\x28\x53\x41\x90\x39\xa1\xa0\x45\xcf\x33\xc7\xa2\xae\xd4\xa7\x71\x03\xfc\x6e\x62\x31\x04\xb0\xd2\x4c\xd5\xdc\x85\x6a\xfc\x79\x28\xde\xe0\x69\xfe\x9c\xe9\xc9\xec\x79\xb3\x83\xf6\x0d\x92\xe4\x8f\x8f\x4a\x1f\x4f\x6a\xbf\x1f\x82\x3f\xc4\xfa\xdd\x17\xb0\x51\x90\x38\x41\x8f\xa7\x04\x1d\x16\x08\xc1\x2c\x3d\xa8\xc4\xa6\x04\x55\x6d\x58\x80\xd4\x61\x95\x28\x8e\x4b\x7e\x60\x81\xba\xd7\x4f\x17\x9a\x45\x84\x34\xf0\x72\xd2\xf0\xcf\xb5\x13\xf0\x8e\xe5\x2e\x7a\xd0\xd6\x1b\xe6\xa2\x49\xe9\x7d\x62\xfd\xd5\x79\xbf\xb8\x77\x62\xd7\xdd\x3c\xff\xf0\x72\xb7\xbb\x6d\xa7\x4c\x7f\xca\x47\xd9\x7f\xd8\x6c\xed\x68\x36\xff\x07\x18\xce\x0b\x6d\xca\x32\xa9\xff\xb3\xc7\xa5\xf1\xe7\xb9\x12\xe0\x0f\x72\x55\x05\x8d\xca\xcd\x6e\xe0\x17\x4c\xfb\xda\xb4\x30\x40\xa9\x36\xb8\x1c\x65\x69\xdc\x86\x0b\x35\xc5\x38\xb7\x35\xf3\x45\xea\xb5\xc6\x28\x17\xea\xd3\xf7\x9f\x2f\xda\x55\xe2\x3c\xe4\x63\x3d\x2e\x8a\x42\x53\x0c\x2d\x53\x7e\xfa\x72\xb9\x33\x74\xd9\x99\xe6\xaf\x35\xd6\x70\xe9\x82\x13\xdd\xfe\x7a\xea\xf2\xb2\xc9\x65\xb2\xc1\xe5\xa4\x64\x3f\x2e\xb4\xc0\x17\x79\x96\xb9\xd9\x35\x89\x1f\x24\xee\x9b\xa7\x33\x5f\x96\xfd\x1b\x00\x00\xff\xff\x83\x27\xe3\xa0\x90\x07\x00\x00")

func initGroovy0JenkinsConfigGroovyBytes() ([]byte, error) {
	return bindataRead(
		_initGroovy0JenkinsConfigGroovy,
		"init-groovy/0-jenkins-config.groovy",
	)
}

func initGroovy0JenkinsConfigGroovy() (*asset, error) {
	bytes, err := initGroovy0JenkinsConfigGroovyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init-groovy/0-jenkins-config.groovy", size: 1936, mode: os.FileMode(420), modTime: time.Unix(1532381258, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jobScriptsInstallDslJobSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xd1\x4e\xdb\x30\x14\x7d\xc6\x5f\x71\xb9\x20\x15\x24\x5c\x57\xdb\x5b\x46\xd9\x26\x06\x08\x86\xd6\x89\x56\x02\x69\xda\x83\x93\xdc\x36\x0e\x8e\x9d\xd9\x4e\x8b\xd4\xe5\xdf\x27\xbb\xd1\xfa\xb2\x4a\x7b\x4a\xe4\x7b\x7c\xce\xb9\xc7\xe7\xe4\x58\x74\xde\x89\x5c\x19\x41\x66\x0d\xb9\xf4\x15\x6b\x9d\x32\x81\xcc\x9a\xb1\x13\x78\x96\x2a\xc0\xd2\x3a\x78\x20\xf3\xaa\x8c\x87\x60\x41\x5b\x59\xb2\x4d\xa5\x34\xc1\x31\x14\x9d\xd3\xc0\x3d\x70\x5e\x91\x2c\x81\x73\x47\xbf\x3a\xf2\x01\xee\x6e\x16\xb0\xdd\xc2\xf8\x73\xab\xa0\xef\xe1\x37\xac\x1c\xb5\x80\xef\x26\x13\x98\x7d\x45\xb8\x02\x51\xd2\x5a\x98\x4e\xeb\x0f\xac\xb4\xec\x88\x8a\xca\x02\x5e\x3f\xde\x67\x49\x54\x99\x55\xd2\xad\x77\xba\xe3\x31\xb2\x23\xaf\x89\x5a\x78\xcf\x4a\x6b\x28\x9a\x5b\x51\x80\xeb\xf9\xd3\x2d\xb4\xce\x06\x2a\x82\xb2\x06\x82\x7d\x25\xc3\x76\x64\x77\x14\x12\xcf\x3f\x31\xc8\x0a\xd7\x35\xf9\x22\xfe\x4f\x4f\xcf\xd2\x1e\xa3\xbd\xe1\x34\xbc\xf7\xbe\x23\x27\x64\xab\xc4\x5b\xa3\x3f\xbe\xb5\x32\x54\xd3\xc2\x9a\x42\x86\x33\x21\x12\xe4\x69\xb7\xee\xad\x22\x5d\x5e\x60\x86\x17\xc3\xf9\xf9\xe8\x3c\x5a\xdc\x38\x15\x08\x6a\x9b\x43\xe9\x75\x0c\x4f\xc2\x52\x69\x62\x85\x0c\x70\x25\x42\xd3\x8a\x28\xf9\x60\xf3\x6f\xb2\x21\xe8\xfb\x71\x84\x5d\x5e\xde\xcc\x1e\xd9\x30\xf8\xe2\x35\xf4\x3d\x8b\x27\xc3\x5a\xcf\x71\x93\xc8\x75\xe0\x7e\x86\x89\xfe\xc0\x34\xba\xf2\x41\xba\x00\x9e\xa8\x4c\xd6\x36\x2a\x54\xd0\x4a\x27\x1b\x0a\xe4\xfc\x20\x33\x8f\xa0\x18\xdf\x3e\x94\xda\xe6\xa2\x70\x24\x03\xf1\xe1\x5d\x78\x6d\x73\x2f\xf2\x4e\xe9\x12\xd9\xae\x0b\x2f\xf0\x7d\x36\x5f\xfc\xef\x2d\xe0\x7c\x69\x5d\x93\x52\x99\x4c\x3f\x1d\x8a\x64\x40\xd5\xde\x9a\xe9\x68\x8b\x7f\xcd\x62\x06\x3f\xb6\x68\x64\x43\x98\x61\x6d\xf3\x08\xc6\x0b\xc0\x48\x87\x59\xfa\x4c\xb0\xff\xd9\x8f\x80\xaf\x87\x8e\x92\x03\x7c\xe1\xb1\x13\x3c\x3d\x7e\x06\xa7\xfb\x26\x20\xfb\x13\x00\x00\xff\xff\x4d\x82\x3a\xd3\x12\x03\x00\x00")

func jobScriptsInstallDslJobShBytes() ([]byte, error) {
	return bindataRead(
		_jobScriptsInstallDslJobSh,
		"job-scripts/install-dsl-job.sh",
	)
}

func jobScriptsInstallDslJobSh() (*asset, error) {
	bytes, err := jobScriptsInstallDslJobShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "job-scripts/install-dsl-job.sh", size: 786, mode: os.FileMode(420), modTime: time.Unix(1532988742, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jobScriptsInstallXmlJobSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x4d\x4f\x1b\x31\x10\x3d\xe3\x5f\xf1\x30\x48\x01\x09\xc7\xa8\xbd\xa5\x04\x8a\x22\x40\xa1\xa8\xa9\x20\x12\xb9\x7a\x77\x87\xc4\xe0\xb5\xb7\xde\xd9\x10\x44\xf7\xbf\x57\xde\x44\xa5\x91\xc8\xcd\xf2\xcc\xbc\x8f\x79\x73\xb0\xaf\x9b\x3a\xea\xcc\x7a\x4d\x7e\x89\xcc\xd4\x0b\x51\x45\xeb\x99\xfc\x52\x88\x03\x3c\x1a\xcb\x78\x0a\x11\xb7\xe4\x5f\xac\xaf\xc1\x01\x2e\x98\x42\xbc\x2e\xac\x23\xec\x23\x6f\xa2\x83\xaa\xa1\xd4\x82\x4c\x01\xa5\x22\xfd\x6e\xa8\x66\xdc\x5c\x4d\xf1\xfe\x8e\xfe\x65\x65\xd1\xb6\xf8\x83\x79\xa4\x0a\xf2\xcb\xe9\x29\x26\x3f\x24\xce\xa1\x0b\x5a\x6a\xdf\x38\xf7\x4d\x14\x41\xec\x51\xbe\x08\x90\xa3\xbb\xf1\xa0\x23\xb5\x7e\xde\xf1\x3e\xaf\x79\xfb\x7d\x29\xf6\x6a\x47\x54\xe1\xab\x28\x82\xa7\x24\x6e\x4e\x8c\xd1\xc3\xfd\x35\xaa\x18\x98\x72\xb6\xc1\x83\xc3\x0b\x79\xb1\x06\xbb\x21\xee\x70\x3e\xed\x91\x22\x8f\x4d\x99\x4d\xd3\x7b\x78\x78\xd4\xf9\xe8\x7d\x08\xee\x8a\xe3\xba\x6e\x28\x6a\x53\x59\xbd\x2a\xdd\xc5\xaa\x32\xbc\x18\xe6\xc1\xe7\x86\x8f\xb4\xee\x5a\xee\xd7\x76\xaf\x2d\xb9\xe2\x44\x0e\xe4\xc9\xe6\xff\xb8\x77\x9c\x24\xbe\x46\xcb\x84\xe7\x90\x61\x55\xba\xb4\x3c\x83\x27\xeb\x48\xe4\x86\x71\xae\xb9\xac\x74\xa2\xbc\x0d\xd9\x4f\x53\x12\xda\xb6\x9f\xda\xce\xce\xae\x26\x77\x62\x53\x98\x95\x0e\x6d\x2b\xd2\xcf\xc6\xd6\x63\x72\x92\xb0\x76\xcc\x0f\x64\x07\xbf\xa3\x9a\x54\xe5\x91\xcc\x5a\xd6\x06\xf2\xb2\x28\xd2\xa2\x92\xce\xed\x09\x29\xfe\x05\x3c\xc3\xaf\xc9\xc3\x74\x7b\x47\x09\x66\xcc\x54\x5e\x78\x53\xd2\x70\x7b\xb4\x07\xa5\x0a\xc3\x46\x65\xd6\x9b\xf8\x86\xef\xbb\xec\xaa\xe5\xe6\x7a\x28\x42\xce\x54\x4a\x4b\x75\xb1\x0c\x70\xf8\x91\x91\xfc\xaf\x69\x14\xd2\x81\xb2\x9a\xbe\x55\x34\x60\x5a\x71\x8a\x47\xfe\x0d\x00\x00\xff\xff\x05\xc4\xa9\x6d\xcc\x02\x00\x00")

func jobScriptsInstallXmlJobShBytes() ([]byte, error) {
	return bindataRead(
		_jobScriptsInstallXmlJobSh,
		"job-scripts/install-xml-job.sh",
	)
}

func jobScriptsInstallXmlJobSh() (*asset, error) {
	bytes, err := jobScriptsInstallXmlJobShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "job-scripts/install-xml-job.sh", size: 716, mode: os.FileMode(420), modTime: time.Unix(1532987293, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jobdslSeedJobDsl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xdd\x4e\xc3\x30\x0c\x85\xef\xf7\x14\xbe\x73\x86\x36\xde\x61\x80\x40\x9a\x2a\x06\xda\x13\x24\xed\x99\x96\x2e\x4b\xaa\xd8\x03\x26\xd4\x77\x47\x6d\xe9\x4f\x6e\xac\xf8\x3b\x76\xbe\x9c\x32\x70\xd4\x7b\xc0\x3e\x39\xc3\x65\x86\x55\x6c\x6b\xc4\x8b\x8f\xb2\xad\x93\x13\x5e\xd3\xef\x8a\x88\xc8\x85\x54\x5e\x0e\xd1\xf0\xc3\xd4\x9a\xda\x05\xbe\x10\x0c\xbf\x15\x87\xa7\x5d\xc1\xeb\x09\x4a\x69\xe3\xe7\x0d\x37\xbc\xa6\x6c\x78\x57\x8c\xac\x5d\xf5\xa5\x82\x94\xd9\x37\xea\x53\x34\xbc\x4f\xee\xe5\x58\x50\x9d\xdc\x86\xf4\x6c\x95\x06\x19\xa1\xa4\x67\x64\x1a\x5c\x86\x31\x2f\x4d\xb0\xf7\x77\x7b\x85\xe1\xe7\x3e\x45\x11\xdf\x63\xa4\xcf\x34\x36\xdb\x2b\x14\x59\x16\xae\x27\x1f\xf0\xd1\x01\xc3\x75\x72\x8f\x95\x04\xde\xd0\xf8\x70\x07\x67\xbd\xbe\x8a\xa2\x59\xce\x57\x12\x16\xb7\xee\xe0\x47\x91\xa3\x0d\xf3\xc2\xf9\xf3\xed\xff\xae\xf6\x2f\x00\x00\xff\xff\x04\xeb\x1c\x98\x64\x01\x00\x00")

func jobdslSeedJobDslBytes() ([]byte, error) {
	return bindataRead(
		_jobdslSeedJobDsl,
		"jobdsl/seed-job-dsl",
	)
}

func jobdslSeedJobDsl() (*asset, error) {
	bytes, err := jobdslSeedJobDslBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jobdsl/seed-job-dsl", size: 356, mode: os.FileMode(420), modTime: time.Unix(1532986279, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginScriptsInstallPluginSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x51\x4f\xdb\x30\x10\x7e\xc6\xbf\xe2\x30\x48\x05\x89\xd4\x68\x7b\x2b\xa5\xdb\x54\x01\xea\xc6\x56\x54\xaa\x81\x34\xed\xc1\x4d\x2e\x89\x87\x6b\x67\xb6\x53\x8a\x58\xfe\xfb\x64\x3b\x55\x19\x42\xdb\x32\x2d\x2f\x71\xec\xf3\x77\xdf\x7d\x77\x5f\xf6\x76\x59\x6d\x0d\x5b\x08\xc5\x50\xad\x60\xc1\x6d\x49\x2a\x23\x94\x43\xb5\x22\x64\x0f\x6e\xb8\x70\x90\x6b\x03\xef\x51\xdd\x09\x65\xc1\x69\x90\x9a\x67\xe4\xbe\x14\x12\x61\x17\xd2\xda\x48\x48\x2c\x24\x49\x89\x3c\x83\x24\x31\xf8\xbd\x46\xeb\xe0\xe2\x6c\x0e\x8f\x8f\xd0\x7f\x57\x09\x68\x1a\xf8\x01\x85\xc1\x0a\xe8\xab\xe3\x63\x98\x7e\xa0\x30\x02\x96\xe1\x8a\xa9\x5a\xca\x13\x92\x69\xb2\x83\x69\xa9\x81\x8e\x2f\x27\x83\x90\x54\xa8\x22\xe4\xfd\x16\xf3\xf6\xfb\x94\xec\x58\x89\x58\xc1\x6b\x92\x69\x85\x9e\x5c\x81\x0e\xc6\xd7\xb3\x73\xa8\x8c\x76\x98\x3a\xa1\x15\x38\x7d\x87\x8a\x44\xb0\x0b\x74\x01\xe7\xc5\x18\x4a\x52\x53\x2f\x17\x73\xbf\x3e\xdd\x3f\x08\x75\xf4\xb6\x84\xc3\xe1\xc4\xda\x1a\x0d\xe3\x95\x60\xeb\xa5\x7c\xb3\xae\xb8\x2b\x4f\x53\xad\x52\xee\x0e\x18\x0b\x21\xb3\x58\xee\xb9\x40\x99\x1d\xd1\x01\x3d\x6a\xf7\x0f\x7b\x87\xa4\xa5\x31\x51\xd6\x71\x29\x3d\x93\x4a\xd6\x85\x50\x41\x97\xab\xb0\x9c\x64\xd0\x34\x6f\xb7\xdf\x9f\xd1\x58\x4f\xb1\x69\x28\x89\xd2\xde\xc2\xd5\xf4\x7a\x0e\x49\x06\xbd\x61\x2b\xc6\x68\x28\x22\x64\x8b\x77\x4a\xff\x0a\x10\xd8\x68\xc8\x36\x10\xbd\xb6\x63\x68\xa0\x37\xd6\xbe\xdf\x2e\x99\x3f\x54\x38\x00\x87\x6b\xe7\xcb\x7d\x12\x41\x6f\x13\xaf\x61\x12\xc4\x1a\xc0\xfe\x56\x39\xfa\xa4\xc7\x91\xcc\x47\xae\x78\x81\x86\xb5\x14\x3f\x61\x8a\xd6\x72\xf3\x10\xe9\xd8\x67\x55\x75\xcb\xc0\x2c\xcf\x71\x86\xd6\x71\xe3\xfc\x00\xdc\x1b\xe1\x70\x23\x6a\xaa\x55\x2e\x0a\x3f\xa0\xb9\x90\x48\x52\xee\x60\xc4\xdc\xb2\x62\xcf\xc4\xe9\xbf\x24\x4e\xbf\xbd\x3d\x1c\x9e\x4d\x2f\xc9\x36\x62\x1c\xb7\x9b\x86\xf8\xfd\xb6\xa1\x37\x7e\x94\x7c\xa2\xae\xf0\x03\x1a\x68\x75\xbd\x46\x88\xc8\xe1\x8b\x77\x59\xe7\x7a\xbe\x9e\x80\x2b\x51\x11\x02\x00\xf0\x1b\x3b\xfb\xe3\xff\x6a\x69\x0f\x98\xe9\xf0\xf2\xcf\x9f\xdd\xbd\x89\xdc\x98\x3c\x02\x78\xa3\x47\xea\xa6\x56\xbf\x76\x9a\x6c\x71\x67\xb5\x52\x1e\x32\x1e\xd4\x86\x07\x9b\xfb\x04\x9d\x0c\x17\x10\x63\xf5\x19\x50\x9b\x1a\x51\x39\xff\x6b\xf8\x87\x96\x1d\x52\x48\x56\x5d\xc7\x3b\x66\x9c\xe3\xda\x91\x5c\xfc\x0c\x00\x00\xff\xff\x79\x8c\xd3\xd8\x9a\x05\x00\x00")

func pluginScriptsInstallPluginShBytes() ([]byte, error) {
	return bindataRead(
		_pluginScriptsInstallPluginSh,
		"plugin-scripts/install-plugin.sh",
	)
}

func pluginScriptsInstallPluginSh() (*asset, error) {
	bytes, err := pluginScriptsInstallPluginShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugin-scripts/install-plugin.sh", size: 1434, mode: os.FileMode(420), modTime: time.Unix(1531432614, 0)}
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
	".DS_Store":                           Ds_store,
	"environment/jenkins-jvm-environment": environmentJenkinsJvmEnvironment,
	"environment/required-plugins":        environmentRequiredPlugins,
	"init-groovy/0-jenkins-config.groovy": initGroovy0JenkinsConfigGroovy,
	"job-scripts/install-dsl-job.sh":      jobScriptsInstallDslJobSh,
	"job-scripts/install-xml-job.sh":      jobScriptsInstallXmlJobSh,
	"jobdsl/seed-job-dsl":                 jobdslSeedJobDsl,
	"plugin-scripts/install-plugin.sh":    pluginScriptsInstallPluginSh,
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
	".DS_Store": &bintree{Ds_store, map[string]*bintree{}},
	"environment": &bintree{nil, map[string]*bintree{
		"jenkins-jvm-environment": &bintree{environmentJenkinsJvmEnvironment, map[string]*bintree{}},
		"required-plugins":        &bintree{environmentRequiredPlugins, map[string]*bintree{}},
	}},
	"init-groovy": &bintree{nil, map[string]*bintree{
		"0-jenkins-config.groovy": &bintree{initGroovy0JenkinsConfigGroovy, map[string]*bintree{}},
	}},
	"job-scripts": &bintree{nil, map[string]*bintree{
		"install-dsl-job.sh": &bintree{jobScriptsInstallDslJobSh, map[string]*bintree{}},
		"install-xml-job.sh": &bintree{jobScriptsInstallXmlJobSh, map[string]*bintree{}},
	}},
	"jobdsl": &bintree{nil, map[string]*bintree{
		"seed-job-dsl": &bintree{jobdslSeedJobDsl, map[string]*bintree{}},
	}},
	"plugin-scripts": &bintree{nil, map[string]*bintree{
		"install-plugin.sh": &bintree{pluginScriptsInstallPluginSh, map[string]*bintree{}},
	}},
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
