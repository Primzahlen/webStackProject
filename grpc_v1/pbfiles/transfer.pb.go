// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: transfer.proto

package pbfiles

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Fixed size message for testing throughout
type BenchmarkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1   string   `protobuf:"bytes,1,opt,name=Field1,proto3" json:"Field1,omitempty"`
	Field9   string   `protobuf:"bytes,9,opt,name=Field9,proto3" json:"Field9,omitempty"`
	Field18  string   `protobuf:"bytes,18,opt,name=Field18,proto3" json:"Field18,omitempty"`
	Field80  bool     `protobuf:"varint,80,opt,name=Field80,proto3" json:"Field80,omitempty"`
	Field81  bool     `protobuf:"varint,81,opt,name=Field81,proto3" json:"Field81,omitempty"`
	Field2   int32    `protobuf:"varint,2,opt,name=Field2,proto3" json:"Field2,omitempty"`
	Field3   int32    `protobuf:"varint,3,opt,name=Field3,proto3" json:"Field3,omitempty"`
	Field280 int32    `protobuf:"varint,280,opt,name=Field280,proto3" json:"Field280,omitempty"`
	Field6   int32    `protobuf:"varint,6,opt,name=Field6,proto3" json:"Field6,omitempty"`
	Field22  uint64   `protobuf:"varint,22,opt,name=Field22,proto3" json:"Field22,omitempty"`
	Field4   string   `protobuf:"bytes,4,opt,name=Field4,proto3" json:"Field4,omitempty"`
	Field5   []uint64 `protobuf:"fixed64,5,rep,packed,name=Field5,proto3" json:"Field5,omitempty"`
	Field59  bool     `protobuf:"varint,59,opt,name=Field59,proto3" json:"Field59,omitempty"`
	Field7   string   `protobuf:"bytes,7,opt,name=Field7,proto3" json:"Field7,omitempty"`
	Field16  int32    `protobuf:"varint,16,opt,name=Field16,proto3" json:"Field16,omitempty"`
	Field130 int32    `protobuf:"varint,130,opt,name=Field130,proto3" json:"Field130,omitempty"`
	Field12  bool     `protobuf:"varint,12,opt,name=Field12,proto3" json:"Field12,omitempty"`
	Field17  bool     `protobuf:"varint,17,opt,name=Field17,proto3" json:"Field17,omitempty"`
	Field13  bool     `protobuf:"varint,13,opt,name=Field13,proto3" json:"Field13,omitempty"`
	Field14  bool     `protobuf:"varint,14,opt,name=Field14,proto3" json:"Field14,omitempty"`
	Field104 int32    `protobuf:"varint,104,opt,name=Field104,proto3" json:"Field104,omitempty"`
	Field100 int32    `protobuf:"varint,100,opt,name=Field100,proto3" json:"Field100,omitempty"`
	Field101 int32    `protobuf:"varint,101,opt,name=Field101,proto3" json:"Field101,omitempty"`
	Field102 string   `protobuf:"bytes,102,opt,name=Field102,proto3" json:"Field102,omitempty"`
	Field103 string   `protobuf:"bytes,103,opt,name=Field103,proto3" json:"Field103,omitempty"`
	Field29  int32    `protobuf:"varint,29,opt,name=Field29,proto3" json:"Field29,omitempty"`
	Field30  bool     `protobuf:"varint,30,opt,name=Field30,proto3" json:"Field30,omitempty"`
	Field60  int32    `protobuf:"varint,60,opt,name=Field60,proto3" json:"Field60,omitempty"`
	Field271 int32    `protobuf:"varint,271,opt,name=Field271,proto3" json:"Field271,omitempty"`
	Field272 int32    `protobuf:"varint,272,opt,name=Field272,proto3" json:"Field272,omitempty"`
	Field150 int32    `protobuf:"varint,150,opt,name=Field150,proto3" json:"Field150,omitempty"`
	Field23  int32    `protobuf:"varint,23,opt,name=Field23,proto3" json:"Field23,omitempty"`
	Field24  bool     `protobuf:"varint,24,opt,name=Field24,proto3" json:"Field24,omitempty"`
	Field25  int32    `protobuf:"varint,25,opt,name=Field25,proto3" json:"Field25,omitempty"`
	Field78  bool     `protobuf:"varint,78,opt,name=Field78,proto3" json:"Field78,omitempty"`
	Field67  int32    `protobuf:"varint,67,opt,name=Field67,proto3" json:"Field67,omitempty"`
	Field68  int32    `protobuf:"varint,68,opt,name=Field68,proto3" json:"Field68,omitempty"`
	Field128 int32    `protobuf:"varint,128,opt,name=Field128,proto3" json:"Field128,omitempty"`
	Field129 string   `protobuf:"bytes,129,opt,name=Field129,proto3" json:"Field129,omitempty"`
	Field131 int32    `protobuf:"varint,131,opt,name=Field131,proto3" json:"Field131,omitempty"`
}

