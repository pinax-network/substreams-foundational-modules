// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sf/substreams/gear/type/v1/types.proto

package typev1

import (
	v1 "github.com/streamingfast/substreams-foundational-modules/vara-common/pb/sf/gear/type/v1"
	_ "github.com/streamingfast/substreams-foundational-modules/vara-common/pb/sf/substreams/v1"
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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number        uint64           `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash          []byte           `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Header        *v1.Header       `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	Extrinsics    []*Extrinsic     `protobuf:"bytes,4,rep,name=extrinsics,proto3" json:"extrinsics,omitempty"`
	Events        []*Event         `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
	DigestItems   []*v1.DigestItem `protobuf:"bytes,6,rep,name=digest_items,json=digestItems,proto3" json:"digest_items,omitempty"`
	Justification []byte           `protobuf:"bytes,7,opt,name=justification,proto3" json:"justification,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Block) GetHeader() *v1.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetExtrinsics() []*Extrinsic {
	if x != nil {
		return x.Extrinsics
	}
	return nil
}

func (x *Block) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *Block) GetDigestItems() []*v1.DigestItem {
	if x != nil {
		return x.DigestItems
	}
	return nil
}

func (x *Block) GetJustification() []byte {
	if x != nil {
		return x.Justification
	}
	return nil
}

type Extrinsic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32        `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Signature *v1.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Call      *Call         `protobuf:"bytes,3,opt,name=call,proto3" json:"call,omitempty"`
}

func (x *Extrinsic) Reset() {
	*x = Extrinsic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extrinsic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extrinsic) ProtoMessage() {}

func (x *Extrinsic) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extrinsic.ProtoReflect.Descriptor instead.
func (*Extrinsic) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *Extrinsic) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Extrinsic) GetSignature() *v1.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Extrinsic) GetCall() *Call {
	if x != nil {
		return x.Call
	}
	return nil
}

type Call struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fields *Fields `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Call) Reset() {
	*x = Call{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Call) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Call) ProtoMessage() {}

func (x *Call) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Call.ProtoReflect.Descriptor instead.
func (*Call) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Call) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Call) GetFields() *Fields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fields *Fields `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetFields() *Fields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Value_Int32
	//	*Value_Bigdecimal
	//	*Value_Bigint
	//	*Value_String_
	//	*Value_Bytes
	//	*Value_Bool
	//	*Value_NilValue
	//	*Value_Array
	//	*Value_Fields
	Type isValue_Type `protobuf_oneof:"type"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{4}
}

func (m *Value) GetType() isValue_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Value) GetInt32() int32 {
	if x, ok := x.GetType().(*Value_Int32); ok {
		return x.Int32
	}
	return 0
}

func (x *Value) GetBigdecimal() string {
	if x, ok := x.GetType().(*Value_Bigdecimal); ok {
		return x.Bigdecimal
	}
	return ""
}

func (x *Value) GetBigint() string {
	if x, ok := x.GetType().(*Value_Bigint); ok {
		return x.Bigint
	}
	return ""
}

func (x *Value) GetString_() string {
	if x, ok := x.GetType().(*Value_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Value) GetBytes() []byte {
	if x, ok := x.GetType().(*Value_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *Value) GetBool() bool {
	if x, ok := x.GetType().(*Value_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Value) GetNilValue() bool {
	if x, ok := x.GetType().(*Value_NilValue); ok {
		return x.NilValue
	}
	return false
}

func (x *Value) GetArray() *Array {
	if x, ok := x.GetType().(*Value_Array); ok {
		return x.Array
	}
	return nil
}

func (x *Value) GetFields() *Fields {
	if x, ok := x.GetType().(*Value_Fields); ok {
		return x.Fields
	}
	return nil
}

type isValue_Type interface {
	isValue_Type()
}

type Value_Int32 struct {
	Int32 int32 `protobuf:"varint,1,opt,name=int32,proto3,oneof"`
}

type Value_Bigdecimal struct {
	Bigdecimal string `protobuf:"bytes,2,opt,name=bigdecimal,proto3,oneof"`
}

type Value_Bigint struct {
	Bigint string `protobuf:"bytes,3,opt,name=bigint,proto3,oneof"`
}

type Value_String_ struct {
	String_ string `protobuf:"bytes,4,opt,name=string,proto3,oneof"`
}

type Value_Bytes struct {
	Bytes []byte `protobuf:"bytes,5,opt,name=bytes,proto3,oneof"`
}

type Value_Bool struct {
	Bool bool `protobuf:"varint,6,opt,name=bool,proto3,oneof"`
}

type Value_NilValue struct {
	NilValue bool `protobuf:"varint,19,opt,name=nil_value,json=nilValue,proto3,oneof"`
}

type Value_Array struct {
	Array *Array `protobuf:"bytes,20,opt,name=array,proto3,oneof"`
}

type Value_Fields struct {
	Fields *Fields `protobuf:"bytes,21,opt,name=fields,proto3,oneof"`
}

func (*Value_Int32) isValue_Type() {}

func (*Value_Bigdecimal) isValue_Type() {}

func (*Value_Bigint) isValue_Type() {}

func (*Value_String_) isValue_Type() {}

func (*Value_Bytes) isValue_Type() {}

func (*Value_Bool) isValue_Type() {}

func (*Value_NilValue) isValue_Type() {}

func (*Value_Array) isValue_Type() {}

func (*Value_Fields) isValue_Type() {}

type Fields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[string]*Value `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Fields) Reset() {
	*x = Fields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fields) ProtoMessage() {}

