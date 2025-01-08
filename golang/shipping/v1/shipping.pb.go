// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: shipping/v1/shipping.proto

package v1

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

type CreateShippingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId       int64                  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Address       string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShippingRequest) Reset() {
	*x = CreateShippingRequest{}
	mi := &file_shipping_v1_shipping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShippingRequest) ProtoMessage() {}

func (x *CreateShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_v1_shipping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShippingRequest.ProtoReflect.Descriptor instead.
func (*CreateShippingRequest) Descriptor() ([]byte, []int) {
	return file_shipping_v1_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShippingRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateShippingRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *CreateShippingRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateShippingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShippingId    int64                  `protobuf:"varint,1,opt,name=shipping_id,json=shippingId,proto3" json:"shipping_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShippingResponse) Reset() {
	*x = CreateShippingResponse{}
	mi := &file_shipping_v1_shipping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShippingResponse) ProtoMessage() {}

func (x *CreateShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_v1_shipping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShippingResponse.ProtoReflect.Descriptor instead.
func (*CreateShippingResponse) Descriptor() ([]byte, []int) {
	return file_shipping_v1_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShippingResponse) GetShippingId() int64 {
	if x != nil {
		return x.ShippingId
	}
	return 0
}

var File_shipping_v1_shipping_proto protoreflect.FileDescriptor

var file_shipping_v1_shipping_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x39, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x32, 0x5f, 0x0a, 0x08, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x22, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x79, 0x69, 0x79,
	0x61, 0x71, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipping_v1_shipping_proto_rawDescOnce sync.Once
	file_shipping_v1_shipping_proto_rawDescData = file_shipping_v1_shipping_proto_rawDesc
)

func file_shipping_v1_shipping_proto_rawDescGZIP() []byte {
	file_shipping_v1_shipping_proto_rawDescOnce.Do(func() {
		file_shipping_v1_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipping_v1_shipping_proto_rawDescData)
	})
	return file_shipping_v1_shipping_proto_rawDescData
}

var file_shipping_v1_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shipping_v1_shipping_proto_goTypes = []any{
	(*CreateShippingRequest)(nil),  // 0: shipping.v1.CreateShippingRequest
	(*CreateShippingResponse)(nil), // 1: shipping.v1.CreateShippingResponse
}
var file_shipping_v1_shipping_proto_depIdxs = []int32{
	0, // 0: shipping.v1.Shipping.Create:input_type -> shipping.v1.CreateShippingRequest
	1, // 1: shipping.v1.Shipping.Create:output_type -> shipping.v1.CreateShippingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shipping_v1_shipping_proto_init() }
func file_shipping_v1_shipping_proto_init() {
	if File_shipping_v1_shipping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shipping_v1_shipping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipping_v1_shipping_proto_goTypes,
		DependencyIndexes: file_shipping_v1_shipping_proto_depIdxs,
		MessageInfos:      file_shipping_v1_shipping_proto_msgTypes,
	}.Build()
	File_shipping_v1_shipping_proto = out.File
	file_shipping_v1_shipping_proto_rawDesc = nil
	file_shipping_v1_shipping_proto_goTypes = nil
	file_shipping_v1_shipping_proto_depIdxs = nil
}
