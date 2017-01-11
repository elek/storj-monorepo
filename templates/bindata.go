// Code generated by go-bindata.
// sources:
// golang.create.tmpl
// golang.delete-all.tmpl
// golang.delete-world.tmpl
// golang.delete.tmpl
// golang.footer.tmpl
// golang.get-all.tmpl
// golang.get-count.tmpl
// golang.get-has.tmpl
// golang.get-last.tmpl
// golang.get-limit-offset.tmpl
// golang.get-paged.tmpl
// golang.get.tmpl
// golang.header.tmpl
// golang.misc.tmpl
// golang.update.tmpl
// DO NOT EDIT!

package templates

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

var _golangCreateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\x41\x6b\xdc\x30\x10\x85\xcf\xd6\xaf\x98\x2e\x14\x6c\xd8\xe8\x07\x04\xf6\x50\x42\x0a\x81\xb0\x90\x6c\x4a\x8f\x46\x59\x8d\x8c\x5a\xaf\xe4\x8e\xc6\xf5\x2e\x42\xff\xbd\x8c\xdd\x0d\xce\xa1\x3d\xe4\x60\xb0\xc7\xe2\x7d\xef\x3d\x4d\xce\x37\x60\xd1\xf9\x80\xb0\x49\xbe\x0b\x86\x47\xc2\x0d\xdc\x94\xa2\xee\x08\x0d\x63\xce\xa0\x0f\xa3\x73\xfe\x0c\xa5\xd4\x39\xc3\x91\xcf\x83\x21\x73\x02\xfd\x85\xba\x04\xa5\x34\x20\xe3\xbf\xb3\x67\xe4\x91\x02\x94\xb2\x05\x24\x92\x27\x52\xa3\x84\x82\xc1\xce\xb2\x6a\x8d\x7c\x8d\xf6\x32\xd3\xb2\xaa\x64\xee\x1d\xe8\x3d\xa2\x4d\xfb\x38\x41\x29\xaa\x6a\xdb\x10\x27\xb8\xdd\xc1\x3e\x4e\x75\xa3\xbf\xbd\xdc\xd5\x8d\x1c\x7d\x93\x93\x77\x1f\x3c\x07\x9c\x40\x7f\xf5\xd8\x5b\xf1\xa4\x54\x75\x8c\x21\x31\xb4\x6d\xe2\x13\xef\xc4\x20\xf9\xc0\x0e\x36\x9f\x7f\x6d\x40\x1f\x9e\x1e\x67\xf9\xf8\xfa\x43\xf7\xb1\x3b\xf0\x89\xeb\xe5\xe8\x16\x72\x06\x43\xdd\x4a\xac\x51\x0b\xc5\x49\x13\xc3\x10\x89\xd3\x12\xd3\x87\x0e\x56\x16\x56\xe9\x55\x25\xe9\x77\x20\xfa\x96\xfc\x6f\x24\xfd\x34\x22\x5d\x9e\xe3\xf4\x1f\x8e\x3e\x1c\x4d\x90\x32\x8d\xb5\x14\x1d\xd4\xae\x37\xcc\x18\xae\xc2\xcd\x6c\xa6\xf2\x6e\xee\xf6\xd3\x0e\x82\xef\x21\xab\xaa\xa2\x85\x1b\x7c\xbf\x85\xef\x64\x86\x7b\xa2\x1a\x89\x1a\x55\x15\x75\xfd\x79\xc5\xad\x6e\x28\xf8\x7e\xe9\xb2\x4f\xb8\x94\xd9\xb6\x84\x69\xb9\xba\xdb\x77\xee\xef\xcf\x78\xfc\x87\xf3\x8f\x78\x6a\xdb\xe1\xe7\x1b\x66\x86\xea\x47\x93\xf8\x21\x24\x24\x7e\xb0\xf5\x87\x53\x8a\xe5\x0e\x59\xc4\x64\x73\x97\xb0\xfa\xe5\x32\xa0\xac\xaf\x60\xdf\xaf\x4f\x51\xab\x8f\x3f\x01\x00\x00\xff\xff\x8d\x11\x59\x35\x0d\x03\x00\x00")

func golangCreateTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangCreateTmpl,
		"golang.create.tmpl",
	)
}

func golangCreateTmpl() (*asset, error) {
	bytes, err := golangCreateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.create.tmpl", size: 781, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteAllTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\xb1\x6a\xeb\x40\x10\x45\xeb\xdd\xaf\xb8\x4f\xf0\x40\x06\x5b\xa4\x08\x29\x02\x2a\x0c\x71\x97\x26\x71\x91\x52\xc8\xd2\xac\xd9\x20\xef\x3a\xa3\x59\xc7\x61\xd8\x7f\x0f\x92\x4d\x92\x2e\xa4\x98\x6a\xe6\x9e\x73\x19\xd5\x15\x7a\x72\x3e\x10\x8a\xd1\xef\x43\x2b\x89\xa9\xc0\x2a\x67\xfb\x40\x03\x09\xa9\xa2\xda\x26\xe7\xfc\x19\x39\x97\xaa\xe8\xe4\x7c\x6c\xb9\x3d\xa0\x5a\xf3\x7e\x44\xce\x0b\x94\x5d\x4c\x41\xe0\x83\xdc\xdd\x2e\x41\xcc\xd3\x44\x5e\xd8\x89\x4e\xa1\x9f\x71\xf6\xa7\x6a\x17\xfb\x8f\xd9\xa2\xd6\x74\x31\x8c\x82\xa6\x19\xe5\x20\xb5\x2a\x8e\xec\x83\x38\x14\xff\xdf\x0a\x54\xdb\xa7\x47\xe4\x6c\x4d\xdc\xbd\x56\x43\xdc\x6f\xe5\x20\xe5\xe5\x74\x09\x55\x38\x4f\x43\x7f\x6a\x87\x44\xdf\x75\xac\x35\x4d\xc3\x34\x5e\x9a\xdc\xd7\x98\xb2\x3d\xfb\x13\x71\xb5\x39\x53\xf7\x6b\xde\x78\x37\x47\xff\xd5\x08\x7e\x80\x5a\x63\x98\x24\x71\xc0\xcd\x12\x2f\xdc\x1e\x37\xcc\x25\x31\x2f\xac\xc9\x76\xea\x9f\x82\x5c\x64\x35\x66\x73\xf5\x1c\xdf\xc7\xb5\x73\xd4\x09\xf5\xe5\x9f\x81\xd7\xdd\x95\x1b\xfc\x60\xb3\x55\xfd\x7a\xe4\x67\x00\x00\x00\xff\xff\x86\x5c\x4a\x69\xb3\x01\x00\x00")

func golangDeleteAllTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteAllTmpl,
		"golang.delete-all.tmpl",
	)
}

