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

var _predictor_test_wiki_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x92\x41\x6f\xd4\x40\x0c\x85\xef\xfb\x2b\xac\x5e\xb8\x64\x57\x15\xa0\x42\x8f\xa5\x88\x15\x82\x6d\xab\xb6\x12\x5c\x9d\x89\x37\xb1\x3a\x33\x8e\xec\xc9\x6e\xf3\xef\xf1\x24\x05\x71\x4a\x34\x63\x3f\xbf\xef\x79\xf6\xd2\x00\x46\x13\x08\x92\x92\xe4\x38\x83\xd2\x91\x54\xa9\x83\x22\x80\x06\xbd\x44\xcc\x7d\x03\x6c\x80\x30\xaa\xf4\x8a\x29\x71\xee\xa1\x1e\x4f\xd8\x13\x70\xe6\xc2\x18\xbd\xb5\xa3\x13\x45\x19\xbd\x17\x0b\xec\x45\xfa\x58\x6f\xe1\xfd\xe5\xe5\x27\x68\x67\x78\x94\x96\xd4\x2f\x94\xc9\x28\x91\x36\xf5\x04\x1e\xf8\x85\xdc\x43\xee\xe0\x07\x65\x78\x1e\x24\x8d\x26\x79\xb7\xf9\x5e\xd6\x99\x56\xb0\x70\xa8\xfa\xdb\x32\x57\xed\x7f\x83\xcf\x5c\x06\xb0\x39\x17\x7c\x85\x28\x62\xb4\x58\x50\x3e\x79\xd1\x51\x25\x41\x19\xdc\x87\x1c\xe1\xd6\xf5\xbb\xae\x9a\xee\x51\xdb\xda\x1a\x24\x46\x0a\x85\x25\x37\x50\x55\xc1\xf0\x48\x65\x6e\xc0\x24\x11\x74\x73\xc6\xc4\xa1\xce\xab\x4d\x01\x47\x6c\x39\x3a\x25\xd9\xaa\x54\x1b\x31\x42\x3b\x71\x2c\x5b\x27\xac\x12\x06\x36\x85\xa1\x46\x76\x42\x65\x6c\x23\x6d\x23\xe5\xde\x2d\xa2\x2a\xce\xb6\x20\xbe\xd0\xbc\x3d\x61\x9c\x08\x12\x8e\xb6\x62\xa3\x13\xa9\x7b\x72\xd0\xdc\xa1\x3a\x20\xb7\x8a\x3a\xef\x36\xcf\x03\xfd\x47\x8b\x55\x22\xcb\x94\x83\xf3\xf9\xd0\x3b\x39\x51\xf2\x44\x6b\xbe\xd7\x8b\x92\xe7\x95\xe5\x0c\x93\xad\x05\x0b\x8b\xe3\xaf\x9b\x78\x67\x75\x7d\xdd\xb4\x50\x7b\x6c\x56\x28\xd9\x6e\xb3\x17\xbf\xb8\xe8\xc3\x45\x7d\x01\x23\x47\xd7\x2b\xd5\x4e\x31\x8f\x8f\xe0\x27\xe7\xe9\xb5\x81\x03\x06\xb8\x7f\x82\xdf\x0d\x7c\x53\xa2\x2f\x4f\x5f\x1b\xb8\xa3\xb2\x7c\xef\x47\xca\xcb\xcf\x83\x5b\x85\xeb\x95\xe9\xc0\x41\xc5\xe4\x58\xe0\x17\xe7\x4e\xce\x06\xfe\x2c\xd4\x17\xe9\x71\xbe\x8d\x5e\xea\xea\x08\xfe\xf0\xf9\xca\xbb\x52\x77\xf5\x71\x6d\xbe\x79\x3c\x54\xab\x81\xcc\x44\x3d\xbd\x30\x70\xf1\x65\x4d\x4a\xee\xf7\x06\x8c\x82\x78\xd5\x5f\xbb\x0d\xf4\x21\xf4\xf2\xf6\x42\xf7\xb7\xb7\x75\xf7\xb9\x50\xee\x76\x9b\x3f\x01\x00\x00\xff\xff\xd4\x5c\x1c\x6b\xde\x02\x00\x00")

func predictor_test_wiki_txt_bytes() ([]byte, error) {
	return bindata_read(
		_predictor_test_wiki_txt,
		"predictor/test/wiki.txt",
	)
}

func predictor_test_wiki_txt() (*asset, error) {
	bytes, err := predictor_test_wiki_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "predictor/test/wiki.txt", size: 734, mode: os.FileMode(420), modTime: time.Unix(1421356853, 0)}
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
	"predictor/test/wiki.txt": predictor_test_wiki_txt,
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
			"wiki.txt": &_bintree_t{predictor_test_wiki_txt, map[string]*_bintree_t{
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

