// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: doc/esManager.proto

package pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_doc_esManager_proto_rawDescGZIP(), []int{0}
}

type ESConnID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ESConnID) Reset() {
	*x = ESConnID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESConnID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESConnID) ProtoMessage() {}

func (x *ESConnID) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESConnID.ProtoReflect.Descriptor instead.
func (*ESConnID) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{1}
}

func (x *ESConnID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Total struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *Total) Reset() {
	*x = Total{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Total) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Total) ProtoMessage() {}

func (x *Total) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Total.ProtoReflect.Descriptor instead.
func (*Total) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{2}
}

func (x *Total) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateESConnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ESConn   string `protobuf:"bytes,1,opt,name=ESConn,proto3" json:"ESConn,omitempty"`
	Version  int64  `protobuf:"varint,2,opt,name=Version,proto3" json:"Version,omitempty"`
	User     string `protobuf:"bytes,3,opt,name=User,proto3" json:"User,omitempty"`
	PassWord string `protobuf:"bytes,4,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
	Describe string `protobuf:"bytes,5,opt,name=Describe,proto3" json:"Describe,omitempty"`
}

func (x *CreateESConnRequest) Reset() {
	*x = CreateESConnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateESConnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateESConnRequest) ProtoMessage() {}

func (x *CreateESConnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateESConnRequest.ProtoReflect.Descriptor instead.
func (*CreateESConnRequest) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{3}
}

func (x *CreateESConnRequest) GetESConn() string {
	if x != nil {
		return x.ESConn
	}
	return ""
}

func (x *CreateESConnRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CreateESConnRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateESConnRequest) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *CreateESConnRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

type UpdateESConnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ESConn   string `protobuf:"bytes,2,opt,name=ESConn,proto3" json:"ESConn,omitempty"`
	Version  int64  `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	User     string `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
	PassWord string `protobuf:"bytes,5,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
	Describe string `protobuf:"bytes,6,opt,name=Describe,proto3" json:"Describe,omitempty"`
}

func (x *UpdateESConnRequest) Reset() {
	*x = UpdateESConnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateESConnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateESConnRequest) ProtoMessage() {}

func (x *UpdateESConnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateESConnRequest.ProtoReflect.Descriptor instead.
func (*UpdateESConnRequest) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateESConnRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateESConnRequest) GetESConn() string {
	if x != nil {
		return x.ESConn
	}
	return ""
}

func (x *UpdateESConnRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UpdateESConnRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UpdateESConnRequest) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *UpdateESConnRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

type ESConnPagingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
}

func (x *ESConnPagingRequest) Reset() {
	*x = ESConnPagingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESConnPagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESConnPagingRequest) ProtoMessage() {}

func (x *ESConnPagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESConnPagingRequest.ProtoReflect.Descriptor instead.
func (*ESConnPagingRequest) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{5}
}

func (x *ESConnPagingRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ESConnPagingRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ESConnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ESConns []*ESConn `protobuf:"bytes,1,rep,name=ESConns,proto3" json:"ESConns,omitempty"`
}

func (x *ESConnResponse) Reset() {
	*x = ESConnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESConnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESConnResponse) ProtoMessage() {}

func (x *ESConnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESConnResponse.ProtoReflect.Descriptor instead.
func (*ESConnResponse) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{6}
}

func (x *ESConnResponse) GetESConns() []*ESConn {
	if x != nil {
		return x.ESConns
	}
	return nil
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{7}
}

func (x *PingResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EsConnID uint64 `protobuf:"varint,1,opt,name=EsConnID,proto3" json:"EsConnID,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{8}
}

func (x *PingRequest) GetEsConnID() uint64 {
	if x != nil {
		return x.EsConnID
	}
	return 0
}

type ESConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ESConn   string `protobuf:"bytes,2,opt,name=ESConn,proto3" json:"ESConn,omitempty"`
	Version  int64  `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	User     string `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
	PassWord string `protobuf:"bytes,5,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
	Describe string `protobuf:"bytes,6,opt,name=Describe,proto3" json:"Describe,omitempty"`
}

func (x *ESConn) Reset() {
	*x = ESConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_esManager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESConn) ProtoMessage() {}

func (x *ESConn) ProtoReflect() protoreflect.Message {
	mi := &file_doc_esManager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESConn.ProtoReflect.Descriptor instead.
func (*ESConn) Descriptor() ([]byte, []int) {
	return file_doc_esManager_proto_rawDescGZIP(), []int{9}
}

func (x *ESConn) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ESConn) GetESConn() string {
	if x != nil {
		return x.ESConn
	}
	return ""
}

func (x *ESConn) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ESConn) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ESConn) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *ESConn) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

var File_doc_esManager_proto protoreflect.FileDescriptor

