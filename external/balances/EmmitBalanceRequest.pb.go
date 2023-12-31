// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: EmmitBalanceRequest.proto

package balances

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

type EmmitBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address  string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *EmmitBalanceRequest) Reset() {
	*x = EmmitBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EmmitBalanceRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmmitBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmmitBalanceRequest) ProtoMessage() {}

func (x *EmmitBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EmmitBalanceRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmmitBalanceRequest.ProtoReflect.Descriptor instead.
func (*EmmitBalanceRequest) Descriptor() ([]byte, []int) {
	return file_EmmitBalanceRequest_proto_rawDescGZIP(), []int{0}
}

func (x *EmmitBalanceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmmitBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EmmitBalanceRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *EmmitBalanceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_EmmitBalanceRequest_proto protoreflect.FileDescriptor

var file_EmmitBalanceRequest_proto_rawDesc = []byte{
	0x0a, 0x19, 0x45, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x13, 0x45,
	0x6d, 0x6d, 0x69, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EmmitBalanceRequest_proto_rawDescOnce sync.Once
	file_EmmitBalanceRequest_proto_rawDescData = file_EmmitBalanceRequest_proto_rawDesc
)

func file_EmmitBalanceRequest_proto_rawDescGZIP() []byte {
	file_EmmitBalanceRequest_proto_rawDescOnce.Do(func() {
		file_EmmitBalanceRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_EmmitBalanceRequest_proto_rawDescData)
	})
	return file_EmmitBalanceRequest_proto_rawDescData
}

var file_EmmitBalanceRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EmmitBalanceRequest_proto_goTypes = []interface{}{
	(*EmmitBalanceRequest)(nil), // 0: EmmitBalanceRequest
}
var file_EmmitBalanceRequest_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EmmitBalanceRequest_proto_init() }
func file_EmmitBalanceRequest_proto_init() {
	if File_EmmitBalanceRequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EmmitBalanceRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmmitBalanceRequest); i {
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
			RawDescriptor: file_EmmitBalanceRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EmmitBalanceRequest_proto_goTypes,
		DependencyIndexes: file_EmmitBalanceRequest_proto_depIdxs,
		MessageInfos:      file_EmmitBalanceRequest_proto_msgTypes,
	}.Build()
	File_EmmitBalanceRequest_proto = out.File
	file_EmmitBalanceRequest_proto_rawDesc = nil
	file_EmmitBalanceRequest_proto_goTypes = nil
	file_EmmitBalanceRequest_proto_depIdxs = nil
}
