package sqlc

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

func sqlc_tmpl_fields_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x57,
		0xdf, 0x6f, 0xdb, 0x36, 0x10, 0x7e, 0xd7, 0x5f, 0x71, 0x30, 0x82, 0x40,
		0x0a, 0x14, 0x25, 0xd8, 0x8a, 0x3e, 0x18, 0xc8, 0x83, 0xd3, 0x3a, 0xab,
		0x07, 0xd7, 0x09, 0x62, 0x65, 0xc1, 0x9e, 0x06, 0x56, 0xa6, 0x5c, 0x6e,
		0x2a, 0xa5, 0x49, 0x74, 0x96, 0xc2, 0xd0, 0xff, 0x3e, 0x1e, 0x49, 0x31,
		0x94, 0x2c, 0x39, 0x36, 0x86, 0x05, 0xcd, 0x43, 0x2c, 0x1e, 0x4f, 0xf7,
		0x7d, 0x77, 0xf7, 0xf1, 0x87, 0x2e, 0x2e, 0x20, 0xfe, 0x34, 0x5b, 0xc2,
		0xcd, 0x6c, 0x3e, 0x85, 0xc7, 0xc9, 0x12, 0x26, 0x0f, 0xf1, 0xed, 0x2f,
		0xd3, 0xc5, 0xf4, 0x7e, 0x12, 0x4f, 0x3f, 0xc2, 0x39, 0x4c, 0x16, 0xbf,
		0xc3, 0xf4, 0xe3, 0x2c, 0x5e, 0x42, 0x7c, 0xab, 0x5d, 0x1f, 0x67, 0xf3,
		0x39, 0x5c, 0x4f, 0x61, 0x7e, 0xbb, 0x8c, 0xe1, 0xf1, 0xd3, 0x74, 0x01,
		0xb3, 0x18, 0xa4, 0xfd, 0x7e, 0x6a, 0xdf, 0xf3, 0xbc, 0xed, 0x16, 0x4e,
		0x8a, 0x92, 0xae, 0x2a, 0x18, 0x5f, 0x41, 0x84, 0x4f, 0x2c, 0x21, 0x82,
		0x56, 0x50, 0xd7, 0x6a, 0x2e, 0xdd, 0xf0, 0x44, 0xcf, 0xe1, 0x93, 0x60,
		0x39, 0x57, 0x53, 0x5e, 0x41, 0x92, 0xbf, 0xc8, 0x9a, 0x42, 0xf5, 0x77,
		0x96, 0x78, 0x1e, 0xfb, 0x56, 0xe4, 0xa5, 0x00, 0xdf, 0x03, 0x18, 0x95,
		0x34, 0xcd, 0x68, 0x22, 0x46, 0xf8, 0x2c, 0xd8, 0x37, 0x3a, 0xf2, 0x02,
		0xcf, 0x7b, 0x22, 0xa5, 0x9a, 0x15, 0xdf, 0x0b, 0xba, 0x14, 0x25, 0xe3,
		0x6b, 0xb8, 0x02, 0xe3, 0x19, 0xc5, 0xd2, 0x78, 0x9b, 0xfa, 0xa3, 0x51,
		0x60, 0x3c, 0xae, 0xf3, 0x3c, 0xdb, 0x9d, 0x4f, 0x49, 0x56, 0xd1, 0xc6,
		0x65, 0xc6, 0xc5, 0xae, 0x07, 0xe3, 0xc2, 0xbf, 0x0c, 0x1c, 0x97, 0xf7,
		0xef, 0x7a, 0x9d, 0xde, 0xbf, 0x73, 0xdc, 0x6e, 0xb2, 0x9c, 0x88, 0x9f,
		0x7f, 0xea, 0xc1, 0xd3, 0x13, 0xfe, 0x65, 0xd4, 0xf6, 0xed, 0x0b, 0x9a,
		0xea, 0x09, 0xd7, 0x37, 0x96, 0xb9, 0xef, 0x3a, 0x62, 0x45, 0xa2, 0x07,
		0xce, 0x9e, 0xfd, 0xcb, 0x10, 0x90, 0x84, 0x2c, 0x0e, 0x7a, 0xc3, 0x8c,
		0x57, 0xb4, 0x14, 0x4b, 0x2a, 0x96, 0x82, 0x16, 0x20, 0x49, 0xd2, 0x32,
		0x25, 0x09, 0x85, 0xad, 0x0c, 0x27, 0x1b, 0x51, 0x12, 0x2e, 0xab, 0x7d,
		0xf2, 0x47, 0x08, 0x27, 0x42, 0xb5, 0x03, 0x5f, 0x52, 0xad, 0x00, 0x90,
		0x2f, 0x61, 0xab, 0x44, 0x74, 0x27, 0xc1, 0xd8, 0xb3, 0x34, 0xfa, 0x9d,
		0xf1, 0x0d, 0xa3, 0xd9, 0x2a, 0x04, 0x6d, 0x9d, 0x33, 0x19, 0x9b, 0x64,
		0xd2, 0x1c, 0xbc, 0xa0, 0x7e, 0xce, 0x4b, 0x8a, 0xc8, 0x1a, 0x8d, 0xf2,
		0x15, 0x86, 0xae, 0x0d, 0xb7, 0x87, 0x62, 0x25, 0x35, 0xf1, 0xd6, 0xdc,
		0x2c, 0xea, 0x10, 0xb7, 0xbd, 0xd0, 0x28, 0x57, 0xf0, 0x19, 0x9c, 0x31,
		0x95, 0x62, 0xd0, 0xc7, 0x24, 0x85, 0x7e, 0x2e, 0x4f, 0x87, 0x54, 0x4a,
		0x65, 0x5f, 0x52, 0xb1, 0x29, 0x39, 0xb0, 0xa8, 0xa2, 0xc2, 0x4f, 0xc3,
		0xa7, 0xc0, 0x53, 0xcb, 0xc6, 0x70, 0x3c, 0x84, 0xe1, 0x06, 0xce, 0x36,
		0x2a, 0xd1, 0xff, 0xcc, 0x70, 0xa7, 0x5e, 0x2e, 0xc3, 0xcd, 0x00, 0xc3,
		0x0b, 0xfc, 0x33, 0x6d, 0xbe, 0xd7, 0x52, 0x25, 0x5f, 0x32, 0xda, 0x6a,
		0xf2, 0x2b, 0x2d, 0xee, 0x63, 0xe8, 0x73, 0x22, 0xd5, 0x5f, 0xa9, 0x95,
		0x1e, 0xf4, 0x7a, 0x78, 0xbb, 0x32, 0xbb, 0x31, 0x3b, 0x8c, 0xcc, 0x67,
		0x18, 0x3e, 0x45, 0x78, 0xb3, 0x2d, 0x59, 0xf8, 0x34, 0x5a, 0x20, 0x9e,
		0x16, 0x57, 0xc5, 0xd6, 0x9c, 0xa5, 0x8c, 0x96, 0xe8, 0x8c, 0x85, 0xe9,
		0xc1, 0x3b, 0xa0, 0x31, 0x15, 0x9c, 0x55, 0x14, 0xeb, 0x21, 0x19, 0xf5,
		0xa7, 0xf0, 0x7a, 0x92, 0x6e, 0x07, 0x4e, 0xa5, 0x83, 0xc8, 0xe7, 0xf9,
		0x3f, 0x48, 0xac, 0xeb, 0xb8, 0xc5, 0x50, 0x63, 0xc0, 0xff, 0xc8, 0x4f,
		0x33, 0x10, 0xa0, 0x7a, 0xf1, 0x96, 0xe0, 0x21, 0xd8, 0x9c, 0xc7, 0x20,
		0xea, 0x5e, 0xad, 0xec, 0xad, 0x9d, 0xee, 0xe4, 0x3e, 0x34, 0xe4, 0xbc,
		0x49, 0x84, 0x62, 0xe7, 0xe4, 0x20, 0x47, 0x16, 0x5a, 0xae, 0x84, 0x46,
		0x88, 0xd2, 0x4c, 0x32, 0x46, 0xaa, 0x17, 0x2f, 0x59, 0x1b, 0xdd, 0xd1,
		0x46, 0x2e, 0x9e, 0x83, 0xba, 0x8b, 0xd6, 0xde, 0xad, 0x62, 0x8c, 0xa9,
		0xf5, 0xd0, 0xdd, 0xba, 0x0a, 0x25, 0x2c, 0x7d, 0x16, 0x5a, 0x61, 0x15,
		0x51, 0x0b, 0x0a, 0x15, 0xf6, 0x44, 0xb2, 0x0d, 0xed, 0x59, 0x7e, 0x1f,
		0x72, 0xbe, 0x62, 0x8a, 0x4f, 0xf3, 0xea, 0xaf, 0x39, 0xe3, 0x43, 0x6f,
		0xb6, 0x59, 0x06, 0x80, 0xbe, 0x9d, 0x08, 0x2f, 0x6a, 0xd5, 0x72, 0x48,
		0xe0, 0x6c, 0x5f, 0x5d, 0x03, 0xbb, 0x7e, 0xfc, 0xa0, 0x5d, 0x20, 0x57,
		0x08, 0xad, 0x09, 0xb4, 0x03, 0x2c, 0x54, 0xf7, 0x21, 0xc1, 0x13, 0x5e,
		0xad, 0xa3, 0x50, 0x99, 0xa7, 0xcf, 0x45, 0x69, 0xcd, 0x38, 0xd0, 0xe6,
		0x49, 0xb9, 0xae, 0xac, 0x19, 0x07, 0xda, 0xfc, 0xe1, 0x2b, 0xcb, 0x56,
		0x63, 0x63, 0x56, 0x03, 0xb4, 0x1f, 0xc3, 0x3e, 0x4d, 0xe4, 0xfe, 0xb4,
		0xe1, 0x21, 0x50, 0x89, 0x65, 0xda, 0x1d, 0x02, 0x91, 0x08, 0x10, 0x45,
		0x91, 0x6d, 0xe3, 0xb6, 0x59, 0xd0, 0x2a, 0x2b, 0x96, 0xc2, 0xa9, 0xc2,
		0x84, 0xab, 0x2b, 0xe0, 0x2c, 0x03, 0x9d, 0xd2, 0x41, 0xaa, 0x57, 0x9e,
		0x5a, 0x83, 0x63, 0xfd, 0x98, 0x44, 0xdc, 0xa6, 0x0f, 0xee, 0x4a, 0x48,
		0x22, 0x3b, 0x68, 0x66, 0x25, 0xa6, 0x79, 0xab, 0x53, 0x53, 0x55, 0x4e,
		0x95, 0x88, 0xaa, 0x20, 0x66, 0x13, 0x9a, 0xaa, 0x61, 0x32, 0xb5, 0x0e,
		0x80, 0x02, 0xab, 0x81, 0xca, 0x1b, 0xcd, 0x0f, 0x40, 0xd9, 0xcc, 0x5b,
		0x29, 0x20, 0x7b, 0x6b, 0x33, 0x3a, 0xa0, 0x56, 0x01, 0x8e, 0x0a, 0x88,
		0xed, 0xbf, 0xa3, 0x81, 0xd3, 0x81, 0xe0, 0x83, 0x4a, 0x6b, 0xe1, 0x74,
		0xf5, 0xd6, 0xc2, 0xeb, 0xaa, 0xae, 0x85, 0xdb, 0xd1, 0x9e, 0xfe, 0xab,
		0x9b, 0x47, 0xb7, 0xee, 0x47, 0xa8, 0x72, 0x52, 0xf9, 0xee, 0xfe, 0xe3,
		0x6a, 0xef, 0xf0, 0x8e, 0xe9, 0x7e, 0xb9, 0xad, 0xda, 0xd7, 0x28, 0x85,
		0x37, 0xd6, 0x3f, 0xda, 0xa2, 0x1a, 0xd7, 0x5b, 0xd5, 0xc1, 0x8a, 0x0e,
		0x56, 0x73, 0xb0, 0x92, 0x43, 0x55, 0xac, 0x8f, 0x5d, 0xc8, 0x13, 0x24,
		0x2e, 0xf7, 0x20, 0x5d, 0x31, 0xb7, 0x54, 0x49, 0xa4, 0x92, 0x3a, 0x22,
		0xd6, 0x67, 0xf2, 0xfd, 0x0b, 0xed, 0x09, 0x28, 0xd7, 0xbd, 0x09, 0x86,
		0x0b, 0x7f, 0x34, 0x6a, 0x2f, 0x22, 0x5d, 0xe9, 0x81, 0x15, 0xd6, 0x90,
		0x38, 0x2e, 0x29, 0xac, 0xed, 0x40, 0x4e, 0x0a, 0xeb, 0xf0, 0x48, 0xf8,
		0x21, 0x20, 0x23, 0xb9, 0xdf, 0x05, 0x6e, 0x3c, 0x3c, 0xc9, 0x3a, 0x47,
		0xc4, 0x11, 0xc1, 0xef, 0x48, 0x49, 0xe5, 0x47, 0x50, 0xe0, 0x9c, 0xa0,
		0x6d, 0xb2, 0x56, 0x6b, 0x9e, 0x3a, 0xd0, 0xe1, 0xfc, 0xbc, 0x7b, 0xa0,
		0x77, 0xce, 0xc1, 0x43, 0x91, 0x07, 0x4e, 0x4b, 0x8c, 0xb3, 0xef, 0xb0,
		0x74, 0xd9, 0x59, 0xe3, 0xf6, 0x9a, 0xc9, 0x27, 0xbe, 0x36, 0xa2, 0x37,
		0xa3, 0xed, 0x6f, 0x78, 0x7e, 0x8e, 0x01, 0x43, 0x86, 0x7a, 0x46, 0xca,
		0xb5, 0x0e, 0xe1, 0xae, 0xf9, 0x6a, 0x1d, 0x1b, 0x16, 0xd6, 0x20, 0xb1,
		0x8e, 0xa9, 0x5e, 0xff, 0xb1, 0xed, 0xa4, 0xb0, 0xf7, 0xd4, 0x76, 0x53,
		0x69, 0x4d, 0x6c, 0xe7, 0x5f, 0xe5, 0x92, 0x4b, 0x42, 0xb8, 0xc7, 0x5f,
		0x4d, 0xff, 0x75, 0xce, 0xad, 0x8b, 0x97, 0xea, 0x93, 0xca, 0xa2, 0xfb,
		0x85, 0x50, 0x39, 0xad, 0x0e, 0xe1, 0xff, 0xbf, 0x14, 0x56, 0xb5, 0xd7,
		0x5c, 0x05, 0x77, 0xee, 0x82, 0x9d, 0xbb, 0xf9, 0x11, 0x55, 0x3f, 0xe0,
		0x02, 0xdf, 0xd6, 0x31, 0xde, 0x16, 0x46, 0xad, 0x37, 0x47, 0x21, 0x18,
		0x03, 0x6e, 0x77, 0x68, 0x90, 0x23, 0xc6, 0xff, 0x44, 0xe2, 0x4e, 0xb0,
		0x9d, 0x0f, 0x34, 0xf3, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9,
		0x97, 0x04, 0xc9, 0x61, 0x11, 0x00, 0x00,
	},
		"sqlc/tmpl/fields.tmpl",
	)
}

