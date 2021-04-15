// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: Chain.proto

package ChainService

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type GetHashContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerId int64  `protobuf:"varint,1,opt,name=managerId,proto3" json:"managerId"`
	Hash      string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash"`
	//目前支持的hash类型  COM CONFIG CONTRACT CONTRACT_MI FACTORING MEDICINE MI PAY SHIPMENT
	BusinessType string `protobuf:"bytes,3,opt,name=businessType,proto3" json:"businessType"`
}

func (x *GetHashContentReq) Reset() {
	*x = GetHashContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHashContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHashContentReq) ProtoMessage() {}

func (x *GetHashContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_Chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHashContentReq.ProtoReflect.Descriptor instead.
func (*GetHashContentReq) Descriptor() ([]byte, []int) {
	return file_Chain_proto_rawDescGZIP(), []int{0}
}

func (x *GetHashContentReq) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *GetHashContentReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetHashContentReq) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

type GetHashContentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//证书内容结构，json
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	//向上追溯用的hash
	LastHash         string `protobuf:"bytes,2,opt,name=lastHash,proto3" json:"lastHash"`
	UploadNode       string `protobuf:"bytes,3,opt,name=uploadNode,proto3" json:"uploadNode"`
	UploaderName     string `protobuf:"bytes,4,opt,name=uploaderName,proto3" json:"uploaderName"`
	UploaderRole     string `protobuf:"bytes,5,opt,name=uploaderRole,proto3" json:"uploaderRole"`
	UploaderRoleName string `protobuf:"bytes,6,opt,name=uploaderRoleName,proto3" json:"uploaderRoleName"`
	UploaderAccount  string `protobuf:"bytes,7,opt,name=uploaderAccount,proto3" json:"uploaderAccount"`
	CurrentHash      string `protobuf:"bytes,8,opt,name=currentHash,proto3" json:"currentHash"`
	// 上链成功时间
	UploadSuccessAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=uploadSuccessAt,proto3" json:"uploadSuccessAt"`
	// 证书描述
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description"`
}

func (x *GetHashContentResp) Reset() {
	*x = GetHashContentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHashContentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHashContentResp) ProtoMessage() {}

func (x *GetHashContentResp) ProtoReflect() protoreflect.Message {
	mi := &file_Chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHashContentResp.ProtoReflect.Descriptor instead.
func (*GetHashContentResp) Descriptor() ([]byte, []int) {
	return file_Chain_proto_rawDescGZIP(), []int{1}
}

func (x *GetHashContentResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GetHashContentResp) GetLastHash() string {
	if x != nil {
		return x.LastHash
	}
	return ""
}

func (x *GetHashContentResp) GetUploadNode() string {
	if x != nil {
		return x.UploadNode
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderName() string {
	if x != nil {
		return x.UploaderName
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderRole() string {
	if x != nil {
		return x.UploaderRole
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderRoleName() string {
	if x != nil {
		return x.UploaderRoleName
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderAccount() string {
	if x != nil {
		return x.UploaderAccount
	}
	return ""
}

func (x *GetHashContentResp) GetCurrentHash() string {
	if x != nil {
		return x.CurrentHash
	}
	return ""
}

func (x *GetHashContentResp) GetUploadSuccessAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadSuccessAt
	}
	return nil
}

func (x *GetHashContentResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_Chain_proto protoreflect.FileDescriptor

var file_Chain_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8c, 0x03, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x44,
	0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x63, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x23, 0x5a, 0x21, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Chain_proto_rawDescOnce sync.Once
	file_Chain_proto_rawDescData = file_Chain_proto_rawDesc
)

func file_Chain_proto_rawDescGZIP() []byte {
	file_Chain_proto_rawDescOnce.Do(func() {
		file_Chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_Chain_proto_rawDescData)
	})
	return file_Chain_proto_rawDescData
}

var file_Chain_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Chain_proto_goTypes = []interface{}{
	(*GetHashContentReq)(nil),     // 0: ChainService.GetHashContentReq
	(*GetHashContentResp)(nil),    // 1: ChainService.GetHashContentResp
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_Chain_proto_depIdxs = []int32{
	2, // 0: ChainService.GetHashContentResp.uploadSuccessAt:type_name -> google.protobuf.Timestamp
	0, // 1: ChainService.ChainService.GetHashContent:input_type -> ChainService.GetHashContentReq
	1, // 2: ChainService.ChainService.GetHashContent:output_type -> ChainService.GetHashContentResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Chain_proto_init() }
func file_Chain_proto_init() {
	if File_Chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHashContentReq); i {
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
		file_Chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHashContentResp); i {
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
			RawDescriptor: file_Chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Chain_proto_goTypes,
		DependencyIndexes: file_Chain_proto_depIdxs,
		MessageInfos:      file_Chain_proto_msgTypes,
	}.Build()
	File_Chain_proto = out.File
	file_Chain_proto_rawDesc = nil
	file_Chain_proto_goTypes = nil
	file_Chain_proto_depIdxs = nil
}