func golangDeleteAllTmpl() (*asset, error) {
	bytes, err := golangDeleteAllTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete-all.tmpl", size: 435, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteWorldTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\xb1\x6a\xc3\x30\x10\x86\x67\xe9\x29\xae\x86\x82\x4d\x13\xd1\xa1\x74\xf3\x10\xda\x6c\x5d\x9a\x0e\x1d\x83\x63\x9f\x83\x8a\x38\x25\xa7\x73\xea\x22\xf4\xee\xc5\x72\x1d\xb2\x76\x10\x02\xa1\xfb\xbe\xfb\xff\x18\xd7\xd0\x61\x6f\x09\xa1\x08\xf6\x48\x8d\x0c\x8c\x05\xac\x53\xd2\xaf\xe8\x50\x70\xe3\x5c\xd9\xca\x08\xad\x27\xc1\x51\xcc\xcb\x7c\x57\x50\xb6\x7e\x20\x01\x4b\xf2\xfc\xb4\x02\x64\x9e\x8e\xe7\x4a\x4f\x48\xa4\x2e\x33\xf4\x2d\xff\xe0\xbb\x9f\x8c\x8e\x5a\x5d\x1a\x86\xfd\x9e\x31\x40\x38\x3b\xb3\xc3\x30\x38\x59\x5e\x6f\xc0\x33\x80\x1b\x3a\x22\x98\x8f\xf7\xb7\x00\x29\x69\x95\x07\x67\x67\x0d\xfe\xf0\x65\x3a\xb6\x17\x64\xb3\x1d\xb1\x2d\x63\x84\x13\x5b\x92\x1e\x8a\xfb\x73\x01\x06\x52\xaa\xb4\xb2\x7d\xfe\x7e\x57\x03\x59\x07\x51\x2b\xc5\x28\x03\x13\x3c\xae\xe0\x93\x9b\xd3\x96\xb9\x44\xe6\x4a\xab\xa4\x27\x41\xde\x61\x51\x64\x9f\xd9\xf9\xef\xb0\xe9\x7b\x6c\x05\xbb\xf2\x9f\x48\x35\x67\x7a\xa8\x97\x78\xd7\x96\xa6\x92\x96\xc1\x3f\x29\x59\xa7\x93\x8e\xf1\xda\xe2\x6f\x00\x00\x00\xff\xff\x4e\xd5\x89\xfb\xa5\x01\x00\x00")

func golangDeleteWorldTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteWorldTmpl,
		"golang.delete-world.tmpl",
	)
}

func golangDeleteWorldTmpl() (*asset, error) {
	bytes, err := golangDeleteWorldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete-world.tmpl", size: 421, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\xb1\x6a\xf3\x30\x14\x85\x67\xe9\x29\xce\x6f\xf8\xc1\x01\xc7\x74\x2e\xa4\x10\x68\xb6\x2e\x6d\x86\x8e\xc6\xb1\xae\x82\x8a\x22\xa5\x57\x72\x9a\x72\xd1\xbb\x17\x3b\x21\xed\x56\xe8\xa0\x49\xe7\x7c\xdf\xe1\x8a\x2c\x61\xc8\xba\x40\xa8\x92\xdb\x87\x3e\x8f\x4c\x15\x96\xa5\xe8\x47\xf2\x94\x49\x04\xed\x76\xb4\xd6\x9d\x51\x4a\x2d\x82\x21\x9f\x8f\x3d\xf7\x07\xb4\x6b\xde\x27\x94\xb2\x40\x6d\xe6\xa8\xc1\x2e\x46\xdf\x80\x98\xa7\x17\x79\xa1\x27\x3c\x05\x33\xf3\xf4\x4f\xd7\x2e\x9a\xcf\x59\x23\x5a\x0d\x31\xa4\x8c\xae\x4b\xf9\x90\x57\x22\x38\xb2\x0b\xd9\xa2\xfa\xff\x5e\xa1\xdd\x3e\x3f\xa1\x14\xad\xe2\xee\xad\xf5\x71\xbf\xcd\x87\x5c\x5f\xa2\x0d\x44\x60\x1d\x79\x73\xea\xfd\x48\xdf\x7b\xb4\x56\x5d\xc7\x94\x2e\x4b\xee\x57\x98\xba\x86\xdd\x89\xb8\xdd\x9c\x69\xf8\xb5\xaf\x9c\x9d\xab\xff\x56\x08\xce\x43\xb4\x52\x4c\x79\xe4\x00\xdb\xfb\x44\x0d\x5e\xb9\x3f\x6e\x98\x6b\x62\x5e\x68\x55\x66\xe1\x10\xc7\x90\x6f\xca\x79\x40\xfb\x12\x3f\xd2\xda\x5a\x1a\x32\x99\xfa\x4f\xdc\xeb\xff\x15\x8f\x07\xdc\x35\x53\x57\x17\x2d\x72\xbb\xec\x57\x00\x00\x00\xff\xff\x62\xe0\x75\x98\xc5\x01\x00\x00")

func golangDeleteTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteTmpl,
		"golang.delete.tmpl",
	)
}

func golangDeleteTmpl() (*asset, error) {
	bytes, err := golangDeleteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete.tmpl", size: 453, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xce\xb1\x0a\xc2\x30\x10\x06\xe0\x39\xf7\x14\x37\xb6\x42\xf3\x10\x22\x6e\x4e\x66\x70\x8d\xc9\xd9\x16\x34\xd6\xe4\xc4\x48\xb8\x77\x97\x22\xa4\x4b\xb6\x83\xfb\xff\x8f\x1f\xf8\xbb\x10\x9a\xcb\x89\x78\x7a\xfa\x84\x73\x60\x8a\x37\xeb\x08\x0b\x94\x32\x60\xb4\x61\x24\xd4\xc7\x77\x70\x09\x45\x40\x95\x82\x7a\x3d\xd6\x27\x05\x8f\x83\x08\x08\xfc\x19\xce\x2d\x46\x55\xbd\x06\x0f\xfb\x66\xf0\xec\x26\x7a\xd8\xae\xc7\xc4\x71\x0e\x63\xab\xe9\xaf\xcd\x66\x05\x01\xd4\x27\xda\xc5\xe4\x8e\x33\xee\xd2\xeb\xae\x4d\xee\xb7\x61\x20\xf0\x0b\x00\x00\xff\xff\x6d\x27\x2d\x7a\xf2\x00\x00\x00")

func golangFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangFooterTmpl,
		"golang.footer.tmpl",
	)
}

