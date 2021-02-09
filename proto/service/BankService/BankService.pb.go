// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: BankService.proto

package BankService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int64 `protobuf:"varint,1,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	PageSize   int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Page) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

//查询银行账号通用接口
type ListBankAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int64  `protobuf:"varint,1,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	AccountType    string `protobuf:"bytes,2,opt,name=accountType,proto3" json:"accountType,omitempty"`
	ManagerId      int64  `protobuf:"varint,3,opt,name=managerId,proto3" json:"managerId,omitempty"`
}

func (x *ListBankAccountReq) Reset() {
	*x = ListBankAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBankAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBankAccountReq) ProtoMessage() {}

func (x *ListBankAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBankAccountReq.ProtoReflect.Descriptor instead.
func (*ListBankAccountReq) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{1}
}

func (x *ListBankAccountReq) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *ListBankAccountReq) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *ListBankAccountReq) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

type ListBankAccountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ListBankAccountVo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListBankAccountResp) Reset() {
	*x = ListBankAccountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBankAccountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBankAccountResp) ProtoMessage() {}

func (x *ListBankAccountResp) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBankAccountResp.ProtoReflect.Descriptor instead.
func (*ListBankAccountResp) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{2}
}

func (x *ListBankAccountResp) GetList() []*ListBankAccountVo {
	if x != nil {
		return x.List
	}
	return nil
}

type ListBankAccountVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int64  `protobuf:"varint,1,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	CardNo         string `protobuf:"bytes,2,opt,name=cardNo,proto3" json:"cardNo,omitempty"`
	AccountType    string `protobuf:"bytes,3,opt,name=accountType,proto3" json:"accountType,omitempty"`
	Bank           string `protobuf:"bytes,4,opt,name=bank,proto3" json:"bank,omitempty"`
	CardOwner      string `protobuf:"bytes,5,opt,name=cardOwner,proto3" json:"cardOwner,omitempty"`
	Toibkn         string `protobuf:"bytes,6,opt,name=toibkn,proto3" json:"toibkn,omitempty"`
	FrozenAmount   int64  `protobuf:"varint,7,opt,name=frozenAmount,proto3" json:"frozenAmount,omitempty"`
	CardStatus     string `protobuf:"bytes,8,opt,name=cardStatus,proto3" json:"cardStatus,omitempty"`
	ValidStatus    string `protobuf:"bytes,9,opt,name=validStatus,proto3" json:"validStatus,omitempty"`
}

func (x *ListBankAccountVo) Reset() {
	*x = ListBankAccountVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBankAccountVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBankAccountVo) ProtoMessage() {}

func (x *ListBankAccountVo) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBankAccountVo.ProtoReflect.Descriptor instead.
func (*ListBankAccountVo) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{3}
}

func (x *ListBankAccountVo) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *ListBankAccountVo) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *ListBankAccountVo) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *ListBankAccountVo) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *ListBankAccountVo) GetCardOwner() string {
	if x != nil {
		return x.CardOwner
	}
	return ""
}

func (x *ListBankAccountVo) GetToibkn() string {
	if x != nil {
		return x.Toibkn
	}
	return ""
}

func (x *ListBankAccountVo) GetFrozenAmount() int64 {
	if x != nil {
		return x.FrozenAmount
	}
	return 0
}

func (x *ListBankAccountVo) GetCardStatus() string {
	if x != nil {
		return x.CardStatus
	}
	return ""
}

func (x *ListBankAccountVo) GetValidStatus() string {
	if x != nil {
		return x.ValidStatus
	}
	return ""
}

//创建银行账号部分
type CreateOneAccountInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List      []*AccountInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	ManagerId int64          `protobuf:"varint,2,opt,name=managerId,proto3" json:"managerId,omitempty"`
}

func (x *CreateOneAccountInfoReq) Reset() {
	*x = CreateOneAccountInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOneAccountInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOneAccountInfoReq) ProtoMessage() {}

func (x *CreateOneAccountInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOneAccountInfoReq.ProtoReflect.Descriptor instead.
func (*CreateOneAccountInfoReq) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOneAccountInfoReq) GetList() []*AccountInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *CreateOneAccountInfoReq) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MibId          int64  `protobuf:"varint,1,opt,name=mibId,proto3" json:"mibId,omitempty"`
	CardNo         string `protobuf:"bytes,2,opt,name=cardNo,proto3" json:"cardNo,omitempty"`
	AccountType    string `protobuf:"bytes,3,opt,name=accountType,proto3" json:"accountType,omitempty"`
	Bank           string `protobuf:"bytes,4,opt,name=bank,proto3" json:"bank,omitempty"`
	OrganizationId int64  `protobuf:"varint,5,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	Toibkn         string `protobuf:"bytes,6,opt,name=toibkn,proto3" json:"toibkn,omitempty"`
	CardOwner      string `protobuf:"bytes,7,opt,name=cardOwner,proto3" json:"cardOwner,omitempty"`
	FrozenAmount   int64  `protobuf:"varint,8,opt,name=frozenAmount,proto3" json:"frozenAmount,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{5}
}

func (x *AccountInfo) GetMibId() int64 {
	if x != nil {
		return x.MibId
	}
	return 0
}

func (x *AccountInfo) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *AccountInfo) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *AccountInfo) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *AccountInfo) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *AccountInfo) GetToibkn() string {
	if x != nil {
		return x.Toibkn
	}
	return ""
}

func (x *AccountInfo) GetCardOwner() string {
	if x != nil {
		return x.CardOwner
	}
	return ""
}

func (x *AccountInfo) GetFrozenAmount() int64 {
	if x != nil {
		return x.FrozenAmount
	}
	return 0
}

type CreateOneAccountInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOneAccountInfoResp) Reset() {
	*x = CreateOneAccountInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOneAccountInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOneAccountInfoResp) ProtoMessage() {}

func (x *CreateOneAccountInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOneAccountInfoResp.ProtoReflect.Descriptor instead.
func (*CreateOneAccountInfoResp) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{6}
}

//更新帐号信息
type UpdateBankAccountListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AccountInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UpdateBankAccountListReq) Reset() {
	*x = UpdateBankAccountListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankAccountListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankAccountListReq) ProtoMessage() {}

func (x *UpdateBankAccountListReq) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankAccountListReq.ProtoReflect.Descriptor instead.
func (*UpdateBankAccountListReq) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBankAccountListReq) GetList() []*AccountInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateBankAccountListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBankAccountListResp) Reset() {
	*x = UpdateBankAccountListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankAccountListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankAccountListResp) ProtoMessage() {}

func (x *UpdateBankAccountListResp) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankAccountListResp.ProtoReflect.Descriptor instead.
func (*UpdateBankAccountListResp) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{8}
}

//账号有效性操作
type UpdateBankAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardNo      string `protobuf:"bytes,1,opt,name=cardNo,proto3" json:"cardNo,omitempty"`
	ValidStatus string `protobuf:"bytes,2,opt,name=validStatus,proto3" json:"validStatus,omitempty"`
	ManagerId   int64  `protobuf:"varint,3,opt,name=managerId,proto3" json:"managerId,omitempty"`
}

func (x *UpdateBankAccountReq) Reset() {
	*x = UpdateBankAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankAccountReq) ProtoMessage() {}

func (x *UpdateBankAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankAccountReq.ProtoReflect.Descriptor instead.
func (*UpdateBankAccountReq) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateBankAccountReq) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *UpdateBankAccountReq) GetValidStatus() string {
	if x != nil {
		return x.ValidStatus
	}
	return ""
}

func (x *UpdateBankAccountReq) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

type UpdateBankAccountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBankAccountResp) Reset() {
	*x = UpdateBankAccountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BankService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankAccountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankAccountResp) ProtoMessage() {}

func (x *UpdateBankAccountResp) ProtoReflect() protoreflect.Message {
	mi := &file_BankService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankAccountResp.ProtoReflect.Descriptor instead.
func (*UpdateBankAccountResp) Descriptor() ([]byte, []int) {
	return file_BankService_proto_rawDescGZIP(), []int{10}
}

var File_BankService_proto protoreflect.FileDescriptor

var file_BankService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xa5,
	0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x56, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61,
	0x72, 0x64, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61,
	0x72, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x61, 0x72, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x69, 0x62,
	0x6b, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x69, 0x62, 0x6b, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x65, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22, 0xf3, 0x01,
	0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x69, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x69,
	0x62, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e,
	0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x69,
	0x62, 0x6b, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x69, 0x62, 0x6b,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x48, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x42, 0x61, 0x6e, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x6e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32,
	0x8c, 0x03, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x54, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e,
	0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5a, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x21, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x22, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x66, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x25, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x21,
	0x5a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BankService_proto_rawDescOnce sync.Once
	file_BankService_proto_rawDescData = file_BankService_proto_rawDesc
)

func file_BankService_proto_rawDescGZIP() []byte {
	file_BankService_proto_rawDescOnce.Do(func() {
		file_BankService_proto_rawDescData = protoimpl.X.CompressGZIP(file_BankService_proto_rawDescData)
	})
	return file_BankService_proto_rawDescData
}

var file_BankService_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_BankService_proto_goTypes = []interface{}{
	(*Page)(nil),                      // 0: BankService.Page
	(*ListBankAccountReq)(nil),        // 1: BankService.ListBankAccountReq
	(*ListBankAccountResp)(nil),       // 2: BankService.ListBankAccountResp
	(*ListBankAccountVo)(nil),         // 3: BankService.ListBankAccountVo
	(*CreateOneAccountInfoReq)(nil),   // 4: BankService.CreateOneAccountInfoReq
	(*AccountInfo)(nil),               // 5: BankService.AccountInfo
	(*CreateOneAccountInfoResp)(nil),  // 6: BankService.CreateOneAccountInfoResp
	(*UpdateBankAccountListReq)(nil),  // 7: BankService.UpdateBankAccountListReq
	(*UpdateBankAccountListResp)(nil), // 8: BankService.UpdateBankAccountListResp
	(*UpdateBankAccountReq)(nil),      // 9: BankService.UpdateBankAccountReq
	(*UpdateBankAccountResp)(nil),     // 10: BankService.UpdateBankAccountResp
}
var file_BankService_proto_depIdxs = []int32{
	3,  // 0: BankService.ListBankAccountResp.list:type_name -> BankService.ListBankAccountVo
	5,  // 1: BankService.CreateOneAccountInfoReq.list:type_name -> BankService.AccountInfo
	5,  // 2: BankService.UpdateBankAccountListReq.list:type_name -> BankService.AccountInfo
	1,  // 3: BankService.BankService.ListBankAccount:input_type -> BankService.ListBankAccountReq
	4,  // 4: BankService.BankService.CreateOneAccountInfo:input_type -> BankService.CreateOneAccountInfoReq
	9,  // 5: BankService.BankService.UpdateBankAccount:input_type -> BankService.UpdateBankAccountReq
	7,  // 6: BankService.BankService.UpdateBankAccountList:input_type -> BankService.UpdateBankAccountListReq
	2,  // 7: BankService.BankService.ListBankAccount:output_type -> BankService.ListBankAccountResp
	6,  // 8: BankService.BankService.CreateOneAccountInfo:output_type -> BankService.CreateOneAccountInfoResp
	10, // 9: BankService.BankService.UpdateBankAccount:output_type -> BankService.UpdateBankAccountResp
	8,  // 10: BankService.BankService.UpdateBankAccountList:output_type -> BankService.UpdateBankAccountListResp
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_BankService_proto_init() }
func file_BankService_proto_init() {
	if File_BankService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BankService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_BankService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBankAccountReq); i {
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
		file_BankService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBankAccountResp); i {
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
		file_BankService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBankAccountVo); i {
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
		file_BankService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOneAccountInfoReq); i {
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
		file_BankService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
		file_BankService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOneAccountInfoResp); i {
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
		file_BankService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankAccountListReq); i {
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
		file_BankService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankAccountListResp); i {
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
		file_BankService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankAccountReq); i {
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
		file_BankService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankAccountResp); i {
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
			RawDescriptor: file_BankService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BankService_proto_goTypes,
		DependencyIndexes: file_BankService_proto_depIdxs,
		MessageInfos:      file_BankService_proto_msgTypes,
	}.Build()
	File_BankService_proto = out.File
	file_BankService_proto_rawDesc = nil
	file_BankService_proto_goTypes = nil
	file_BankService_proto_depIdxs = nil
}
