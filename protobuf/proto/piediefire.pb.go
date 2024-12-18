// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: proto/piediefire.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_piediefire_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_piediefire_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_piediefire_proto_rawDescGZIP(), []int{0}
}

type DataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Beef          map[string]int32       `protobuf:"bytes,1,rep,name=beef,proto3" json:"beef,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	mi := &file_proto_piediefire_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_piediefire_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_proto_piediefire_proto_rawDescGZIP(), []int{1}
}

func (x *DataResponse) GetBeef() map[string]int32 {
	if x != nil {
		return x.Beef
	}
	return nil
}

var File_proto_piediefire_proto protoreflect.FileDescriptor

var file_proto_piediefire_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x65, 0x64, 0x69, 0x65, 0x66, 0x69,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x79, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x65, 0x65, 0x66, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x62, 0x65, 0x65, 0x66, 0x1a, 0x37, 0x0a, 0x09, 0x42, 0x65, 0x65,
	0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0x39, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a,
	0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_piediefire_proto_rawDescOnce sync.Once
	file_proto_piediefire_proto_rawDescData = file_proto_piediefire_proto_rawDesc
)

func file_proto_piediefire_proto_rawDescGZIP() []byte {
	file_proto_piediefire_proto_rawDescOnce.Do(func() {
		file_proto_piediefire_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_piediefire_proto_rawDescData)
	})
	return file_proto_piediefire_proto_rawDescData
}

var file_proto_piediefire_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_piediefire_proto_goTypes = []any{
	(*Empty)(nil),        // 0: data.Empty
	(*DataResponse)(nil), // 1: data.DataResponse
	nil,                  // 2: data.DataResponse.BeefEntry
}
var file_proto_piediefire_proto_depIdxs = []int32{
	2, // 0: data.DataResponse.beef:type_name -> data.DataResponse.BeefEntry
	0, // 1: data.DataService.GetData:input_type -> data.Empty
	1, // 2: data.DataService.GetData:output_type -> data.DataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_piediefire_proto_init() }
func file_proto_piediefire_proto_init() {
	if File_proto_piediefire_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_piediefire_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_piediefire_proto_goTypes,
		DependencyIndexes: file_proto_piediefire_proto_depIdxs,
		MessageInfos:      file_proto_piediefire_proto_msgTypes,
	}.Build()
	File_proto_piediefire_proto = out.File
	file_proto_piediefire_proto_rawDesc = nil
	file_proto_piediefire_proto_goTypes = nil
	file_proto_piediefire_proto_depIdxs = nil
}