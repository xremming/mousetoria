// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: database.proto

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

type InsertTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int64 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *InsertTransactionResponse) Reset() {
	*x = InsertTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertTransactionResponse) ProtoMessage() {}

func (x *InsertTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertTransactionResponse.ProtoReflect.Descriptor instead.
func (*InsertTransactionResponse) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{0}
}

func (x *InsertTransactionResponse) GetTransactionId() int64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int64 `protobuf:"varint,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{1}
}

func (x *GetTransactionRequest) GetTransactionId() int64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

type GetAllTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit *int64 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetAllTransactionsRequest) Reset() {
	*x = GetAllTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTransactionsRequest) ProtoMessage() {}

func (x *GetAllTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetAllTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllTransactionsRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

var File_database_proto protoreflect.FileDescriptor

var file_database_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42,
	0x0a, 0x19, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x40, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x32, 0xc3, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x72, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67,
	0x2f, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_proto_rawDescOnce sync.Once
	file_database_proto_rawDescData = file_database_proto_rawDesc
)

func file_database_proto_rawDescGZIP() []byte {
	file_database_proto_rawDescOnce.Do(func() {
		file_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_proto_rawDescData)
	})
	return file_database_proto_rawDescData
}

var file_database_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_database_proto_goTypes = []interface{}{
	(*InsertTransactionResponse)(nil), // 0: InsertTransactionResponse
	(*GetTransactionRequest)(nil),     // 1: GetTransactionRequest
	(*GetAllTransactionsRequest)(nil), // 2: GetAllTransactionsRequest
	(*Transaction)(nil),               // 3: Transaction
}
var file_database_proto_depIdxs = []int32{
	3, // 0: Database.InsertTransaction:input_type -> Transaction
	1, // 1: Database.GetTransaction:input_type -> GetTransactionRequest
	2, // 2: Database.GetAllTransactions:input_type -> GetAllTransactionsRequest
	0, // 3: Database.InsertTransaction:output_type -> InsertTransactionResponse
	3, // 4: Database.GetTransaction:output_type -> Transaction
	3, // 5: Database.GetAllTransactions:output_type -> Transaction
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_database_proto_init() }
func file_database_proto_init() {
	if File_database_proto != nil {
		return
	}
	file_ledger_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertTransactionResponse); i {
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
		file_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRequest); i {
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
		file_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTransactionsRequest); i {
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
	file_database_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_database_proto_goTypes,
		DependencyIndexes: file_database_proto_depIdxs,
		MessageInfos:      file_database_proto_msgTypes,
	}.Build()
	File_database_proto = out.File
	file_database_proto_rawDesc = nil
	file_database_proto_goTypes = nil
	file_database_proto_depIdxs = nil
}
