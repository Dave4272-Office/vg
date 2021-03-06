// Code generated by go-bindata.
// sources:
// data/bash
// data/fish
// data/sh
// data/zsh
// DO NOT EDIT!

package cmd

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

var _dataBash = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\x5d\xeb\xa2\x40\x14\xc6\xef\xe7\x53\x9c\x66\xe7\xa2\x5a\x44\xd8\x5b\x71\x77\x63\x8b\x08\x2a\xc5\xcd\x0a\x22\xa6\xc9\xa6\x14\x26\x0d\x1d\xc5\x65\x6b\x3f\xfb\xe2\x28\x95\x31\x4b\x37\x7b\xf3\xef\x22\xf4\xf8\x9c\xe7\xbc\xfc\xce\xa7\x8e\xb9\x8f\x62\x73\xcf\xb2\x10\xa1\xe2\xd4\xed\xc1\x6f\x04\x00\xcd\x7f\xf5\x13\x49\xc0\x04\xf8\x3f\x47\x74\x6d\x69\x82\xfe\x23\xa8\x34\xf6\x2e\x60\x19\x07\x4c\x0c\x0c\x51\x0c\xfd\xb2\xdf\x03\x1e\x84\x09\xe0\x5f\x3c\xc3\x60\x59\xc0\x33\x16\xec\x5a\x49\xfe\x6b\x52\xfe\x26\x29\xe3\x12\x3e\x97\x60\xb5\x03\x79\x13\xb8\xc1\x97\xaf\x60\x1e\x78\x61\xc6\xb9\x10\x48\xc5\x28\x2d\x4e\x80\xc9\x77\x5c\xbf\x6e\x00\x13\x55\x19\x43\xc7\x6e\xaa\x6c\xe1\x7a\x55\x46\x46\xde\xd2\xac\xb5\x9a\x12\xdd\x10\xaa\x5c\x69\xca\x45\xc2\x0e\xf7\xcd\xb9\x9e\x33\x73\x17\xf4\x87\x33\x9b\x0d\xe6\x43\x9b\xd0\xe5\xc4\x5b\xf8\x83\xe9\xd8\xa1\x8e\x37\x19\x4f\xe6\x83\x29\x6d\x6b\x54\x1a\x2f\x98\x00\x4c\xba\xc5\xa9\x7e\x34\x8c\x2c\xe4\x42\x40\x45\xa6\x87\x55\x31\xd7\x1b\x2d\x9f\xdc\xdc\xd5\xd0\xc6\xb8\x69\xe2\x92\x26\xe7\x8b\xa4\x41\x72\x3e\xb3\xf8\xd1\x4c\xcd\x49\xb0\x4c\x52\x5e\x46\xd2\x26\xdf\xd0\xbf\xf1\xea\xe8\x7e\x08\xb8\xd1\x51\xc1\x72\x57\xc3\x1a\x15\xd1\xac\x0a\xc3\xd6\x02\x19\xf2\xf8\x6e\xab\xdd\xa7\x32\x41\x0f\x4d\xb5\x5b\x96\xcb\x84\xb2\x40\x46\x05\x93\x5c\x7d\x3a\x46\xff\xef\x8c\x94\x28\xe5\x32\x4f\x63\x20\x77\x52\x15\xf0\x6a\xac\x0d\x74\xaa\xc9\x5a\xf7\x82\xc1\xfe\x03\x1a\xea\xb0\x7d\x1e\xf1\xfd\xdd\xd9\x44\x73\x87\x2f\x12\xac\xab\x63\xc1\x6b\x47\xe8\x18\xa1\xbf\x01\x00\x00\xff\xff\xf0\x27\xb1\x97\x4b\x04\x00\x00")

func dataBashBytes() ([]byte, error) {
	return bindataRead(
		_dataBash,
		"data/bash",
	)
}

