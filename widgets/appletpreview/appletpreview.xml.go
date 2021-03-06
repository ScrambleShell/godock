package appletpreview

import (
	"bytes"
	"compress/gzip"
	"io"
)

// appletpreviewXML returns raw, uncompressed file data.
func appletpreviewXML() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x5a,
		0x41, 0x53, 0xdb, 0x3c, 0x10, 0xbd, 0xf3, 0x2b, 0xfc, 0xe9, 0xfa, 0x4d,
		0x08, 0x81, 0x1e, 0x38, 0x38, 0x61, 0xa6, 0x07, 0x28, 0x33, 0xbd, 0x95,
		0xb6, 0xc7, 0x8c, 0x22, 0x6d, 0x62, 0x35, 0x8a, 0xe4, 0x4a, 0x0a, 0x09,
		0xfd, 0xf5, 0x95, 0x71, 0x52, 0x9c, 0x44, 0x72, 0x64, 0x99, 0xa1, 0x98,
		0x72, 0x6b, 0x9d, 0xdd, 0xd5, 0xbe, 0xe7, 0xd5, 0xea, 0xad, 0x4c, 0x7a,
		0xb5, 0x5e, 0xf0, 0xe4, 0x1e, 0x94, 0x66, 0x52, 0x0c, 0xd1, 0xe0, 0xf4,
		0x0c, 0x25, 0x20, 0x88, 0xa4, 0x4c, 0xcc, 0x86, 0xe8, 0xeb, 0xdd, 0x75,
		0xef, 0x12, 0x5d, 0x8d, 0x4e, 0xd2, 0xff, 0x7a, 0xbd, 0xe4, 0x06, 0x04,
		0x28, 0x6c, 0x80, 0x26, 0x2b, 0x66, 0xb2, 0x64, 0xc6, 0x31, 0x85, 0xe4,
		0xe2, 0x74, 0x70, 0x79, 0x7a, 0x91, 0xf4, 0x7a, 0xd6, 0x88, 0x09, 0x03,
		0x6a, 0x8a, 0x09, 0x8c, 0x4e, 0x92, 0x24, 0x55, 0xf0, 0x73, 0xc9, 0x14,
		0xe8, 0x84, 0xb3, 0xc9, 0x10, 0xcd, 0xcc, 0xfc, 0x7f, 0xf4, 0xb4, 0xd0,
		0x85, 0x5d, 0xa8, 0xff, 0x68, 0x26, 0x27, 0x3f, 0x80, 0x98, 0x84, 0x70,
		0xac, 0xf5, 0x10, 0xdd, 0x98, 0xf9, 0x47, 0xb9, 0x46, 0x09, 0xa3, 0x43,
		0xb4, 0x62, 0x74, 0x06, 0x06, 0x15, 0x56, 0xd6, 0x2e, 0x57, 0x32, 0x07,
		0x65, 0x1e, 0x12, 0x81, 0x17, 0x30, 0x44, 0xf7, 0x4c, 0xb3, 0x09, 0x07,
		0x34, 0xba, 0x53, 0x4b, 0x48, 0xfb, 0xdb, 0x5f, 0xdd, 0xc6, 0x04, 0x8b,
		0xf1, 0x54, 0x92, 0xa5, 0x46, 0xa3, 0x6b, 0xcc, 0xf5, 0x51, 0x7b, 0xa9,
		0x18, 0x08, 0x83, 0x8d, 0x4d, 0x15, 0x8d, 0x6c, 0xce, 0x86, 0x11, 0xcc,
		0x8f, 0x39, 0xe9, 0x1c, 0x13, 0xcb, 0x1a, 0x1a, 0x5d, 0x1e, 0x58, 0x92,
		0x8c, 0x71, 0x5a, 0xfe, 0xdb, 0x85, 0xf8, 0x5a, 0xd9, 0x00, 0x25, 0xe6,
		0x0c, 0x2c, 0xa9, 0x0a, 0x6d, 0x6d, 0x1b, 0xe2, 0x8e, 0xc1, 0xee, 0xf2,
		0xe1, 0x78, 0x02, 0x7c, 0xbc, 0xc6, 0x9c, 0xcd, 0x2c, 0x01, 0x67, 0x21,
		0x2e, 0x3a, 0xc3, 0x54, 0xae, 0xc6, 0xe6, 0x21, 0xb7, 0xb9, 0xc9, 0xa5,
		0x71, 0xfa, 0xec, 0xf0, 0xe0, 0xe6, 0xe2, 0xcf, 0xdb, 0x9f, 0xc8, 0xf5,
		0xa7, 0x7d, 0x32, 0xa2, 0x08, 0x89, 0x25, 0x25, 0xba, 0x30, 0xfc, 0x70,
		0x8f, 0x40, 0xe6, 0x4c, 0xc0, 0x00, 0xed, 0x3b, 0x44, 0x42, 0x6e, 0x03,
		0xdb, 0x9f, 0xbd, 0x1b, 0xc1, 0xe7, 0xa2, 0x5e, 0x4a, 0x0c, 0x86, 0x99,
		0x22, 0x33, 0x87, 0x5b, 0x0b, 0x1c, 0x6d, 0xb1, 0xb8, 0xfc, 0x97, 0x1a,
		0xc6, 0x0b, 0xac, 0xe6, 0xcb, 0x3c, 0x64, 0xfd, 0xb4, 0x5f, 0x82, 0x76,
		0xfe, 0x66, 0x37, 0xfe, 0xdc, 0xee, 0xfc, 0xb0, 0x75, 0x61, 0x9d, 0x63,
		0x41, 0x23, 0x93, 0x9e, 0x32, 0xce, 0xe3, 0xe8, 0xca, 0x31, 0xa5, 0x8f,
		0xdd, 0xe9, 0x43, 0x63, 0x4f, 0xa9, 0x59, 0x59, 0xf0, 0xce, 0x46, 0x50,
		0xe5, 0xc8, 0x4b, 0x44, 0xda, 0xf7, 0x54, 0x53, 0x64, 0x95, 0x69, 0xbb,
		0x07, 0xe1, 0x0e, 0xd6, 0xa6, 0x0b, 0x95, 0x26, 0xe4, 0x58, 0x67, 0xb6,
		0x37, 0xe2, 0xd8, 0x77, 0xf7, 0x5e, 0xaa, 0xc1, 0x9e, 0x64, 0xbe, 0x39,
		0x82, 0x40, 0xd0, 0xf8, 0x42, 0x1f, 0xfc, 0xdd, 0x42, 0xbf, 0x5d, 0xe0,
		0x19, 0x54, 0x0a, 0xfd, 0x96, 0x14, 0x49, 0xbd, 0xc6, 0x5a, 0x7d, 0x03,
		0xb5, 0xf6, 0x3c, 0x15, 0x73, 0xfe, 0xdc, 0x15, 0xe3, 0x63, 0xd6, 0xcf,
		0x6a, 0x34, 0xa3, 0x71, 0x6c, 0x46, 0x1d, 0x13, 0x1e, 0x1e, 0x9c, 0x1c,
		0x44, 0xc9, 0xa7, 0xf3, 0x8e, 0xcb, 0x27, 0xbc, 0x34, 0x99, 0xdc, 0x97,
		0xbc, 0xad, 0x81, 0xb4, 0x05, 0xe3, 0xf2, 0x7f, 0x3f, 0x94, 0xba, 0xac,
		0x9f, 0xd8, 0xaf, 0x4e, 0x88, 0xf4, 0x77, 0xe9, 0xf4, 0xcf, 0x4b, 0xa7,
		0x0e, 0x1f, 0x84, 0x35, 0x5c, 0x84, 0x1f, 0x84, 0x87, 0xf8, 0x0f, 0x8c,
		0xca, 0x26, 0x90, 0x14, 0x6f, 0x6f, 0x7b, 0x77, 0xc3, 0x0c, 0x2c, 0x76,
		0x36, 0x78, 0x9a, 0x73, 0x4c, 0x20, 0x93, 0x9c, 0x82, 0xea, 0x7b, 0x63,
		0xed, 0x2f, 0x76, 0x48, 0x72, 0x63, 0x72, 0x9b, 0x91, 0xda, 0xa8, 0x79,
		0xee, 0x91, 0xb8, 0x83, 0xe5, 0xd8, 0x9d, 0xdb, 0x17, 0xa2, 0x24, 0xe7,
		0x40, 0xbf, 0x33, 0x41, 0xe5, 0x6a, 0xd3, 0x15, 0x37, 0xcf, 0x6e, 0xc5,
		0x54, 0xbe, 0xf8, 0x15, 0x5c, 0xc0, 0xdd, 0xd8, 0x37, 0x06, 0xab, 0x5c,
		0x2a, 0x53, 0x66, 0x7b, 0xbf, 0xf9, 0xdf, 0x5e, 0xb6, 0x51, 0x19, 0xc7,
		0x66, 0xed, 0xce, 0xdc, 0x9d, 0x7d, 0xf5, 0x66, 0xcf, 0x91, 0x73, 0x74,
		0xde, 0x6d, 0x72, 0x77, 0xf9, 0x36, 0xba, 0xe3, 0xf3, 0x13, 0x70, 0x84,
		0x04, 0x0a, 0xb6, 0xda, 0x58, 0x5e, 0xac, 0x52, 0x3c, 0x7b, 0x8d, 0x67,
		0xb1, 0x17, 0x96, 0x1b, 0x5a, 0x45, 0x5e, 0x54, 0xc0, 0x79, 0x90, 0xb5,
		0x46, 0xf7, 0x1c, 0x08, 0x5d, 0x31, 0x9a, 0x09, 0x06, 0x77, 0x8c, 0x95,
		0xc2, 0x2d, 0xbc, 0x6b, 0xaf, 0xdd, 0xf7, 0x5c, 0x6b, 0xf4, 0x4a, 0x19,
		0xba, 0x4e, 0xb3, 0xb8, 0x16, 0x6f, 0xa4, 0x5b, 0x5c, 0x01, 0xc2, 0xb5,
		0x8b, 0xcb, 0x3b, 0x58, 0x2b, 0x97, 0xe8, 0xeb, 0x25, 0x99, 0x4f, 0x33,
		0xbf, 0x05, 0x9d, 0xf7, 0x5a, 0x66, 0x8a, 0xca, 0xa7, 0xab, 0x5c, 0x41,
		0x71, 0x24, 0x95, 0x4f, 0x3a, 0x32, 0x1e, 0x64, 0x9b, 0xbd, 0x46, 0xa0,
		0xf8, 0x74, 0xd9, 0xd4, 0x3b, 0xe0, 0x33, 0x59, 0x18, 0xaf, 0x6e, 0x6e,
		0x2b, 0xd7, 0x80, 0x1b, 0x6e, 0xcb, 0x27, 0x1d, 0xeb, 0xa8, 0xb6, 0x9b,
		0xce, 0x98, 0x18, 0x6b, 0x83, 0xad, 0x72, 0x19, 0x0d, 0x02, 0x36, 0xb6,
		0x37, 0x06, 0x14, 0x5b, 0xad, 0x55, 0x04, 0x23, 0xf3, 0x96, 0x11, 0x26,
		0xd2, 0x18, 0xb9, 0x08, 0x0b, 0x52, 0xdf, 0x9f, 0x6b, 0x1a, 0x54, 0x98,
		0xa6, 0xdf, 0xcd, 0xd6, 0xa9, 0xef, 0x03, 0xd7, 0xeb, 0x7e, 0x43, 0xdc,
		0x0e, 0xbe, 0x07, 0x1f, 0xdf, 0x8f, 0x7a, 0xbe, 0xf4, 0xe8, 0x1a, 0x33,
		0xe3, 0xc5, 0xcf, 0x65, 0xa1, 0x33, 0x4a, 0xec, 0x58, 0xe6, 0x20, 0xcd,
		0x3f, 0x96, 0x3d, 0xc1, 0x48, 0xfb, 0x95, 0x3f, 0x16, 0xf9, 0x1d, 0x00,
		0x00, 0xff, 0xff, 0x10, 0x74, 0x72, 0xc7, 0x85, 0x22, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}
