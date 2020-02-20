// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/exchange/exchange.proto

package exchange

import (
	context "context"
	fmt "fmt"
	proto "github.com/tron-us/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("protos/exchange/exchange.proto", fileDescriptor_b8e4cc6fcea5eff3) }

var fileDescriptor_b8e4cc6fcea5eff3 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x40, 0x17, 0x04, 0x91, 0x07, 0x40, 0x1e, 0x10, 0xcd, 0x00, 0x85, 0x81, 0xb1, 0x48, 0xb0,
	0x23, 0x45, 0x85, 0x99, 0x4f, 0x81, 0x18, 0x10, 0x3a, 0xcc, 0x41, 0x3c, 0xf4, 0x1c, 0x7c, 0x17,
	0x3e, 0x7e, 0x22, 0xff, 0x0a, 0xa9, 0xd8, 0x49, 0xeb, 0xba, 0xdd, 0x12, 0xbd, 0x97, 0xe7, 0x53,
	0xce, 0x6a, 0xaf, 0xf1, 0x4e, 0x1c, 0x1f, 0xe3, 0xb7, 0xa9, 0x81, 0xde, 0xb1, 0x7b, 0x18, 0x4d,
	0x81, 0x2e, 0xe2, 0x7b, 0x79, 0xb4, 0xcc, 0x7c, 0x9e, 0x20, 0x33, 0xc4, 0x2f, 0x4e, 0x7e, 0xd7,
	0x54, 0x71, 0x11, 0x90, 0xbe, 0x57, 0x5b, 0x57, 0x1e, 0x1b, 0xf0, 0xf8, 0x60, 0xa5, 0x7e, 0xf5,
	0xf0, 0xa5, 0x87, 0xa3, 0xee, 0x88, 0x04, 0xdd, 0xe0, 0x47, 0x8b, 0x2c, 0xe5, 0xc1, 0x0a, 0x83,
	0x1b, 0x47, 0x8c, 0xba, 0x52, 0x45, 0x17, 0x1c, 0xf4, 0x7a, 0x5a, 0x2a, 0x73, 0x28, 0x24, 0x6e,
	0xd5, 0x66, 0xa8, 0x9f, 0x63, 0xe3, 0xd8, 0x8a, 0xde, 0x5f, 0x38, 0x37, 0x90, 0x98, 0x1b, 0x2e,
	0x17, 0x42, 0xf4, 0x4c, 0x6d, 0xc4, 0xda, 0x6e, 0x2f, 0x27, 0x99, 0x41, 0x86, 0xf4, 0x43, 0x8d,
	0x1d, 0xbd, 0x59, 0x3f, 0xc9, 0x0c, 0x35, 0x4f, 0x32, 0x43, 0xa5, 0x42, 0x88, 0x3e, 0xaa, 0xed,
	0xeb, 0x16, 0xfd, 0xcf, 0x9d, 0x07, 0x62, 0x30, 0x62, 0x1d, 0xe9, 0x99, 0x7f, 0x9c, 0xb2, 0x18,
	0x3e, 0x5c, 0xa5, 0x84, 0xf4, 0x93, 0xda, 0xa9, 0x8c, 0xd8, 0x4f, 0x10, 0xac, 0x8c, 0x71, 0x2d,
	0xc9, 0x25, 0x8d, 0x6b, 0xb0, 0x34, 0xbb, 0xe6, 0xc4, 0xc8, 0xac, 0x79, 0xc1, 0xf8, 0xcf, 0xbf,
	0xac, 0x4f, 0xaf, 0xd4, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x58, 0x0c, 0xa2, 0xa6,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	PrepareWithdraw(ctx context.Context, in *PrepareWithdrawRequest, opts ...grpc.CallOption) (*PrepareWithdrawResponse, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	PrepareDeposit(ctx context.Context, in *PrepareDepositRequest, opts ...grpc.CallOption) (*PrepareDepositResponse, error)
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error)
	QueryTransaction(ctx context.Context, in *QueryTransactionRequest, opts ...grpc.CallOption) (*QueryTransactionResponse, error)
	ActivateAccountOnChain(ctx context.Context, in *ActivateAccountRequest, opts ...grpc.CallOption) (*ActivateAccountResponse, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) PrepareWithdraw(ctx context.Context, in *PrepareWithdrawRequest, opts ...grpc.CallOption) (*PrepareWithdrawResponse, error) {
	out := new(PrepareWithdrawResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/PrepareWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) PrepareDeposit(ctx context.Context, in *PrepareDepositRequest, opts ...grpc.CallOption) (*PrepareDepositResponse, error) {
	out := new(PrepareDepositResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/PrepareDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error) {
	out := new(ConfirmDepositResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/ConfirmDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) QueryTransaction(ctx context.Context, in *QueryTransactionRequest, opts ...grpc.CallOption) (*QueryTransactionResponse, error) {
	out := new(QueryTransactionResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/QueryTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) ActivateAccountOnChain(ctx context.Context, in *ActivateAccountRequest, opts ...grpc.CallOption) (*ActivateAccountResponse, error) {
	out := new(ActivateAccountResponse)
	err := c.cc.Invoke(ctx, "/exchange.Exchange/ActivateAccountOnChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	PrepareWithdraw(context.Context, *PrepareWithdrawRequest) (*PrepareWithdrawResponse, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	PrepareDeposit(context.Context, *PrepareDepositRequest) (*PrepareDepositResponse, error)
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	ConfirmDeposit(context.Context, *ConfirmDepositRequest) (*ConfirmDepositResponse, error)
	QueryTransaction(context.Context, *QueryTransactionRequest) (*QueryTransactionResponse, error)
	ActivateAccountOnChain(context.Context, *ActivateAccountRequest) (*ActivateAccountResponse, error)
}

// UnimplementedExchangeServer can be embedded to have forward compatible implementations.
type UnimplementedExchangeServer struct {
}

func (*UnimplementedExchangeServer) PrepareWithdraw(ctx context.Context, req *PrepareWithdrawRequest) (*PrepareWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareWithdraw not implemented")
}
func (*UnimplementedExchangeServer) Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (*UnimplementedExchangeServer) PrepareDeposit(ctx context.Context, req *PrepareDepositRequest) (*PrepareDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDeposit not implemented")
}
func (*UnimplementedExchangeServer) Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedExchangeServer) ConfirmDeposit(ctx context.Context, req *ConfirmDepositRequest) (*ConfirmDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeposit not implemented")
}
func (*UnimplementedExchangeServer) QueryTransaction(ctx context.Context, req *QueryTransactionRequest) (*QueryTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTransaction not implemented")
}
func (*UnimplementedExchangeServer) ActivateAccountOnChain(ctx context.Context, req *ActivateAccountRequest) (*ActivateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateAccountOnChain not implemented")
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_PrepareWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).PrepareWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/PrepareWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).PrepareWithdraw(ctx, req.(*PrepareWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_PrepareDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).PrepareDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/PrepareDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).PrepareDeposit(ctx, req.(*PrepareDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_ConfirmDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).ConfirmDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/ConfirmDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).ConfirmDeposit(ctx, req.(*ConfirmDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_QueryTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).QueryTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/QueryTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).QueryTransaction(ctx, req.(*QueryTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_ActivateAccountOnChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).ActivateAccountOnChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.Exchange/ActivateAccountOnChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).ActivateAccountOnChain(ctx, req.(*ActivateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareWithdraw",
			Handler:    _Exchange_PrepareWithdraw_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Exchange_Withdraw_Handler,
		},
		{
			MethodName: "PrepareDeposit",
			Handler:    _Exchange_PrepareDeposit_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Exchange_Deposit_Handler,
		},
		{
			MethodName: "ConfirmDeposit",
			Handler:    _Exchange_ConfirmDeposit_Handler,
		},
		{
			MethodName: "QueryTransaction",
			Handler:    _Exchange_QueryTransaction_Handler,
		},
		{
			MethodName: "ActivateAccountOnChain",
			Handler:    _Exchange_ActivateAccountOnChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/exchange/exchange.proto",
}