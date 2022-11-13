// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: timestamp.proto

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

type TimeOfDay int32

const (
	// Night is the time between 00:00 and 05:59
	TimeOfDay_NIGHT TimeOfDay = 0
	// Morning is the time between 06:00 and 11:59
	TimeOfDay_MORNING TimeOfDay = 1
	// Afternoon is the time between 12:00 and 17:59
	TimeOfDay_AFTERNOON TimeOfDay = 2
	// Evening is the time between 18:00 and 23:59
	TimeOfDay_EVENING TimeOfDay = 3
)

// Enum value maps for TimeOfDay.
var (
	TimeOfDay_name = map[int32]string{
		0: "NIGHT",
		1: "MORNING",
		2: "AFTERNOON",
		3: "EVENING",
	}
	TimeOfDay_value = map[string]int32{
		"NIGHT":     0,
		"MORNING":   1,
		"AFTERNOON": 2,
		"EVENING":   3,
	}
)

func (x TimeOfDay) Enum() *TimeOfDay {
	p := new(TimeOfDay)
	*p = x
	return p
}

func (x TimeOfDay) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeOfDay) Descriptor() protoreflect.EnumDescriptor {
	return file_timestamp_proto_enumTypes[0].Descriptor()
}

func (TimeOfDay) Type() protoreflect.EnumType {
	return &file_timestamp_proto_enumTypes[0]
}

func (x TimeOfDay) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeOfDay.Descriptor instead.
func (TimeOfDay) EnumDescriptor() ([]byte, []int) {
	return file_timestamp_proto_rawDescGZIP(), []int{0}
}

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day       int32     `protobuf:"zigzag32,1,opt,name=day,proto3" json:"day,omitempty"`
	TimeOfDay TimeOfDay `protobuf:"varint,2,opt,name=time_of_day,json=timeOfDay,proto3,enum=TimeOfDay" json:"time_of_day,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timestamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_timestamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_timestamp_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Timestamp) GetTimeOfDay() TimeOfDay {
	if x != nil {
		return x.TimeOfDay
	}
	return TimeOfDay_NIGHT
}

var File_timestamp_proto protoreflect.FileDescriptor

var file_timestamp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x64, 0x61, 0x79,
	0x12, 0x2a, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61,
	0x79, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x2a, 0x3f, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x49, 0x47,
	0x48, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x46, 0x54, 0x45, 0x52, 0x4e, 0x4f, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x45, 0x56, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x72, 0x65, 0x6d,
	0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timestamp_proto_rawDescOnce sync.Once
	file_timestamp_proto_rawDescData = file_timestamp_proto_rawDesc
)

func file_timestamp_proto_rawDescGZIP() []byte {
	file_timestamp_proto_rawDescOnce.Do(func() {
		file_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_timestamp_proto_rawDescData)
	})
	return file_timestamp_proto_rawDescData
}

var file_timestamp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_timestamp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_timestamp_proto_goTypes = []interface{}{
	(TimeOfDay)(0),    // 0: TimeOfDay
	(*Timestamp)(nil), // 1: Timestamp
}
var file_timestamp_proto_depIdxs = []int32{
	0, // 0: Timestamp.time_of_day:type_name -> TimeOfDay
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_timestamp_proto_init() }
func file_timestamp_proto_init() {
	if File_timestamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timestamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
			RawDescriptor: file_timestamp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timestamp_proto_goTypes,
		DependencyIndexes: file_timestamp_proto_depIdxs,
		EnumInfos:         file_timestamp_proto_enumTypes,
		MessageInfos:      file_timestamp_proto_msgTypes,
	}.Build()
	File_timestamp_proto = out.File
	file_timestamp_proto_rawDesc = nil
	file_timestamp_proto_goTypes = nil
	file_timestamp_proto_depIdxs = nil
}