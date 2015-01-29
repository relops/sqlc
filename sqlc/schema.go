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
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x58,
		0x5f, 0x6f, 0xdb, 0x36, 0x10, 0x7f, 0xd7, 0xa7, 0x38, 0x18, 0x45, 0x20,
		0x07, 0x8a, 0x12, 0x6c, 0x45, 0x1f, 0x0c, 0xe4, 0xc1, 0x69, 0x9c, 0xd5,
		0x83, 0xeb, 0x04, 0xb1, 0xb2, 0x60, 0x4f, 0x03, 0x23, 0x53, 0x29, 0x37,
		0x55, 0xf6, 0x24, 0x3a, 0x4b, 0x61, 0xe8, 0xbb, 0xef, 0x8e, 0xa4, 0x64,
		0x4a, 0xa2, 0x1c, 0x7b, 0xc3, 0x8a, 0xe5, 0xa1, 0x96, 0x8e, 0xf7, 0xf7,
		0x77, 0xbf, 0x23, 0xa9, 0x9e, 0x9f, 0x43, 0xf4, 0x69, 0xba, 0x80, 0x9b,
		0xe9, 0x6c, 0x02, 0x8f, 0xe3, 0x05, 0x8c, 0x1f, 0xa2, 0xdb, 0x9f, 0x26,
		0xf3, 0xc9, 0xfd, 0x38, 0x9a, 0x5c, 0xc3, 0x19, 0x8c, 0xe7, 0xbf, 0xc2,
		0xe4, 0x7a, 0x1a, 0x2d, 0x20, 0xba, 0xd5, 0xaa, 0x8f, 0xd3, 0xd9, 0x0c,
		0xae, 0x26, 0x30, 0xbb, 0x5d, 0x44, 0xf0, 0xf8, 0x69, 0x32, 0x87, 0x69,
		0x04, 0x28, 0xbf, 0x9f, 0xd4, 0x76, 0x9e, 0xb7, 0xdd, 0xc2, 0xbb, 0x75,
		0xce, 0x97, 0x05, 0x8c, 0x2e, 0x21, 0xa4, 0x27, 0x11, 0x33, 0xc9, 0x0b,
		0x28, 0x4b, 0xb5, 0x96, 0x6c, 0xb2, 0x58, 0xaf, 0xd1, 0x93, 0x14, 0xab,
		0x4c, 0x2d, 0x79, 0x6b, 0x16, 0xff, 0xc1, 0x9e, 0x39, 0x14, 0x7f, 0xa6,
		0xb1, 0xe7, 0x89, 0xaf, 0xeb, 0x55, 0x2e, 0xc1, 0xf7, 0x00, 0x06, 0x4b,
		0x26, 0xd9, 0x13, 0x2b, 0xf8, 0x39, 0x2e, 0x0d, 0x48, 0x90, 0xf3, 0x24,
		0xe5, 0xb1, 0x54, 0xcf, 0x52, 0x7c, 0xe5, 0x03, 0x6f, 0xe8, 0x79, 0x2f,
		0x2c, 0x57, 0xea, 0xf2, 0xdb, 0x9a, 0x5f, 0xad, 0x56, 0x29, 0x5c, 0x82,
		0xd1, 0x0b, 0x23, 0x14, 0xdd, 0x26, 0x7e, 0xc2, 0xd2, 0x82, 0x0f, 0x8d,
		0xca, 0x35, 0xe6, 0xd4, 0x55, 0x21, 0x6f, 0xe1, 0x43, 0x26, 0x5e, 0xfd,
		0x8b, 0x00, 0x2e, 0x86, 0xb6, 0x32, 0x2d, 0x1d, 0x6c, 0x70, 0x93, 0xae,
		0x98, 0xfc, 0xf1, 0x07, 0x47, 0x0e, 0x7a, 0xc1, 0xbf, 0x08, 0x9b, 0xba,
		0x1f, 0xde, 0xf7, 0xe8, 0x7e, 0x78, 0x6f, 0xeb, 0x4e, 0x33, 0xd9, 0xd5,
		0x13, 0x99, 0xf4, 0x2f, 0x6c, 0x15, 0x97, 0x33, 0x91, 0x29, 0x57, 0xb5,
		0xda, 0x7c, 0x93, 0xa6, 0x6e, 0x98, 0x10, 0xe5, 0xb0, 0x5a, 0xdd, 0x96,
		0xb6, 0xbe, 0x1b, 0x33, 0x5a, 0x61, 0x4f, 0xa9, 0x02, 0xa9, 0xab, 0xef,
		0x86, 0xcd, 0xb6, 0x21, 0x8d, 0xa6, 0x5d, 0x2f, 0x7a, 0x55, 0x6a, 0x06,
		0x32, 0x87, 0x95, 0xab, 0xf4, 0xbd, 0x56, 0x4e, 0x44, 0x2b, 0x0b, 0x85,
		0x65, 0x47, 0x7f, 0x5f, 0x0c, 0x87, 0xc5, 0x42, 0xe6, 0x22, 0x7b, 0xee,
		0x37, 0xd1, 0xeb, 0x4d, 0x9b, 0x68, 0x2f, 0x6a, 0x51, 0x03, 0xb1, 0x3e,
		0xff, 0x83, 0x41, 0xa5, 0xe1, 0xf6, 0xd6, 0xa1, 0x2e, 0x0e, 0x11, 0x69,
		0xc3, 0x34, 0x2b, 0x78, 0x2e, 0x17, 0x5c, 0x2e, 0x24, 0x5f, 0x03, 0x12,
		0x87, 0xe7, 0x09, 0x8b, 0x39, 0x6c, 0xd1, 0x1d, 0x4a, 0xfd, 0x88, 0x72,
		0xb8, 0x11, 0x3c, 0x5d, 0x06, 0xbb, 0x55, 0xcc, 0x67, 0x67, 0xf8, 0x79,
		0x95, 0x73, 0x32, 0x46, 0x03, 0x1c, 0xf9, 0x9c, 0x65, 0x38, 0xd7, 0xef,
		0x7e, 0x0b, 0xe0, 0x9d, 0x54, 0x83, 0x4f, 0x51, 0xd4, 0xd0, 0x2b, 0x7f,
		0xb4, 0x29, 0xc8, 0xf0, 0x0e, 0xb3, 0x13, 0xaf, 0x28, 0xf4, 0x5b, 0xef,
		0x26, 0x90, 0x96, 0xce, 0x04, 0x86, 0x63, 0x29, 0x8a, 0x7b, 0xa3, 0xf1,
		0x6c, 0x49, 0xae, 0x4b, 0x53, 0xcc, 0xc3, 0x1a, 0xb7, 0x0f, 0xfe, 0x0f,
		0x8a, 0xa9, 0x0d, 0xbf, 0x4b, 0x31, 0x3d, 0xd1, 0x76, 0xc5, 0xec, 0x0d,
		0x4d, 0x3b, 0x29, 0xf8, 0x02, 0x4e, 0x85, 0xc2, 0x64, 0xe8, 0xca, 0x24,
		0x01, 0x77, 0x2e, 0x2f, 0x87, 0x40, 0xab, 0xe0, 0xca, 0xb9, 0xdc, 0xe4,
		0x19, 0x88, 0x90, 0x70, 0x4b, 0xd0, 0x72, 0xe8, 0xa9, 0x2d, 0xdd, 0x24,
		0x79, 0x48, 0x8a, 0x1b, 0x38, 0xdd, 0xa8, 0x4a, 0xff, 0x75, 0x8a, 0x1d,
		0xc0, 0xec, 0x14, 0x37, 0x7d, 0x29, 0x9e, 0xd3, 0x9f, 0x61, 0xc6, 0xbd,
		0x1e, 0x07, 0xea, 0x7f, 0x83, 0x17, 0x6f, 0x34, 0xd9, 0x95, 0xa2, 0x9f,
		0x31, 0x9c, 0xb0, 0x42, 0x0d, 0xe2, 0xd0, 0xa9, 0xe1, 0x75, 0x99, 0x79,
		0x63, 0x8e, 0x3f, 0x2c, 0xa8, 0x3f, 0x7c, 0x42, 0xe1, 0xcd, 0x99, 0x59,
		0x87, 0x4f, 0xc2, 0x39, 0xc5, 0xd3, 0xf4, 0x2a, 0xc4, 0x73, 0x26, 0x12,
		0xc1, 0x73, 0x52, 0x26, 0x64, 0x1c, 0xf1, 0x0e, 0xe8, 0x4c, 0x01, 0xa7,
		0x05, 0x27, 0x3c, 0x30, 0x23, 0x77, 0x09, 0x6f, 0x17, 0x69, 0xb7, 0xe0,
		0x04, 0x15, 0xe4, 0x6a, 0xb6, 0xfa, 0x8b, 0x12, 0x6b, 0x2b, 0x6e, 0xc9,
		0xd5, 0x08, 0xe8, 0x5f, 0xca, 0x4f, 0x67, 0x20, 0x41, 0xf5, 0xe2, 0x7b,
		0x06, 0x0f, 0xa0, 0xae, 0x79, 0x04, 0xb2, 0x74, 0x72, 0x65, 0x2f, 0x76,
		0xba, 0x93, 0xfb, 0xa2, 0x51, 0xce, 0x9b, 0x58, 0xaa, 0xec, 0xac, 0x1a,
		0xf0, 0xad, 0x0e, 0x8d, 0xa3, 0x50, 0x11, 0x11, 0xc5, 0x2c, 0x15, 0xac,
		0xd8, 0x69, 0x21, 0x36, 0xba, 0xa3, 0x15, 0x5d, 0x3c, 0x2b, 0x6a, 0x37,
		0x5a, 0x73, 0x83, 0xdb, 0x6d, 0x6e, 0x9d, 0xcd, 0x6b, 0xad, 0x88, 0xa5,
		0x2f, 0x6a, 0x35, 0xb1, 0xd6, 0x61, 0x23, 0x14, 0x31, 0xec, 0x85, 0xa5,
		0x1b, 0xee, 0x98, 0xbf, 0x8f, 0xab, 0x6c, 0x29, 0x54, 0x3e, 0x95, 0xe9,
		0xcf, 0x2b, 0x91, 0xf5, 0x59, 0x36, 0xb3, 0x1c, 0x02, 0xe9, 0xb6, 0x3c,
		0xec, 0xd8, 0xaa, 0xe9, 0x10, 0xc3, 0xe9, 0x3e, 0x5c, 0x87, 0xf5, 0xfc,
		0xf8, 0xc3, 0x26, 0x40, 0x36, 0x11, 0x1a, 0x0b, 0x24, 0x07, 0x98, 0xab,
		0xee, 0x43, 0x4c, 0xd7, 0x4f, 0x35, 0x47, 0x81, 0x12, 0x4f, 0x5e, 0xd7,
		0x79, 0x2d, 0xa6, 0x17, 0x2d, 0x1e, 0xe7, 0xcf, 0x45, 0x2d, 0xa6, 0x17,
		0x2d, 0xfe, 0xf8, 0x45, 0xa4, 0xcb, 0x91, 0x11, 0xab, 0x17, 0x92, 0x1f,
		0x93, 0x7d, 0x12, 0xe3, 0x06, 0xb5, 0xc9, 0x02, 0xe0, 0x18, 0xcb, 0xb4,
		0x3b, 0x00, 0x86, 0x11, 0x20, 0x0c, 0xc3, 0xc6, 0x49, 0xb4, 0xa3, 0xb7,
		0x48, 0xe0, 0x44, 0xc5, 0x84, 0xcb, 0x4b, 0xc8, 0x44, 0x0a, 0xba, 0xa4,
		0x83, 0x58, 0xaf, 0x34, 0x35, 0x07, 0x47, 0xfa, 0x31, 0x0e, 0xb3, 0xba,
		0x7c, 0xb0, 0x27, 0x21, 0x0e, 0xeb, 0x97, 0x6a, 0x15, 0x63, 0x1a, 0xab,
		0x16, 0xa6, 0x0a, 0x4e, 0x55, 0x88, 0x42, 0x90, 0xaa, 0x09, 0x0c, 0x6a,
		0x54, 0x4c, 0xa9, 0x1d, 0x10, 0xc1, 0x4a, 0xe0, 0x78, 0xff, 0xfe, 0x1f,
		0xa4, 0x6c, 0xd6, 0x6b, 0x2a, 0x50, 0xf6, 0xb5, 0xcc, 0xf0, 0x80, 0xd7,
		0x0c, 0xb0, 0x58, 0xc0, 0xea, 0xfe, 0x5b, 0x1c, 0x38, 0xe9, 0x71, 0xde,
		0xcb, 0xb4, 0x46, 0x9c, 0x36, 0xdf, 0x1a, 0xf1, 0xda, 0xac, 0x6b, 0xc4,
		0x6d, 0x71, 0x4f, 0xff, 0x95, 0xd5, 0xa3, 0x8d, 0xfb, 0x11, 0xac, 0x1c,
		0x17, 0xbe, 0xbd, 0xff, 0xd8, 0xdc, 0x3b, 0xbc, 0x63, 0xba, 0x5f, 0x76,
		0xab, 0xf6, 0x35, 0x4a, 0xc5, 0x1b, 0xe9, 0x1f, 0x2d, 0x51, 0x8d, 0x73,
		0xa2, 0xda, 0x8b, 0x68, 0x2f, 0x9a, 0xbd, 0x48, 0xf6, 0xa1, 0x58, 0x1e,
		0x3b, 0xc8, 0x63, 0x4a, 0x1c, 0xf7, 0x20, 0x8d, 0x98, 0x0d, 0x55, 0x1c,
		0xaa, 0xa2, 0x8e, 0xf0, 0xf5, 0x99, 0x7d, 0x7b, 0xe2, 0x0e, 0x87, 0x38,
		0xf7, 0xc6, 0x19, 0x0d, 0xfe, 0x60, 0xd0, 0x1c, 0x22, 0x8d, 0x74, 0xcf,
		0x84, 0x55, 0x49, 0x1c, 0x57, 0x14, 0x61, 0xdb, 0x53, 0x93, 0x8a, 0x75,
		0xb8, 0x27, 0xfa, 0xd8, 0x40, 0x4f, 0xf6, 0xb7, 0x87, 0xed, 0x8f, 0x4e,
		0xb2, 0xd6, 0x11, 0x71, 0x84, 0xf3, 0x3b, 0x96, 0x73, 0xfc, 0xf8, 0x1d,
		0x5a, 0x27, 0x68, 0x33, 0xd9, 0x9a, 0x6b, 0x9e, 0x3a, 0xd0, 0xe1, 0xec,
		0xac, 0x7d, 0xa0, 0xb7, 0xce, 0xc1, 0x43, 0x23, 0xf7, 0x9c, 0x96, 0xe4,
		0x67, 0xdf, 0x61, 0x69, 0x67, 0x57, 0x0b, 0xb7, 0x57, 0x02, 0x9f, 0xb2,
		0x67, 0x43, 0x7a, 0xf3, 0xb6, 0xfd, 0x85, 0xce, 0xcf, 0x11, 0x90, 0xcb,
		0x40, 0xaf, 0x20, 0x5d, 0xcb, 0x00, 0xee, 0xaa, 0xff, 0x52, 0x19, 0x99,
		0x2c, 0x6a, 0x01, 0xc6, 0x3a, 0x06, 0x3d, 0xf7, 0xb1, 0x6d, 0x95, 0xb0,
		0xf7, 0xd4, 0xb6, 0x4b, 0x69, 0x2c, 0x6c, 0x67, 0x5f, 0x70, 0xe4, 0xe2,
		0x00, 0xee, 0xe9, 0x57, 0xa7, 0xff, 0x76, 0xce, 0x8d, 0x8b, 0x97, 0xea,
		0x93, 0xaa, 0xa2, 0xfd, 0x89, 0x50, 0x58, 0xad, 0x0e, 0xe0, 0xbf, 0xbf,
		0x14, 0x16, 0xa5, 0x57, 0x5d, 0x05, 0x3b, 0x77, 0xc1, 0xd6, 0xdd, 0xfc,
		0x08, 0xd4, 0x0f, 0xb8, 0xc0, 0x37, 0x79, 0x4c, 0xb7, 0x85, 0x41, 0xc3,
		0x72, 0x10, 0x80, 0x11, 0xd0, 0x76, 0x47, 0x02, 0x7c, 0x13, 0xd9, 0xef,
		0x94, 0xb8, 0xe5, 0xac, 0xf3, 0x85, 0x66, 0x1e, 0xff, 0x0e, 0x00, 0x00,
		0xff, 0xff, 0x1a, 0x58, 0xf9, 0x36, 0xfe, 0x13, 0x00, 0x00,
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
