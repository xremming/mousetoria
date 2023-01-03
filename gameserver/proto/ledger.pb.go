// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.11
// source: ledger.proto

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

type AccountGroup int32

const (
	AccountGroup_INVALID     AccountGroup = 0
	AccountGroup_ASSETS      AccountGroup = 1
	AccountGroup_LIABILITIES AccountGroup = 2
	AccountGroup_EQUITY      AccountGroup = 3
	AccountGroup_REVENUES    AccountGroup = 4
	AccountGroup_EXPENSES    AccountGroup = 5
)

// Enum value maps for AccountGroup.
var (
	AccountGroup_name = map[int32]string{
		0: "INVALID",
		1: "ASSETS",
		2: "LIABILITIES",
		3: "EQUITY",
		4: "REVENUES",
		5: "EXPENSES",
	}
	AccountGroup_value = map[string]int32{
		"INVALID":     0,
		"ASSETS":      1,
		"LIABILITIES": 2,
		"EQUITY":      3,
		"REVENUES":    4,
		"EXPENSES":    5,
	}
)

func (x AccountGroup) Enum() *AccountGroup {
	p := new(AccountGroup)
	*p = x
	return p
}

func (x AccountGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_ledger_proto_enumTypes[0].Descriptor()
}

func (AccountGroup) Type() protoreflect.EnumType {
	return &file_ledger_proto_enumTypes[0]
}

func (x AccountGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountGroup.Descriptor instead.
func (AccountGroup) EnumDescriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{0}
}

type AccountNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group  AccountGroup `protobuf:"varint,1,opt,name=group,proto3,enum=AccountGroup" json:"group,omitempty"`
	Number int32        `protobuf:"zigzag32,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AccountNumber) Reset() {
	*x = AccountNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountNumber) ProtoMessage() {}

func (x *AccountNumber) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountNumber.ProtoReflect.Descriptor instead.
func (*AccountNumber) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *AccountNumber) GetGroup() AccountGroup {
	if x != nil {
		return x.Group
	}
	return AccountGroup_INVALID
}

func (x *AccountNumber) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber *AccountNumber `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Value         int32          `protobuf:"zigzag32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{1}
}

func (x *Record) GetAccountNumber() *AccountNumber {
	if x != nil {
		return x.AccountNumber
	}
	return nil
}

func (x *Record) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp     *Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Comment       *string    `protobuf:"bytes,2,opt,name=comment,proto3,oneof" json:"comment,omitempty"`
	RecordsDebit  []*Record  `protobuf:"bytes,3,rep,name=records_debit,json=recordsDebit,proto3" json:"records_debit,omitempty"`
	RecordsCredit []*Record  `protobuf:"bytes,4,rep,name=records_credit,json=recordsCredit,proto3" json:"records_credit,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetTimestamp() *Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Transaction) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

func (x *Transaction) GetRecordsDebit() []*Record {
	if x != nil {
		return x.RecordsDebit
	}
	return nil
}

func (x *Transaction) GetRecordsCredit() []*Record {
	if x != nil {
		return x.RecordsCredit
	}
	return nil
}

var File_ledger_proto protoreflect.FileDescriptor

var file_ledger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4c, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x55, 0x0a,
	0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a,
	0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x65, 0x62, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x65, 0x62, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x0e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0d, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x60, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x53, 0x53, 0x45, 0x54, 0x53, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x49, 0x45, 0x53, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55, 0x49, 0x54, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x55, 0x45, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x58, 0x50, 0x45, 0x4e, 0x53, 0x45, 0x53, 0x10, 0x05, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x72, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67,
	0x2f, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledger_proto_rawDescOnce sync.Once
	file_ledger_proto_rawDescData = file_ledger_proto_rawDesc
)

func file_ledger_proto_rawDescGZIP() []byte {
	file_ledger_proto_rawDescOnce.Do(func() {
		file_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledger_proto_rawDescData)
	})
	return file_ledger_proto_rawDescData
}

var file_ledger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ledger_proto_goTypes = []interface{}{
	(AccountGroup)(0),     // 0: AccountGroup
	(*AccountNumber)(nil), // 1: AccountNumber
	(*Record)(nil),        // 2: Record
	(*Transaction)(nil),   // 3: Transaction
	(*Timestamp)(nil),     // 4: Timestamp
}
var file_ledger_proto_depIdxs = []int32{
	0, // 0: AccountNumber.group:type_name -> AccountGroup
	1, // 1: Record.account_number:type_name -> AccountNumber
	4, // 2: Transaction.timestamp:type_name -> Timestamp
	2, // 3: Transaction.records_debit:type_name -> Record
	2, // 4: Transaction.records_credit:type_name -> Record
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ledger_proto_init() }
func file_ledger_proto_init() {
	if File_ledger_proto != nil {
		return
	}
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountNumber); i {
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
		file_ledger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_ledger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
	file_ledger_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ledger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledger_proto_goTypes,
		DependencyIndexes: file_ledger_proto_depIdxs,
		EnumInfos:         file_ledger_proto_enumTypes,
		MessageInfos:      file_ledger_proto_msgTypes,
	}.Build()
	File_ledger_proto = out.File
	file_ledger_proto_rawDesc = nil
	file_ledger_proto_goTypes = nil
	file_ledger_proto_depIdxs = nil
}
