// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: doc/monitoringManager.proto

package pb

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

//-------------------
// 请求
//-------------------
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[0]
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
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{0}
}

type Total struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *Total) Reset() {
	*x = Total{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Total) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Total) ProtoMessage() {}

func (x *Total) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[1]
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
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{1}
}

func (x *Total) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AlertRuleID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AlertRuleID) Reset() {
	*x = AlertRuleID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRuleID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRuleID) ProtoMessage() {}

func (x *AlertRuleID) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRuleID.ProtoReflect.Descriptor instead.
func (*AlertRuleID) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{2}
}

func (x *AlertRuleID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type AlertRuleCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Type string `protobuf:"bytes,6,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *AlertRuleCountRequest) Reset() {
	*x = AlertRuleCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRuleCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRuleCountRequest) ProtoMessage() {}

func (x *AlertRuleCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRuleCountRequest.ProtoReflect.Descriptor instead.
func (*AlertRuleCountRequest) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{3}
}

func (x *AlertRuleCountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertRuleCountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AlertRulePagingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	OrderKey string `protobuf:"bytes,3,opt,name=OrderKey,proto3" json:"OrderKey,omitempty"`
	Order    string `protobuf:"bytes,4,opt,name=Order,proto3" json:"Order,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *AlertRulePagingRequest) Reset() {
	*x = AlertRulePagingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRulePagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRulePagingRequest) ProtoMessage() {}

func (x *AlertRulePagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRulePagingRequest.ProtoReflect.Descriptor instead.
func (*AlertRulePagingRequest) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{4}
}

func (x *AlertRulePagingRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AlertRulePagingRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AlertRulePagingRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *AlertRulePagingRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *AlertRulePagingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertRulePagingRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AlertRulePagingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertRules []*AlertRule `protobuf:"bytes,1,rep,name=AlertRules,proto3" json:"AlertRules,omitempty"`
}

func (x *AlertRulePagingResponse) Reset() {
	*x = AlertRulePagingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRulePagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRulePagingResponse) ProtoMessage() {}

func (x *AlertRulePagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRulePagingResponse.ProtoReflect.Descriptor instead.
func (*AlertRulePagingResponse) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{5}
}

func (x *AlertRulePagingResponse) GetAlertRules() []*AlertRule {
	if x != nil {
		return x.AlertRules
	}
	return nil
}

type AlertRuleDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type     string            `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Group    string            `protobuf:"bytes,4,opt,name=Group,proto3" json:"Group,omitempty"`
	Tag      string            `protobuf:"bytes,5,opt,name=Tag,proto3" json:"Tag,omitempty"`
	To       int64             `protobuf:"varint,6,opt,name=To,proto3" json:"To,omitempty"`
	Expr     string            `protobuf:"bytes,7,opt,name=Expr,proto3" json:"Expr,omitempty"`
	Operator string            `protobuf:"bytes,8,opt,name=Operator,proto3" json:"Operator,omitempty"`
	Value    string            `protobuf:"bytes,9,opt,name=Value,proto3" json:"Value,omitempty"`
	For      string            `protobuf:"bytes,10,opt,name=For,proto3" json:"For,omitempty"`
	Summary  string            `protobuf:"bytes,11,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Describe string            `protobuf:"bytes,12,opt,name=Describe,proto3" json:"Describe,omitempty"`
	IsWrite  bool              `protobuf:"varint,13,opt,name=IsWrite,proto3" json:"IsWrite,omitempty"`
	Labels   map[string]string `protobuf:"bytes,14,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AlertRuleDetailResponse) Reset() {
	*x = AlertRuleDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRuleDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRuleDetailResponse) ProtoMessage() {}

func (x *AlertRuleDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRuleDetailResponse.ProtoReflect.Descriptor instead.
func (*AlertRuleDetailResponse) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{6}
}

func (x *AlertRuleDetailResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AlertRuleDetailResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *AlertRuleDetailResponse) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetFor() string {
	if x != nil {
		return x.For
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *AlertRuleDetailResponse) GetIsWrite() bool {
	if x != nil {
		return x.IsWrite
	}
	return false
}

func (x *AlertRuleDetailResponse) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type AlertRuleLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]string `protobuf:"bytes,1,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AlertRuleLabelsResponse) Reset() {
	*x = AlertRuleLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRuleLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRuleLabelsResponse) ProtoMessage() {}

func (x *AlertRuleLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRuleLabelsResponse.ProtoReflect.Descriptor instead.
func (*AlertRuleLabelsResponse) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{7}
}

