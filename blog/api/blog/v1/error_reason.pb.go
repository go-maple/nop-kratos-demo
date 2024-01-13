// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/blog/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	ErrorReason_Success                     ErrorReason = 0
	ErrorReason_USER_NOT_FOUND              ErrorReason = 1
	ErrorReason_CONTENT_MISSING             ErrorReason = 2
	ErrorReason_Global_Service_Unauthorized ErrorReason = 3
	ErrorReason_Global_Service_Forbidden    ErrorReason = 4
	ErrorReason_Global_Service_NotFound     ErrorReason = 5
	ErrorReason_INVALID_PARAMS              ErrorReason = 100
	ErrorReason_BLOG_INVALID_ID             ErrorReason = 200
	// create article tx failed
	ErrorReason_CreateArticleFailed ErrorReason = 1000
	ErrorReason_CreateError         ErrorReason = 1001
	ErrorReason_GetUserIdFailed     ErrorReason = 30002
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:     "Success",
		1:     "USER_NOT_FOUND",
		2:     "CONTENT_MISSING",
		3:     "Global_Service_Unauthorized",
		4:     "Global_Service_Forbidden",
		5:     "Global_Service_NotFound",
		100:   "INVALID_PARAMS",
		200:   "BLOG_INVALID_ID",
		1000:  "CreateArticleFailed",
		1001:  "CreateError",
		30002: "GetUserIdFailed",
	}
	ErrorReason_value = map[string]int32{
		"Success":                     0,
		"USER_NOT_FOUND":              1,
		"CONTENT_MISSING":             2,
		"Global_Service_Unauthorized": 3,
		"Global_Service_Forbidden":    4,
		"Global_Service_NotFound":     5,
		"INVALID_PARAMS":              100,
		"BLOG_INVALID_ID":             200,
		"CreateArticleFailed":         1000,
		"CreateError":                 1001,
		"GetUserIdFailed":             30002,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_blog_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_blog_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_blog_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_api_blog_v1_error_reason_proto protoreflect.FileDescriptor

var file_api_blog_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xb6, 0x02, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x25, 0x0a, 0x1b, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x22, 0x0a, 0x18, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x46, 0x6f,
	0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12,
	0x21, 0x0a, 0x17, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45,
	0x94, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x53, 0x10, 0x64, 0x12, 0x14, 0x0a, 0x0f, 0x42, 0x4c, 0x4f, 0x47, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xc8, 0x01, 0x12, 0x18, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0xe8, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xe9, 0x07, 0x12, 0x1b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0xb2, 0xea, 0x01, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x70,
	0x6c, 0x65, 0x2f, 0x6e, 0x6f, 0x70, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blog_v1_error_reason_proto_rawDescOnce sync.Once
	file_api_blog_v1_error_reason_proto_rawDescData = file_api_blog_v1_error_reason_proto_rawDesc
)

func file_api_blog_v1_error_reason_proto_rawDescGZIP() []byte {
	file_api_blog_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_api_blog_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blog_v1_error_reason_proto_rawDescData)
	})
	return file_api_blog_v1_error_reason_proto_rawDescData
}

var file_api_blog_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_blog_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.blog.v1.ErrorReason
}
var file_api_blog_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_blog_v1_error_reason_proto_init() }
func file_api_blog_v1_error_reason_proto_init() {
	if File_api_blog_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_blog_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_blog_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_api_blog_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_api_blog_v1_error_reason_proto_enumTypes,
	}.Build()
	File_api_blog_v1_error_reason_proto = out.File
	file_api_blog_v1_error_reason_proto_rawDesc = nil
	file_api_blog_v1_error_reason_proto_goTypes = nil
	file_api_blog_v1_error_reason_proto_depIdxs = nil
}
