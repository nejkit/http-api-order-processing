// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: orders_enums.proto

package orders

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

type Direction int32

const (
	Direction_DIRECTION_TYPE_NONE Direction = 0
	Direction_DIRECTION_TYPE_SELL Direction = 1
	Direction_DIRECTION_TYPE_BUY  Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "DIRECTION_TYPE_NONE",
		1: "DIRECTION_TYPE_SELL",
		2: "DIRECTION_TYPE_BUY",
	}
	Direction_value = map[string]int32{
		"DIRECTION_TYPE_NONE": 0,
		"DIRECTION_TYPE_SELL": 1,
		"DIRECTION_TYPE_BUY":  2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_enums_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_orders_enums_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_orders_enums_proto_rawDescGZIP(), []int{0}
}

type OrderType int32

const (
	OrderType_ORDER_TYPE_NONE   OrderType = 0
	OrderType_ORDER_TYPE_MARKET OrderType = 1
	OrderType_ORDER_TYPE_LIMIT  OrderType = 2
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "ORDER_TYPE_NONE",
		1: "ORDER_TYPE_MARKET",
		2: "ORDER_TYPE_LIMIT",
	}
	OrderType_value = map[string]int32{
		"ORDER_TYPE_NONE":   0,
		"ORDER_TYPE_MARKET": 1,
		"ORDER_TYPE_LIMIT":  2,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_enums_proto_enumTypes[1].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_orders_enums_proto_enumTypes[1]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_orders_enums_proto_rawDescGZIP(), []int{1}
}

type OrderState int32

const (
	OrderState_ORDER_STATE_NONE       OrderState = 0
	OrderState_ORDER_STATE_NEW        OrderState = 1
	OrderState_ORDER_STATE_IN_PROCESS OrderState = 2
	OrderState_ORDER_STATE_PART_FILL  OrderState = 3
	OrderState_ORDER_STATE_FILL       OrderState = 4
	OrderState_ORDER_STATE_DONE       OrderState = 5
	OrderState_ORDER_STATE_REJECT     OrderState = 6
)

// Enum value maps for OrderState.
var (
	OrderState_name = map[int32]string{
		0: "ORDER_STATE_NONE",
		1: "ORDER_STATE_NEW",
		2: "ORDER_STATE_IN_PROCESS",
		3: "ORDER_STATE_PART_FILL",
		4: "ORDER_STATE_FILL",
		5: "ORDER_STATE_DONE",
		6: "ORDER_STATE_REJECT",
	}
	OrderState_value = map[string]int32{
		"ORDER_STATE_NONE":       0,
		"ORDER_STATE_NEW":        1,
		"ORDER_STATE_IN_PROCESS": 2,
		"ORDER_STATE_PART_FILL":  3,
		"ORDER_STATE_FILL":       4,
		"ORDER_STATE_DONE":       5,
		"ORDER_STATE_REJECT":     6,
	}
)

func (x OrderState) Enum() *OrderState {
	p := new(OrderState)
	*p = x
	return p
}

func (x OrderState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderState) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_enums_proto_enumTypes[2].Descriptor()
}

func (OrderState) Type() protoreflect.EnumType {
	return &file_orders_enums_proto_enumTypes[2]
}

func (x OrderState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderState.Descriptor instead.
func (OrderState) EnumDescriptor() ([]byte, []int) {
	return file_orders_enums_proto_rawDescGZIP(), []int{2}
}

type MatchState int32

const (
	MatchState_MATCH_STATE_NONE        MatchState = 0
	MatchState_MATCH_STATE_NEW         MatchState = 1
	MatchState_MATCH_STATE_IN_PROGRESS MatchState = 2
	MatchState_MATCH_STATE_DONE        MatchState = 3
	MatchState_MATCH_STATE_REJECT      MatchState = 4
)

// Enum value maps for MatchState.
var (
	MatchState_name = map[int32]string{
		0: "MATCH_STATE_NONE",
		1: "MATCH_STATE_NEW",
		2: "MATCH_STATE_IN_PROGRESS",
		3: "MATCH_STATE_DONE",
		4: "MATCH_STATE_REJECT",
	}
	MatchState_value = map[string]int32{
		"MATCH_STATE_NONE":        0,
		"MATCH_STATE_NEW":         1,
		"MATCH_STATE_IN_PROGRESS": 2,
		"MATCH_STATE_DONE":        3,
		"MATCH_STATE_REJECT":      4,
	}
)

func (x MatchState) Enum() *MatchState {
	p := new(MatchState)
	*p = x
	return p
}

func (x MatchState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchState) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_enums_proto_enumTypes[3].Descriptor()
}

func (MatchState) Type() protoreflect.EnumType {
	return &file_orders_enums_proto_enumTypes[3]
}

func (x MatchState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchState.Descriptor instead.
func (MatchState) EnumDescriptor() ([]byte, []int) {
	return file_orders_enums_proto_rawDescGZIP(), []int{3}
}

var File_orders_enums_proto protoreflect.FileDescriptor

var file_orders_enums_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2a, 0x55, 0x0a, 0x09,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55,
	0x59, 0x10, 0x02, 0x2a, 0x4d, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54,
	0x10, 0x02, 0x2a, 0xb2, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x5f, 0x46, 0x49, 0x4c,
	0x4c, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x05, 0x12,
	0x16, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x06, 0x2a, 0x82, 0x01, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x04, 0x42, 0x09, 0x5a, 0x07,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_enums_proto_rawDescOnce sync.Once
	file_orders_enums_proto_rawDescData = file_orders_enums_proto_rawDesc
)

func file_orders_enums_proto_rawDescGZIP() []byte {
	file_orders_enums_proto_rawDescOnce.Do(func() {
		file_orders_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_enums_proto_rawDescData)
	})
	return file_orders_enums_proto_rawDescData
}

var file_orders_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_orders_enums_proto_goTypes = []interface{}{
	(Direction)(0),  // 0: orders.Direction
	(OrderType)(0),  // 1: orders.OrderType
	(OrderState)(0), // 2: orders.OrderState
	(MatchState)(0), // 3: orders.MatchState
}
var file_orders_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orders_enums_proto_init() }
func file_orders_enums_proto_init() {
	if File_orders_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orders_enums_proto_goTypes,
		DependencyIndexes: file_orders_enums_proto_depIdxs,
		EnumInfos:         file_orders_enums_proto_enumTypes,
	}.Build()
	File_orders_enums_proto = out.File
	file_orders_enums_proto_rawDesc = nil
	file_orders_enums_proto_goTypes = nil
	file_orders_enums_proto_depIdxs = nil
}
