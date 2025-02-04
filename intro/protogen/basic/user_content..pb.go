// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: proto/basic/user_content..proto

package basic

import (
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

type UserContent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserContentId uint32                 `protobuf:"varint,1,opt,name=user_content_id,proto3" json:"user_content_id,omitempty"`
	Slug          string                 `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	HtmlContent   string                 `protobuf:"bytes,4,opt,name=html_content,proto3" json:"html_content,omitempty"`
	AuthorId      uint32                 `protobuf:"varint,5,opt,name=author_id,proto3" json:"author_id,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserContent) Reset() {
	*x = UserContent{}
	mi := &file_proto_basic_user_content__proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserContent) ProtoMessage() {}

func (x *UserContent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_content__proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserContent.ProtoReflect.Descriptor instead.
func (*UserContent) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_content__proto_rawDescGZIP(), []int{0}
}

func (x *UserContent) GetUserContentId() uint32 {
	if x != nil {
		return x.UserContentId
	}
	return 0
}

func (x *UserContent) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *UserContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserContent) GetHtmlContent() string {
	if x != nil {
		return x.HtmlContent
	}
	return ""
}

func (x *UserContent) GetAuthorId() uint32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *UserContent) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_proto_basic_user_content__proto protoreflect.FileDescriptor

var file_proto_basic_user_content__proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x74, 0x6d,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x61, 0x72, 0x63, 0x69, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x3b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_user_content__proto_rawDescOnce sync.Once
	file_proto_basic_user_content__proto_rawDescData = file_proto_basic_user_content__proto_rawDesc
)

func file_proto_basic_user_content__proto_rawDescGZIP() []byte {
	file_proto_basic_user_content__proto_rawDescOnce.Do(func() {
		file_proto_basic_user_content__proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_user_content__proto_rawDescData)
	})
	return file_proto_basic_user_content__proto_rawDescData
}

var file_proto_basic_user_content__proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_basic_user_content__proto_goTypes = []any{
	(*UserContent)(nil), // 0: UserContent
}
var file_proto_basic_user_content__proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_basic_user_content__proto_init() }
func file_proto_basic_user_content__proto_init() {
	if File_proto_basic_user_content__proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_basic_user_content__proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_user_content__proto_goTypes,
		DependencyIndexes: file_proto_basic_user_content__proto_depIdxs,
		MessageInfos:      file_proto_basic_user_content__proto_msgTypes,
	}.Build()
	File_proto_basic_user_content__proto = out.File
	file_proto_basic_user_content__proto_rawDesc = nil
	file_proto_basic_user_content__proto_goTypes = nil
	file_proto_basic_user_content__proto_depIdxs = nil
}