func dataBash() (*asset, error) {
	bytes, err := dataBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/bash", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x7b\x6f\xdb\x38\x12\xff\x3f\x9f\x62\xea\x1a\x88\xbd\x80\xe2\x4b\xee\xbf\x16\x39\x20\xdd\xf4\xda\x00\x6d\x13\x24\xed\x2e\x0e\xbd\x56\xa0\xa5\xb1\x4c\x84\x26\xb5\x24\xa5\xc4\x8b\x7c\xf8\x03\x49\x3d\x48\x89\x4e\xb2\x57\x77\x0d\x04\x49\xc8\x79\x71\x1e\xbf\x19\xd2\x2f\x5f\x2c\x96\x94\x2f\x56\x54\xad\x17\x07\x07\xab\x8a\x67\x9a\x0a\x0e\x75\x71\x00\x00\x40\x57\xc0\x85\x86\x4c\x54\x5c\xc3\x94\xc8\xa2\x86\x7f\xc1\x22\xc7\x7a\xc1\x2b\xc6\x2c\x89\xf9\x64\x62\xb3\x21\x3c\x6f\xb9\xcc\x47\xa2\xae\x24\xb7\xff\x22\xcf\xed\xef\x97\x80\xd9\x5a\x38\x31\x5f\x8f\xbf\xd9\x35\x75\x47\x75\xb6\x0e\xd7\xac\x44\xa2\x10\x28\xa7\xba\x5b\x31\x9f\xb4\x2e\x52\xb3\xd8\xd2\x1f\x1d\x25\x43\x1e\x92\x69\x5a\x13\x8d\x23\xbe\x76\xe3\x11\xde\x1c\x77\x72\x47\xb6\x2c\x4b\x96\x97\x24\xbb\x25\x05\xaa\x11\x4b\x64\xab\xd1\xa2\xb4\x14\xdb\x88\x0a\xbb\xfe\x88\x7d\x05\x13\x4b\xc2\xde\xde\x63\x36\x62\xee\xb7\x1e\xe1\xaf\xca\x42\x92\x7c\x7c\xb8\xe1\xba\x25\x3e\xfc\xe5\x30\x20\xec\x63\xec\x34\x74\xa1\x35\x3f\x5e\xe6\xb4\x31\x3a\x88\x3b\xde\x2e\xd7\x05\x30\xca\x6f\x1d\x6b\xc0\x39\x70\x1a\x5d\xc1\x57\x48\xfe\x84\xc9\xf4\xb7\x8b\xeb\xcf\x5f\xce\x3e\xbc\xbb\x9c\x40\x7f\x2a\x9b\x4f\x6f\xa5\x14\xf2\x15\x7c\x12\x50\x53\xa9\x2b\xc2\x0a\x01\x77\x42\xde\xaa\x92\x64\x08\x54\xb9\x94\x40\x20\x1a\xf4\x1a\x61\x23\x36\xc8\xfb\xb4\xaa\x0b\x58\x23\x2b\x63\x81\x74\x29\x0c\xc7\x41\x12\x67\x39\xf4\xb6\xa4\x57\x67\x9f\xdf\x2f\x94\xcc\x16\x91\xa3\x04\x09\xe3\x0e\xc2\x77\x1d\x24\x92\x62\x56\x60\xc7\x3a\xf3\x4a\x70\x0e\x09\xd3\x70\xe2\xb1\x2b\xd4\x60\xa3\x7e\xf2\x0d\x66\x4b\xa2\x90\x93\x0d\xc2\xf4\xea\xf7\xf3\x79\x28\xcb\x10\xf6\xbe\x99\xd5\x85\xad\xb1\x1b\xd4\x9a\xf2\x42\x35\xa9\x73\x62\x53\x67\xde\x73\x24\xc5\x3d\x74\x66\xc3\xb4\x13\x70\xd0\xc5\xe0\xcc\x59\x4e\x79\xe1\x79\xa7\xe3\xaf\x0b\xe9\x3b\xed\xfa\xf2\xf2\x33\x4c\xdf\x5f\x7e\x7c\xbb\x38\xea\x42\x16\x57\x66\x1d\x0c\xd3\xba\x90\x5f\x8f\xbf\x2d\x3c\xd1\x1d\xb9\x60\x79\x5a\x88\x92\xe8\x35\x4c\xdf\x5d\x3a\x72\x2b\xba\x10\xe1\x01\xd2\x5e\xe8\xe5\x87\xf3\x96\xb4\x67\x6f\xd1\xa7\x8f\xd4\xbb\xcb\x37\x17\x9f\x26\x03\x37\xc7\x44\xbd\xb9\xf8\x04\x8e\xda\x79\x84\x29\x7c\x1e\x4f\xcc\x26\x03\xc7\xe3\xa0\x8d\x45\x18\xda\xf4\xdf\x17\x37\xef\x61\x6a\xfe\x7c\x8a\xf2\xe6\x3d\xcc\x6c\xa8\x12\x05\xaf\x2c\x07\x3c\x40\x56\x69\x48\x96\x70\x92\x78\xc1\x76\x48\x92\xae\x08\x63\x4b\x92\xdd\xda\x24\xf9\xa3\x42\xb9\x6d\xb3\x64\xee\xa5\xe5\x74\x48\x7d\x0a\x5a\x56\x18\xf1\x59\xeb\xf0\x30\xb4\xaf\xa2\x2e\xd8\xed\xc5\xb8\x94\xb8\xbb\x1a\x1f\x0f\x8a\xd5\x78\xd7\x92\x6d\x6e\x73\x2a\x21\x29\xe3\xe5\x1c\x52\xb8\xd8\x06\xd2\x9d\x1d\x8d\x0e\x6b\xc4\x2e\xc0\x4a\xcf\x2f\x6e\xce\xde\x7c\x78\x9b\x5e\x5d\x5f\x7e\xbc\xfa\xec\x27\x54\x8b\x17\x0a\x92\x0c\x4c\x0b\x4e\x4b\x29\x36\xa5\x86\xd4\xe4\xa5\xb7\x30\xe2\x80\xd8\x66\x57\x8e\x26\x7d\x67\xbd\x05\x73\x98\x84\x98\xbf\x4b\x7a\x0b\x72\xed\xef\xba\xe8\xfa\xe9\x35\x12\x16\x01\xb9\x01\x68\xfd\x7d\x78\x1d\x69\xc8\x23\xbc\xee\xf1\xe9\xbc\xa5\x1e\x20\x54\x1f\xd2\x2e\x5f\x22\x85\x1a\x52\xb9\xc8\xef\xcc\xdb\x86\x2e\x46\xd5\x95\xec\x0f\xe6\x0a\x46\xa3\x1f\x24\xd3\x30\xc4\x4f\x71\x60\x3c\x29\xc2\xa2\x42\x18\x63\xbb\xbf\x98\x86\x4e\xc0\x5d\x88\xf5\x04\xc9\x63\x04\x83\x78\x8c\xb7\xad\x09\xb1\x3c\xed\x87\xae\x67\x34\xd4\x27\xf2\x78\x4f\xb9\x1c\xe6\xf3\x78\x2a\x0c\x12\xda\xaf\x4b\x18\xf4\xfc\x41\xcf\xed\xa2\xe6\xcd\x6c\xe1\x74\xe9\x5a\xfc\xc1\x23\xa3\xc9\x6b\x68\x18\x19\x55\x1a\x1e\xec\x35\xa0\x90\x58\xc2\xe4\x7b\x4f\xf7\xdf\xe9\x24\x7a\x1f\xb0\xfa\x46\xd3\xcc\x38\x2a\xfe\xe0\x69\xe3\xc9\x6c\x3f\xef\x7d\x39\x38\xd8\x0f\x4c\x51\x36\x5c\x5f\xca\xdc\x41\x40\x38\x76\x14\x02\x0a\xa3\xbd\x82\x82\xea\x75\xb5\x3c\xca\xc4\x66\xf1\x0e\xf5\x8d\x96\x48\x36\x8b\xba\xf0\x44\x5c\x23\x13\x24\x37\x32\x4c\x40\xd5\x1a\x19\x83\x92\x48\x0d\x62\xf5\x97\xa5\x36\xae\xc2\x9a\x30\x48\x12\x27\xcb\x54\x20\x3c\x80\x12\x95\xcc\x70\x18\xa0\xc0\x39\x13\x48\x48\xe8\x89\x17\xa7\x63\x9a\x01\xf8\x5e\x07\x40\x18\xd0\xfa\xd1\xeb\x67\xf6\x31\x49\x3c\x94\x83\xbb\xc9\x33\x6a\xcc\x2f\xa0\xff\x88\x0a\xd4\x5a\x54\x2c\x07\x55\x62\x46\x57\x5b\x20\x5d\xf6\xea\x35\xd1\xed\xee\x12\x01\xef\x31\xab\x34\xe6\xa3\x96\x10\xb9\x1e\xc5\x5b\xc2\xcf\x4e\x35\x13\xcf\xa0\xce\x5a\xa5\x4a\x13\x5d\xa9\x34\x13\x39\xc2\xd4\xfd\xb3\xd7\x10\x3f\x23\x72\x9e\x57\xa6\x9e\x39\xc3\x88\xda\x0b\x4c\xa5\x45\x7f\x7b\x4b\x12\xc1\x93\x9a\x48\x4a\x96\x0c\xe1\xea\xf7\xf3\xd6\x70\x27\x05\x92\x84\xaa\xa4\x09\x59\xa2\xaa\xa5\xd2\x54\x57\x46\xd8\xae\x27\x81\xe6\x4d\x60\x45\x79\x0e\x84\x83\xd1\x96\xb4\xc9\x69\x07\x1c\x16\x00\x43\xbf\x95\x4a\x21\xb4\xbd\xdd\xd8\xfd\xbb\x35\x65\x08\x1a\x95\xc9\xb4\x01\x91\x71\xd7\xc4\x47\x75\x4b\x96\xac\x60\x32\x24\xf5\xae\x22\xe1\xb0\x64\xd4\x73\xbc\xeb\x5d\x99\xda\xcb\xd5\x2c\x23\xfa\x71\x29\xf3\x40\x0c\x5d\xc1\x12\x0b\xca\x2d\x90\xda\x23\xfd\xd1\x37\xce\xd7\x20\x64\x73\x82\x88\xaa\x17\xa7\x5e\x72\xbe\x0e\x9a\x40\x34\xf0\x63\x11\x61\xe3\x1a\xf0\x7b\x61\x19\x6e\xbf\x04\xbd\xa6\x0a\x94\x96\xb4\x54\x16\xf0\x18\x51\x1a\xec\x6d\x2b\x13\x9b\x52\x70\xe4\x1a\x56\x52\x6c\xec\xa6\x59\x3f\x0a\x1b\xd4\x20\x1e\xb3\xe6\x05\x68\xb0\xfc\x00\x0a\x73\x38\x54\x0f\x8b\xaf\xdf\x17\xdf\x7e\x99\x3e\x3c\x1c\xce\x43\xac\x19\x27\xe4\xc1\x81\xb1\x80\xa1\x49\xcd\xcc\x38\x20\x49\xf0\x3e\x63\x95\xa2\x75\x64\x8f\xc3\xc4\xfa\x77\x96\xba\x19\x87\x57\x9b\x25\xca\x54\xac\xd2\x6c\x93\xa7\x44\x16\x2a\xbd\x13\xa9\x28\xb5\x9a\xc3\x29\x1c\x4f\x20\xb9\x27\xbd\x4b\xed\x27\xc9\x61\x72\xd6\xae\x90\x06\xa7\x68\x16\xeb\xfc\x93\xfd\xe8\xef\xdf\x24\x5a\xfd\xbf\xae\x09\x2f\xd0\x3a\xdb\x28\x33\x20\x9e\x53\x89\x99\x16\x72\x0b\x5a\xb8\x9e\x24\x33\x6f\x51\xac\xec\x62\x33\x87\xec\xdb\xc2\x1e\x02\x5b\x0b\xcf\xfb\x15\xa3\x37\xab\xa4\x34\x39\xf2\xf3\x9c\xd4\x0e\x36\x7d\x90\xae\x71\x23\x6a\x54\x20\x38\x9a\xba\xda\x54\x4c\xd3\x92\x79\x87\xb7\xa3\x0d\x61\xcc\x58\x48\x25\x64\x82\x6b\xe4\x5a\xed\xc9\x20\xe4\xaa\x92\x08\x9e\x41\x67\x70\x27\x49\x59\xa2\x84\x95\x90\x90\x63\xe9\x5a\x1a\xe5\x4a\x13\xc6\x5c\x61\xe5\x58\x22\xcf\x91\x67\x14\x15\x50\x6e\xd7\xa2\x33\x25\x57\x1a\x49\x6e\xc7\x0d\xe4\xb9\x90\xfb\xb2\xda\x34\x2d\xf0\xad\x7e\x7b\x5f\x0a\xe5\xa2\xa8\x32\x49\x4b\x3d\x6a\xc4\x35\x61\x09\xe6\xad\xb5\x6d\xa8\xed\x1c\xb3\x27\xab\xbc\x67\xcd\xce\x2a\xdb\xfd\xbd\xf1\xc0\xd1\xb0\x2d\xcc\x44\xa5\x15\xcd\x31\x9a\xf0\xf3\x3d\x59\x64\x27\x8d\xc0\x4f\xef\xcd\x0a\x59\x8a\x4a\x03\xe1\xdb\xd6\xac\x3d\xa9\xb3\xaf\xdd\x81\xba\x5f\x25\x5a\x00\xe2\x39\x20\xb7\xbd\x98\x0c\x32\x9b\x51\x7e\x0b\x54\xb7\x78\xd0\xc6\xa5\xc3\x84\xc9\xc1\x4b\xd8\x9b\x71\xdd\xdb\xa1\x35\xee\xb3\x69\x17\x6d\x60\xcc\x36\x25\x8c\xfe\x89\x2e\xc5\x55\x4b\x6a\x5a\xbb\x2d\x05\x02\x19\x4a\x4d\x28\xdf\x3b\x2a\x58\x1f\x04\x7e\xfb\x60\x56\x9e\x80\x24\xe7\x32\xaa\x7c\x5f\xed\xc9\x1c\xa5\x87\xe6\x28\x6d\x31\x08\xef\xa9\xb2\xd3\x78\x67\xc5\xbe\x80\x88\x89\x8c\xb0\x0b\x07\x32\x4e\xe9\x45\x8b\x38\x04\x9a\xce\xe2\x9a\xf7\x56\x54\xd2\x46\x45\x6d\x95\xc6\x8d\x01\x99\xb6\x8e\xf6\x1d\x18\x03\xcd\xbf\x59\xe0\x6a\x3d\xf1\xd1\x82\xb5\xc5\x3b\xb7\x3e\xea\x68\x9e\x0d\x7b\xca\xdc\xe0\x41\xd3\xcb\x5c\xc2\x0d\xb0\x55\x66\x1a\xd1\xc2\x51\x75\x29\xda\xa7\xaf\xcd\xdc\x7d\x3b\x46\xa1\xae\xca\x10\x80\x6d\x79\x2b\x2f\x51\x29\x77\xa1\xda\x27\xc4\x36\xe3\xbb\xa7\xf7\x66\x2d\xee\x80\xf2\x95\x68\x50\xcd\xaa\x6c\xab\x66\xdf\xc7\xae\x38\x6d\x33\xb4\x51\xff\xa5\x5b\x19\x64\xe9\x4f\x9c\x63\x2a\xee\x03\x86\xb3\xc2\xac\xa8\x38\x82\xf6\xf6\x78\x8d\x59\x1f\x2a\x0b\x3b\x36\x77\xf6\x65\x97\x7b\x20\xf1\xed\x6a\x56\xc2\xf9\x60\x49\x39\x91\x5b\x8b\xfe\xd2\x3e\x50\x18\xfc\xff\x89\x7d\xb9\x46\xa9\xcc\x25\xad\xb7\xeb\x4a\x52\xae\xdb\x32\x76\x9b\xfe\x8b\x88\x9b\x1d\xb6\xa2\x02\x22\x4d\x81\x51\x5e\x4c\x82\xfb\xa6\xd5\x5d\x17\xa9\xdd\x4a\x9b\x16\xd2\xdd\xff\xb2\x4d\x0e\xb3\x66\x91\x51\x8e\x90\x88\x32\x9b\x47\xbe\x1f\xcb\x36\xf9\x1c\x92\x42\xc3\xf1\xf0\x35\xaf\xfd\x8a\x19\x4e\x2d\xd5\xd7\x93\x6f\x83\x07\xbd\xe6\x56\xfc\x8f\xe8\x6b\xb8\x7f\x6f\x3e\x8e\x5d\x96\x0b\xd4\x69\xf0\x0d\xa2\xbd\xd7\xb6\x4a\xed\x97\x0a\x90\x6c\xc8\x7d\x8e\xa5\x5e\xc3\x3f\x21\xd9\x50\xde\xfd\xad\xb7\x25\x42\xde\x5c\x83\xd4\x69\xc0\x75\x7a\x0a\x0f\xee\xe5\x2d\xa9\xe1\xf0\x7b\x21\x44\xc1\xf0\xa8\x10\x8c\xf0\xe2\x48\xc8\x62\x71\xf8\xb4\xba\x13\x4f\xdd\xc9\x73\xd5\xfd\x9f\xba\x62\x47\x63\x7f\xef\xd1\x9e\x54\x17\xd7\xb5\x23\xaa\x5d\x89\x87\xf1\x1d\x45\x7d\xf4\xad\xd4\x0e\x79\xcd\x57\x66\x4f\x09\xdb\xfd\xb0\x1d\xa9\xe0\x1d\xf5\xd3\xdd\x60\x5d\xd5\x4e\x66\xcd\x7b\x6e\x7c\x18\xde\x25\xa4\xb9\x61\xfd\x90\x8c\x0e\xe6\x9d\x94\xc3\xd9\x23\xee\x9d\x1f\xfe\x15\xc9\xfe\x94\x33\x12\x3e\xf0\xf5\xfc\xf0\xe0\x7f\x01\x00\x00\xff\xff\x0f\xad\x00\xbd\xf0\x22\x00\x00")