func golangFooterTmpl() (*asset, error) {
	bytes, err := golangFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.footer.tmpl", size: 242, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetAllTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\x4d\xeb\xdc\x20\x10\xc6\xcf\xfa\x29\xa6\x81\x82\x81\xfd\xe7\x03\x6c\xc9\xa1\x94\xd2\x4b\x29\x6c\x73\xe8\x31\xb8\x71\x0c\x16\x57\xd3\x89\xd9\x17\x06\xbf\x7b\x31\xd9\xdd\xf6\x50\x16\x7a\xf0\xa0\xce\x3c\xcf\x6f\x1e\x65\x7e\x03\x83\xd6\x05\x84\x6a\x76\x63\xd0\x69\x21\xac\xe0\x2d\x67\xf9\x05\x13\x33\x34\xdd\x62\xad\xbb\x42\xce\x8a\x19\x86\x74\x9d\x34\xe9\x13\x34\x1f\x69\x9c\x21\xe7\x1a\x14\xc5\xcb\x0c\xcc\x30\x7b\x37\x60\xb4\xd0\x7c\x8f\x17\xc8\x79\x07\x48\x54\x56\xa4\x5a\x16\x1b\x0c\x66\xd5\x95\x7f\x7b\x1e\xa3\xb9\x55\x90\x33\x4b\x31\xc4\x30\x27\xe8\xfb\x39\x9d\x52\xcb\x0c\x13\xb9\x90\x2c\x54\xef\x7f\x55\xd0\x74\x87\xaf\x90\xb3\x14\xf1\xf8\xb3\xf1\x71\xec\xd2\x29\xa9\xad\x74\x57\xbc\xad\x43\x6f\xce\xda\x2f\xf8\x87\x4c\x4a\xd1\xf7\x05\x6e\x23\xd9\xb7\x50\x9a\x0d\xb9\x33\x52\x73\x58\x90\x6e\xaf\x14\x8a\x80\x70\x76\x6d\x7d\xd7\x42\x70\x1e\x58\x0a\x41\x98\x16\x0a\x65\xbb\x83\x1f\xa4\xa7\xcf\x44\x0a\x89\x6a\x29\xb2\x14\x06\x2d\x12\x6c\xa6\xcd\x27\x1f\x67\x54\x85\xc2\xc6\xe7\xe1\x37\xbc\x26\x55\xaf\x4a\xcc\xe0\x82\x4b\x01\x2f\x8f\xc4\xa4\x10\xc5\xad\x7d\x14\x77\x83\x0e\x25\x74\x6d\x0c\x45\x0b\xca\x7a\x9d\x12\x86\xb5\xbc\x5e\x27\x14\xff\x20\x7c\x81\x58\x18\xc5\xfa\x5c\x2d\xe8\x69\xc2\x60\xd4\x96\x4f\x31\xa1\xf1\xc1\xb1\x0d\x73\x57\xde\x3f\x71\x8a\x4e\xfd\xe1\xbf\x02\xb9\x5f\x6e\x26\xc1\x79\x99\x25\xf3\xf3\x23\xfc\x0e\x00\x00\xff\xff\x88\x0b\x9a\x08\x7c\x02\x00\x00")

func golangGetAllTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetAllTmpl,
		"golang.get-all.tmpl",
	)
}

func golangGetAllTmpl() (*asset, error) {
	bytes, err := golangGetAllTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-all.tmpl", size: 636, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetCountTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8f\x3f\x4f\xc3\x30\x10\xc5\x67\xdf\xa7\x78\x44\x02\x25\x52\x1a\x31\x20\xb6\x0e\x08\xb1\xb1\x94\x0c\x8c\x55\x9a\x38\x95\x51\x6a\x97\x8b\x5d\x5a\x9d\xee\xbb\x23\x07\xf1\x67\x63\xf0\xe4\xf7\x7b\xf7\x7e\x22\x2b\x0c\x76\x74\xde\xa2\x98\xdd\xde\x77\x31\xb1\x2d\xb0\x52\xa5\xc7\x90\x7c\x14\x41\xd3\xa6\x71\x74\x67\xa8\x96\x22\xe8\xe3\xf9\xd8\x71\x77\x40\xf3\xc0\xfb\x19\xaa\x15\xca\x3e\x27\xe1\x7c\xbc\xbf\xab\x61\x99\xf3\x0b\x5c\x51\x2e\xb7\x7e\x58\xda\xe8\xef\xa5\x5d\x18\x2e\x05\x54\x85\x4c\x1f\xfc\x1c\xb1\xdd\xce\xf1\x10\xd7\x22\x38\xb2\xf3\x71\x44\x71\xfd\x5e\xa0\x69\x37\xcf\x50\x25\x13\x76\x6f\xcd\x14\xf6\x6d\x3c\xc4\xf2\x2b\x5a\x43\x04\xa3\xb3\xd3\x70\xea\xa6\x64\x7f\xd7\x10\x99\xbc\x60\x8d\xcc\x0c\xec\x4e\x96\x9b\x4d\xb2\x7c\x79\x09\x1f\xff\xb1\x4d\xdb\x77\xbe\xbc\x59\x74\x2a\x32\x6e\x5c\x64\xae\xd6\xf0\x6e\x82\x90\x31\x6c\x63\x62\x8f\xdb\x1a\xaf\xdc\x1d\x9f\x98\x4b\xcb\x5c\x91\x51\xa2\xef\xbf\x05\xae\x33\x41\x4a\x22\x3f\xfe\x9f\x01\x00\x00\xff\xff\x26\xd9\x70\xe1\x69\x01\x00\x00")

func golangGetCountTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetCountTmpl,
		"golang.get-count.tmpl",
	)
}

func golangGetCountTmpl() (*asset, error) {
	bytes, err := golangGetCountTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-count.tmpl", size: 361, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetHasTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8f\xbd\x6e\xeb\x30\x0c\x85\x67\xf1\x29\xce\x35\x70\x0b\x1b\x70\xfc\x06\x19\x3a\x14\xe8\xd0\x25\xf5\xd0\x31\x50\x62\x29\x51\xe1\x48\x2e\x25\xa7\x09\x08\xbe\x7b\xe1\x14\xfd\xd9\x3a\x70\x22\xbf\xc3\xf3\x89\xac\x30\x38\x1f\xa2\x43\x95\xc3\x21\xda\x32\xb3\xab\xb0\x52\xa5\x47\x9b\x45\xd0\xf5\xb3\xf7\xe1\x02\xd5\x5a\x04\xfb\x72\x99\x2c\xdb\x13\xba\x7b\x3e\x64\xa8\x36\xa8\x8f\x36\x63\x97\xd2\xd8\xc2\x31\x2f\x93\xb8\xa1\x25\xd7\xc5\xe1\x16\x44\xbf\x9f\xec\xd2\x70\xad\xa0\x2a\x64\xf6\x29\xe6\x82\xed\x36\x97\x53\x59\x8b\x60\xe2\x10\x8b\x47\xf5\xff\xad\x42\xd7\x6f\x9e\xa0\x4a\x26\xed\x5e\xbb\x31\x1d\xfa\x72\x2a\xf5\xe7\x69\x0b\x11\xf8\xe0\xc6\xe1\x6c\xc7\xd9\xfd\x54\x21\x32\x4b\x83\x35\x16\x66\xe0\x70\x76\xdc\x6d\x66\xc7\xd7\xe7\xf4\xfe\x17\xdb\xf5\x7b\x1b\xeb\xbb\xa3\xcd\x0d\x99\xe0\x6f\x2a\xff\xd6\x88\x61\x84\x90\x31\xec\xca\xcc\x11\xde\x8e\xd9\xb5\x78\x61\x3b\x3d\x30\xd7\x8e\xb9\x21\xa3\xf4\xb5\x3e\xda\xdc\x2e\x08\x29\x89\x7c\xeb\x7f\x04\x00\x00\xff\xff\xa2\xc5\x75\x44\x63\x01\x00\x00")

func golangGetHasTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetHasTmpl,
		"golang.get-has.tmpl",
	)
}

