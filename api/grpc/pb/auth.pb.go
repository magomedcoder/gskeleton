// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: auth.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Login struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Login) Reset() {
	*x = Login{}
	mi := &file_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type Login_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Login_Request) Reset() {
	*x = Login_Request{}
	mi := &file_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Login_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login_Request) ProtoMessage() {}

func (x *Login_Request) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login_Request.ProtoReflect.Descriptor instead.
func (*Login_Request) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Login_Request) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Login_Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Login_Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Login_Response) Reset() {
	*x = Login_Response{}
	mi := &file_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Login_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login_Response) ProtoMessage() {}

func (x *Login_Response) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login_Response.ProtoReflect.Descriptor instead.
func (*Login_Response) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Login_Response) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

const file_auth_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"auth.proto\x12\x13gskeleton.grpc.user\"x\n" +
	"\x05Login\x1aA\n" +
	"\aRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x1a,\n" +
	"\bResponse\x12 \n" +
	"\vaccessToken\x18\x01 \x01(\tR\vaccessToken2a\n" +
	"\vAuthService\x12R\n" +
	"\x05Login\x12\".gskeleton.grpc.user.Login.Request\x1a#.gskeleton.grpc.user.Login.Response\"\x00B/Z-github.com/magomedcoder/gskeleton/api/grpc/pbb\x06proto3"

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData []byte
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)))
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auth_proto_goTypes = []any{
	(*Login)(nil),          // 0: gskeleton.grpc.user.Login
	(*Login_Request)(nil),  // 1: gskeleton.grpc.user.Login.Request
	(*Login_Response)(nil), // 2: gskeleton.grpc.user.Login.Response
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: gskeleton.grpc.user.AuthService.Login:input_type -> gskeleton.grpc.user.Login.Request
	2, // 1: gskeleton.grpc.user.AuthService.Login:output_type -> gskeleton.grpc.user.Login.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