func dataFishBytes() ([]byte, error) {
	return bindataRead(
		_dataFish,
		"data/fish",
	)
}

func dataFish() (*asset, error) {
	bytes, err := dataFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/fish", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x6d\x4f\xe3\x46\x10\xfe\xee\x5f\x31\xec\x59\x77\x70\x12\x58\xf4\x23\xc8\x2d\x41\x20\x40\xe2\x48\x94\xe4\x5a\x55\xf4\xce\xda\xd8\x1b\xdb\x8a\xe3\x75\x77\xd7\x06\x4a\xf2\xdf\x2b\xef\xc6\xeb\xb7\xcd\x4b\xdb\x2b\x9f\xe2\x9d\xb7\x9d\x99\x67\x9e\x59\x3e\x1c\x39\xb3\x38\x75\x78\x64\x59\x45\x78\x7c\x02\xef\x16\x00\x80\xe7\x15\x21\x20\xfb\x0a\x59\x6b\xcb\x2a\x3f\xb4\xc4\xc7\x9c\x00\xb2\xcf\x11\xc4\xa9\x3c\x28\xff\xe2\x34\x16\x27\xfa\xab\x72\xe0\x95\xc7\x80\xec\xf7\xab\x8b\x9f\xd6\xa8\x25\xbe\xbc\xd4\x9f\xd8\x17\x71\x81\x05\x31\xd8\x57\xa2\xbd\x3e\x02\xb2\xc3\x4b\x2d\xdc\xeb\xc7\x0f\x32\xec\x2f\x70\x48\xb8\xc1\x4f\x2d\x3c\xe0\x3e\x5c\x30\xfa\x66\xbc\x8c\x94\xec\xf5\x10\x26\x74\x86\x93\xdb\x57\xe2\x1b\x9c\xd4\xc2\xbd\x7e\xf2\x2c\x64\x38\x30\x95\x65\x23\xd9\xeb\xe1\x73\xdb\xd6\xa7\xcb\x25\x4e\x03\xa8\xf0\x51\x9e\x11\x8e\xfd\x0a\x28\xb2\xe9\x1a\x2d\x45\x08\x8d\xe2\x6f\xd4\x8b\x10\x92\x38\x5d\x68\x8b\xba\xb0\xda\x2e\x9e\xc3\x33\x9c\xfe\x05\xc8\xfe\xf5\x61\x3c\xfd\x3a\x78\xbc\x1b\x22\xf8\x76\x09\x22\x22\x35\xea\x88\x1f\x51\xb8\x65\x8c\xb2\x0b\x78\xa2\x50\xc4\x4c\xe4\x38\x09\x29\xbc\x50\xb6\xe0\x19\xf6\x09\xc4\x5c\xc5\x27\x80\x45\x69\x0c\x4b\xba\x24\xa9\xd0\x2e\x8a\x10\x22\x92\x64\x8d\xc6\x6b\x11\x23\x22\x67\x29\x9c\xcb\x83\x79\xac\xc0\x1f\x34\x6f\xe4\x8d\x06\xd3\x7b\x87\x33\xdf\x91\x63\x62\xb5\x30\xdb\x49\x25\xdd\x9d\x4a\x11\x36\x30\x5c\x45\x54\xb5\x7d\xcd\x28\x13\xa0\x6d\xe5\xa1\xfe\x72\xed\xe3\x22\x94\xe3\x37\x21\x42\xc4\x69\xc8\x65\x99\x4f\xcc\xa6\x96\xae\x1a\x1a\xa8\x58\x71\x1a\x42\xe3\x5e\xe6\x90\x32\x4f\xd7\x7e\xaf\x0f\xc6\xc3\xe1\xf4\xe2\xd4\xbe\x1f\x7e\xb9\x75\xce\x74\xdd\xd7\x8e\xdd\x8d\xa5\x3c\x79\xb5\xe5\xf0\xf1\x66\x34\x39\x77\xed\xd1\xe4\x7c\xbb\xc6\xdd\x70\x13\x51\xfd\xa8\x22\x85\x74\xbd\xc3\xeb\x60\x7a\xef\x4d\xee\x5d\xbb\xfc\x61\x75\x2a\x7f\x37\xbc\x7e\x78\x32\x00\xc8\x1c\xfc\xfa\xe1\xc9\x55\x26\x2a\x5c\xc2\xc9\x41\x26\xa6\x24\x4a\x72\x6d\xf5\x53\xcd\xae\x37\xc7\x49\x32\xc3\xfe\x42\x35\xf0\xcf\x9c\xb0\xb7\xaa\x83\x27\x8d\xdb\x23\xbb\xa3\x8f\xc0\x05\x24\x58\x4e\xb6\x66\x53\x15\xaf\xdd\xbe\x0b\xe3\xed\xb6\xe6\x67\x76\x62\x02\xe6\x26\xf7\xce\x50\x54\x59\x6f\x94\xa4\xb3\x4d\x1b\x2e\x64\x87\xd0\xb6\x31\xf7\x6e\x1e\x26\x83\xeb\xc7\x5b\x6f\x34\x1e\x7e\x19\x4d\xb7\xa6\x59\xa2\x08\x1d\xd7\x76\x27\x50\x62\x0a\xb5\xae\xb8\x5c\x04\x31\x83\xd3\x6c\xdb\xd0\x76\x74\x14\x4a\xac\x2e\x71\x8d\x09\x4e\x34\x57\xd5\x53\xfa\x9f\xb8\xea\x5f\x91\x54\x87\x21\x4c\x24\xd5\x98\xf0\x9b\x4a\x7b\xd7\x8c\xcb\x59\xec\x8d\xa7\xa1\xbf\x06\xc0\x5b\x06\xbc\x6c\x07\x59\x03\x08\xdd\x80\x6a\x72\xd5\xbd\xf2\x94\x93\x06\xf5\xf4\xa9\xc3\x3c\xf6\xa6\x81\x34\x92\x4a\x87\xd5\x1a\x7d\x95\x7b\xd9\xd0\xd4\xf3\x7e\x33\x0f\xea\xf8\x0f\xda\x50\x6d\x00\xc8\x4b\xb6\x64\x2d\x04\x34\x56\x15\xb4\x77\x75\xfd\xee\x68\x20\xa1\x33\xfd\x46\xf5\xab\xf6\x48\x99\xb7\x19\x7c\xfc\x08\x47\xf0\xae\x36\x3b\x17\xb0\x82\x90\x91\x0c\xd0\xf7\x5a\xeb\x0f\x1b\xc1\xcf\xe0\x04\xa4\x70\xd2\x3c\x49\x2e\x61\x7d\xc8\x0a\xac\xd7\x6a\x2e\x68\x7f\xb7\x56\x08\xa7\xa9\xc7\x28\x15\xae\x3d\xfa\xed\x46\x0a\x5e\xa2\x38\x21\x92\x3f\x3b\x2a\x08\x8e\x5c\x40\xb2\x5b\x01\xed\x34\x74\xde\x57\x6f\x6c\x37\x73\x87\x53\xf2\xe2\x69\x15\x2f\xc5\x4b\xe2\xda\xc7\x3e\x16\xbb\x5d\xb5\x9f\x54\x8a\xe9\x25\xb3\xb7\xab\xba\x5a\xc9\x14\xfa\x31\x54\x16\xfb\xd0\x07\xbd\x07\x98\xc1\x53\xcb\xa2\x81\x9d\x1a\x5b\x26\x64\x7d\x00\x11\xc5\x1c\xb8\x60\x71\xc6\x25\x6c\x13\xcc\x05\x64\x58\x44\x25\x8c\x32\x9a\x92\x54\xc0\x9c\xd1\xa5\x14\x96\xe7\x67\xdd\x17\x7f\xdd\xb5\x63\xc5\x58\xfd\x5e\xad\x80\x93\x00\x3e\xf1\x95\xf3\xfc\xdd\xf9\xf6\xd9\x5e\xad\x3e\xa9\xd2\x05\x34\x25\x7a\x76\x37\x2f\x59\x8d\x0a\x9a\x04\x75\x96\x8d\xd5\x74\xf8\x73\xac\x43\xf5\x7d\x76\xfd\x9a\x05\x8a\x58\x75\x1c\xb5\xdc\x29\x84\x44\xc0\x69\x0e\x61\x2c\xa2\x7c\x76\xe6\xd3\xa5\x73\x47\xc4\x44\x30\x82\x97\x4e\x11\x36\x5c\x8c\x49\x42\x71\x50\xfa\x28\x2b\xc4\x23\x92\x24\x90\x61\x26\x80\xce\xff\xb1\x57\x7d\x67\x26\x9d\x76\x47\xb5\x55\x90\xcd\xb8\x3e\xb7\x2b\x20\xf1\xd4\xd5\x33\xed\x30\x34\x6e\xed\x95\xb6\x89\x65\x86\x9d\x41\x49\xcd\x76\xe7\xbf\x99\xc3\xd8\xb7\x49\xaa\xbf\xd3\x1c\x78\x44\xf3\x24\x00\x9e\x11\x3f\x9e\xbf\x01\xd6\x3c\x26\x22\x2c\x2a\xe9\x8c\x00\x79\x25\x7e\x2e\x48\xd0\xdb\xab\xf5\x05\x76\xef\xd5\xad\xb8\xfa\x31\xc0\xd2\x5c\xab\x82\x7b\x3e\x0d\x88\x6b\xff\xf2\x3f\x35\xf3\x90\x0e\x35\x2a\x61\x37\x2e\x25\x59\xd9\x40\xcb\xd6\xdf\x01\x00\x00\xff\xff\xe6\x47\x88\xda\x4c\x10\x00\x00")