func golangGetHasTmpl() (*asset, error) {
	bytes, err := golangGetHasTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-has.tmpl", size: 355, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetLastTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x90\xbf\x6a\xc3\x30\x10\xc6\x67\xdd\x53\x5c\x0d\x05\x1b\x12\x4d\xa5\x5b\xa6\xd2\x2d\x4b\x92\x42\xc7\xa0\xd8\x67\xa3\xd6\x91\xdc\xf3\xa5\x4d\x7a\xe8\xdd\x8b\x62\x52\x9c\x41\x68\xf9\xfe\xfc\xbe\x53\x5d\x62\x43\xad\x0f\x84\xc5\xe8\xbb\xe0\xe4\xc4\x54\xe0\x32\x25\xe8\x48\xd6\x6e\x14\x55\xb4\x5b\x92\x13\x07\xfb\x76\x19\x08\x53\x2a\x6b\x39\x63\x1d\x83\xd0\x59\xec\xcb\xf4\x2f\x70\xf8\x44\x1f\xe4\xf9\xa9\xc2\x52\x15\x07\xc7\xee\x78\x33\x62\x4a\x0b\x24\xe6\xfc\x22\x57\x90\x4b\x29\x34\xd7\x16\x98\x13\x1c\x62\x73\xb9\x96\x2b\x98\x3a\x86\x51\x70\xbf\x1f\xe5\x28\xab\x9c\xc8\x3e\x48\x8b\xc5\xe3\x57\x81\x76\xb7\x59\x63\x4a\x60\xe2\xe1\xc3\xf6\xb1\xdb\xc9\x51\xca\x49\x9a\x41\x2a\x00\xa3\x8a\x3e\x78\x99\x21\x80\xc9\x08\x2b\xcc\x9e\x86\xfd\x37\xb1\xdd\x9c\x88\x2f\xdb\xf8\x33\xf7\xda\x5d\xed\x42\x9e\xe0\x9a\x86\x63\x8b\x65\xdb\x3b\x11\x0a\xb7\xa4\x0a\x53\xaa\xc0\xf8\xf6\xba\xe8\x61\x85\xc1\xf7\xa8\x60\x0c\x4f\x45\xaa\xf8\x4b\x1c\xef\xb6\xbf\xb3\x1b\x5e\x99\x4b\x62\xae\xc0\x24\x98\x69\x1d\x77\x77\xd2\xe0\x7b\x48\xa0\xfa\x7f\xa0\xbf\x00\x00\x00\xff\xff\x4d\xc0\x11\x01\xa2\x01\x00\x00")

func golangGetLastTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetLastTmpl,
		"golang.get-last.tmpl",
	)
}

func golangGetLastTmpl() (*asset, error) {
	bytes, err := golangGetLastTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-last.tmpl", size: 418, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetLimitOffsetTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\x41\x8b\xd4\x30\x14\xc7\xcf\xc9\xa7\xf8\x5b\x10\x5a\x98\xed\x49\x3c\xac\xf4\x20\x22\x5e\x44\x58\xe7\xe0\x71\xc8\x36\x2f\x43\x24\x4d\xea\x4b\xba\x33\x4b\xcc\x77\x97\xb4\x3b\x83\xe2\x20\x78\x28\xb4\xcd\xcb\xef\xff\xcb\x7b\xc9\xf9\x0e\x9a\x8c\xf5\x84\x26\xda\xa3\x57\x69\x61\x6a\x70\x57\x8a\xfc\x44\x29\x67\xf4\xfb\xc5\x18\x7b\x46\x29\x6d\xce\x18\xd3\x79\x56\xac\x26\xf4\xef\xf9\x18\x4b\xd9\xc1\xd9\xc9\x26\x58\x9f\x76\x08\xc6\x44\x5a\xdf\xdf\xbe\xe9\xd0\x72\x38\x45\xe4\x8c\xe8\xec\x48\xc1\xa0\xff\x1a\x4e\xa8\x5b\x88\xb9\x3e\x81\x65\x4d\x27\xaf\xd7\x38\xf9\xbb\xca\x63\xd0\xcf\x0d\x4a\xc9\x52\x8c\xc1\xc7\x84\xc3\x21\xa6\x29\x0d\x39\x63\x66\xeb\x93\x41\xf3\xfa\x47\x83\x7e\xff\xf0\x19\xa5\x48\x11\x1e\xbf\xf7\x2e\x1c\xf7\x69\x4a\xed\x56\xba\xab\xd1\xc6\x92\xd3\x4f\xca\x2d\xb4\x09\xe3\x27\xc6\x30\x4d\x0a\xa5\xac\xde\x17\xe7\x4e\x4a\x71\x38\x54\xe1\xcd\xee\x7e\x40\x25\x6a\xb6\x4f\xc4\xfd\xc3\x42\xfc\x7c\x1b\xfb\x07\xf5\x2f\xa8\xb0\x66\xc5\xbd\x1a\xe0\xad\x43\x96\x42\x30\xa5\x85\x7d\xfd\xdc\xe1\x1b\xab\xf9\x23\x73\x4b\xcc\x9d\x14\x45\x0a\x4d\x86\x18\x9b\x48\xff\xc1\x85\x48\x6d\x35\x33\xe1\xfa\xf3\x0b\x9d\x53\xdb\xad\xa4\x9c\x61\xbd\x4d\x9e\x4e\x97\xce\x4a\x21\x6a\xda\x70\x29\xde\x8f\xca\xd7\xa1\x29\xad\x39\x18\xb4\xc6\xa9\x94\xc8\xaf\xe5\x1d\x4a\xe9\xa4\xb8\x61\xf8\x0f\xc5\xea\x28\xd6\xb1\x0e\x50\xf3\x4c\x5e\xb7\x5b\xcf\x6a\x08\x1f\x2f\x1e\xdb\x61\x5e\xc8\xf7\x57\x9d\xca\xe9\xde\xfd\x57\x43\x5e\x16\xb7\x10\x6f\x9d\x2c\x32\xe7\xeb\x8d\xf9\x15\x00\x00\xff\xff\x1f\xe9\x76\x5c\xbc\x02\x00\x00")

func golangGetLimitOffsetTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetLimitOffsetTmpl,
		"golang.get-limit-offset.tmpl",
	)
}

