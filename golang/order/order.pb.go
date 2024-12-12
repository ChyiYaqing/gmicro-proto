// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateOrderRequest struct {
	UserId               int64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderItems           []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateOrderRequest) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

type OrderItem struct {
	ProductCode          string   `protobuf:"bytes,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	UintPrice            float32  `protobuf:"fixed32,2,opt,name=uint_price,json=uintPrice,proto3" json:"uint_price,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{1}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetProductCode() string {
	if m != nil {
		return m.ProductCode
	}
	return ""
}

func (m *OrderItem) GetUintPrice() float32 {
	if m != nil {
		return m.UintPrice
	}
	return 0
}

func (m *OrderItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type CreateOrderResponse struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{2}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type GetOrderRequest struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderRequest) Reset()         { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()    {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{3}
}

func (m *GetOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderRequest.Unmarshal(m, b)
}
func (m *GetOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderRequest.Merge(m, src)
}
func (m *GetOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderRequest.Size(m)
}
func (m *GetOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderRequest proto.InternalMessageInfo

func (m *GetOrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type GetOrderResponse struct {
	UserId               int64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderItems           []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetOrderResponse) Reset()         { *m = GetOrderResponse{} }
func (m *GetOrderResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderResponse) ProtoMessage()    {}
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{4}
}

func (m *GetOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResponse.Unmarshal(m, b)
}
func (m *GetOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResponse.Merge(m, src)
}
func (m *GetOrderResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderResponse.Size(m)
}
func (m *GetOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResponse proto.InternalMessageInfo

func (m *GetOrderResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetOrderResponse) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "CreateOrderRequest")
	proto.RegisterType((*OrderItem)(nil), "OrderItem")
	proto.RegisterType((*CreateOrderResponse)(nil), "CreateOrderResponse")
	proto.RegisterType((*GetOrderRequest)(nil), "GetOrderRequest")
	proto.RegisterType((*GetOrderResponse)(nil), "GetOrderResponse")
}

func init() {
	proto.RegisterFile("order/order.proto", fileDescriptor_fa47a2077d8980ed)
}

var fileDescriptor_fa47a2077d8980ed = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0x29, 0x0d, 0x5f, 0x83, 0x89, 0xb0, 0x98, 0x58, 0x49, 0x4c, 0x6a, 0x4f, 0x8d, 0xe2,
	0x62, 0x30, 0xfe, 0x01, 0x39, 0x10, 0x4e, 0x9a, 0x9e, 0x0c, 0x97, 0x06, 0xba, 0x13, 0xd8, 0xc4,
	0x76, 0xcb, 0x76, 0xf6, 0xc0, 0xbf, 0x37, 0x5d, 0x0a, 0x06, 0x34, 0x9e, 0xbc, 0x6c, 0x76, 0xde,
	0xf9, 0xc8, 0xfb, 0xec, 0x0e, 0xf4, 0x95, 0x16, 0xa8, 0xc7, 0xf6, 0xe4, 0xb9, 0x56, 0xa4, 0x82,
	0x05, 0xb0, 0xa9, 0xc6, 0x25, 0xe1, 0x5b, 0x29, 0x46, 0xb8, 0x35, 0x58, 0x10, 0xbb, 0x86, 0x96,
	0x29, 0x50, 0xc7, 0x52, 0x78, 0x8e, 0xef, 0x84, 0x6e, 0xd4, 0x2c, 0xc3, 0xb9, 0x60, 0x0f, 0xd0,
	0xb5, 0xdd, 0xb1, 0x24, 0x4c, 0x0b, 0xaf, 0xee, 0xbb, 0x61, 0x77, 0x02, 0xdc, 0x36, 0xcf, 0x09,
	0xd3, 0x08, 0xd4, 0xe1, 0x5a, 0x04, 0x12, 0x3a, 0xc7, 0x04, 0xbb, 0x83, 0x8b, 0x5c, 0x2b, 0x61,
	0x12, 0x8a, 0x13, 0x25, 0xd0, 0xce, 0xed, 0x44, 0xdd, 0x4a, 0x9b, 0x2a, 0x81, 0xec, 0x16, 0xc0,
	0xc8, 0x8c, 0xe2, 0x5c, 0xcb, 0x04, 0xbd, 0xba, 0xef, 0x84, 0xf5, 0xa8, 0x53, 0x2a, 0xef, 0xa5,
	0xc0, 0x86, 0xd0, 0xde, 0x9a, 0x65, 0x46, 0x92, 0x76, 0x9e, 0xeb, 0x3b, 0x61, 0x23, 0x3a, 0xc6,
	0xc1, 0x13, 0x0c, 0x4e, 0x30, 0x8a, 0x5c, 0x65, 0x05, 0xb2, 0x1b, 0x68, 0x57, 0x76, 0x0f, 0x20,
	0xad, 0xbd, 0x3f, 0x11, 0x8c, 0xe0, 0x72, 0x86, 0x74, 0x42, 0xfd, 0x47, 0xf5, 0x07, 0xf4, 0xbe,
	0xab, 0xab, 0xe1, 0xff, 0xf2, 0x48, 0x93, 0x4f, 0x68, 0xd8, 0x04, 0x7b, 0x81, 0xe6, 0x1e, 0x81,
	0x0d, 0xf8, 0xcf, 0x2f, 0x19, 0x5e, 0xf1, 0x5f, 0x00, 0x83, 0x1a, 0x1b, 0x81, 0x3b, 0x43, 0x62,
	0x3d, 0x7e, 0x46, 0x33, 0xec, 0xf3, 0x73, 0xc7, 0x41, 0xed, 0xf5, 0x7e, 0x11, 0xae, 0x25, 0x6d,
	0xcc, 0x8a, 0x27, 0x2a, 0x1d, 0x27, 0x9b, 0x9d, 0xdc, 0x2d, 0xb7, 0x32, 0x5b, 0x8f, 0xd7, 0xa9,
	0x4c, 0xb4, 0x7a, 0xb4, 0x4b, 0xb1, 0x5f, 0x90, 0x55, 0xd3, 0x06, 0xcf, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x67, 0x00, 0xea, 0x39, 0x36, 0x02, 0x00, 0x00,
}