func (x *BenchmarkMessage) Reset() {
	*x = BenchmarkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkMessage) ProtoMessage() {}

func (x *BenchmarkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkMessage.ProtoReflect.Descriptor instead.
func (*BenchmarkMessage) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *BenchmarkMessage) GetField1() string {
	if x != nil {
		return x.Field1
	}
	return ""
}

func (x *BenchmarkMessage) GetField9() string {
	if x != nil {
		return x.Field9
	}
	return ""
}

func (x *BenchmarkMessage) GetField18() string {
	if x != nil {
		return x.Field18
	}
	return ""
}

func (x *BenchmarkMessage) GetField80() bool {
	if x != nil {
		return x.Field80
	}
	return false
}

func (x *BenchmarkMessage) GetField81() bool {
	if x != nil {
		return x.Field81
	}
	return false
}

func (x *BenchmarkMessage) GetField2() int32 {
	if x != nil {
		return x.Field2
	}
	return 0
}

func (x *BenchmarkMessage) GetField3() int32 {
	if x != nil {
		return x.Field3
	}
	return 0
}

func (x *BenchmarkMessage) GetField280() int32 {
	if x != nil {
		return x.Field280
	}
	return 0
}

func (x *BenchmarkMessage) GetField6() int32 {
	if x != nil {
		return x.Field6
	}
	return 0
}

func (x *BenchmarkMessage) GetField22() uint64 {
	if x != nil {
		return x.Field22
	}
	return 0
}

func (x *BenchmarkMessage) GetField4() string {
	if x != nil {
		return x.Field4
	}
	return ""
}

func (x *BenchmarkMessage) GetField5() []uint64 {
	if x != nil {
		return x.Field5
	}
	return nil
}

func (x *BenchmarkMessage) GetField59() bool {
	if x != nil {
		return x.Field59
	}
	return false
}

func (x *BenchmarkMessage) GetField7() string {
	if x != nil {
		return x.Field7
	}
	return ""
}

func (x *BenchmarkMessage) GetField16() int32 {
	if x != nil {
		return x.Field16
	}
	return 0
}

func (x *BenchmarkMessage) GetField130() int32 {
	if x != nil {
		return x.Field130
	}
	return 0
}

func (x *BenchmarkMessage) GetField12() bool {
	if x != nil {
		return x.Field12
	}
	return false
}

func (x *BenchmarkMessage) GetField17() bool {
	if x != nil {
		return x.Field17
	}
	return false
}

func (x *BenchmarkMessage) GetField13() bool {
	if x != nil {
		return x.Field13
	}
	return false
}

func (x *BenchmarkMessage) GetField14() bool {
	if x != nil {
		return x.Field14
	}
	return false
}

func (x *BenchmarkMessage) GetField104() int32 {
	if x != nil {
		return x.Field104
	}
	return 0
}

func (x *BenchmarkMessage) GetField100() int32 {
	if x != nil {
		return x.Field100
	}
	return 0
}