func golangGetLimitOffsetTmpl() (*asset, error) {
	bytes, err := golangGetLimitOffsetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-limit-offset.tmpl", size: 700, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetPagedTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\x41\xab\xd4\x30\x14\x85\xd7\xc9\xaf\xb8\x06\x84\x16\x3a\xe5\xb9\x7d\x52\x41\x44\xdc\x88\xfa\x9c\x85\x0b\x91\x21\xaf\xbd\x29\xd1\x36\x19\x6f\x6f\xdf\xcc\x23\xe6\xbf\x4b\x92\x99\x71\x40\x41\x5d\x0c\x4c\xd3\x93\x73\xbf\x93\x9c\x86\xb0\x81\x01\x8d\x75\x08\x6a\xb1\xa3\xd3\xbc\x12\x2a\xd8\xc4\x28\xdf\x20\x87\x00\xed\x76\x35\xc6\x1e\x21\xc6\x2a\x04\xe8\xf9\xb8\xd7\xa4\x67\x68\x5f\xd2\xb8\x40\x8c\x0d\xf4\xec\xbf\xa1\x83\x85\xc9\xba\xb1\x81\xc9\xce\x96\xc1\x3a\xae\xa1\x22\x7f\x58\x20\x04\x58\x26\xdb\xa3\x37\xd0\x7e\xf4\x87\xab\x3d\x7e\xe5\xcb\x36\x24\x4a\x3f\x4f\xb5\x4c\x48\xe8\x86\xcc\x20\xaf\xf9\xee\xfd\xf0\xa8\x20\xc6\x20\x85\x35\xe7\xb9\x5d\x07\x4a\x41\x90\x42\x9c\x17\x40\xdd\x28\x29\xa2\x94\xa2\xf7\x6e\x61\xd8\xed\x16\x9e\xb9\x0b\x01\xf6\x64\x1d\x1b\x50\x4f\xbf\x2b\x68\xb7\x77\x6f\x21\x46\x29\xfc\xfd\xd7\x76\xf2\xe3\x96\x67\xae\x8a\xb4\x49\xcc\xc6\xe2\x34\x3c\xe8\x69\xc5\x53\xd6\x1f\xd0\xfb\x79\xd6\x10\x63\x19\x74\x8a\x5a\x4b\x29\x76\xbb\x94\xb4\x84\xb8\xed\x20\x39\x0e\x64\x1f\x90\xda\xbb\x15\xe9\xf1\x5f\x6c\x7f\x73\x4d\x11\x93\xdf\x93\x0e\x9c\x9d\x72\x40\x42\x5e\xc9\xa5\xc7\x06\x94\x6a\xe0\x13\xe9\xfd\x6b\xa2\x0a\x89\xea\x14\x58\x0c\x68\x90\xa0\xd0\xb4\xaf\x26\xbf\x60\x95\xf0\x8c\xbf\x2c\xbe\xc3\x23\x57\x75\x76\x0b\x01\xac\xb3\xec\xf0\x70\xbe\x17\x29\x44\x9a\xd8\x9d\xc5\xdb\x5e\xbb\x74\xe9\x7a\x18\xc8\x1b\xa8\xcc\xa4\x99\xd1\x65\x79\x0d\x31\xd6\x52\xfc\x81\xf2\x2f\x98\x89\x53\xe4\x62\x74\xa0\xf7\x7b\x74\x43\x55\x0e\x2f\x0d\xa2\xf1\xcc\x52\x02\x9d\xdc\x6f\x2f\x48\xc9\xa7\x7e\xfe\xbf\x07\x93\x8d\x4a\x31\x5f\xc0\x4d\xde\x92\x16\xd0\xe5\xd1\x75\xea\x50\x79\x9b\xf9\x7f\x95\xb3\x03\x33\x73\xbb\xcd\xb5\xc9\xd2\xcf\x59\xb6\x79\xf6\xa5\x4d\x5f\xc6\x07\x3d\xe2\xf0\xde\x9d\x4e\x22\x4a\x11\x01\xa7\x05\xaf\xca\x58\x4c\xca\xff\x02\x72\x42\x2d\x91\x2f\xa2\x26\xb1\xcb\x28\x43\xb8\x34\xff\x67\x00\x00\x00\xff\xff\x8b\xd7\xe8\x93\x99\x03\x00\x00")

func golangGetPagedTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetPagedTmpl,
		"golang.get-paged.tmpl",
	)
}

func golangGetPagedTmpl() (*asset, error) {
	bytes, err := golangGetPagedTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-paged.tmpl", size: 921, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x90\x3f\x6f\xe3\x30\x0c\xc5\x67\xe9\x53\xf0\x0c\x1c\x60\x03\x89\xbe\x81\x87\x1b\x82\x5b\x0e\x07\x24\x1e\x3a\x06\x4a\x44\x19\x2a\x1c\x29\xa1\xe5\xfc\x29\xc1\xef\x5e\xc8\x6d\x93\x76\x69\x07\x2e\x0f\x7c\x3f\x3e\x3e\xe6\x25\x38\xf4\x21\x22\x54\x63\xe8\xa3\xcd\x13\x61\x05\x4b\x11\xfd\x17\x33\x33\x98\x6e\xf2\x3e\x5c\x41\xa4\x66\x86\x7d\xbe\x1e\x2d\xd9\x03\x98\x3f\xd4\x8f\x20\xd2\x40\x91\xdf\xb5\x4d\xba\x80\xc8\x02\x90\xa8\x4c\xa2\x46\x17\x3e\x46\x37\x03\xf5\xe7\x63\xbb\xe4\x6e\x15\x88\xb0\x56\xfb\x14\xc7\x0c\xdb\xed\x98\x0f\xb9\x2d\x34\x0a\x31\x7b\xa8\x7e\x9f\x2a\x30\xdd\xfa\x1f\x88\x68\x95\x76\xcf\x66\x48\x7d\x97\x0f\xb9\x7e\x5b\x5d\x00\x33\xf8\x80\x83\x3b\xdb\x61\xc2\x47\x24\xad\x15\x33\x84\x18\xf2\x47\x24\xad\x4a\xa4\x16\x0a\xc4\x51\x38\x23\x99\xf5\x84\x74\xdb\xa4\xcb\x4f\x30\xd3\xed\x6d\x2c\x3f\x5a\xe7\x28\x79\xa8\xfd\x60\x73\xc6\x38\xa3\x9b\xf9\x9c\x0a\x7e\x7e\xb9\x6d\x61\x3c\x0d\x66\x45\xf4\x3f\x6d\xd2\x65\x04\xd6\x4a\x11\xe6\x89\x62\xa1\xbf\x20\xa5\x47\x47\x31\x0c\x5a\xc9\xdd\xfb\xab\x2d\xca\xb7\x8e\x27\xb2\xc7\x15\x51\x8d\x44\xcd\x6c\x7d\x2c\x5a\xea\xbf\x92\x45\x33\xdf\x7b\x7f\x0d\x00\x00\xff\xff\xb1\xbd\x81\xdb\xe4\x01\x00\x00")

func golangGetTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetTmpl,
		"golang.get.tmpl",
	)
}