func sqlc_tmpl_schema_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x54,
		0x5b, 0x6f, 0x9b, 0x30, 0x14, 0x7e, 0xe7, 0x57, 0x1c, 0xa1, 0x6a, 0x82,
		0x29, 0x23, 0xef, 0x91, 0xf2, 0xc0, 0x54, 0xda, 0x22, 0x31, 0xb2, 0x15,
		0xda, 0x6a, 0x9a, 0xa6, 0xca, 0x21, 0x26, 0x61, 0xe5, 0x36, 0x6c, 0xba,
		0x45, 0x11, 0xff, 0x7d, 0xbe, 0x40, 0x42, 0x8a, 0x69, 0xd0, 0xfa, 0x50,
		0x11, 0xfb, 0x9c, 0xef, 0x76, 0x6c, 0xcf, 0xe7, 0x10, 0xde, 0xb9, 0x01,
		0xdc, 0xb8, 0x9e, 0x03, 0x4f, 0x76, 0x00, 0xf6, 0x43, 0xb8, 0xba, 0x75,
		0x7c, 0xe7, 0xde, 0x0e, 0x9d, 0x6b, 0xf8, 0x04, 0xb6, 0xff, 0x1d, 0x9c,
		0x6b, 0x37, 0x0c, 0x20, 0x5c, 0xc9, 0xd2, 0x27, 0xd7, 0xf3, 0xe0, 0xb3,
		0x03, 0xde, 0x2a, 0x08, 0xe1, 0xe9, 0xce, 0xf1, 0xc1, 0x0d, 0x81, 0xad,
		0xdf, 0x3b, 0xc7, 0x3e, 0x8d, 0xc1, 0xda, 0x21, 0x1c, 0x0e, 0x60, 0x7d,
		0xad, 0x8a, 0x57, 0x9c, 0xa3, 0x3c, 0xc2, 0x56, 0x98, 0x64, 0x98, 0x50,
		0x94, 0x95, 0xd0, 0x34, 0xf0, 0x10, 0xb8, 0xfe, 0x2d, 0x90, 0xdf, 0x69,
		0x04, 0x8f, 0xce, 0x7d, 0xe0, 0xae, 0xfc, 0xb7, 0xe5, 0x8f, 0xb8, 0x22,
		0x49, 0x91, 0xb3, 0x62, 0x4d, 0x63, 0x5b, 0x57, 0x74, 0x5f, 0x62, 0x02,
		0x8b, 0x25, 0x58, 0xa1, 0xf8, 0xe2, 0xeb, 0x25, 0x8a, 0x5e, 0xd0, 0x16,
		0xcb, 0xd6, 0xf6, 0x9b, 0xaf, 0x27, 0x59, 0x59, 0x54, 0x14, 0x0c, 0x0d,
		0x40, 0xdf, 0x26, 0x74, 0x57, 0xaf, 0xad, 0xa8, 0xc8, 0xe6, 0x64, 0x57,
		0x53, 0xfc, 0x6b, 0xce, 0x59, 0xc5, 0x3f, 0x9d, 0xef, 0x13, 0x5a, 0x25,
		0xf9, 0x96, 0xe8, 0x9a, 0xa9, 0x69, 0x51, 0x91, 0x13, 0x0a, 0x41, 0xb4,
		0xc3, 0x19, 0x82, 0x25, 0xe8, 0x1c, 0xb7, 0xfd, 0xd5, 0x34, 0xba, 0xd0,
		0x51, 0xa1, 0x9c, 0x91, 0x5c, 0x3d, 0xcf, 0x98, 0x22, 0xa9, 0x06, 0xad,
		0xd3, 0x56, 0x0e, 0x97, 0xc8, 0xb5, 0xd0, 0xc2, 0x2b, 0xfe, 0xe0, 0x8a,
		0x55, 0x58, 0x3e, 0xca, 0xb8, 0x24, 0x60, 0x2c, 0x75, 0x44, 0xe1, 0xc0,
		0x18, 0xcf, 0x41, 0x62, 0x0e, 0xc2, 0x0a, 0x6f, 0x12, 0x9c, 0x6e, 0x04,
		0x0c, 0x48, 0x88, 0x87, 0xb2, 0xe4, 0x10, 0xf1, 0x09, 0x82, 0x29, 0xb6,
		0x78, 0x12, 0xb1, 0x48, 0x80, 0x2d, 0x89, 0x1e, 0x59, 0x8f, 0xf3, 0x8d,
		0xec, 0x45, 0x69, 0x82, 0x08, 0x48, 0x53, 0x1a, 0xd3, 0x14, 0xd7, 0x79,
		0x04, 0x06, 0x85, 0x8f, 0x4a, 0x5d, 0x26, 0xb8, 0x24, 0xc0, 0x29, 0x8e,
		0x28, 0x77, 0x61, 0x98, 0x70, 0x98, 0xd0, 0x22, 0x13, 0x61, 0xc5, 0x92,
		0x45, 0x98, 0xaa, 0x30, 0xad, 0xab, 0xbc, 0xdd, 0x9a, 0xc2, 0xcb, 0x3f,
		0x94, 0x10, 0xc2, 0xe6, 0xb7, 0x9a, 0xf9, 0x88, 0x13, 0xbc, 0x31, 0xa8,
		0xd5, 0xd1, 0xcd, 0xc4, 0x3c, 0x86, 0x58, 0xba, 0x39, 0x85, 0xcf, 0x26,
		0x06, 0x6a, 0xd9, 0x4c, 0xc9, 0x71, 0xf2, 0xdd, 0xa7, 0xff, 0xa0, 0xec,
		0xe7, 0x05, 0x53, 0x26, 0x37, 0x32, 0xbb, 0x05, 0x50, 0x4b, 0xb9, 0x31,
		0xeb, 0x7a, 0x8e, 0xf3, 0x6b, 0x27, 0xb8, 0x40, 0x7c, 0xab, 0x99, 0xe4,
		0x8c, 0xd7, 0x2b, 0xa3, 0xa4, 0x96, 0xc0, 0x9a, 0x02, 0xf2, 0x05, 0xed,
		0xd7, 0x58, 0x81, 0x94, 0xc4, 0x1d, 0x0a, 0x2c, 0xd9, 0x85, 0xd0, 0x41,
		0x26, 0xd1, 0x12, 0x8c, 0x0c, 0x84, 0x2b, 0x07, 0x9c, 0x12, 0x7c, 0x5e,
		0xdd, 0xc9, 0x69, 0x7d, 0xcd, 0xf9, 0xdf, 0xe0, 0x4e, 0xed, 0x65, 0xaa,
		0xdd, 0x15, 0xbf, 0xa4, 0x5b, 0xbe, 0x0c, 0xec, 0xdd, 0xc0, 0x71, 0xf2,
		0xb7, 0xbb, 0x12, 0x46, 0xce, 0xb7, 0xcf, 0x86, 0xad, 0xac, 0x1b, 0x1c,
		0xbb, 0xb7, 0x55, 0x46, 0x7f, 0x68, 0x47, 0xd6, 0x59, 0x0b, 0x4d, 0xac,
		0x50, 0xee, 0x09, 0x3e, 0x93, 0x1f, 0xc3, 0xd3, 0x24, 0x3b, 0x7b, 0x97,
		0x0c, 0xc8, 0xd3, 0xc3, 0x42, 0xff, 0xf1, 0x53, 0x48, 0x18, 0x08, 0xeb,
		0xaf, 0x4f, 0x3f, 0x86, 0xc3, 0xc7, 0x62, 0xcc, 0x8b, 0xae, 0x3c, 0x98,
		0xba, 0x39, 0x3c, 0x9a, 0x62, 0x6a, 0x3d, 0x87, 0xef, 0xbf, 0x86, 0xaf,
		0xa8, 0x82, 0xe7, 0x67, 0xf5, 0x6b, 0xb8, 0x1c, 0xbb, 0x66, 0xb2, 0x4d,
		0xa9, 0x74, 0xb4, 0x89, 0x85, 0xa5, 0xfd, 0xe7, 0x9b, 0xba, 0x50, 0xe5,
		0x34, 0x22, 0xfa, 0xbd, 0xa4, 0x7a, 0xa1, 0x9c, 0x45, 0xc4, 0xcd, 0x88,
		0x44, 0xbc, 0xe4, 0x85, 0xa5, 0xb2, 0xec, 0x86, 0x79, 0x5c, 0x3b, 0x5c,
		0x08, 0x11, 0xd4, 0x59, 0xf4, 0x19, 0x1b, 0xed, 0x5f, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xd4, 0xa6, 0xda, 0x3b, 0xc5, 0x07, 0x00, 0x00,
	},
		"sqlc/tmpl/schema.tmpl",
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
	"sqlc/tmpl/fields.tmpl": sqlc_tmpl_fields_tmpl,
	"sqlc/tmpl/schema.tmpl": sqlc_tmpl_schema_tmpl,
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
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
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
	"sqlc": &_bintree_t{nil, map[string]*_bintree_t{
		"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
			"fields.tmpl": &_bintree_t{sqlc_tmpl_fields_tmpl, map[string]*_bintree_t{
			}},
			"schema.tmpl": &_bintree_t{sqlc_tmpl_schema_tmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}
