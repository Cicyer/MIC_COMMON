// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package BcAndBankService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BcAndBankServiceClient is the client API for BcAndBankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BcAndBankServiceClient interface {
	// --------单纯上链接口 start---------
	// 上链 configs 信息
	BcUploadConfigs(ctx context.Context, in *UploadChainConfigs, opts ...grpc.CallOption) (*BcResponses, error)
	// 上链 企业信息s
	BcUploadCompanies(ctx context.Context, in *UploadChainCompanies, opts ...grpc.CallOption) (*BcResponses, error)
	// 医疗机构上链
	BcUploadMis(ctx context.Context, in *UploadChainMis, opts ...grpc.CallOption) (*BcResponses, error)
	// 合同信息上链
	BcUploadContracts(ctx context.Context, in *UploadChainContracts, opts ...grpc.CallOption) (*BcResponses, error)
	// 子合同信息上链
	BcUploadContractMis(ctx context.Context, in *UploadChainContractMis, opts ...grpc.CallOption) (*BcResponses, error)
	// 配送计划上链
	BcUploadShipments(ctx context.Context, in *UploadChainShipments, opts ...grpc.CallOption) (*BcResponses, error)
	// 配送单上链
	BcUploadShipmentOrders(ctx context.Context, in *UploadChainShipmentOrders, opts ...grpc.CallOption) (*BcResponses, error)
	// 支付单上链, 确权
	BcUploadPayOrders(ctx context.Context, in *UploadChainPayOrders, opts ...grpc.CallOption) (*BcResponses, error)
	// 保理单上链（更新）
	BcUploadFactoringOrders(ctx context.Context, in *UploadChainFactoringOrders, opts ...grpc.CallOption) (*BcResponses, error)
	// --------上链 + 银行 组合接口 start---------
	// 发起付款 一笔付款单内可以包含多笔 支付单(接收参数: 多笔 payOrder + 一个 银行请求)
	// 查所有的支付单, 看是否符合要求
	// 返回:: 链上数据hash集合 加 银行返回
	BcAndBankPayOrderAction(ctx context.Context, in *BcAndBankPayOrderActionRequest, opts ...grpc.CallOption) (*BcAndBankActionResp, error)
	// 发起保理(多笔支付单)
	// 保理单上链, 支付单上链后才上链, 查所有的支付单: 确认 支付单 处于 待支付/失败/(私有状态 未保理) 中; 调银行
	BcAndBankFactoringOrdersAction(ctx context.Context, in *BcAndBankFactoringActionRequest, opts ...grpc.CallOption) (*BcAndBankActionResp, error)
	// 发起保理还款(以保理单为准)
	// 查保理单, 处于 已保理 状态 ->
	BcAndBankFactoringRepayAction(ctx context.Context, in *BcAndBankFactoringRepayActionRequest, opts ...grpc.CallOption) (*BcAndBankActionResp, error)
	// --------危险动作: 强制上链接口 start---------
	// 强制配送计划上链
	BcForceUploadShipments(ctx context.Context, in *UploadChainShipments, opts ...grpc.CallOption) (*BcResponses, error)
	// 强制配送单上链
	BcForceUploadShipmentOrders(ctx context.Context, in *UploadChainShipmentOrders, opts ...grpc.CallOption) (*BcResponses, error)
	// 强制支付单上链
	BcForceUploadPayOrders(ctx context.Context, in *UploadChainPayOrders, opts ...grpc.CallOption) (*BcResponses, error)
	// 强制Factoring上链
	BcForceUploadFactoring(ctx context.Context, in *UploadChainFactoringOrders, opts ...grpc.CallOption) (*BcResponses, error)
	// 根据 txID 查询 链上交易, 返回 解密后的 json.Marshal 字符串
	BcQueryByTxID(ctx context.Context, in *QueryTxIDReq, opts ...grpc.CallOption) (*QueryTxIDRsp, error)
}

type bcAndBankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBcAndBankServiceClient(cc grpc.ClientConnInterface) BcAndBankServiceClient {
	return &bcAndBankServiceClient{cc}
}