func golangGetTmpl() (*asset, error) {
	bytes, err := golangGetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get.tmpl", size: 484, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x57\x5d\x8f\x9b\x38\x14\x7d\xc6\xbf\xe2\x2e\x6a\x2b\x18\xa5\x44\xfb\x9a\x55\x1e\x9a\x49\x76\x35\x52\x77\xd2\xce\x50\xad\x56\xaa\x34\x75\xc0\x30\x74\xc0\x26\xb6\x49\xa8\x58\xfe\xfb\xca\xc6\x80\xc9\x24\xb3\x53\xed\x53\xc0\xbe\xf7\xdc\x73\x3f\x8e\x71\xe6\x73\xf8\xf0\x25\xdc\xfe\xb1\xb9\xdd\xdc\x7d\x08\x37\x6b\x58\xfd\x0d\x29\x2b\x9f\xd2\x20\xa3\x73\x51\xe2\x88\x14\x8c\x3e\x91\x1f\x29\x9b\xc7\xbb\x3a\x38\xfc\x8a\xe6\x73\x58\x6f\xe1\x76\x1b\xc2\x66\x7d\x13\x06\x08\x95\x38\x7a\xc2\x29\x81\xa6\x81\xe0\x93\x79\x6e\x5b\x84\xb2\xa2\x64\x5c\x82\x87\x1c\x37\x62\x54\x92\x5a\xba\xc8\x71\x63\x2c\xf1\x0e\x0b\x32\x17\xfb\x5c\xbd\x27\x85\x5e\x96\x59\x41\xd4\xaf\x90\x3c\x62\xf4\xe0\x22\xd4\x34\xef\x81\x63\x9a\x12\x08\x36\xb5\xe4\xf8\x46\xc3\x09\x05\xed\xa8\x50\x1f\xd4\x23\xb8\xd3\xa8\xae\x76\x23\x34\x56\x66\x3e\x52\x64\x3f\x71\x72\x20\x54\x42\xc4\x68\x9c\xc9\x8c\x51\x9c\x43\x66\xb0\x12\xce\x0a\x88\x70\x25\x32\x9a\xc2\xae\xca\xf2\x18\x12\x9c\xe5\x15\x27\x02\x1d\x30\x87\x07\x58\x82\x61\x14\xdc\x48\x86\x91\x5e\xfd\x8b\xe3\x72\xc3\x39\x2c\x21\xa9\x68\xe4\x11\xce\x81\x70\xce\xb8\xdf\xfd\x40\xc3\x89\xac\x38\x55\x6f\xad\x76\xf8\xc8\xd2\x94\xf0\xce\x3a\x61\xbc\xc0\x52\xa1\x66\x34\x9d\x01\xe6\xa9\x80\x20\x08\x32\x2a\x09\x4f\x70\x44\x9a\xd6\xd7\x3e\xb7\xec\x08\x4b\x50\x65\x09\x6e\xd9\x11\x21\xf9\xa3\x24\x10\xf3\xec\x40\x38\x0c\xc6\xd0\x20\x67\x53\x93\xc8\xdb\x57\x84\xff\x78\x11\x14\x3c\xb1\xcf\x83\x3b\x22\xaa\x5c\xce\x0c\x5f\xe4\x7c\x56\x7e\xaf\xf2\xbe\xd2\xee\xec\x28\x4e\x9c\xef\xd8\xf1\x35\xfe\xbd\x3b\x6a\x4d\x2a\xeb\x95\x72\xa8\x22\xa9\x72\xd0\xbb\xeb\x15\x72\xe2\xdd\x9f\x44\x3e\xb2\x58\x28\x3b\x55\x2f\xd8\x96\x84\x7a\x5d\xde\x33\x10\xac\xe2\x11\x31\x91\x7c\xf0\xe2\x1d\x5c\xad\x57\x9a\x51\xdf\x82\x06\x39\x62\x9f\x3f\xc4\xbb\x6e\x75\xb1\x04\x85\x7d\x06\xc5\x47\x4e\x96\x68\x9b\x5f\x96\x40\xb3\x5c\x79\x3a\xa6\x73\x34\xcb\x67\x7d\x9b\x55\x7f\x7d\xe4\xb4\xc8\x89\x49\xd2\x77\xb1\x0b\x01\x86\xb7\x8e\x7a\x06\xcd\x30\x09\xae\x73\x26\x88\xe7\x3b\x0e\x72\x14\x4e\x6b\xdc\x7d\x34\x50\xe8\x68\x2a\xdb\x4f\x19\x4d\x3d\xff\xb7\x9f\x22\x86\x1c\x35\x31\x59\x51\xe6\x30\x16\xd0\x11\xc7\x4c\x46\x8f\xfd\xcc\x34\xb6\xa0\xd6\x19\xce\x49\x64\xc4\x14\x61\xd1\x89\xf7\x16\x17\x04\xfe\x81\x92\x67\x54\x26\xe0\xbe\xdd\xbb\xd0\xb6\x0b\x95\x99\x42\x5e\x02\x25\xc7\xc1\xac\x1d\x93\x98\x08\x95\x44\x5b\xaa\x8a\xad\x91\x55\x76\x0f\x33\x38\xc9\x50\x0f\x6c\xd3\x4c\xe2\x04\xd0\xb6\xcf\x93\x7e\x21\x6b\xa7\xb5\x95\x6e\x3d\xaa\x2e\xe1\x2a\x97\x8b\x0b\x45\x4b\x0a\x19\x6c\xd4\xa8\x24\x9e\x5b\x51\x51\x95\xea\x24\x20\x71\x5f\xa6\xb7\xc2\x9d\x99\x67\x5f\x17\x17\xf5\x28\xef\xd6\x2b\xc5\x69\xbd\x5a\x98\x44\x66\xc8\x19\xe7\x75\xa1\xcb\x3f\x43\x4e\x3b\x53\xf1\x86\xf1\xf5\xd8\xee\xbb\x9a\x51\x1f\xcc\x10\x80\x37\x1d\x56\x83\xde\xd3\x63\xbb\xef\xc1\x7a\xd5\x4f\x8c\x7f\x06\x47\x8f\xb2\x12\x64\x58\xcf\x2c\x1c\x59\x0f\x03\x6f\x30\x56\x24\xcd\xa8\xf7\xd3\x63\x3e\x26\x1c\xd6\xca\x56\xd6\x0b\x90\xf5\x4c\x3f\x0d\xc9\xaa\x10\x47\x8e\xcb\xb0\xf6\x64\xed\x4f\xd2\xd6\xea\x0e\x6b\x4b\xdd\xb2\xee\x84\x12\xd6\x68\xc4\x18\x33\x53\xdb\x61\xed\xc3\x35\x2b\x8a\x4c\xfe\x67\x85\x64\x1d\xc8\x3a\xe8\x8d\xfd\xe7\x38\x77\x2c\xcf\x77\x38\x7a\x7a\x25\xd2\x68\xae\xb1\x2e\xa8\x04\x35\x0d\xbc\x89\x77\x3a\xb9\xc5\xf2\xb9\x56\xc4\x7a\xe5\x76\x93\x08\x6f\x64\x7d\xd9\x2c\xac\x07\x33\x35\x30\x97\x0d\x6f\x8a\x32\xd7\xa6\x5d\x41\x27\x0e\x6d\x6b\x55\xd7\x0c\x6e\xf7\x73\x32\x2f\x27\x5e\x3e\xe4\x2c\xbd\x97\x85\xf4\x84\x2c\xa6\xdf\xa0\x20\x08\x60\x72\x60\x37\xdd\x57\xd6\x88\xfd\xa3\xe5\x37\x38\xf8\xc8\x66\x67\x8a\x33\xe5\x36\x1c\x91\xc8\x39\x25\x33\x50\x3d\x39\x56\xec\x63\xf5\x6a\x0a\x3c\xb6\xf0\xdd\x64\xa3\xd1\x4a\x5c\x40\xa7\xc9\x93\x38\x8b\xce\xd8\x5a\xd1\x47\x4b\x57\xaf\xde\xa7\x55\x23\x7c\xa6\x78\x43\x08\x1f\xee\xa3\x47\x52\x60\xcf\x37\x65\xb3\xc8\x7c\x53\xf4\xbb\xed\xfb\xcf\x1f\xa1\x6d\xbf\xbd\x8c\x34\x08\xa7\x97\x85\x0f\x83\x2c\x4e\x73\x34\xb3\xd4\x91\x1e\xa5\xf8\x73\x39\x76\x3e\x43\x8e\x43\xcb\x06\xf0\x0b\x62\xbd\xd8\xb2\x0b\x93\xf1\xf2\x15\xa0\x41\xce\x7c\x0e\xe1\x76\xbd\x5d\x00\x27\x34\x26\x1c\xca\x1c\x47\xe4\x91\xe5\x31\xe1\x42\x9f\x53\xe6\xa6\x64\x1d\x55\xdd\x8a\xe7\x8a\x7d\xbe\xf8\x4a\xdf\x8a\xaf\x54\x81\xab\xc7\xbd\x3b\x83\x71\x1e\x7d\x93\x9c\xf5\x29\x50\x2a\x33\x52\xbe\xd7\xf9\x89\x5e\x7a\x26\xdd\x41\x78\x13\x99\x99\x4d\xbb\x2a\xd6\xa1\xf0\x7b\x46\xf2\x78\xbc\x85\x1a\x77\x5d\x91\xd0\x54\xc9\xa2\x60\x08\x65\x09\x04\xdb\xb2\xbb\x7c\xde\x50\x41\xb8\x1c\x61\x86\xc0\xc1\x35\x27\x58\x92\x8e\x6a\x8f\x7b\x8e\xc2\x25\xa4\x09\xa1\x2b\x3d\x94\x36\xd6\x94\xd7\xb4\x4e\x1d\xc5\x2f\xa5\xbe\x9d\xe7\xe4\x1c\x3b\xbd\xf9\x0a\x76\x67\x40\xfe\x17\xb1\x67\x75\x57\x8b\x6f\x12\xab\x85\x53\xb0\xb1\x91\xc9\xf3\x4e\xc2\xc1\xee\x14\xb4\xc3\x38\x8f\x5d\x7f\xb0\x0f\xa3\x89\xb9\x7f\x8a\xda\x80\x91\xea\x74\xbd\x39\x2c\xe0\x30\x80\x7b\xd5\xc9\xb6\x0f\x11\xcb\xab\x82\xaa\x18\xd6\x71\xd2\x63\xe9\x7f\x34\xd7\xda\x42\xfd\xa1\xb1\x60\xae\x4e\x71\x0e\x38\xaf\x14\x84\x25\x33\x68\x40\x0b\xa9\x82\xe5\xb9\xcf\xbd\xbe\x2e\x9a\xd7\x2a\x38\xa0\xcb\x1c\xef\x48\xe2\xf9\xa7\x21\x47\x96\xef\x2a\xb0\xc4\xf6\x7e\xda\xc4\x7f\x03\x00\x00\xff\xff\x63\x84\x8d\x52\x55\x0e\x00\x00")

func golangHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangHeaderTmpl,
		"golang.header.tmpl",
	)
}

func golangHeaderTmpl() (*asset, error) {
	bytes, err := golangHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.header.tmpl", size: 3669, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangMiscTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4b\xc3\x40\x10\x85\xcf\xcd\xaf\x78\x84\x08\x2a\xa6\x3f\xa0\xe0\xa5\x07\x41\x04\x0f\x5a\x3c\x77\xcd\x4e\xc2\x48\xba\x8d\x9b\x8d\x50\xc6\xfd\xef\x32\xbb\xb2\x56\x7a\x5b\xde\x7c\xef\xbd\x99\x15\x69\x61\xa9\x67\x47\xa8\x2d\x75\x63\x8d\x18\xab\x7e\x71\x1d\xae\x8f\xef\x1f\xb8\x15\xc1\xfa\x85\x3a\xe2\x2f\xf2\x5b\x33\x13\x62\x7c\x3c\x4c\xe3\x0d\x74\xf0\xca\x83\x33\x61\xf1\xaa\x26\x61\x7b\xb4\x27\x0d\x10\x01\x39\x8b\x36\xc6\xaa\x3a\x6f\x98\x83\x5f\xba\x90\x3a\xc2\x69\xa2\xe4\x79\x36\x87\xe4\xcf\x33\x88\x9a\xbd\x71\x03\x61\xfd\xc0\x34\xda\x59\xe9\x95\x08\xb8\x2f\xf0\xb9\xef\xb7\x2b\x8b\x3b\x4d\x4d\x4f\xa5\x77\x66\x98\xd3\x12\xab\x7d\x09\x6d\xf8\x0e\x4d\xc0\xe6\xfe\x6f\x9c\xe9\x86\x2f\xd2\x9e\x48\xaf\xd9\xe8\xf3\xcd\x8c\x0b\xe1\x1b\x93\x67\x17\x7a\xd4\x57\x9f\x75\x86\x32\xbd\xd7\x0d\xdb\x72\x74\xd1\xab\x7f\x7f\xf1\x13\x00\x00\xff\xff\x1b\x36\x80\x57\x6d\x01\x00\x00")

func golangMiscTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangMiscTmpl,
		"golang.misc.tmpl",
	)
}