func (x *AlertRuleLabelsResponse) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type AlertRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Group    string `protobuf:"bytes,4,opt,name=Group,proto3" json:"Group,omitempty"`
	Tag      string `protobuf:"bytes,5,opt,name=Tag,proto3" json:"Tag,omitempty"`
	To       int64  `protobuf:"varint,6,opt,name=To,proto3" json:"To,omitempty"`
	Expr     string `protobuf:"bytes,7,opt,name=Expr,proto3" json:"Expr,omitempty"`
	Operator string `protobuf:"bytes,8,opt,name=Operator,proto3" json:"Operator,omitempty"`
	Value    string `protobuf:"bytes,9,opt,name=Value,proto3" json:"Value,omitempty"`
	For      string `protobuf:"bytes,10,opt,name=For,proto3" json:"For,omitempty"`
	Summary  string `protobuf:"bytes,11,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Describe string `protobuf:"bytes,12,opt,name=Describe,proto3" json:"Describe,omitempty"`
	IsWrite  bool   `protobuf:"varint,13,opt,name=IsWrite,proto3" json:"IsWrite,omitempty"`
}

func (x *AlertRule) Reset() {
	*x = AlertRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_monitoringManager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRule) ProtoMessage() {}

func (x *AlertRule) ProtoReflect() protoreflect.Message {
	mi := &file_doc_monitoringManager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRule.ProtoReflect.Descriptor instead.
func (*AlertRule) Descriptor() ([]byte, []int) {
	return file_doc_monitoringManager_proto_rawDescGZIP(), []int{8}
}

func (x *AlertRule) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AlertRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertRule) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AlertRule) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *AlertRule) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlertRule) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *AlertRule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *AlertRule) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *AlertRule) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AlertRule) GetFor() string {
	if x != nil {
		return x.For
	}
	return ""
}

func (x *AlertRule) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *AlertRule) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *AlertRule) GetIsWrite() bool {
	if x != nil {
		return x.IsWrite
	}
	return false
}

var File_doc_monitoringManager_proto protoreflect.FileDescriptor

var file_doc_monitoringManager_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x6f, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x3f, 0x0a, 0x15, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x16, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x17, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb7, 0x03, 0x0a,
	0x17, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x45, 0x78, 0x70, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x45, 0x78, 0x70, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x46, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x49, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x17, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70,
	0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa3, 0x02, 0x0a, 0x09, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x45, 0x78, 0x70, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x45, 0x78, 0x70, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x46, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x32, 0xe9,
	0x02, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x49, 0x44, 0x1a, 0x25, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70,
	0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x19, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x1a, 0x25, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e,
	0x0a, 0x0f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x24, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doc_monitoringManager_proto_rawDescOnce sync.Once
	file_doc_monitoringManager_proto_rawDescData = file_doc_monitoringManager_proto_rawDesc
)

func file_doc_monitoringManager_proto_rawDescGZIP() []byte {
	file_doc_monitoringManager_proto_rawDescOnce.Do(func() {
		file_doc_monitoringManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_doc_monitoringManager_proto_rawDescData)
	})
	return file_doc_monitoringManager_proto_rawDescData
}

var file_doc_monitoringManager_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_doc_monitoringManager_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: monitoringpb.Empty
	(*Total)(nil),                   // 1: monitoringpb.Total
	(*AlertRuleID)(nil),             // 2: monitoringpb.AlertRuleID
	(*AlertRuleCountRequest)(nil),   // 3: monitoringpb.AlertRuleCountRequest
	(*AlertRulePagingRequest)(nil),  // 4: monitoringpb.AlertRulePagingRequest
	(*AlertRulePagingResponse)(nil), // 5: monitoringpb.AlertRulePagingResponse
	(*AlertRuleDetailResponse)(nil), // 6: monitoringpb.AlertRuleDetailResponse
	(*AlertRuleLabelsResponse)(nil), // 7: monitoringpb.AlertRuleLabelsResponse
	(*AlertRule)(nil),               // 8: monitoringpb.AlertRule
	nil,                             // 9: monitoringpb.AlertRuleDetailResponse.LabelsEntry
	nil,                             // 10: monitoringpb.AlertRuleLabelsResponse.LabelsEntry
}
var file_doc_monitoringManager_proto_depIdxs = []int32{
	8,  // 0: monitoringpb.AlertRulePagingResponse.AlertRules:type_name -> monitoringpb.AlertRule
	9,  // 1: monitoringpb.AlertRuleDetailResponse.Labels:type_name -> monitoringpb.AlertRuleDetailResponse.LabelsEntry
	10, // 2: monitoringpb.AlertRuleLabelsResponse.Labels:type_name -> monitoringpb.AlertRuleLabelsResponse.LabelsEntry
	2,  // 3: monitoringpb.MonitoringManager.AlertRuleDetail:input_type -> monitoringpb.AlertRuleID
	2,  // 4: monitoringpb.MonitoringManager.AlertRuleLabels:input_type -> monitoringpb.AlertRuleID
	4,  // 5: monitoringpb.MonitoringManager.AlertRulePaging:input_type -> monitoringpb.AlertRulePagingRequest
	3,  // 6: monitoringpb.MonitoringManager.AlertRuleCount:input_type -> monitoringpb.AlertRuleCountRequest
	6,  // 7: monitoringpb.MonitoringManager.AlertRuleDetail:output_type -> monitoringpb.AlertRuleDetailResponse
	7,  // 8: monitoringpb.MonitoringManager.AlertRuleLabels:output_type -> monitoringpb.AlertRuleLabelsResponse
	5,  // 9: monitoringpb.MonitoringManager.AlertRulePaging:output_type -> monitoringpb.AlertRulePagingResponse
	1,  // 10: monitoringpb.MonitoringManager.AlertRuleCount:output_type -> monitoringpb.Total
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_doc_monitoringManager_proto_init() }
func file_doc_monitoringManager_proto_init() {
	if File_doc_monitoringManager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doc_monitoringManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_doc_monitoringManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_doc_monitoringManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRuleID); i {
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
		file_doc_monitoringManager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRuleCountRequest); i {
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
		file_doc_monitoringManager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRulePagingRequest); i {
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
		file_doc_monitoringManager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRulePagingResponse); i {
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
		file_doc_monitoringManager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRuleDetailResponse); i {
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
		file_doc_monitoringManager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRuleLabelsResponse); i {
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
		file_doc_monitoringManager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRule); i {
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
			RawDescriptor: file_doc_monitoringManager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doc_monitoringManager_proto_goTypes,
		DependencyIndexes: file_doc_monitoringManager_proto_depIdxs,
		MessageInfos:      file_doc_monitoringManager_proto_msgTypes,
	}.Build()
	File_doc_monitoringManager_proto = out.File
	file_doc_monitoringManager_proto_rawDesc = nil
	file_doc_monitoringManager_proto_goTypes = nil
	file_doc_monitoringManager_proto_depIdxs = nil
}