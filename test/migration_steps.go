package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

func test_db_001_initial_schema_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x72, 0x0e,
		0x72, 0x75, 0x0c, 0x71, 0x55, 0x08, 0x71, 0x74, 0xf2, 0x71, 0x55, 0xf0,
		0x74, 0x53, 0xf0, 0xf3, 0x0f, 0x51, 0x70, 0x8d, 0xf0, 0x0c, 0x0e, 0x09,
		0x56, 0x48, 0xcb, 0xcf, 0x57, 0xd0, 0xe0, 0xe2, 0x4c, 0x4a, 0xac, 0x52,
		0x08, 0x71, 0x8d, 0x08, 0x51, 0x08, 0x08, 0xf2, 0xf4, 0x75, 0x0c, 0x8a,
		0x54, 0xf0, 0x76, 0x8d, 0xd4, 0x01, 0x09, 0x17, 0x81, 0x85, 0xb9, 0x34,
		0xad, 0xb9, 0xb8, 0x5c, 0x5c, 0x7d, 0x5c, 0x81, 0xa6, 0xb8, 0x05, 0xf9,
		0xfb, 0x82, 0xb4, 0x59, 0x73, 0x01, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61,
		0xe1, 0xbe, 0x13, 0x57, 0x00, 0x00, 0x00,
	},
		"test/db/001_initial_schema.sql",
	)
}

func test_db_002_populate_table_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xf2, 0xf4,
		0x0b, 0x76, 0x0d, 0x0a, 0x51, 0xf0, 0xf4, 0x0b, 0xf1, 0x57, 0x48, 0xcb,
		0xcf, 0x57, 0xd0, 0x48, 0x4a, 0xac, 0xd2, 0x51, 0x48, 0x4a, 0x2c, 0xd2,
		0x54, 0x08, 0x73, 0xf4, 0x09, 0x75, 0x0d, 0x56, 0xd0, 0x50, 0x2a, 0x2c,
		0x2d, 0xad, 0x50, 0xd2, 0x51, 0x50, 0x4a, 0xcf, 0x2f, 0x2a, 0x50, 0xd2,
		0xb4, 0xe6, 0x02, 0x04, 0x00, 0x00, 0xff, 0xff, 0xda, 0x67, 0x8e, 0xdf,
		0x34, 0x00, 0x00, 0x00,
	},
		"test/db/002_populate_table.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"test/db/001_initial_schema.sql": test_db_001_initial_schema_sql,
	"test/db/002_populate_table.sql": test_db_002_populate_table_sql,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"test": &_bintree_t{nil, map[string]*_bintree_t{
		"db": &_bintree_t{nil, map[string]*_bintree_t{
			"001_initial_schema.sql": &_bintree_t{test_db_001_initial_schema_sql, map[string]*_bintree_t{
			}},
			"002_populate_table.sql": &_bintree_t{test_db_002_populate_table_sql, map[string]*_bintree_t{
			}},
		}},
	}},
}}