func dataShBytes() ([]byte, error) {
	return bindataRead(
		_dataSh,
		"data/sh",
	)
}

func dataSh() (*asset, error) {
	bytes, err := dataShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataZsh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x5d\x8b\xdb\x30\x10\x7c\xf7\xaf\xd8\xea\x0e\x9a\x3c\x84\xfb\x05\x85\x96\xf6\xa0\x85\xbe\x95\x96\x42\x29\x61\x23\xad\x6d\x11\x45\x52\xb5\x2b\xe7\x72\xc7\xfd\xf7\x62\xe5\x53\xae\xa1\x6f\x61\xb4\x9a\x99\xcd\x8c\x75\xf7\xe6\x61\x63\xfd\xc3\x33\xf7\x4d\xb3\x5e\x0f\xdd\x3a\x91\x0b\x68\x16\x4b\x78\x69\x00\x00\x68\x40\x07\xea\x7e\x31\x74\xc7\x9f\xab\x15\xf7\xe4\x1c\x3c\x73\xbf\x54\xcd\x6b\xd3\xe8\x3e\xee\xcd\xba\xcd\x5e\x8b\x0d\x9e\xdf\x2d\xee\x5f\x26\xd0\xaf\xf7\xbf\x5f\x41\x15\x72\xcc\x12\xd6\xa8\xc5\x0e\x28\xa4\x96\x27\x49\x1d\x76\xd1\x51\x99\xbd\xe8\x8e\x18\x1a\x03\xe7\xe1\x0a\xd4\x26\xa2\xde\x62\x47\x5c\xc1\x86\x66\xa7\x0d\xb1\xa4\x70\xa8\x30\xf2\x9c\x53\x3d\x36\xae\x57\x01\x9d\x0b\x1b\x74\x8f\x4f\xa4\x2b\xb8\x27\x17\x2b\xc0\x7a\x2b\x05\xb8\xab\xa0\x6f\x24\x62\x7d\x57\x5b\x74\xd6\x6f\x27\x00\x4b\x0d\x04\x8d\xee\x8b\x67\x41\x57\xdb\xd9\x85\x81\x7e\x90\x37\x21\x4d\xc4\xfe\x64\x4a\x87\x59\x35\x26\xc9\xb5\x57\x16\x94\x5c\x0f\x65\x6f\x67\xd4\xb2\xff\xc7\x6a\x8e\x5d\x42\x53\xff\x69\x03\x25\xb6\xc1\x4f\x1c\xad\x7e\x82\xfa\x70\x8a\x02\x10\x38\x92\xb6\xad\xd5\x30\xd8\x24\x19\x5d\x17\x60\x1f\xd2\x96\x23\x6a\x52\x75\xc2\x35\xc7\xa7\x4b\xa0\x20\x3d\x81\xce\x29\x91\x97\x79\x9a\x49\xf8\x13\x33\xb0\x4f\x18\x23\x25\x68\x43\x02\x43\x11\xa4\x47\x81\xd3\xe2\x5c\xd8\x0d\x45\xf2\x86\xbc\xb6\xc4\x60\x7d\xc1\x66\x94\xca\x25\x42\x03\xa1\x85\xa1\xc4\xa1\x6e\xdb\x54\xeb\x3e\x3e\xc5\xc0\x47\xf3\xac\x93\x8d\x72\xd4\xe5\x3e\x64\x67\x60\x43\xa5\x74\x2b\x32\x67\xbd\xf3\x86\xe5\x1b\x53\xd7\x4a\xd6\xac\x9f\xc9\x45\xc0\x4d\xc8\x02\xe8\x0f\xe3\xd1\x0e\xbd\x51\xd7\x66\xd6\xe3\x1f\x13\x95\x1c\xfc\x58\x7b\xdc\xb8\x31\x92\xeb\x3a\x23\x3c\x46\x0d\x56\x40\x42\xe5\xc2\xd8\x44\x5a\x42\x3a\xa8\xb9\x8e\x8f\xd4\x5f\xc7\x8b\xff\x89\xe6\xc8\x6a\xf9\x96\xee\xd2\xad\x9a\xee\x7b\x29\x1d\xcf\x9b\x80\x36\x85\x5d\x39\xba\xc9\x42\xde\x72\x21\x23\x03\x12\xd4\xb9\xb5\xe5\x5d\x0a\xbb\x68\xa8\x85\xe9\x0b\x03\x43\xd7\xfc\x0d\x00\x00\xff\xff\x36\x38\x3a\x01\xf5\x04\x00\x00")

func dataZshBytes() ([]byte, error) {
	return bindataRead(
		_dataZsh,
		"data/zsh",
	)
}

func dataZsh() (*asset, error) {
	bytes, err := dataZshBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/zsh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/bash": dataBash,
	"data/fish": dataFish,
	"data/sh": dataSh,
	"data/zsh": dataZsh,
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
	"data": &bintree{nil, map[string]*bintree{
		"bash": &bintree{dataBash, map[string]*bintree{}},
		"fish": &bintree{dataFish, map[string]*bintree{}},
		"sh": &bintree{dataSh, map[string]*bintree{}},
		"zsh": &bintree{dataZsh, map[string]*bintree{}},
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

