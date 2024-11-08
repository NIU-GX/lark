// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: pb_wallet/wallet.proto

package pb_wallet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Wallet_GetBalances_FullMethodName = "/pb_wallet.Wallet/GetBalances"
	Wallet_Exchange_FullMethodName    = "/pb_wallet.Wallet/Exchange"
	Wallet_Transfer_FullMethodName    = "/pb_wallet.Wallet/Transfer"
	Wallet_Recharge_FullMethodName    = "/pb_wallet.Wallet/Recharge"
	Wallet_Withdraw_FullMethodName    = "/pb_wallet.Wallet/Withdraw"
)

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	GetBalances(ctx context.Context, in *GetBalancesReq, opts ...grpc.CallOption) (*GetBalancesResp, error)
	Exchange(ctx context.Context, in *ExchangeReq, opts ...grpc.CallOption) (*ExchangeResp, error)
	Transfer(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error)
	Recharge(ctx context.Context, in *RechargeReq, opts ...grpc.CallOption) (*RechargeResp, error)
	Withdraw(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawResp, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GetBalances(ctx context.Context, in *GetBalancesReq, opts ...grpc.CallOption) (*GetBalancesResp, error) {
	out := new(GetBalancesResp)
	err := c.cc.Invoke(ctx, Wallet_GetBalances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Exchange(ctx context.Context, in *ExchangeReq, opts ...grpc.CallOption) (*ExchangeResp, error) {
	out := new(ExchangeResp)
	err := c.cc.Invoke(ctx, Wallet_Exchange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Transfer(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error) {
	out := new(TransferResp)
	err := c.cc.Invoke(ctx, Wallet_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Recharge(ctx context.Context, in *RechargeReq, opts ...grpc.CallOption) (*RechargeResp, error) {
	out := new(RechargeResp)
	err := c.cc.Invoke(ctx, Wallet_Recharge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Withdraw(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawResp, error) {
	out := new(WithdrawResp)
	err := c.cc.Invoke(ctx, Wallet_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	GetBalances(context.Context, *GetBalancesReq) (*GetBalancesResp, error)
	Exchange(context.Context, *ExchangeReq) (*ExchangeResp, error)
	Transfer(context.Context, *TransferReq) (*TransferResp, error)
	Recharge(context.Context, *RechargeReq) (*RechargeResp, error)
	Withdraw(context.Context, *WithdrawReq) (*WithdrawResp, error)
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) GetBalances(context.Context, *GetBalancesReq) (*GetBalancesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalances not implemented")
}
func (UnimplementedWalletServer) Exchange(context.Context, *ExchangeReq) (*ExchangeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedWalletServer) Transfer(context.Context, *TransferReq) (*TransferResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedWalletServer) Recharge(context.Context, *RechargeReq) (*RechargeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recharge not implemented")
}
func (UnimplementedWalletServer) Withdraw(context.Context, *WithdrawReq) (*WithdrawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_GetBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalancesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetBalances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetBalances(ctx, req.(*GetBalancesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_Exchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Exchange(ctx, req.(*ExchangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Transfer(ctx, req.(*TransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Recharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Recharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_Recharge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Recharge(ctx, req.(*RechargeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Withdraw(ctx, req.(*WithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_wallet.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalances",
			Handler:    _Wallet_GetBalances_Handler,
		},
		{
			MethodName: "Exchange",
			Handler:    _Wallet_Exchange_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Wallet_Transfer_Handler,
		},
		{
			MethodName: "Recharge",
			Handler:    _Wallet_Recharge_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Wallet_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_wallet/wallet.proto",
}
