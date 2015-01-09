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
		0x5f, 0x6f, 0xdb, 0x36, 0x10, 0x7f, 0xd7, 0xa7, 0x38, 0x18, 0x41, 0x20,
		0x05, 0x8a, 0x12, 0x6c, 0x45, 0x1f, 0x0c, 0xe4, 0xc1, 0x69, 0x9d, 0xd5,
		0x83, 0xeb, 0x04, 0xb1, 0xb3, 0x60, 0x4f, 0x03, 0x23, 0x53, 0x2e, 0x37,
		0x55, 0xf2, 0x24, 0x3a, 0x4b, 0x61, 0xe8, 0xbb, 0xef, 0x8e, 0xa4, 0x18,
		0x4a, 0x96, 0x1c, 0x1b, 0xc3, 0x8a, 0xe6, 0xa1, 0x96, 0x8e, 0xf7, 0xe7,
		0x77, 0x77, 0xbf, 0x23, 0xa9, 0x5e, 0x5c, 0xc0, 0xe2, 0xd3, 0x64, 0x0e,
		0x37, 0x93, 0xe9, 0x18, 0x1e, 0x47, 0x73, 0x18, 0x3d, 0x2c, 0x6e, 0x7f,
		0x19, 0xcf, 0xc6, 0xf7, 0xa3, 0xc5, 0xf8, 0x23, 0x9c, 0xc3, 0x68, 0xf6,
		0x3b, 0x8c, 0x3f, 0x4e, 0x16, 0x73, 0x58, 0xdc, 0x6a, 0xd5, 0xc7, 0xc9,
		0x74, 0x0a, 0xd7, 0x63, 0x98, 0xde, 0xce, 0x17, 0xf0, 0xf8, 0x69, 0x3c,
		0x83, 0xc9, 0x02, 0x50, 0x7e, 0x3f, 0xb6, 0x76, 0x9e, 0xb7, 0xdd, 0xc2,
		0xc9, 0xba, 0xe0, 0xcb, 0x12, 0x86, 0x57, 0x10, 0xd1, 0x93, 0x88, 0x99,
		0xe4, 0x25, 0x54, 0x95, 0x5a, 0x4b, 0x36, 0x59, 0xac, 0xd7, 0xe8, 0x49,
		0x8a, 0x3c, 0x53, 0x4b, 0xde, 0x9a, 0xc5, 0x7f, 0xb1, 0x15, 0x87, 0xf2,
		0xef, 0x34, 0xf6, 0x3c, 0xf1, 0x75, 0x9d, 0x17, 0x12, 0x7c, 0x0f, 0x60,
		0xb0, 0x64, 0x92, 0x3d, 0xb1, 0x92, 0x5f, 0xe0, 0xd2, 0x80, 0x04, 0x05,
		0x4f, 0x52, 0x1e, 0x4b, 0xf5, 0x2c, 0xc5, 0x57, 0x3e, 0xf0, 0x02, 0xcf,
		0x7b, 0x66, 0x85, 0x52, 0x97, 0xdf, 0xd6, 0xfc, 0x3a, 0xcf, 0x53, 0xb8,
		0x02, 0xa3, 0x17, 0x2d, 0x50, 0x74, 0x9b, 0xf8, 0x09, 0x4b, 0x4b, 0x1e,
		0x18, 0x95, 0x9b, 0x34, 0x67, 0xf2, 0xe7, 0x9f, 0x3a, 0xb4, 0xf4, 0x82,
		0x7f, 0x19, 0x05, 0x0d, 0xdd, 0xf7, 0xef, 0x7a, 0x74, 0xdf, 0xbf, 0x73,
		0x75, 0x27, 0x99, 0xdc, 0xd5, 0x13, 0x99, 0xf4, 0x2f, 0x5d, 0x95, 0x2e,
		0x67, 0x22, 0x53, 0xae, 0xac, 0xda, 0x6c, 0x93, 0xa6, 0xdd, 0x89, 0x60,
		0x1d, 0xa2, 0x7a, 0x75, 0x5b, 0xb9, 0xfa, 0xbd, 0x59, 0xd5, 0x26, 0x26,
		0x95, 0x0e, 0xab, 0x2e, 0x48, 0x7b, 0xad, 0x3a, 0x33, 0xad, 0x2d, 0x54,
		0x8e, 0x3b, 0xfa, 0xfb, 0x62, 0x74, 0x58, 0xcc, 0x65, 0x21, 0xb2, 0x55,
		0xbf, 0x89, 0x5e, 0x6f, 0xda, 0x2c, 0x90, 0x0f, 0xbb, 0x16, 0xb4, 0xc2,
		0x9e, 0x52, 0x4e, 0xab, 0xaf, 0xfa, 0x7d, 0xfe, 0x07, 0x83, 0x5a, 0xa3,
		0xdb, 0x1b, 0x71, 0x2e, 0x7a, 0xc8, 0xc4, 0x8b, 0x7f, 0x19, 0x02, 0x35,
		0x0c, 0xe9, 0x47, 0xda, 0x30, 0xc9, 0x4a, 0x5e, 0xc8, 0x39, 0x97, 0x73,
		0xc9, 0xd7, 0x80, 0x0d, 0xe5, 0x45, 0xc2, 0x62, 0x0e, 0x5b, 0x74, 0x87,
		0xdc, 0x2f, 0x58, 0x86, 0x04, 0x3f, 0xf9, 0x23, 0x84, 0x13, 0xa9, 0x26,
		0x80, 0x8c, 0x14, 0xfb, 0x01, 0xd0, 0x88, 0xa6, 0x43, 0x46, 0x77, 0x18,
		0x4c, 0xbc, 0xa0, 0xd0, 0x6f, 0xbd, 0xdf, 0x08, 0x9e, 0x2e, 0x43, 0xd0,
		0xd2, 0xa9, 0x40, 0xdf, 0x2c, 0x45, 0x71, 0xf0, 0x1a, 0xf5, 0x73, 0x5e,
		0x70, 0x8a, 0xac, 0xa3, 0xf1, 0x6c, 0x49, 0xae, 0x2b, 0x83, 0xed, 0x61,
		0x8d, 0x73, 0xc4, 0xbf, 0x37, 0x36, 0x1b, 0xb5, 0x0f, 0xdb, 0xde, 0xd0,
		0xb4, 0x43, 0x80, 0x2f, 0xe0, 0x4c, 0xa8, 0x14, 0x83, 0x2e, 0x24, 0x09,
		0x74, 0x63, 0x79, 0x3e, 0xa4, 0x52, 0x2a, 0xfb, 0x82, 0xcb, 0x4d, 0x91,
		0x81, 0x88, 0x70, 0xc1, 0x4f, 0xd0, 0x32, 0xf0, 0xd4, 0x56, 0x65, 0x40,
		0x1e, 0x02, 0x71, 0x03, 0x67, 0x1b, 0x95, 0xe9, 0x7f, 0x86, 0xb8, 0x53,
		0x30, 0x17, 0xe2, 0xa6, 0x0f, 0xe2, 0x05, 0xfd, 0x99, 0x46, 0xdf, 0x6b,
		0xb2, 0x12, 0xdd, 0x1b, 0x6d, 0x7e, 0xa3, 0xc9, 0x5d, 0x10, 0xfd, 0x8c,
		0x21, 0xff, 0x4b, 0x35, 0x26, 0x41, 0xa7, 0x86, 0xb7, 0x4b, 0xb4, 0x1b,
		0xb3, 0xad, 0x63, 0x42, 0xfd, 0xe1, 0x13, 0x0a, 0x6f, 0xce, 0x02, 0x1b,
		0x3e, 0x89, 0x66, 0x14, 0x4f, 0xd3, 0xab, 0x14, 0xab, 0x4c, 0x24, 0x82,
		0x17, 0xa4, 0x4c, 0x95, 0xe9, 0x88, 0x77, 0x40, 0x67, 0x4a, 0x38, 0x2b,
		0x39, 0xd5, 0x03, 0x11, 0x75, 0xa7, 0xf0, 0x76, 0x92, 0x6e, 0x0b, 0x4e,
		0x51, 0x41, 0xe6, 0xd3, 0xfc, 0x1f, 0x02, 0xd6, 0x56, 0xdc, 0x92, 0xab,
		0x21, 0xd0, 0xbf, 0x84, 0x4f, 0x23, 0x90, 0xa0, 0x7a, 0xf1, 0x3d, 0x83,
		0x87, 0x60, 0x73, 0x1e, 0x82, 0xac, 0x3a, 0xb9, 0xb2, 0xb7, 0x76, 0xba,
		0x93, 0xfb, 0xa2, 0x11, 0xe6, 0x4d, 0x2c, 0x15, 0x3a, 0x27, 0x07, 0x7c,
		0xb3, 0xa1, 0x71, 0x14, 0x6a, 0x22, 0xa2, 0x98, 0xa5, 0x82, 0x95, 0xaf,
		0x5a, 0x58, 0x1b, 0xdd, 0xd1, 0x9a, 0x2e, 0x9e, 0x13, 0x75, 0x37, 0x5a,
		0x73, 0xbf, 0x5a, 0x90, 0x4f, 0xcd, 0x87, 0xf6, 0xe6, 0xb5, 0x56, 0xc4,
		0xd2, 0x17, 0x10, 0x4b, 0xac, 0x75, 0xd4, 0x08, 0x45, 0x0c, 0x7b, 0x66,
		0xe9, 0x86, 0x77, 0xcc, 0xdf, 0x87, 0x3c, 0x5b, 0x0a, 0x85, 0xa7, 0x36,
		0xfd, 0x35, 0x17, 0x59, 0x9f, 0x65, 0x13, 0x65, 0x00, 0xa4, 0xdb, 0xf2,
		0xf0, 0xca, 0x56, 0x4d, 0x87, 0x18, 0xce, 0xf6, 0xd5, 0x35, 0xb0, 0xf3,
		0xe3, 0x07, 0xcd, 0x02, 0xb9, 0x44, 0x68, 0x2c, 0x90, 0x1c, 0x60, 0xa6,
		0xba, 0x0f, 0x31, 0x5d, 0xab, 0xd4, 0x1c, 0x85, 0x4a, 0x3c, 0x7e, 0x59,
		0x17, 0x56, 0x4c, 0x2f, 0x5a, 0x3c, 0x2a, 0x56, 0xa5, 0x15, 0xd3, 0x8b,
		0x16, 0x7f, 0xf8, 0x22, 0xd2, 0xe5, 0xd0, 0x88, 0xd5, 0x0b, 0xc9, 0x8f,
		0x41, 0x9f, 0xc4, 0xb8, 0x41, 0x6d, 0xb2, 0x10, 0x38, 0xc6, 0x32, 0xed,
		0x0e, 0x81, 0x61, 0x04, 0x88, 0xa2, 0xc8, 0xb6, 0x71, 0x5b, 0x0f, 0xb4,
		0xca, 0x4a, 0x24, 0x70, 0xaa, 0x62, 0xc2, 0xd5, 0x15, 0x64, 0x22, 0x05,
		0x9d, 0xd2, 0x41, 0xac, 0x57, 0x9a, 0x9a, 0x83, 0x43, 0xfd, 0x18, 0x47,
		0x99, 0x4d, 0x1f, 0xdc, 0x49, 0x88, 0x23, 0xfb, 0x52, 0xaf, 0x62, 0x4c,
		0x63, 0xd5, 0xaa, 0xa9, 0x2a, 0xa7, 0x4a, 0x44, 0x55, 0x90, 0xb2, 0x09,
		0x4d, 0xd5, 0x28, 0x99, 0x4a, 0x3b, 0x20, 0x82, 0x55, 0xc0, 0xf1, 0x5e,
		0xf9, 0x03, 0x40, 0x36, 0xeb, 0x96, 0x0a, 0x84, 0xde, 0xca, 0x0c, 0x0f,
		0xb8, 0x65, 0x80, 0xc3, 0x02, 0x66, 0xfb, 0xef, 0x70, 0xe0, 0xb4, 0xc7,
		0x79, 0x2f, 0xd3, 0x1a, 0x71, 0xda, 0x7c, 0x6b, 0xc4, 0x6b, 0xb3, 0xae,
		0x11, 0xb7, 0xc5, 0x3d, 0xfd, 0x57, 0xd5, 0x8f, 0x6e, 0xdd, 0x8f, 0x60,
		0xe5, 0xa8, 0xf4, 0xdd, 0xfd, 0xc7, 0xe5, 0xde, 0xe1, 0x1d, 0xd3, 0xfd,
		0x72, 0x5b, 0xb5, 0xaf, 0x51, 0x2a, 0xde, 0x50, 0xff, 0x68, 0x89, 0x6a,
		0x5c, 0x67, 0x55, 0x7b, 0x2b, 0xda, 0x5b, 0xcd, 0xde, 0x4a, 0xf6, 0x55,
		0xb1, 0x3a, 0x76, 0x90, 0x47, 0x04, 0x1c, 0xf7, 0x20, 0x5d, 0x31, 0xb7,
		0x54, 0x71, 0xa4, 0x92, 0x3a, 0xc2, 0xd7, 0x67, 0xf6, 0xed, 0x89, 0x77,
		0x38, 0xc4, 0xb9, 0x37, 0xce, 0x68, 0xf0, 0x07, 0x83, 0xe6, 0x10, 0xe9,
		0x4a, 0xf7, 0x4c, 0x58, 0x0d, 0xe2, 0xb8, 0xa4, 0xa8, 0xb6, 0x3d, 0x39,
		0xa9, 0x58, 0x87, 0x7b, 0xa2, 0x4f, 0x01, 0xf4, 0xe4, 0x7e, 0x19, 0xb8,
		0xfe, 0xe8, 0x24, 0x6b, 0x1d, 0x11, 0x47, 0x38, 0xbf, 0x63, 0x05, 0xc7,
		0x4f, 0xc6, 0xc0, 0x39, 0x41, 0x9b, 0x60, 0x2d, 0xd7, 0x3c, 0x75, 0xa0,
		0xc3, 0xf9, 0x79, 0xfb, 0x40, 0x6f, 0x9d, 0x83, 0x87, 0x46, 0xee, 0x39,
		0x2d, 0xc9, 0xcf, 0xbe, 0xc3, 0xd2, 0x45, 0x67, 0x85, 0xdb, 0x6b, 0x81,
		0x4f, 0xd9, 0xca, 0x90, 0xde, 0xbc, 0x6d, 0x7f, 0xa3, 0xf3, 0x73, 0x08,
		0xe4, 0x32, 0xd4, 0x2b, 0x48, 0xd7, 0x2a, 0x84, 0xbb, 0xfa, 0xbf, 0x0a,
		0x86, 0x06, 0x85, 0x15, 0x60, 0xac, 0x63, 0xaa, 0xd7, 0x7d, 0x6c, 0x3b,
		0x29, 0xec, 0x3d, 0xb5, 0xdd, 0x54, 0x1a, 0x0b, 0xdb, 0xe9, 0x17, 0x1c,
		0xb9, 0x38, 0x84, 0x7b, 0xfa, 0xd5, 0xf0, 0xdf, 0xc6, 0xdc, 0xb8, 0x78,
		0xa9, 0x3e, 0xa9, 0x2c, 0xda, 0x9f, 0x08, 0xa5, 0xd3, 0xea, 0x10, 0xfe,
		0xff, 0x4b, 0x61, 0x59, 0x79, 0xf5, 0x55, 0x70, 0xe7, 0x2e, 0xd8, 0xba,
		0x9b, 0x1f, 0x51, 0xf5, 0x03, 0x2e, 0xf0, 0x4d, 0x1e, 0xd3, 0x6d, 0x61,
		0xd0, 0xb0, 0x1c, 0x84, 0x60, 0x04, 0xb4, 0xdd, 0x91, 0x00, 0xdf, 0x44,
		0xf6, 0x27, 0x01, 0x77, 0x9c, 0xed, 0x7c, 0xa1, 0x99, 0xc7, 0x7f, 0x03,
		0x00, 0x00, 0xff, 0xff, 0xee, 0xe3, 0x46, 0x2b, 0xd6, 0x12, 0x00, 0x00,
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
			"schema.tmpl": &_bintree_t{sqlc_tmpl_schema_tmpl, map[string]*_bintree_t{
			}},
			"fields.tmpl": &_bintree_t{sqlc_tmpl_fields_tmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}