func golangMiscTmpl() (*asset, error) {
	bytes, err := golangMiscTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.misc.tmpl", size: 365, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\xc1\x6e\xe3\x36\x10\x3d\x4b\x5f\x31\x15\x50\x40\x42\x1c\xa2\xe8\x71\x01\x1d\x82\x6d\x5a\x04\x48\x8d\x24\xee\xb6\x87\xc5\xc2\x60\xac\xa1\xc0\x40\x26\x6d\x72\x14\x3b\x10\xf4\xef\xc5\x90\xb2\x23\xc5\x4e\xda\xa2\x97\x3d\xf8\x40\x73\xe6\xcd\xcc\x7b\xc3\xa7\xae\xbb\x84\x0a\x95\x36\x08\x99\xd7\xb5\x91\xd4\x3a\xcc\xe0\xb2\xef\xd3\x2f\x9b\x4a\x12\x76\x1d\x88\x45\xab\x94\xde\x43\xdf\xe7\x5d\x07\x2b\xda\x6f\xa4\x93\x6b\x10\x57\xae\xf6\xd0\xf7\x33\x68\x43\x24\x84\x50\x72\xed\x8a\x44\xcc\x8d\x87\xb9\x5c\x23\xf4\x7d\x01\x9c\x3d\xa4\x3e\x20\xb5\xce\x84\x64\x74\x8e\x7f\xd6\x15\x29\x37\x83\xa6\x0a\xd5\xd3\x71\x67\x8f\xb6\x7a\x09\x4d\x75\x69\xb2\xb2\xc6\x13\x2c\x97\x9e\xd6\xb4\xdc\x38\x54\x7a\x5f\x32\xb0\xd3\x86\x14\x64\x3f\x6e\x33\x10\x8b\xfb\xdb\xbb\x70\x03\x7d\xff\x26\xc3\x87\x59\xce\x65\x1c\xa7\x4c\xd3\xe4\x59\x3a\xd8\x7a\xf8\xfa\xed\xf1\x85\x30\x1e\x9f\x65\xd3\x22\xff\xa5\x0d\xa1\x53\x72\x85\x1d\x47\x76\x1d\x68\x05\xe2\xce\x7a\x4d\xda\x1a\xd9\x5c\xb9\xba\x5d\xa3\x21\x1f\x6a\x3f\xc1\xa7\x92\x89\x69\xd0\x1c\x09\x63\xbc\x80\x55\x82\xdc\x6c\xd0\x54\x79\x3c\xcf\x38\x50\x69\x6c\xaa\x70\x3e\xc6\x17\xa1\xca\x91\x98\x44\x56\xd5\xa1\x08\xa3\xab\xd6\xac\x72\xc3\x24\x7b\x72\xda\xd4\x05\x74\x69\x92\x6c\x47\xf0\x5b\x3f\x03\x0e\x10\x42\x14\x69\x92\x30\xb3\x1f\xf5\x9c\x3c\x5d\x5c\x9c\x41\xc8\x20\x4c\x22\x0e\xc1\x47\x86\xb3\x01\xf7\x6d\x82\x27\xb7\xb2\xe6\x59\xdc\x90\x95\xf9\x53\xf1\x4e\x54\x36\x83\x6c\xd4\x18\x36\x1e\x63\x17\xff\xba\xfe\x1b\x00\x53\x85\xfc\x41\x1c\x27\x4d\x8d\xd3\xbd\x94\x8f\x0d\xfe\xca\x34\xc7\x71\xb5\x1a\x16\x58\x30\xfa\xb0\xad\xf0\x43\x09\x46\x37\x81\xca\x11\xdf\x79\xc6\x31\x9f\x6d\xd3\xae\x79\x7b\x33\x2e\xfa\x9e\x98\xa7\xa0\x22\x5c\xe5\x45\xc1\xdd\x1d\x34\xed\x63\xa3\x51\x92\x39\x62\xe5\xe7\x76\x17\x1a\x5b\x2e\x8d\xdd\xb1\xc2\x73\xbb\xcb\x0b\xf1\xe5\x8f\xcf\xf9\x74\x15\x46\xf3\x5d\xb5\x64\x47\x33\x9d\xb4\x3c\x74\xc0\x0d\x7f\xb0\x7c\xe2\xc6\x68\xfa\x53\x36\x93\xa5\xeb\x5f\xd7\xdc\x58\x3a\x29\xa5\x15\x2f\x77\xbe\xf5\x05\x94\x25\xfc\x14\x18\x73\xf1\x79\x1b\xdd\xcc\xe0\x2f\x27\x37\xd7\xce\xe5\x6a\x4d\xe2\x9a\xdf\xb9\xca\x33\x83\x58\x01\xd9\x83\x71\x48\x82\x06\xa5\x27\xb0\x06\xe3\x03\xc8\xce\x72\x74\x6c\xe1\xdc\xe6\x5e\xfe\x9f\x87\x15\x0a\x44\x87\x60\xc2\x27\xee\x02\x17\xc3\xc3\xca\xb7\xfe\xeb\xa7\x61\xd6\xcb\x9f\xbf\x15\x70\x31\x35\x95\x34\xb1\x8f\x4f\xa2\xb1\xf5\x82\xd6\x94\xc7\xab\xd9\xe0\x1b\x61\x41\xe3\x0c\x46\xd3\xc8\x00\x8f\x0e\xb2\x68\x37\x1b\xeb\xc8\xc7\x1b\x6d\xea\x70\xc9\xee\x58\x02\xe3\x56\x4e\x3f\xa3\x13\xf7\x2d\xba\x97\x07\xbb\x3b\x83\x2f\x16\x2b\x69\xd8\x64\x65\x55\x39\xab\x20\x57\x8d\x24\x62\xe3\x89\x98\x45\x1c\x59\xab\xe0\xb9\x65\x09\x7e\xdb\xb0\x24\x73\xfb\x60\x77\xfe\x44\x38\xa3\x9b\xa0\xc1\x10\x3f\x7a\x10\xe7\xe4\x45\xe7\x5e\x25\xe3\x37\x1c\xf4\x58\x46\x7f\x9f\x4c\x70\xbd\xc7\xd5\x39\x76\xfe\x6b\xa1\x83\xad\xd7\x48\x01\xec\xc4\xd3\x7f\x43\x5a\xdc\xdf\x06\x1a\xc7\xc2\x64\xf9\xcd\xef\x77\xb7\x37\xd7\xbf\x14\x90\x05\x09\x0f\x00\xef\xaf\xc9\xc7\x3a\xfc\x63\xfe\xf7\x23\xcc\xc1\x38\x86\x40\x6e\xc9\xd5\x93\xcf\x31\x83\xf7\x93\x2f\xf1\xdf\x01\x00\x00\xff\xff\x4e\xc7\xca\xc1\x23\x08\x00\x00")

func golangUpdateTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangUpdateTmpl,
		"golang.update.tmpl",
	)
}

func golangUpdateTmpl() (*asset, error) {
	bytes, err := golangUpdateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.update.tmpl", size: 2083, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
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
	"golang.create.tmpl": golangCreateTmpl,
	"golang.delete-all.tmpl": golangDeleteAllTmpl,
	"golang.delete-world.tmpl": golangDeleteWorldTmpl,
	"golang.delete.tmpl": golangDeleteTmpl,
	"golang.footer.tmpl": golangFooterTmpl,
	"golang.get-all.tmpl": golangGetAllTmpl,
	"golang.get-count.tmpl": golangGetCountTmpl,
	"golang.get-has.tmpl": golangGetHasTmpl,
	"golang.get-last.tmpl": golangGetLastTmpl,
	"golang.get-limit-offset.tmpl": golangGetLimitOffsetTmpl,
	"golang.get-paged.tmpl": golangGetPagedTmpl,
	"golang.get.tmpl": golangGetTmpl,
	"golang.header.tmpl": golangHeaderTmpl,
	"golang.misc.tmpl": golangMiscTmpl,
	"golang.update.tmpl": golangUpdateTmpl,
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
	"golang.create.tmpl": &bintree{golangCreateTmpl, map[string]*bintree{}},
	"golang.delete-all.tmpl": &bintree{golangDeleteAllTmpl, map[string]*bintree{}},
	"golang.delete-world.tmpl": &bintree{golangDeleteWorldTmpl, map[string]*bintree{}},
	"golang.delete.tmpl": &bintree{golangDeleteTmpl, map[string]*bintree{}},
	"golang.footer.tmpl": &bintree{golangFooterTmpl, map[string]*bintree{}},
	"golang.get-all.tmpl": &bintree{golangGetAllTmpl, map[string]*bintree{}},
	"golang.get-count.tmpl": &bintree{golangGetCountTmpl, map[string]*bintree{}},
	"golang.get-has.tmpl": &bintree{golangGetHasTmpl, map[string]*bintree{}},
	"golang.get-last.tmpl": &bintree{golangGetLastTmpl, map[string]*bintree{}},
	"golang.get-limit-offset.tmpl": &bintree{golangGetLimitOffsetTmpl, map[string]*bintree{}},
	"golang.get-paged.tmpl": &bintree{golangGetPagedTmpl, map[string]*bintree{}},
	"golang.get.tmpl": &bintree{golangGetTmpl, map[string]*bintree{}},
	"golang.header.tmpl": &bintree{golangHeaderTmpl, map[string]*bintree{}},
	"golang.misc.tmpl": &bintree{golangMiscTmpl, map[string]*bintree{}},
	"golang.update.tmpl": &bintree{golangUpdateTmpl, map[string]*bintree{}},
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

