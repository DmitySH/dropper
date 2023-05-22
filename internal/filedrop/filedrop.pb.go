// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: api/dropper/filedrop.proto

package filedrop

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkData []byte `protobuf:"bytes,1,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dropper_filedrop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dropper_filedrop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_api_dropper_filedrop_proto_rawDescGZIP(), []int{0}
}

func (x *FileRequest) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

var File_api_dropper_filedrop_proto protoreflect.FileDescriptor

var file_api_dropper_filedrop_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x64, 0x72, 0x6f, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x32, 0x7e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x36, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x64,
	0x72, 0x6f, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x30,
	0x01, 0x42, 0x24, 0x5a, 0x22, 0x64, 0x6d, 0x69, 0x74, 0x79, 0x73, 0x68, 0x2f, 0x79, 0x6f, 0x75,
	0x72, 0x2d, 0x64, 0x72, 0x6f, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x64, 0x72, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dropper_filedrop_proto_rawDescOnce sync.Once
	file_api_dropper_filedrop_proto_rawDescData = file_api_dropper_filedrop_proto_rawDesc
)

func file_api_dropper_filedrop_proto_rawDescGZIP() []byte {
	file_api_dropper_filedrop_proto_rawDescOnce.Do(func() {
		file_api_dropper_filedrop_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dropper_filedrop_proto_rawDescData)
	})
	return file_api_dropper_filedrop_proto_rawDescData
}

var file_api_dropper_filedrop_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_dropper_filedrop_proto_goTypes = []interface{}{
	(*FileRequest)(nil),   // 0: filedrop.FileRequest
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_api_dropper_filedrop_proto_depIdxs = []int32{
	1, // 0: filedrop.FileDrop.Ping:input_type -> google.protobuf.Empty
	1, // 1: filedrop.FileDrop.GetFile:input_type -> google.protobuf.Empty
	1, // 2: filedrop.FileDrop.Ping:output_type -> google.protobuf.Empty
	0, // 3: filedrop.FileDrop.GetFile:output_type -> filedrop.FileRequest
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_dropper_filedrop_proto_init() }
func file_api_dropper_filedrop_proto_init() {
	if File_api_dropper_filedrop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dropper_filedrop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_dropper_filedrop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dropper_filedrop_proto_goTypes,
		DependencyIndexes: file_api_dropper_filedrop_proto_depIdxs,
		MessageInfos:      file_api_dropper_filedrop_proto_msgTypes,
	}.Build()
	File_api_dropper_filedrop_proto = out.File
	file_api_dropper_filedrop_proto_rawDesc = nil
	file_api_dropper_filedrop_proto_goTypes = nil
	file_api_dropper_filedrop_proto_depIdxs = nil
}
