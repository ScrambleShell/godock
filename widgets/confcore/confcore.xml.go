package confcore

import (
	"bytes"
	"compress/gzip"
	"io"
)

// confcoreXML returns raw, uncompressed file data.
func confcoreXML() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xbc, 0x55,
		0x4d, 0x6f, 0xdb, 0x30, 0x0c, 0xbd, 0xf7, 0x57, 0x68, 0xba, 0x0e, 0xae,
		0x97, 0xf5, 0xd2, 0x83, 0xed, 0x62, 0x28, 0xd0, 0x62, 0xd8, 0x30, 0x0c,
		0x6b, 0xb6, 0x1d, 0x0d, 0x45, 0x62, 0x63, 0x2d, 0x8a, 0xe4, 0x51, 0x74,
		0x9b, 0xfc, 0xfb, 0xc9, 0x1f, 0x49, 0x9c, 0xc4, 0x49, 0x93, 0x15, 0xdb,
		0x25, 0x81, 0x44, 0xf2, 0x91, 0x7c, 0x7c, 0xa2, 0x93, 0x9b, 0xc5, 0xdc,
		0xb0, 0x27, 0x40, 0xaf, 0x9d, 0x4d, 0xf9, 0xe8, 0xf2, 0x1d, 0x67, 0x60,
		0xa5, 0x53, 0xda, 0x4e, 0x53, 0xfe, 0x7d, 0x7c, 0x17, 0x5d, 0xf3, 0x9b,
		0xec, 0x22, 0x79, 0x13, 0x45, 0xec, 0x1e, 0x2c, 0xa0, 0x20, 0x50, 0xec,
		0x59, 0x53, 0xc1, 0xa6, 0x46, 0x28, 0x60, 0x57, 0x97, 0xa3, 0xeb, 0xcb,
		0x2b, 0x16, 0x45, 0xc1, 0x49, 0x5b, 0x02, 0x7c, 0x14, 0x12, 0xb2, 0x0b,
		0xc6, 0x12, 0x84, 0xdf, 0x95, 0x46, 0xf0, 0xcc, 0xe8, 0x49, 0xca, 0xa7,
		0x34, 0x7b, 0xcb, 0x37, 0x89, 0xae, 0x42, 0xa2, 0xb8, 0x71, 0x73, 0x93,
		0x5f, 0x20, 0x89, 0x49, 0x23, 0xbc, 0x4f, 0xf9, 0x3d, 0xcd, 0x3e, 0x6b,
		0x4f, 0x0f, 0xe4, 0x10, 0x38, 0xd3, 0x2a, 0xe5, 0x73, 0xa7, 0xc0, 0xf0,
		0xda, 0x35, 0x38, 0x4b, 0x67, 0xaa, 0xb9, 0xf5, 0xed, 0x29, 0x9c, 0xeb,
		0xb2, 0xda, 0xbb, 0xc8, 0x8a, 0x39, 0xb0, 0x4f, 0xb0, 0x6c, 0x4a, 0xe9,
		0xcc, 0xad, 0x89, 0xd1, 0xb2, 0x84, 0x50, 0x81, 0x2c, 0x04, 0x0a, 0x44,
		0xb1, 0x6c, 0x33, 0x0f, 0x02, 0x7c, 0x94, 0xce, 0x1e, 0x44, 0xb8, 0x57,
		0xb3, 0xaf, 0x7a, 0x31, 0xa9, 0x1e, 0x8f, 0x00, 0x7c, 0xa9, 0x7f, 0x5e,
		0x55, 0xc2, 0xd8, 0x39, 0x43, 0xba, 0x3c, 0x03, 0x24, 0x89, 0x7b, 0xbc,
		0x24, 0x71, 0xcb, 0xe8, 0x30, 0xb9, 0x0f, 0x12, 0x9d, 0x31, 0xa0, 0x7e,
		0x6a, 0xab, 0xdc, 0x73, 0xcb, 0xf0, 0xb3, 0x56, 0x53, 0xa0, 0x15, 0xc5,
		0x25, 0xba, 0x12, 0x90, 0x96, 0xac, 0x2e, 0x26, 0xe5, 0x4f, 0xda, 0xeb,
		0x89, 0x01, 0x9e, 0x8d, 0xb1, 0x82, 0x24, 0x5e, 0x59, 0x87, 0x9d, 0xa5,
		0xb0, 0xf9, 0xa3, 0x93, 0x95, 0xe7, 0xd9, 0x9d, 0x30, 0xfe, 0x45, 0x7f,
		0x5f, 0x88, 0x50, 0x45, 0x5e, 0xf7, 0xc5, 0x33, 0x20, 0x59, 0x80, 0x8a,
		0xb4, 0xdd, 0x8b, 0x92, 0x85, 0x36, 0x6a, 0x4d, 0xc6, 0x5e, 0x4f, 0x63,
		0x04, 0xf8, 0xa1, 0xa1, 0xeb, 0x86, 0xc2, 0x89, 0xaf, 0x9c, 0xcf, 0xec,
		0xe7, 0x85, 0x9e, 0x4e, 0x0d, 0x29, 0x84, 0x3f, 0x37, 0xa4, 0x53, 0x79,
		0xf3, 0x77, 0x52, 0x0a, 0x08, 0x8f, 0x0f, 0x7d, 0xbe, 0x6e, 0x67, 0x90,
		0xef, 0x63, 0x91, 0xd2, 0x68, 0x39, 0x13, 0xe7, 0xc5, 0x62, 0x65, 0xc0,
		0xe7, 0x45, 0x78, 0xe7, 0xa7, 0x77, 0xe6, 0x41, 0xa0, 0x2c, 0xf2, 0x56,
		0xa1, 0x3c, 0x7b, 0x3f, 0x18, 0xd3, 0x4c, 0x98, 0x35, 0xfb, 0xc3, 0x0a,
		0x13, 0x35, 0xc7, 0x3a, 0xd4, 0x84, 0x49, 0x87, 0x7d, 0xd1, 0x9b, 0xe7,
		0x21, 0x01, 0x3c, 0xac, 0x7d, 0x1b, 0x15, 0x6c, 0x42, 0xe3, 0x5e, 0x9a,
		0x78, 0x4b, 0x49, 0xbb, 0xca, 0x3a, 0xae, 0xae, 0xdb, 0xb6, 0x81, 0x06,
		0xbd, 0x6d, 0xa6, 0xde, 0x15, 0x5b, 0x95, 0x0d, 0x01, 0x0e, 0x83, 0xde,
		0x82, 0x31, 0xdf, 0xc0, 0x86, 0x41, 0x00, 0x76, 0x2b, 0xa5, 0x05, 0x0e,
		0xf7, 0x0d, 0x6c, 0xbc, 0x87, 0x21, 0x88, 0x50, 0x4f, 0x2a, 0x02, 0xbf,
		0x6b, 0xea, 0x1b, 0x3b, 0xd2, 0xcb, 0x16, 0x33, 0x1b, 0x25, 0xf1, 0xda,
		0xb4, 0x87, 0x18, 0x1f, 0x82, 0xdc, 0xe3, 0x69, 0x7b, 0xab, 0xfc, 0x13,
		0x32, 0xeb, 0xbd, 0xb9, 0x4b, 0xe6, 0x8e, 0x92, 0x48, 0x53, 0x50, 0x2b,
		0x23, 0x14, 0xd6, 0x1b, 0x41, 0xb5, 0x76, 0x53, 0xbe, 0x84, 0xf0, 0xca,
		0x3e, 0x94, 0xa5, 0x01, 0x1a, 0x52, 0xd6, 0x10, 0x8e, 0x77, 0x48, 0x9d,
		0x1e, 0x73, 0xad, 0x1a, 0x92, 0x0e, 0x04, 0xfe, 0xc5, 0x34, 0xc7, 0xb0,
		0xa0, 0xcd, 0x2c, 0xb1, 0xbb, 0xa5, 0x70, 0x3b, 0x7a, 0xed, 0x50, 0xe7,
		0x02, 0x67, 0x55, 0xd9, 0x3c, 0xa1, 0xff, 0x34, 0xd4, 0x6d, 0x87, 0x9e,
		0x71, 0x63, 0x48, 0xe2, 0xde, 0x57, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff,
		0xff, 0x64, 0x19, 0xf0, 0x3a, 0x4e, 0x08, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}
