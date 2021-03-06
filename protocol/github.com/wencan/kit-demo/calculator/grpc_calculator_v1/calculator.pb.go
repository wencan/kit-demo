// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package grpc_calculator_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CalculatorAddRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorAddRequest) Reset()         { *m = CalculatorAddRequest{} }
func (m *CalculatorAddRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorAddRequest) ProtoMessage()    {}
func (*CalculatorAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *CalculatorAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorAddRequest.Unmarshal(m, b)
}
func (m *CalculatorAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorAddRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorAddRequest.Merge(m, src)
}
func (m *CalculatorAddRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorAddRequest.Size(m)
}
func (m *CalculatorAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorAddRequest proto.InternalMessageInfo

func (m *CalculatorAddRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *CalculatorAddRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type CalculatorSubRequest struct {
	C                    int32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	D                    int32    `protobuf:"varint,2,opt,name=d,proto3" json:"d,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorSubRequest) Reset()         { *m = CalculatorSubRequest{} }
func (m *CalculatorSubRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorSubRequest) ProtoMessage()    {}
func (*CalculatorSubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *CalculatorSubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorSubRequest.Unmarshal(m, b)
}
func (m *CalculatorSubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorSubRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorSubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorSubRequest.Merge(m, src)
}
func (m *CalculatorSubRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorSubRequest.Size(m)
}
func (m *CalculatorSubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorSubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorSubRequest proto.InternalMessageInfo

func (m *CalculatorSubRequest) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *CalculatorSubRequest) GetD() int32 {
	if m != nil {
		return m.D
	}
	return 0
}

type CalculatorMulRequest struct {
	E                    int32    `protobuf:"varint,1,opt,name=e,proto3" json:"e,omitempty"`
	F                    int32    `protobuf:"varint,2,opt,name=f,proto3" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorMulRequest) Reset()         { *m = CalculatorMulRequest{} }
func (m *CalculatorMulRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorMulRequest) ProtoMessage()    {}
func (*CalculatorMulRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *CalculatorMulRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorMulRequest.Unmarshal(m, b)
}
func (m *CalculatorMulRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorMulRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorMulRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorMulRequest.Merge(m, src)
}
func (m *CalculatorMulRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorMulRequest.Size(m)
}
func (m *CalculatorMulRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorMulRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorMulRequest proto.InternalMessageInfo

func (m *CalculatorMulRequest) GetE() int32 {
	if m != nil {
		return m.E
	}
	return 0
}

func (m *CalculatorMulRequest) GetF() int32 {
	if m != nil {
		return m.F
	}
	return 0
}

type CalculatorDivRequest struct {
	M                    int32    `protobuf:"varint,1,opt,name=m,proto3" json:"m,omitempty"`
	N                    int32    `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorDivRequest) Reset()         { *m = CalculatorDivRequest{} }
func (m *CalculatorDivRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorDivRequest) ProtoMessage()    {}
func (*CalculatorDivRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{3}
}

func (m *CalculatorDivRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorDivRequest.Unmarshal(m, b)
}
func (m *CalculatorDivRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorDivRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorDivRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorDivRequest.Merge(m, src)
}
func (m *CalculatorDivRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorDivRequest.Size(m)
}
func (m *CalculatorDivRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorDivRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorDivRequest proto.InternalMessageInfo

func (m *CalculatorDivRequest) GetM() int32 {
	if m != nil {
		return m.M
	}
	return 0
}

func (m *CalculatorDivRequest) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

type CalculatorInt32Response struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorInt32Response) Reset()         { *m = CalculatorInt32Response{} }
func (m *CalculatorInt32Response) String() string { return proto.CompactTextString(m) }
func (*CalculatorInt32Response) ProtoMessage()    {}
func (*CalculatorInt32Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{4}
}

func (m *CalculatorInt32Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorInt32Response.Unmarshal(m, b)
}
func (m *CalculatorInt32Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorInt32Response.Marshal(b, m, deterministic)
}
func (m *CalculatorInt32Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorInt32Response.Merge(m, src)
}
func (m *CalculatorInt32Response) XXX_Size() int {
	return xxx_messageInfo_CalculatorInt32Response.Size(m)
}
func (m *CalculatorInt32Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorInt32Response.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorInt32Response proto.InternalMessageInfo

func (m *CalculatorInt32Response) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatorFloatResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorFloatResponse) Reset()         { *m = CalculatorFloatResponse{} }
func (m *CalculatorFloatResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorFloatResponse) ProtoMessage()    {}
func (*CalculatorFloatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{5}
}

func (m *CalculatorFloatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorFloatResponse.Unmarshal(m, b)
}
func (m *CalculatorFloatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorFloatResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorFloatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorFloatResponse.Merge(m, src)
}
func (m *CalculatorFloatResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorFloatResponse.Size(m)
}
func (m *CalculatorFloatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorFloatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorFloatResponse proto.InternalMessageInfo

func (m *CalculatorFloatResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculatorAddRequest)(nil), "grpc.calculator.v1.CalculatorAddRequest")
	proto.RegisterType((*CalculatorSubRequest)(nil), "grpc.calculator.v1.CalculatorSubRequest")
	proto.RegisterType((*CalculatorMulRequest)(nil), "grpc.calculator.v1.CalculatorMulRequest")
	proto.RegisterType((*CalculatorDivRequest)(nil), "grpc.calculator.v1.CalculatorDivRequest")
	proto.RegisterType((*CalculatorInt32Response)(nil), "grpc.calculator.v1.CalculatorInt32Response")
	proto.RegisterType((*CalculatorFloatResponse)(nil), "grpc.calculator.v1.CalculatorFloatResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd3, 0x4d, 0x4b, 0xfb, 0x30,
	0x1c, 0x07, 0x70, 0xba, 0xf1, 0xdf, 0x21, 0xfc, 0x41, 0x09, 0x3e, 0x0c, 0x4f, 0xb2, 0xd3, 0x40,
	0x4c, 0xe9, 0x76, 0xf1, 0xba, 0x39, 0x14, 0x0f, 0x83, 0xb1, 0x81, 0x07, 0x19, 0x8c, 0x3c, 0x6d,
	0x16, 0xd3, 0xa4, 0xb6, 0x49, 0xc5, 0xbb, 0xaf, 0xc6, 0x57, 0x29, 0xd1, 0xc5, 0x34, 0x38, 0x2d,
	0x3b, 0x7e, 0x43, 0x3f, 0xdf, 0x96, 0xdf, 0xaf, 0x01, 0x87, 0x14, 0x0b, 0x6a, 0x04, 0xd6, 0xaa,
	0x40, 0x79, 0xa1, 0xb4, 0x82, 0x70, 0x53, 0xe4, 0x14, 0xd5, 0x8e, 0xab, 0xa4, 0x37, 0x00, 0x47,
	0xd7, 0xdf, 0x07, 0x23, 0xc6, 0xe6, 0xfc, 0xd9, 0xf0, 0x52, 0xc3, 0xff, 0x20, 0xc2, 0xdd, 0xe8,
	0x3c, 0xea, 0xff, 0x9b, 0x47, 0xd8, 0x26, 0xd2, 0x6d, 0x7d, 0x25, 0x12, 0x9a, 0x85, 0x21, 0x35,
	0x43, 0x9d, 0xa1, 0x36, 0x31, 0x67, 0x58, 0x68, 0xa6, 0x46, 0xd4, 0x0c, 0x77, 0x86, 0xdb, 0xb4,
	0x76, 0x66, 0x1d, 0x9a, 0x49, 0x5a, 0xd5, 0x4c, 0xe6, 0x4c, 0x66, 0x93, 0x74, 0x46, 0xf6, 0x12,
	0x70, 0xea, 0xcd, 0x9d, 0xd4, 0xc3, 0xc1, 0x9c, 0x97, 0xb9, 0x92, 0x25, 0x87, 0x27, 0xa0, 0x53,
	0xf0, 0xd2, 0x08, 0xbd, 0xb5, 0xdb, 0x14, 0x92, 0x1b, 0xa1, 0xb0, 0xfe, 0x85, 0xb4, 0x1c, 0x19,
	0xbc, 0xb5, 0x01, 0xf0, 0x06, 0x2e, 0x41, 0x7b, 0xc4, 0x18, 0xec, 0xa3, 0x9f, 0x03, 0x46, 0xbb,
	0xa6, 0x7b, 0x76, 0xf1, 0xf7, 0x93, 0xe1, 0x77, 0x2f, 0x41, 0x7b, 0x61, 0x48, 0x53, 0xbb, 0xdf,
	0xc3, 0xde, 0xed, 0x53, 0x23, 0x9a, 0xda, 0xfd, 0xc6, 0xf6, 0x6e, 0x9f, 0xa4, 0x55, 0x53, 0xbb,
	0xdf, 0x6d, 0x53, 0x7b, 0xb0, 0x9e, 0xf1, 0x2b, 0x38, 0x4e, 0xd5, 0x0e, 0x30, 0x3e, 0xf0, 0x62,
	0x66, 0x7f, 0xfd, 0x59, 0xf4, 0x70, 0xb5, 0x49, 0xf5, 0xa3, 0x21, 0x88, 0xaa, 0x2c, 0x7e, 0xe1,
	0x92, 0x62, 0x19, 0x3f, 0xa5, 0xfa, 0x92, 0xf1, 0x4c, 0xc5, 0x1e, 0xc7, 0xb6, 0x6c, 0xe5, 0xf3,
	0xaa, 0x4a, 0xde, 0x5b, 0xf0, 0xd6, 0xbe, 0xc1, 0x57, 0xa2, 0xfb, 0x84, 0x74, 0x3e, 0xaf, 0xd4,
	0xf0, 0x23, 0x00, 0x00, 0xff, 0xff, 0x11, 0xc6, 0xeb, 0x05, 0x66, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *CalculatorAddRequest, opts ...grpc.CallOption) (*CalculatorInt32Response, error)
	Sub(ctx context.Context, in *CalculatorSubRequest, opts ...grpc.CallOption) (*CalculatorInt32Response, error)
	Mul(ctx context.Context, in *CalculatorMulRequest, opts ...grpc.CallOption) (*CalculatorInt32Response, error)
	Div(ctx context.Context, in *CalculatorDivRequest, opts ...grpc.CallOption) (*CalculatorFloatResponse, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *CalculatorAddRequest, opts ...grpc.CallOption) (*CalculatorInt32Response, error) {
	out := new(CalculatorInt32Response)
	err := c.cc.Invoke(ctx, "/grpc.calculator.v1.Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Sub(ctx context.Context, in *CalculatorSubRequest, opts ...grpc.CallOption) (*CalculatorInt32Response, error) {
	out := new(CalculatorInt32Response)
	err := c.cc.Invoke(ctx, "/grpc.calculator.v1.Calculator/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Mul(ctx context.Context, in *CalculatorMulRequest, opts ...grpc.CallOption) (*CalculatorInt32Response, error) {
	out := new(CalculatorInt32Response)
	err := c.cc.Invoke(ctx, "/grpc.calculator.v1.Calculator/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Div(ctx context.Context, in *CalculatorDivRequest, opts ...grpc.CallOption) (*CalculatorFloatResponse, error) {
	out := new(CalculatorFloatResponse)
	err := c.cc.Invoke(ctx, "/grpc.calculator.v1.Calculator/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Add(context.Context, *CalculatorAddRequest) (*CalculatorInt32Response, error)
	Sub(context.Context, *CalculatorSubRequest) (*CalculatorInt32Response, error)
	Mul(context.Context, *CalculatorMulRequest) (*CalculatorInt32Response, error)
	Div(context.Context, *CalculatorDivRequest) (*CalculatorFloatResponse, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Add(ctx context.Context, req *CalculatorAddRequest) (*CalculatorInt32Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculatorServer) Sub(ctx context.Context, req *CalculatorSubRequest) (*CalculatorInt32Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (*UnimplementedCalculatorServer) Mul(ctx context.Context, req *CalculatorMulRequest) (*CalculatorInt32Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (*UnimplementedCalculatorServer) Div(ctx context.Context, req *CalculatorDivRequest) (*CalculatorFloatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.calculator.v1.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*CalculatorAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorSubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.calculator.v1.Calculator/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sub(ctx, req.(*CalculatorSubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorMulRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.calculator.v1.Calculator/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Mul(ctx, req.(*CalculatorMulRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorDivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.calculator.v1.Calculator/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Div(ctx, req.(*CalculatorDivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.calculator.v1.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Calculator_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _Calculator_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _Calculator_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
