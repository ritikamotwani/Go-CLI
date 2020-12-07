// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// data/.DS_Store
// data/hello-world.js
// data/hello-world.txt
// data/hello_world.py
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

var _dataDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xc1\x4a\x03\x31\x14\x45\xef\x4b\x67\x31\x20\x68\x96\x2e\xf3\x03\x0a\xfe\x41\x28\x75\xe1\xda\x0f\x50\x6a\x05\x91\x40\x44\x2b\xea\x6e\xfe\x5c\x89\xb9\xa5\xad\x9d\x32\x20\x88\xa3\xdc\x03\xe5\x14\x32\x37\x09\x59\x64\xde\x1b\x00\x36\x7d\x5e\x9c\x01\x1e\x40\x8b\x6a\x2b\x7f\x7a\x68\xf9\xdb\xc1\xd1\x93\x12\xfe\x9c\x23\xe3\x04\x2f\xc8\x78\x44\xea\x9f\x6b\x74\x94\xbd\x1f\xe2\x0e\xb7\x48\x48\x5b\xfb\x5f\xe0\x14\xf7\x78\xba\x48\xf9\x66\x9e\xf2\xbc\x1c\xd2\xfb\x17\x00\x1c\xed\xcd\x2e\xf1\x8a\xe5\x40\x7a\x73\xe5\xab\xad\xf4\x03\xde\x06\xb2\x42\x08\x21\xc4\xf7\xb0\xaa\xf6\xe0\xb7\x37\x22\x84\x18\x1d\xe5\x7e\x08\x74\xa4\xbb\x6a\xe3\xb8\xa3\x9b\x8d\x8c\xa7\x03\x1d\xe9\xae\xda\xf8\x9c\xa3\x1b\xba\xa5\x3d\x1d\xe8\x48\x77\xd5\xbc\xb4\x8c\xcd\x87\x71\xe5\x55\xf3\x62\x9e\x0e\x74\xfc\x99\xb3\x11\xe2\xaf\x33\xa9\xf2\xe5\xfd\x7f\xbe\xbf\xff\x17\x42\xfc\x63\xac\x99\x5d\xce\xa6\xeb\x86\x60\x07\xc7\x42\xe0\x7a\x15\x60\x21\x80\x9e\x22\xc0\xd5\x8f\x85\xc7\x58\x8f\xab\x10\x10\x62\x64\x7c\x04\x00\x00\xff\xff\x02\x0c\x9f\xaf\x04\x18\x00\x00")

func dataDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_dataDs_store,
		"data/.DS_Store",
	)
}

func dataDs_store() (*asset, error) {
	bytes, err := dataDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1607373907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHelloWorldJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\x41\x6a\xc4\x30\x0c\x45\xf7\x3e\x85\x0a\x05\x2b\x10\x92\x40\x77\x33\xb8\xbb\x42\x6f\xd0\xed\x98\x19\x91\x09\xa8\x96\x62\xab\x10\x08\xb9\x7b\x71\x93\x66\x27\x3e\xef\xff\xa7\xbb\xa4\x62\x40\x8b\x66\x2a\x05\x02\x64\x9a\x7f\xa6\x4c\xe8\x8f\xc8\x37\x57\xb7\x33\x51\x15\xc2\x3f\x89\x67\xac\x92\x0d\x02\xbc\x0d\xc3\x70\x75\x2e\xaa\x76\x23\x19\xfa\xde\xb7\x80\x99\xe6\x16\x32\x95\x06\xc2\x3b\xac\x0e\xea\xdd\x15\x4a\x0f\xf4\x9f\xc4\x2c\xf0\x25\x99\x1f\x2f\xd5\xb1\x35\x47\x9b\xa7\x62\x94\xb0\xce\xb6\x80\x67\xb3\xca\x84\xa9\x63\x19\xf1\xf6\xb1\xc4\x6f\x65\xfa\x7b\x69\xe7\xa7\x34\x42\x34\x78\x9a\xe9\xa5\xef\x59\xee\x91\x9f\x52\xec\xf2\xba\xd6\xa1\xed\x76\x18\x7e\x03\x00\x00\xff\xff\x54\x8e\xe1\xad\xee\x00\x00\x00")

func dataHelloWorldJsBytes() ([]byte, error) {
	return bindataRead(
		_dataHelloWorldJs,
		"data/hello-world.js",
	)
}

func dataHelloWorldJs() (*asset, error) {
	bytes, err := dataHelloWorldJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/hello-world.js", size: 238, mode: os.FileMode(420), modTime: time.Unix(1607372574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHelloWorldTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xb1\x0a\xc2\x30\x10\x06\xe0\xfd\x9e\xe2\xb7\x2e\x4a\x87\x3a\xb7\xc1\x45\x10\x47\x37\x71\x0c\x49\x90\x83\x24\x07\x97\xbb\x49\x7c\x77\xa1\xfb\xb7\x2c\x78\x8b\x2b\xee\xac\xc3\x70\x9b\x67\x3c\x55\x3e\x1a\x1b\xd1\x91\x7b\xaa\x9e\x0b\x02\xcb\x30\x2d\xb1\x5d\x89\xb8\x1b\x5a\xe4\x7e\x3a\xe3\x4b\x00\x30\x2c\xaf\x6b\x12\x37\x84\x80\xe9\x51\x6a\x15\xbc\x44\x6b\x3e\x4c\xdb\x0e\xb4\x98\x6b\xc7\x65\xa3\xdf\x3f\x00\x00\xff\xff\x77\x1a\x3f\xae\x6d\x00\x00\x00")

func dataHelloWorldTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataHelloWorldTxt,
		"data/hello-world.txt",
	)
}

func dataHelloWorldTxt() (*asset, error) {
	bytes, err := dataHelloWorldTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/hello-world.txt", size: 109, mode: os.FileMode(420), modTime: time.Unix(1607367053, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHello_worldPy = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x28\xca\xcc\x2b\xd1\x50\xf2\x48\xcd\xc9\xc9\x57\x08\xcf\x2f\xca\x49\x51\xd2\x04\x04\x00\x00\xff\xff\x4b\x10\x5f\x7a\x14\x00\x00\x00")

func dataHello_worldPyBytes() ([]byte, error) {
	return bindataRead(
		_dataHello_worldPy,
		"data/hello_world.py",
	)
}

func dataHello_worldPy() (*asset, error) {
	bytes, err := dataHello_worldPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/hello_world.py", size: 20, mode: os.FileMode(420), modTime: time.Unix(1607363160, 0)}
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
	"data/.DS_Store":       dataDs_store,
	"data/hello-world.js":  dataHelloWorldJs,
	"data/hello-world.txt": dataHelloWorldTxt,
	"data/hello_world.py":  dataHello_worldPy,
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
		".DS_Store":       &bintree{dataDs_store, map[string]*bintree{}},
		"hello-world.js":  &bintree{dataHelloWorldJs, map[string]*bintree{}},
		"hello-world.txt": &bintree{dataHelloWorldTxt, map[string]*bintree{}},
		"hello_world.py":  &bintree{dataHello_worldPy, map[string]*bintree{}},
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
