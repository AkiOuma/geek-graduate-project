// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: api/user/v1/user_error.proto

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

type UserServiceErrorReason int32

const (
	UserServiceErrorReason_UNKNOWN_ERROR  UserServiceErrorReason = 0
	UserServiceErrorReason_USER_EXIST     UserServiceErrorReason = 1
	UserServiceErrorReason_USER_NOT_FOUND UserServiceErrorReason = 2
)

// Enum value maps for UserServiceErrorReason.
var (
	UserServiceErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "USER_EXIST",
		2: "USER_NOT_FOUND",
	}
	UserServiceErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":  0,
		"USER_EXIST":     1,
		"USER_NOT_FOUND": 2,
	}
)

func (x UserServiceErrorReason) Enum() *UserServiceErrorReason {
	p := new(UserServiceErrorReason)
	*p = x
	return p
}

func (x UserServiceErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserServiceErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_user_error_proto_enumTypes[0].Descriptor()
}

func (UserServiceErrorReason) Type() protoreflect.EnumType {
	return &file_api_user_v1_user_error_proto_enumTypes[0]
}

func (x UserServiceErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserServiceErrorReason.Descriptor instead.
func (UserServiceErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_user_error_proto_rawDescGZIP(), []int{0}
}

var File_api_user_v1_user_error_proto protoreflect.FileDescriptor

var file_api_user_v1_user_error_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x55, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x02, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x28, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2d,
	0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_user_error_proto_rawDescOnce sync.Once
	file_api_user_v1_user_error_proto_rawDescData = file_api_user_v1_user_error_proto_rawDesc
)

func file_api_user_v1_user_error_proto_rawDescGZIP() []byte {
	file_api_user_v1_user_error_proto_rawDescOnce.Do(func() {
		file_api_user_v1_user_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_user_error_proto_rawDescData)
	})
	return file_api_user_v1_user_error_proto_rawDescData
}

var file_api_user_v1_user_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_v1_user_error_proto_goTypes = []interface{}{
	(UserServiceErrorReason)(0), // 0: api.user.v1.UserServiceErrorReason
}
var file_api_user_v1_user_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_user_error_proto_init() }
func file_api_user_v1_user_error_proto_init() {
	if File_api_user_v1_user_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_v1_user_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_v1_user_error_proto_goTypes,
		DependencyIndexes: file_api_user_v1_user_error_proto_depIdxs,
		EnumInfos:         file_api_user_v1_user_error_proto_enumTypes,
	}.Build()
	File_api_user_v1_user_error_proto = out.File
	file_api_user_v1_user_error_proto_rawDesc = nil
	file_api_user_v1_user_error_proto_goTypes = nil
	file_api_user_v1_user_error_proto_depIdxs = nil
}
