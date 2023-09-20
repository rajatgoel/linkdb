// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: linkdb/v1/messages.proto

package linkdbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DiskHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Written when the disk is initially formatted to make sure
	// these parameters do not change when disk is re-opened.
	BlockSize int64      `protobuf:"varint,1,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	Capacity  int64      `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Metadata  *anypb.Any `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *DiskHeader) Reset() {
	*x = DiskHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkdb_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskHeader) ProtoMessage() {}

func (x *DiskHeader) ProtoReflect() protoreflect.Message {
	mi := &file_linkdb_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskHeader.ProtoReflect.Descriptor instead.
func (*DiskHeader) Descriptor() ([]byte, []int) {
	return file_linkdb_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *DiskHeader) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *DiskHeader) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *DiskHeader) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset        int64      `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Key           string     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Version       int64      `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	ValueLength   int64      `protobuf:"varint,4,opt,name=value_length,json=valueLength,proto3" json:"value_length,omitempty"`       // 0 if no value
	ValueChecksum uint64     `protobuf:"varint,5,opt,name=value_checksum,json=valueChecksum,proto3" json:"value_checksum,omitempty"` // value checksum
	Metadata      *anypb.Any `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkdb_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_linkdb_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_linkdb_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *KeyValue) GetValueLength() int64 {
	if x != nil {
		return x.ValueLength
	}
	return 0
}

func (x *KeyValue) GetValueChecksum() uint64 {
	if x != nil {
		return x.ValueChecksum
	}
	return 0
}

func (x *KeyValue) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Gap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset    int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	GapLength int64 `protobuf:"varint,2,opt,name=gap_length,json=gapLength,proto3" json:"gap_length,omitempty"`
}

func (x *Gap) Reset() {
	*x = Gap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkdb_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gap) ProtoMessage() {}

func (x *Gap) ProtoReflect() protoreflect.Message {
	mi := &file_linkdb_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gap.ProtoReflect.Descriptor instead.
func (*Gap) Descriptor() ([]byte, []int) {
	return file_linkdb_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Gap) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Gap) GetGapLength() int64 {
	if x != nil {
		return x.GapLength
	}
	return 0
}

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Wrapped:
	//
	//	*Wrapper_KeyValue
	//	*Wrapper_Gap
	Wrapped isWrapper_Wrapped `protobuf_oneof:"wrapped"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkdb_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_linkdb_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_linkdb_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (m *Wrapper) GetWrapped() isWrapper_Wrapped {
	if m != nil {
		return m.Wrapped
	}
	return nil
}

func (x *Wrapper) GetKeyValue() *KeyValue {
	if x, ok := x.GetWrapped().(*Wrapper_KeyValue); ok {
		return x.KeyValue
	}
	return nil
}

func (x *Wrapper) GetGap() *Gap {
	if x, ok := x.GetWrapped().(*Wrapper_Gap); ok {
		return x.Gap
	}
	return nil
}

type isWrapper_Wrapped interface {
	isWrapper_Wrapped()
}

type Wrapper_KeyValue struct {
	KeyValue *KeyValue `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3,oneof"`
}

type Wrapper_Gap struct {
	Gap *Gap `protobuf:"bytes,2,opt,name=gap,proto3,oneof"`
}

func (*Wrapper_KeyValue) isWrapper_Wrapped() {}

func (*Wrapper_Gap) isWrapper_Wrapped() {}

var File_linkdb_v1_messages_proto protoreflect.FileDescriptor

var file_linkdb_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x69, 0x6e, 0x6b,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x79, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xca, 0x01, 0x0a, 0x08,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x25, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x03, 0x47, 0x61, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x70, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x61, 0x70,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x6c, 0x0a, 0x07, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x67, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x61, 0x70, 0x48, 0x00, 0x52, 0x03, 0x67, 0x61, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x42, 0x97, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6a, 0x61, 0x74, 0x67, 0x6f, 0x65, 0x6c, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x64, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x2f,
	0x76, 0x31, 0x3b, 0x6c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4c, 0x58,
	0x58, 0xaa, 0x02, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09,
	0x4c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4c, 0x69, 0x6e, 0x6b,
	0x64, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x64, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_linkdb_v1_messages_proto_rawDescOnce sync.Once
	file_linkdb_v1_messages_proto_rawDescData = file_linkdb_v1_messages_proto_rawDesc
)

func file_linkdb_v1_messages_proto_rawDescGZIP() []byte {
	file_linkdb_v1_messages_proto_rawDescOnce.Do(func() {
		file_linkdb_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_linkdb_v1_messages_proto_rawDescData)
	})
	return file_linkdb_v1_messages_proto_rawDescData
}

var file_linkdb_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_linkdb_v1_messages_proto_goTypes = []interface{}{
	(*DiskHeader)(nil), // 0: linkdb.v1.DiskHeader
	(*KeyValue)(nil),   // 1: linkdb.v1.KeyValue
	(*Gap)(nil),        // 2: linkdb.v1.Gap
	(*Wrapper)(nil),    // 3: linkdb.v1.Wrapper
	(*anypb.Any)(nil),  // 4: google.protobuf.Any
}
var file_linkdb_v1_messages_proto_depIdxs = []int32{
	4, // 0: linkdb.v1.DiskHeader.metadata:type_name -> google.protobuf.Any
	4, // 1: linkdb.v1.KeyValue.metadata:type_name -> google.protobuf.Any
	1, // 2: linkdb.v1.Wrapper.key_value:type_name -> linkdb.v1.KeyValue
	2, // 3: linkdb.v1.Wrapper.gap:type_name -> linkdb.v1.Gap
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_linkdb_v1_messages_proto_init() }
func file_linkdb_v1_messages_proto_init() {
	if File_linkdb_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_linkdb_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskHeader); i {
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
		file_linkdb_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_linkdb_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gap); i {
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
		file_linkdb_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
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
	file_linkdb_v1_messages_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Wrapper_KeyValue)(nil),
		(*Wrapper_Gap)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_linkdb_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_linkdb_v1_messages_proto_goTypes,
		DependencyIndexes: file_linkdb_v1_messages_proto_depIdxs,
		MessageInfos:      file_linkdb_v1_messages_proto_msgTypes,
	}.Build()
	File_linkdb_v1_messages_proto = out.File
	file_linkdb_v1_messages_proto_rawDesc = nil
	file_linkdb_v1_messages_proto_goTypes = nil
	file_linkdb_v1_messages_proto_depIdxs = nil
}