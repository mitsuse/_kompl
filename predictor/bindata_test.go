package predictor

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _predictor_test_lorem_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x90\xcb\x71\xf3\x30\x0c\x84\xef\x7f\x15\x5b\x80\x47\x55\xfc\xb9\xe5\x9a\x02\x10\x72\xad\x60\x86\x2f\x93\x80\xc7\xe5\x07\x8c\x92\x9b\x28\x00\xfb\xf8\xde\xfb\x64\x85\x8e\xe5\x15\xb9\x97\x3e\xb1\xd4\x20\x95\x76\x43\xea\x6d\x31\x19\xcd\x27\x24\xeb\xd0\xa5\x49\xdb\x09\x16\x8d\xe9\x62\x8e\x0b\x50\x7d\xd5\x9e\x61\xac\x23\xae\xb5\x25\xcd\x9a\xbd\x19\xdc\x50\xe4\x33\xf4\x41\xbb\xb4\x89\x2a\x67\x13\x48\xd1\x87\xcb\x81\x0f\x03\x9b\xd6\x10\x47\xd5\xfd\xf1\x8c\xa7\xd4\x1b\x1e\xae\x0b\xad\x2f\x9b\x9e\xc1\x17\x67\x52\x13\xd3\xde\xe0\xa5\x48\x4d\xfd\x52\xde\x4b\x11\x6a\x3b\xfd\x48\xea\x88\x65\x50\x22\x79\x8d\x4c\xfd\x6a\x10\x56\x76\xe0\xff\x96\x14\x37\x42\xa7\x47\x92\xab\xac\x36\x4c\x8e\xc9\x2f\xb6\xcc\x19\xcd\xe3\xc7\xb3\x17\x1f\x61\xc7\x88\x13\x4d\xc1\xb5\x88\xa4\xa5\xfc\x21\x8a\x42\x8e\xbb\x9f\x2a\x86\xb6\x03\x61\xc8\x8c\x87\xcf\x03\x6f\xaf\xc4\x61\xf4\xcd\x31\x18\xf4\x94\x84\x29\xf6\x92\x0f\xcd\x62\xfb\x22\x5a\x8c\xd9\x35\xb3\x6d\x8a\x9b\x54\x98\x26\x2f\x43\x76\x6f\xf4\xfb\x3d\x30\x0b\x32\x17\xe7\x9e\xd6\x5e\x76\x0c\xd9\x80\x34\x70\xac\x5f\xae\x5e\x8f\x7f\xdf\x01\x00\x00\xff\xff\xf8\xa9\x89\x99\xbf\x01\x00\x00")

func predictor_test_lorem_txt_bytes() ([]byte, error) {
	return bindata_read(
		_predictor_test_lorem_txt,
		"predictor/test/lorem.txt",
	)
}

func predictor_test_lorem_txt() (*asset, error) {
	bytes, err := predictor_test_lorem_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "predictor/test/lorem.txt", size: 447, mode: os.FileMode(420), modTime: time.Unix(1421275717, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"predictor/test/lorem.txt": predictor_test_lorem_txt,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"predictor": &_bintree_t{nil, map[string]*_bintree_t{
		"test": &_bintree_t{nil, map[string]*_bintree_t{
			"lorem.txt": &_bintree_t{predictor_test_lorem_txt, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