func (x *BenchmarkMessage) GetField101() int32 {
	if x != nil {
		return x.Field101
	}
	return 0
}

func (x *BenchmarkMessage) GetField102() string {
	if x != nil {
		return x.Field102
	}
	return ""
}

func (x *BenchmarkMessage) GetField103() string {
	if x != nil {
		return x.Field103
	}
	return ""
}

func (x *BenchmarkMessage) GetField29() int32 {
	if x != nil {
		return x.Field29
	}
	return 0
}

func (x *BenchmarkMessage) GetField30() bool {
	if x != nil {
		return x.Field30
	}
	return false
}

func (x *BenchmarkMessage) GetField60() int32 {
	if x != nil {
		return x.Field60
	}
	return 0
}

func (x *BenchmarkMessage) GetField271() int32 {
	if x != nil {
		return x.Field271
	}
	return 0
}

func (x *BenchmarkMessage) GetField272() int32 {
	if x != nil {
		return x.Field272
	}
	return 0
}

func (x *BenchmarkMessage) GetField150() int32 {
	if x != nil {
		return x.Field150
	}
	return 0
}

func (x *BenchmarkMessage) GetField23() int32 {
	if x != nil {
		return x.Field23
	}
	return 0
}

func (x *BenchmarkMessage) GetField24() bool {
	if x != nil {
		return x.Field24
	}
	return false
}

func (x *BenchmarkMessage) GetField25() int32 {
	if x != nil {
		return x.Field25
	}
	return 0
}

func (x *BenchmarkMessage) GetField78() bool {
	if x != nil {
		return x.Field78
	}
	return false
}

func (x *BenchmarkMessage) GetField67() int32 {
	if x != nil {
		return x.Field67
	}
	return 0
}

func (x *BenchmarkMessage) GetField68() int32 {
	if x != nil {
		return x.Field68
	}
	return 0
}

func (x *BenchmarkMessage) GetField128() int32 {
	if x != nil {
		return x.Field128
	}
	return 0
}

func (x *BenchmarkMessage) GetField129() string {
	if x != nil {
		return x.Field129
	}
	return ""
}

func (x *BenchmarkMessage) GetField131() int32 {
	if x != nil {
		return x.Field131
	}
	return 0
}

var File_transfer_proto protoreflect.FileDescriptor

var file_transfer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xb4, 0x08, 0x0a, 0x10, 0x42, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x39,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x12, 0x18,
	0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x38, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x38, 0x30, 0x18, 0x50, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x38, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x31, 0x18, 0x51, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x31, 0x12, 0x16, 0x0a, 0x06,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x1b, 0x0a, 0x08,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x38, 0x30, 0x18, 0x98, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x38, 0x30, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x36, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x06, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x35, 0x39, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x35, 0x39, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x12, 0x18, 0x0a,
	0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x33, 0x30, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x33, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x12, 0x18,
	0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x37, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x33, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x34, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x30, 0x30, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x30, 0x30, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x31,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x31,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x32, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x32, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x33, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x39, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x12, 0x18, 0x0a, 0x07,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x30, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x36, 0x30, 0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x37, 0x31, 0x18, 0x8f, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x37, 0x31, 0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x32, 0x18,
	0x90, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x32,
	0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x30, 0x18, 0x96, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x30, 0x12, 0x18, 0x0a,
	0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x33, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x34, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x34, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x35, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x37, 0x38, 0x18, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x37, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x37,
	0x18, 0x43, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x37, 0x12,
	0x18, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x38, 0x18, 0x44, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x38, 0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x32, 0x38, 0x18, 0x80, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x32, 0x38, 0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x32, 0x39, 0x18, 0x81, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x32, 0x39, 0x12, 0x1b, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x31, 0x18,
	0x83, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x31,
	0x32, 0xb1, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x42, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x16, 0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_proto_rawDescOnce sync.Once
	file_transfer_proto_rawDescData = file_transfer_proto_rawDesc
)

func file_transfer_proto_rawDescGZIP() []byte {
	file_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_proto_rawDescData)
	})
	return file_transfer_proto_rawDescData
}