func (x *Fields) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fields.ProtoReflect.Descriptor instead.
func (*Fields) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{5}
}

func (x *Fields) GetMap() map[string]*Value {
	if x != nil {
		return x.Map
	}
	return nil
}

type Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Value `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Array) Reset() {
	*x = Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Array) ProtoMessage() {}

func (x *Array) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_gear_type_v1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Array.ProtoReflect.Descriptor instead.
func (*Array) Descriptor() ([]byte, []int) {
	return file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP(), []int{6}
}

func (x *Array) GetItems() []*Value {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_sf_substreams_gear_type_v1_types_proto protoreflect.FileDescriptor

var file_sf_substreams_gear_type_v1_types_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x67, 0x65, 0x61, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x73, 0x66, 0x2f, 0x67, 0x65, 0x61, 0x72, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcc, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e,
	0x73, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69,
	0x63, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x73, 0x12, 0x39, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61,
	0x72, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x66, 0x2e, 0x67, 0x65, 0x61, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6a, 0x75, 0x73, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x95,
	0x01, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x67,
	0x65, 0x61, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x34, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76,
	0x61, 0x72, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x22, 0x56, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x57,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x66,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xc3, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x20, 0x0a, 0x0a, 0x62, 0x69, 0x67,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x62, 0x69, 0x67, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x62,
	0x69, 0x67, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x62,
	0x69, 0x67, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a,
	0x09, 0x6e, 0x69, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x08, 0x6e, 0x69, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x05,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x66,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00,
	0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x48, 0x00, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa2, 0x01,
	0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x59, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x40, 0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0xa4, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x61, 0x72, 0x61, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x76, 0x61, 0x72, 0x61, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x67, 0x65, 0x61,
	0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x53, 0x53, 0x56, 0x54, 0xaa, 0x02, 0x1a, 0x53, 0x66, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x61, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x5c, 0x56, 0x61, 0x72, 0x61, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x26, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x5c, 0x56, 0x61, 0x72, 0x61, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x53, 0x66, 0x3a,
	0x3a, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x61, 0x72,
	0x61, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_gear_type_v1_types_proto_rawDescOnce sync.Once
	file_sf_substreams_gear_type_v1_types_proto_rawDescData = file_sf_substreams_gear_type_v1_types_proto_rawDesc
)

func file_sf_substreams_gear_type_v1_types_proto_rawDescGZIP() []byte {
	file_sf_substreams_gear_type_v1_types_proto_rawDescOnce.Do(func() {
		file_sf_substreams_gear_type_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_gear_type_v1_types_proto_rawDescData)
	})
	return file_sf_substreams_gear_type_v1_types_proto_rawDescData
}

var file_sf_substreams_gear_type_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sf_substreams_gear_type_v1_types_proto_goTypes = []interface{}{
	(*Block)(nil),         // 0: sf.substreams.vara.type.v1.Block
	(*Extrinsic)(nil),     // 1: sf.substreams.vara.type.v1.Extrinsic
	(*Call)(nil),          // 2: sf.substreams.vara.type.v1.Call
	(*Event)(nil),         // 3: sf.substreams.vara.type.v1.Event
	(*Value)(nil),         // 4: sf.substreams.vara.type.v1.Value
	(*Fields)(nil),        // 5: sf.substreams.vara.type.v1.Fields
	(*Array)(nil),         // 6: sf.substreams.vara.type.v1.Array
	nil,                   // 7: sf.substreams.vara.type.v1.Fields.MapEntry
	(*v1.Header)(nil),     // 8: sf.gear.type.v1.Header
	(*v1.DigestItem)(nil), // 9: sf.gear.type.v1.DigestItem
	(*v1.Signature)(nil),  // 10: sf.gear.type.v1.Signature
}
var file_sf_substreams_gear_type_v1_types_proto_depIdxs = []int32{
	8,  // 0: sf.substreams.vara.type.v1.Block.header:type_name -> sf.gear.type.v1.Header
	1,  // 1: sf.substreams.vara.type.v1.Block.extrinsics:type_name -> sf.substreams.vara.type.v1.Extrinsic
	3,  // 2: sf.substreams.vara.type.v1.Block.events:type_name -> sf.substreams.vara.type.v1.Event
	9,  // 3: sf.substreams.vara.type.v1.Block.digest_items:type_name -> sf.gear.type.v1.DigestItem
	10, // 4: sf.substreams.vara.type.v1.Extrinsic.signature:type_name -> sf.gear.type.v1.Signature
	2,  // 5: sf.substreams.vara.type.v1.Extrinsic.call:type_name -> sf.substreams.vara.type.v1.Call
	5,  // 6: sf.substreams.vara.type.v1.Call.fields:type_name -> sf.substreams.vara.type.v1.Fields
	5,  // 7: sf.substreams.vara.type.v1.Event.fields:type_name -> sf.substreams.vara.type.v1.Fields
	6,  // 8: sf.substreams.vara.type.v1.Value.array:type_name -> sf.substreams.vara.type.v1.Array
	5,  // 9: sf.substreams.vara.type.v1.Value.fields:type_name -> sf.substreams.vara.type.v1.Fields
	7,  // 10: sf.substreams.vara.type.v1.Fields.map:type_name -> sf.substreams.vara.type.v1.Fields.MapEntry
	4,  // 11: sf.substreams.vara.type.v1.Array.items:type_name -> sf.substreams.vara.type.v1.Value
	4,  // 12: sf.substreams.vara.type.v1.Fields.MapEntry.value:type_name -> sf.substreams.vara.type.v1.Value
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_sf_substreams_gear_type_v1_types_proto_init() }
func file_sf_substreams_gear_type_v1_types_proto_init() {
	if File_sf_substreams_gear_type_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extrinsic); i {
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
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Call); i {
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
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fields); i {
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
		file_sf_substreams_gear_type_v1_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Array); i {
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
	file_sf_substreams_gear_type_v1_types_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Value_Int32)(nil),
		(*Value_Bigdecimal)(nil),
		(*Value_Bigint)(nil),
		(*Value_String_)(nil),
		(*Value_Bytes)(nil),
		(*Value_Bool)(nil),
		(*Value_NilValue)(nil),
		(*Value_Array)(nil),
		(*Value_Fields)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_substreams_gear_type_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_gear_type_v1_types_proto_goTypes,
		DependencyIndexes: file_sf_substreams_gear_type_v1_types_proto_depIdxs,
		MessageInfos:      file_sf_substreams_gear_type_v1_types_proto_msgTypes,
	}.Build()
	File_sf_substreams_gear_type_v1_types_proto = out.File
	file_sf_substreams_gear_type_v1_types_proto_rawDesc = nil
	file_sf_substreams_gear_type_v1_types_proto_goTypes = nil
	file_sf_substreams_gear_type_v1_types_proto_depIdxs = nil
}