func (c *bcAndBankServiceClient) BcUploadConfigs(ctx context.Context, in *UploadChainConfigs, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadCompanies(ctx context.Context, in *UploadChainCompanies, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadMis(ctx context.Context, in *UploadChainMis, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadMis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadContracts(ctx context.Context, in *UploadChainContracts, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadContractMis(ctx context.Context, in *UploadChainContractMis, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadContractMis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadShipments(ctx context.Context, in *UploadChainShipments, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadShipments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadShipmentOrders(ctx context.Context, in *UploadChainShipmentOrders, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadShipmentOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadPayOrders(ctx context.Context, in *UploadChainPayOrders, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadPayOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcUploadFactoringOrders(ctx context.Context, in *UploadChainFactoringOrders, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcUploadFactoringOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcAndBankPayOrderAction(ctx context.Context, in *BcAndBankPayOrderActionRequest, opts ...grpc.CallOption) (*BcAndBankActionResp, error) {
	out := new(BcAndBankActionResp)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcAndBankPayOrderAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcAndBankFactoringOrdersAction(ctx context.Context, in *BcAndBankFactoringActionRequest, opts ...grpc.CallOption) (*BcAndBankActionResp, error) {
	out := new(BcAndBankActionResp)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcAndBankFactoringOrdersAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcAndBankFactoringRepayAction(ctx context.Context, in *BcAndBankFactoringRepayActionRequest, opts ...grpc.CallOption) (*BcAndBankActionResp, error) {
	out := new(BcAndBankActionResp)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcAndBankFactoringRepayAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcForceUploadShipments(ctx context.Context, in *UploadChainShipments, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcForceUploadShipments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcForceUploadShipmentOrders(ctx context.Context, in *UploadChainShipmentOrders, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcForceUploadShipmentOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcForceUploadPayOrders(ctx context.Context, in *UploadChainPayOrders, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcForceUploadPayOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcForceUploadFactoring(ctx context.Context, in *UploadChainFactoringOrders, opts ...grpc.CallOption) (*BcResponses, error) {
	out := new(BcResponses)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcForceUploadFactoring", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bcAndBankServiceClient) BcQueryByTxID(ctx context.Context, in *QueryTxIDReq, opts ...grpc.CallOption) (*QueryTxIDRsp, error) {
	out := new(QueryTxIDRsp)
	err := c.cc.Invoke(ctx, "/BcAndBankService.BcAndBankService/BcQueryByTxID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BcAndBankServiceServer is the server API for BcAndBankService service.
// All implementations must embed UnimplementedBcAndBankServiceServer
// for forward compatibility
type BcAndBankServiceServer interface {
	// --------单纯上链接口 start---------
	// 上链 configs 信息
	BcUploadConfigs(context.Context, *UploadChainConfigs) (*BcResponses, error)
	// 上链 企业信息s
	BcUploadCompanies(context.Context, *UploadChainCompanies) (*BcResponses, error)
	// 医疗机构上链
	BcUploadMis(context.Context, *UploadChainMis) (*BcResponses, error)
	// 合同信息上链
	BcUploadContracts(context.Context, *UploadChainContracts) (*BcResponses, error)
	// 子合同信息上链
	BcUploadContractMis(context.Context, *UploadChainContractMis) (*BcResponses, error)
	// 配送计划上链
	BcUploadShipments(context.Context, *UploadChainShipments) (*BcResponses, error)
	// 配送单上链
	BcUploadShipmentOrders(context.Context, *UploadChainShipmentOrders) (*BcResponses, error)
	// 支付单上链, 确权
	BcUploadPayOrders(context.Context, *UploadChainPayOrders) (*BcResponses, error)
	// 保理单上链（更新）
	BcUploadFactoringOrders(context.Context, *UploadChainFactoringOrders) (*BcResponses, error)
	// --------上链 + 银行 组合接口 start---------
	// 发起付款 一笔付款单内可以包含多笔 支付单(接收参数: 多笔 payOrder + 一个 银行请求)
	// 查所有的支付单, 看是否符合要求
	// 返回:: 链上数据hash集合 加 银行返回
	BcAndBankPayOrderAction(context.Context, *BcAndBankPayOrderActionRequest) (*BcAndBankActionResp, error)
	// 发起保理(多笔支付单)
	// 保理单上链, 支付单上链后才上链, 查所有的支付单: 确认 支付单 处于 待支付/失败/(私有状态 未保理) 中; 调银行
	BcAndBankFactoringOrdersAction(context.Context, *BcAndBankFactoringActionRequest) (*BcAndBankActionResp, error)
	// 发起保理还款(以保理单为准)
	// 查保理单, 处于 已保理 状态 ->
	BcAndBankFactoringRepayAction(context.Context, *BcAndBankFactoringRepayActionRequest) (*BcAndBankActionResp, error)
	// --------危险动作: 强制上链接口 start---------
	// 强制配送计划上链
	BcForceUploadShipments(context.Context, *UploadChainShipments) (*BcResponses, error)
	// 强制配送单上链
	BcForceUploadShipmentOrders(context.Context, *UploadChainShipmentOrders) (*BcResponses, error)
	// 强制支付单上链
	BcForceUploadPayOrders(context.Context, *UploadChainPayOrders) (*BcResponses, error)
	// 强制Factoring上链
	BcForceUploadFactoring(context.Context, *UploadChainFactoringOrders) (*BcResponses, error)
	// 根据 txID 查询 链上交易, 返回 解密后的 json.Marshal 字符串
	BcQueryByTxID(context.Context, *QueryTxIDReq) (*QueryTxIDRsp, error)
	mustEmbedUnimplementedBcAndBankServiceServer()
}

// UnimplementedBcAndBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBcAndBankServiceServer struct {
}

func (UnimplementedBcAndBankServiceServer) BcUploadConfigs(context.Context, *UploadChainConfigs) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadConfigs not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadCompanies(context.Context, *UploadChainCompanies) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadCompanies not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadMis(context.Context, *UploadChainMis) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadMis not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadContracts(context.Context, *UploadChainContracts) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadContracts not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadContractMis(context.Context, *UploadChainContractMis) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadContractMis not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadShipments(context.Context, *UploadChainShipments) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadShipments not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadShipmentOrders(context.Context, *UploadChainShipmentOrders) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadShipmentOrders not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadPayOrders(context.Context, *UploadChainPayOrders) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadPayOrders not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcUploadFactoringOrders(context.Context, *UploadChainFactoringOrders) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcUploadFactoringOrders not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcAndBankPayOrderAction(context.Context, *BcAndBankPayOrderActionRequest) (*BcAndBankActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcAndBankPayOrderAction not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcAndBankFactoringOrdersAction(context.Context, *BcAndBankFactoringActionRequest) (*BcAndBankActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcAndBankFactoringOrdersAction not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcAndBankFactoringRepayAction(context.Context, *BcAndBankFactoringRepayActionRequest) (*BcAndBankActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcAndBankFactoringRepayAction not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcForceUploadShipments(context.Context, *UploadChainShipments) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcForceUploadShipments not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcForceUploadShipmentOrders(context.Context, *UploadChainShipmentOrders) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcForceUploadShipmentOrders not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcForceUploadPayOrders(context.Context, *UploadChainPayOrders) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcForceUploadPayOrders not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcForceUploadFactoring(context.Context, *UploadChainFactoringOrders) (*BcResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcForceUploadFactoring not implemented")
}
func (UnimplementedBcAndBankServiceServer) BcQueryByTxID(context.Context, *QueryTxIDReq) (*QueryTxIDRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BcQueryByTxID not implemented")
}
func (UnimplementedBcAndBankServiceServer) mustEmbedUnimplementedBcAndBankServiceServer() {}

// UnsafeBcAndBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BcAndBankServiceServer will
// result in compilation errors.
type UnsafeBcAndBankServiceServer interface {
	mustEmbedUnimplementedBcAndBankServiceServer()
}

func RegisterBcAndBankServiceServer(s grpc.ServiceRegistrar, srv BcAndBankServiceServer) {
	s.RegisterService(&BcAndBankService_ServiceDesc, srv)
}

func _BcAndBankService_BcUploadConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainConfigs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadConfigs(ctx, req.(*UploadChainConfigs))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainCompanies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadCompanies(ctx, req.(*UploadChainCompanies))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadMis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainMis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadMis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadMis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadMis(ctx, req.(*UploadChainMis))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainContracts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadContracts(ctx, req.(*UploadChainContracts))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadContractMis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainContractMis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadContractMis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadContractMis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadContractMis(ctx, req.(*UploadChainContractMis))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainShipments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadShipments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadShipments(ctx, req.(*UploadChainShipments))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadShipmentOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainShipmentOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadShipmentOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadShipmentOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadShipmentOrders(ctx, req.(*UploadChainShipmentOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadPayOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainPayOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadPayOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadPayOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadPayOrders(ctx, req.(*UploadChainPayOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcUploadFactoringOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainFactoringOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcUploadFactoringOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcUploadFactoringOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcUploadFactoringOrders(ctx, req.(*UploadChainFactoringOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcAndBankPayOrderAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BcAndBankPayOrderActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcAndBankPayOrderAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcAndBankPayOrderAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcAndBankPayOrderAction(ctx, req.(*BcAndBankPayOrderActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcAndBankFactoringOrdersAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BcAndBankFactoringActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcAndBankFactoringOrdersAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcAndBankFactoringOrdersAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcAndBankFactoringOrdersAction(ctx, req.(*BcAndBankFactoringActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcAndBankFactoringRepayAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BcAndBankFactoringRepayActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcAndBankFactoringRepayAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcAndBankFactoringRepayAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcAndBankFactoringRepayAction(ctx, req.(*BcAndBankFactoringRepayActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcForceUploadShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainShipments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcForceUploadShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcForceUploadShipments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcForceUploadShipments(ctx, req.(*UploadChainShipments))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcForceUploadShipmentOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainShipmentOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcForceUploadShipmentOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcForceUploadShipmentOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcForceUploadShipmentOrders(ctx, req.(*UploadChainShipmentOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcForceUploadPayOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainPayOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcForceUploadPayOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcForceUploadPayOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcForceUploadPayOrders(ctx, req.(*UploadChainPayOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcForceUploadFactoring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadChainFactoringOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcForceUploadFactoring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcForceUploadFactoring",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcForceUploadFactoring(ctx, req.(*UploadChainFactoringOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _BcAndBankService_BcQueryByTxID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTxIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BcAndBankServiceServer).BcQueryByTxID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BcAndBankService.BcAndBankService/BcQueryByTxID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BcAndBankServiceServer).BcQueryByTxID(ctx, req.(*QueryTxIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BcAndBankService_ServiceDesc is the grpc.ServiceDesc for BcAndBankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BcAndBankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BcAndBankService.BcAndBankService",
	HandlerType: (*BcAndBankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BcUploadConfigs",
			Handler:    _BcAndBankService_BcUploadConfigs_Handler,
		},
		{
			MethodName: "BcUploadCompanies",
			Handler:    _BcAndBankService_BcUploadCompanies_Handler,
		},
		{
			MethodName: "BcUploadMis",
			Handler:    _BcAndBankService_BcUploadMis_Handler,
		},
		{
			MethodName: "BcUploadContracts",
			Handler:    _BcAndBankService_BcUploadContracts_Handler,
		},
		{
			MethodName: "BcUploadContractMis",
			Handler:    _BcAndBankService_BcUploadContractMis_Handler,
		},
		{
			MethodName: "BcUploadShipments",
			Handler:    _BcAndBankService_BcUploadShipments_Handler,
		},
		{
			MethodName: "BcUploadShipmentOrders",
			Handler:    _BcAndBankService_BcUploadShipmentOrders_Handler,
		},
		{
			MethodName: "BcUploadPayOrders",
			Handler:    _BcAndBankService_BcUploadPayOrders_Handler,
		},
		{
			MethodName: "BcUploadFactoringOrders",
			Handler:    _BcAndBankService_BcUploadFactoringOrders_Handler,
		},
		{
			MethodName: "BcAndBankPayOrderAction",
			Handler:    _BcAndBankService_BcAndBankPayOrderAction_Handler,
		},
		{
			MethodName: "BcAndBankFactoringOrdersAction",
			Handler:    _BcAndBankService_BcAndBankFactoringOrdersAction_Handler,
		},
		{
			MethodName: "BcAndBankFactoringRepayAction",
			Handler:    _BcAndBankService_BcAndBankFactoringRepayAction_Handler,
		},
		{
			MethodName: "BcForceUploadShipments",
			Handler:    _BcAndBankService_BcForceUploadShipments_Handler,
		},
		{
			MethodName: "BcForceUploadShipmentOrders",
			Handler:    _BcAndBankService_BcForceUploadShipmentOrders_Handler,
		},
		{
			MethodName: "BcForceUploadPayOrders",
			Handler:    _BcAndBankService_BcForceUploadPayOrders_Handler,
		},
		{
			MethodName: "BcForceUploadFactoring",
			Handler:    _BcAndBankService_BcForceUploadFactoring_Handler,
		},
		{
			MethodName: "BcQueryByTxID",
			Handler:    _BcAndBankService_BcQueryByTxID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BcAndBankService.proto",
}
