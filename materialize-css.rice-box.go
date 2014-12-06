package docor

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `materialize.css`,
		FileModTime: time.Unix(1417800419, 0),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `materialize.min.css`,
		FileModTime: time.Unix(1417632188, 0),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `roboto.css`,
		FileModTime: time.Unix(1417799022, 0),
		Content:     string([]byte{0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x54, 0x68, 0x69, 0x6e, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x54, 0x68, 0x69, 0x6e, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x4a, 0x7a, 0x6f, 0x36, 0x32, 0x49, 0x33, 0x39, 0x6a, 0x63, 0x30, 0x67, 0x51, 0x52, 0x72, 0x62, 0x6e, 0x64, 0x4e, 0x36, 0x6e, 0x66, 0x65, 0x73, 0x5a, 0x57, 0x32, 0x78, 0x4f, 0x51, 0x2d, 0x78, 0x73, 0x4e, 0x71, 0x4f, 0x34, 0x37, 0x6d, 0x35, 0x35, 0x44, 0x41, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x48, 0x67, 0x6f, 0x31, 0x33, 0x6b, 0x2d, 0x74, 0x66, 0x53, 0x70, 0x6e, 0x30, 0x71, 0x69, 0x31, 0x53, 0x46, 0x64, 0x55, 0x66, 0x61, 0x43, 0x57, 0x63, 0x79, 0x6e, 0x66, 0x5f, 0x63, 0x44, 0x78, 0x58, 0x77, 0x43, 0x4c, 0x78, 0x69, 0x69, 0x78, 0x47, 0x31, 0x63, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x34, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x7a, 0x4e, 0x37, 0x47, 0x42, 0x46, 0x77, 0x66, 0x4d, 0x50, 0x34, 0x75, 0x41, 0x36, 0x41, 0x52, 0x30, 0x48, 0x43, 0x6f, 0x4c, 0x51, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x52, 0x78, 0x5a, 0x4a, 0x64, 0x6e, 0x7a, 0x65, 0x6f, 0x33, 0x52, 0x35, 0x7a, 0x53, 0x65, 0x78, 0x67, 0x65, 0x38, 0x55, 0x55, 0x61, 0x43, 0x57, 0x63, 0x79, 0x6e, 0x66, 0x5f, 0x63, 0x44, 0x78, 0x58, 0x77, 0x43, 0x4c, 0x78, 0x69, 0x69, 0x78, 0x47, 0x31, 0x63, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x37, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x42, 0x6f, 0x6c, 0x64, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x42, 0x6f, 0x6c, 0x64, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x64, 0x2d, 0x36, 0x49, 0x59, 0x70, 0x6c, 0x4f, 0x46, 0x6f, 0x63, 0x43, 0x61, 0x63, 0x4b, 0x7a, 0x78, 0x77, 0x58, 0x53, 0x4f, 0x4b, 0x43, 0x57, 0x63, 0x79, 0x6e, 0x66, 0x5f, 0x63, 0x44, 0x78, 0x58, 0x77, 0x43, 0x4c, 0x78, 0x69, 0x69, 0x78, 0x47, 0x31, 0x63, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x54, 0x68, 0x69, 0x6e, 0x20, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x54, 0x68, 0x69, 0x6e, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x31, 0x32, 0x6d, 0x45, 0x34, 0x6a, 0x66, 0x4d, 0x53, 0x42, 0x54, 0x6d, 0x67, 0x2d, 0x38, 0x31, 0x45, 0x69, 0x53, 0x2d, 0x59, 0x53, 0x33, 0x55, 0x53, 0x42, 0x6e, 0x53, 0x76, 0x70, 0x6b, 0x6f, 0x70, 0x51, 0x61, 0x55, 0x52, 0x2d, 0x32, 0x72, 0x37, 0x69, 0x55, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x37, 0x6d, 0x38, 0x6c, 0x37, 0x54, 0x6c, 0x46, 0x4f, 0x2d, 0x53, 0x33, 0x56, 0x6b, 0x68, 0x48, 0x75, 0x52, 0x30, 0x61, 0x74, 0x35, 0x30, 0x45, 0x41, 0x56, 0x78, 0x74, 0x30, 0x47, 0x30, 0x62, 0x69, 0x45, 0x6e, 0x74, 0x70, 0x34, 0x33, 0x51, 0x74, 0x36, 0x45, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x34, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x57, 0x34, 0x77, 0x44, 0x73, 0x42, 0x55, 0x6c, 0x75, 0x79, 0x77, 0x30, 0x74, 0x4b, 0x33, 0x74, 0x79, 0x6b, 0x68, 0x58, 0x45, 0x66, 0x65, 0x73, 0x5a, 0x57, 0x32, 0x78, 0x4f, 0x51, 0x2d, 0x78, 0x73, 0x4e, 0x71, 0x4f, 0x34, 0x37, 0x6d, 0x35, 0x35, 0x44, 0x41, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x20, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x4f, 0x4c, 0x66, 0x66, 0x47, 0x42, 0x54, 0x61, 0x46, 0x30, 0x58, 0x46, 0x4f, 0x57, 0x31, 0x67, 0x6e, 0x75, 0x48, 0x46, 0x30, 0x5a, 0x30, 0x45, 0x41, 0x56, 0x78, 0x74, 0x30, 0x47, 0x30, 0x62, 0x69, 0x45, 0x6e, 0x74, 0x70, 0x34, 0x33, 0x51, 0x74, 0x36, 0x45, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0x40, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x27, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3a, 0x20, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x3b, 0xa, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x37, 0x30, 0x30, 0x3b, 0xa, 0x20, 0x20, 0x73, 0x72, 0x63, 0x3a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x20, 0x42, 0x6f, 0x6c, 0x64, 0x20, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x28, 0x27, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2d, 0x42, 0x6f, 0x6c, 0x64, 0x49, 0x74, 0x61, 0x6c, 0x69, 0x63, 0x27, 0x29, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x74, 0x36, 0x4e, 0x64, 0x34, 0x63, 0x66, 0x50, 0x52, 0x68, 0x5a, 0x50, 0x34, 0x34, 0x51, 0x35, 0x51, 0x41, 0x6a, 0x63, 0x43, 0x35, 0x30, 0x45, 0x41, 0x56, 0x78, 0x74, 0x30, 0x47, 0x30, 0x62, 0x69, 0x45, 0x6e, 0x74, 0x70, 0x34, 0x33, 0x51, 0x74, 0x36, 0x45, 0x2e, 0x74, 0x74, 0x66, 0x29, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x27, 0x74, 0x72, 0x75, 0x65, 0x74, 0x79, 0x70, 0x65, 0x27, 0x29, 0x3b, 0xa, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1417800419, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // materialize.css
			file3, // materialize.min.css
			file4, // roboto.css

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`materialize/css`, &embedded.EmbeddedBox{
		Name: `materialize/css`,
		Time: time.Unix(1417800434, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"materialize.css":     file2,
			"materialize.min.css": file3,
			"roboto.css":          file4,
		},
	})
}