var file_doc_esManager_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x6f, 0x63, 0x2f, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x73, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1a, 0x0a, 0x08, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x57, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x57, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x22, 0xa3, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x53, 0x43, 0x6f, 0x6e,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x53, 0x43, 0x6f,
	0x6e, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x38, 0x0a,
	0x0e, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x07, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x07,
	0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x22, 0x38, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x29, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x45, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x45, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x22, 0x96, 0x01, 0x0a,
	0x06, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x53, 0x43, 0x6f, 0x6e,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x32, 0xa6, 0x03, 0x0a, 0x10, 0x45, 0x73, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x45, 0x53,
	0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x65, 0x73, 0x70,
	0x62, 0x2e, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x53, 0x43,
	0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x45,
	0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x0b, 0x2e, 0x65, 0x73, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x65,
	0x73, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x53, 0x43, 0x6f, 0x6e,
	0x6e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x53, 0x43, 0x6f, 0x6e,
	0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x6c, 0x6c, 0x12, 0x0b,
	0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x65, 0x73,
	0x70, 0x62, 0x2e, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x53, 0x43, 0x6f, 0x6e,
	0x6e, 0x12, 0x19, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x65,
	0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x19, 0x2e, 0x65, 0x73, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x53, 0x43, 0x6f, 0x6e,
	0x6e, 0x12, 0x0e, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x49,
	0x44, 0x1a, 0x0b, 0x2e, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doc_esManager_proto_rawDescOnce sync.Once
	file_doc_esManager_proto_rawDescData = file_doc_esManager_proto_rawDesc
)

func file_doc_esManager_proto_rawDescGZIP() []byte {
	file_doc_esManager_proto_rawDescOnce.Do(func() {
		file_doc_esManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_doc_esManager_proto_rawDescData)
	})
	return file_doc_esManager_proto_rawDescData
}

var file_doc_esManager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_doc_esManager_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: espb.Empty
	(*ESConnID)(nil),            // 1: espb.ESConnID
	(*Total)(nil),               // 2: espb.Total
	(*CreateESConnRequest)(nil), // 3: espb.CreateESConnRequest
	(*UpdateESConnRequest)(nil), // 4: espb.UpdateESConnRequest
	(*ESConnPagingRequest)(nil), // 5: espb.ESConnPagingRequest
	(*ESConnResponse)(nil),      // 6: espb.ESConnResponse
	(*PingResponse)(nil),        // 7: espb.PingResponse
	(*PingRequest)(nil),         // 8: espb.PingRequest
	(*ESConn)(nil),              // 9: espb.ESConn
	(*anypb.Any)(nil),           // 10: google.protobuf.Any
}
var file_doc_esManager_proto_depIdxs = []int32{
	9,  // 0: espb.ESConnResponse.ESConns:type_name -> espb.ESConn
	10, // 1: espb.PingResponse.Data:type_name -> google.protobuf.Any
	5,  // 2: espb.EsManagerService.ESConnPaging:input_type -> espb.ESConnPagingRequest
	0,  // 3: espb.EsManagerService.ESConnTotal:input_type -> espb.Empty
	8,  // 4: espb.EsManagerService.Ping:input_type -> espb.PingRequest
	1,  // 5: espb.EsManagerService.ESConnDetail:input_type -> espb.ESConnID
	0,  // 6: espb.EsManagerService.ESConnAll:input_type -> espb.Empty
	3,  // 7: espb.EsManagerService.CreateESConn:input_type -> espb.CreateESConnRequest
	4,  // 8: espb.EsManagerService.UpdateESConn:input_type -> espb.UpdateESConnRequest
	1,  // 9: espb.EsManagerService.DeleteESConn:input_type -> espb.ESConnID
	6,  // 10: espb.EsManagerService.ESConnPaging:output_type -> espb.ESConnResponse
	2,  // 11: espb.EsManagerService.ESConnTotal:output_type -> espb.Total
	7,  // 12: espb.EsManagerService.Ping:output_type -> espb.PingResponse
	9,  // 13: espb.EsManagerService.ESConnDetail:output_type -> espb.ESConn
	6,  // 14: espb.EsManagerService.ESConnAll:output_type -> espb.ESConnResponse
	0,  // 15: espb.EsManagerService.CreateESConn:output_type -> espb.Empty
	0,  // 16: espb.EsManagerService.UpdateESConn:output_type -> espb.Empty
	0,  // 17: espb.EsManagerService.DeleteESConn:output_type -> espb.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_doc_esManager_proto_init() }
func file_doc_esManager_proto_init() {
	if File_doc_esManager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doc_esManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_doc_esManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESConnID); i {
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
		file_doc_esManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Total); i {
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
		file_doc_esManager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateESConnRequest); i {
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
		file_doc_esManager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateESConnRequest); i {
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
		file_doc_esManager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESConnPagingRequest); i {
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
		file_doc_esManager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESConnResponse); i {
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
		file_doc_esManager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_doc_esManager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_doc_esManager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESConn); i {
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
			RawDescriptor: file_doc_esManager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doc_esManager_proto_goTypes,
		DependencyIndexes: file_doc_esManager_proto_depIdxs,
		MessageInfos:      file_doc_esManager_proto_msgTypes,
	}.Build()
	File_doc_esManager_proto = out.File
	file_doc_esManager_proto_rawDesc = nil
	file_doc_esManager_proto_goTypes = nil
	file_doc_esManager_proto_depIdxs = nil
}