var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transfer_proto_goTypes = []interface{}{
	(*BenchmarkMessage)(nil), // 0: pbfiles.BenchmarkMessage
}
var file_transfer_proto_depIdxs = []int32{
	0, // 0: pbfiles.TransferService.DataTransmission:input_type -> pbfiles.BenchmarkMessage
	0, // 1: pbfiles.TransferService.DataTransmissionStream:input_type -> pbfiles.BenchmarkMessage
	0, // 2: pbfiles.TransferService.DataTransmission:output_type -> pbfiles.BenchmarkMessage
	0, // 3: pbfiles.TransferService.DataTransmissionStream:output_type -> pbfiles.BenchmarkMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_proto_init() }
func file_transfer_proto_init() {
	if File_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkMessage); i {
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
			RawDescriptor: file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_proto_msgTypes,
	}.Build()
	File_transfer_proto = out.File
	file_transfer_proto_rawDesc = nil
	file_transfer_proto_goTypes = nil
	file_transfer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferServiceClient interface {
	DataTransmission(ctx context.Context, in *BenchmarkMessage, opts ...grpc.CallOption) (*BenchmarkMessage, error)
	DataTransmissionStream(ctx context.Context, in *BenchmarkMessage, opts ...grpc.CallOption) (TransferService_DataTransmissionStreamClient, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) DataTransmission(ctx context.Context, in *BenchmarkMessage, opts ...grpc.CallOption) (*BenchmarkMessage, error) {
	out := new(BenchmarkMessage)
	err := c.cc.Invoke(ctx, "/pbfiles.TransferService/DataTransmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) DataTransmissionStream(ctx context.Context, in *BenchmarkMessage, opts ...grpc.CallOption) (TransferService_DataTransmissionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TransferService_serviceDesc.Streams[0], "/pbfiles.TransferService/DataTransmissionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferServiceDataTransmissionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransferService_DataTransmissionStreamClient interface {
	Recv() (*BenchmarkMessage, error)
	grpc.ClientStream
}

type transferServiceDataTransmissionStreamClient struct {
	grpc.ClientStream
}

func (x *transferServiceDataTransmissionStreamClient) Recv() (*BenchmarkMessage, error) {
	m := new(BenchmarkMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServiceServer is the server API for TransferService service.
type TransferServiceServer interface {
	DataTransmission(context.Context, *BenchmarkMessage) (*BenchmarkMessage, error)
	DataTransmissionStream(*BenchmarkMessage, TransferService_DataTransmissionStreamServer) error
}

// UnimplementedTransferServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransferServiceServer struct {
}

func (*UnimplementedTransferServiceServer) DataTransmission(context.Context, *BenchmarkMessage) (*BenchmarkMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataTransmission not implemented")
}
func (*UnimplementedTransferServiceServer) DataTransmissionStream(*BenchmarkMessage, TransferService_DataTransmissionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DataTransmissionStream not implemented")
}

func RegisterTransferServiceServer(s *grpc.Server, srv TransferServiceServer) {
	s.RegisterService(&_TransferService_serviceDesc, srv)
}

func _TransferService_DataTransmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchmarkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).DataTransmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfiles.TransferService/DataTransmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).DataTransmission(ctx, req.(*BenchmarkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_DataTransmissionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BenchmarkMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransferServiceServer).DataTransmissionStream(m, &transferServiceDataTransmissionStreamServer{stream})
}

type TransferService_DataTransmissionStreamServer interface {
	Send(*BenchmarkMessage) error
	grpc.ServerStream
}

type transferServiceDataTransmissionStreamServer struct {
	grpc.ServerStream
}

func (x *transferServiceDataTransmissionStreamServer) Send(m *BenchmarkMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _TransferService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbfiles.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataTransmission",
			Handler:    _TransferService_DataTransmission_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DataTransmissionStream",
			Handler:       _TransferService_DataTransmissionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transfer.proto",
}
