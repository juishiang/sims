// Code generated by go-bindata. DO NOT EDIT.
// sources:
// easy.dat
// hard.dat
// impossible.dat

package main


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataEasydat = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xf7\xb0\xe2\x54\xf1\x4b\xcc\x4d\xe5\x54\xf5\xcc\x2b\x28\x2d\x89\x36" +
	"\xb2\x32\xd0\x31\x88\xb5\x31\xb2\x32\xd4\x31\xb1\x43\x16\x34\x8c\x45\xe6\x19\xa1\xf0\x8c\x63\x39\x55\xfd\x4b\x4b" +
	"\xd0\xf4\x1b\xd9\xa1\x88\x1a\xc6\x72\xc5\xbb\x58\x71\xba\x96\xa5\xe6\x95\xc4\x1b\x70\x1a\x72\x1a\x80\xa1\x21\xa7" +
	"\x01\x92\xb8\x21\x27\x4c\x06\x55\xdc\x08\x26\x06\x22\x91\xc4\x8d\x11\xa6\x70\x1a\x72\x01\x02\x00\x00\xff\xff\x7d" +
	"\xc7\x16\x70\xce\x00\x00\x00")

func bindataEasydatBytes() ([]byte, error) {
	return bindataRead(
		_bindataEasydat,
		"easy.dat",
	)
}



func bindataEasydat() (*asset, error) {
	bytes, err := bindataEasydatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "easy.dat",
		size: 206,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1567674257, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataHarddat = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xf7\xb0\xe2\x54\xf1\x4b\xcc\x4d\xe5\x54\xf5\xcc\x2b\x28\x2d\x89\x36" +
	"\xb2\x32\xd0\x31\x88\xb5\x31\xb2\x32\xd4\x31\xb1\x43\x16\x34\x8c\x45\xe6\x19\xa1\xf0\x8c\x63\x39\x55\xfd\x4b\x4b" +
	"\xd0\xf4\x1b\xd9\xa1\x88\x1a\xc6\x72\xc5\xbb\x58\x71\xba\x96\xa5\xe6\x95\xc4\x1b\x70\x1a\x82\x21\x88\x36\x40\x12" +
	"\x87\x88\x40\xe5\x90\xc4\x8d\x20\x2a\x21\x18\x49\xdc\x18\xaa\xde\x00\x2c\x0e\x08\x00\x00\xff\xff\xdd\x0b\x6c\xb8" +
	"\xce\x00\x00\x00")

func bindataHarddatBytes() ([]byte, error) {
	return bindataRead(
		_bindataHarddat,
		"hard.dat",
	)
}



func bindataHarddat() (*asset, error) {
	bytes, err := bindataHarddatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "hard.dat",
		size: 206,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1567674257, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataImpossibledat = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xf7\xb0\xe2\x54\xf1\x4b\xcc\x4d\xe5\x54\xf5\xcc\x2b\x28\x2d\x89\x36" +
	"\xb2\x32\xd0\x31\x88\xb5\x31\xb2\x32\xd4\x31\xb1\x43\x16\x34\x8c\x45\xe6\x19\xa1\xf0\x8c\x63\x39\x55\xfd\x4b\x4b" +
	"\xd0\xf4\x1b\xd9\xa1\x88\x1a\xc6\x72\xc5\xbb\x58\x71\xba\x96\xa5\xe6\x95\xc4\x1b\x70\x1a\x72\xc2\x31\x92\x38\x42" +
	"\x14\x55\xdc\x08\x22\x02\x86\x86\x48\xe2\xc6\x10\x11\x88\x1e\x2e\x40\x00\x00\x00\xff\xff\x49\x76\x16\xe1\xce\x00" +
	"\x00\x00")

func bindataImpossibledatBytes() ([]byte, error) {
	return bindataRead(
		_bindataImpossibledat,
		"impossible.dat",
	)
}



func bindataImpossibledat() (*asset, error) {
	bytes, err := bindataImpossibledatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "impossible.dat",
		size: 206,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1567674257, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"easy.dat":       bindataEasydat,
	"hard.dat":       bindataHarddat,
	"impossible.dat": bindataImpossibledat,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"easy.dat": {Func: bindataEasydat, Children: map[string]*bintree{}},
	"hard.dat": {Func: bindataHarddat, Children: map[string]*bintree{}},
	"impossible.dat": {Func: bindataImpossibledat, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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