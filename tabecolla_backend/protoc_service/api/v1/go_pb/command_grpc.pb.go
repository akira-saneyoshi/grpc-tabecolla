//
// CommandService用のParam型とResult型を定義したprotoファイル
//

// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: go_proto/command.proto

// パッケージの宣言

package go_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserCommand_Create_FullMethodName = "/go_protoc.UserCommand/Create"
	UserCommand_Update_FullMethodName = "/go_protoc.UserCommand/Update"
	UserCommand_Delete_FullMethodName = "/go_protoc.UserCommand/Delete"
)

// UserCommandClient is the client API for UserCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *****************************************
// ユーザの更新サービス型の定義
// *****************************************
// ユーザ更新サービス型
type UserCommandClient interface {
	// ユーザを追加した結果を返す
	Create(ctx context.Context, in *UserUpParam, opts ...grpc.CallOption) (*UserUpResult, error)
	// ユーザを更新した結果を返す
	Update(ctx context.Context, in *UserUpParam, opts ...grpc.CallOption) (*UserUpResult, error)
	// ユーザを削除した結果を返す
	Delete(ctx context.Context, in *UserUpParam, opts ...grpc.CallOption) (*UserUpResult, error)
}

type userCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCommandClient(cc grpc.ClientConnInterface) UserCommandClient {
	return &userCommandClient{cc}
}

func (c *userCommandClient) Create(ctx context.Context, in *UserUpParam, opts ...grpc.CallOption) (*UserUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserUpResult)
	err := c.cc.Invoke(ctx, UserCommand_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCommandClient) Update(ctx context.Context, in *UserUpParam, opts ...grpc.CallOption) (*UserUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserUpResult)
	err := c.cc.Invoke(ctx, UserCommand_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCommandClient) Delete(ctx context.Context, in *UserUpParam, opts ...grpc.CallOption) (*UserUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserUpResult)
	err := c.cc.Invoke(ctx, UserCommand_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCommandServer is the server API for UserCommand service.
// All implementations must embed UnimplementedUserCommandServer
// for forward compatibility
//
// *****************************************
// ユーザの更新サービス型の定義
// *****************************************
// ユーザ更新サービス型
type UserCommandServer interface {
	// ユーザを追加した結果を返す
	Create(context.Context, *UserUpParam) (*UserUpResult, error)
	// ユーザを更新した結果を返す
	Update(context.Context, *UserUpParam) (*UserUpResult, error)
	// ユーザを削除した結果を返す
	Delete(context.Context, *UserUpParam) (*UserUpResult, error)
	mustEmbedUnimplementedUserCommandServer()
}

// UnimplementedUserCommandServer must be embedded to have forward compatible implementations.
type UnimplementedUserCommandServer struct {
}

func (UnimplementedUserCommandServer) Create(context.Context, *UserUpParam) (*UserUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserCommandServer) Update(context.Context, *UserUpParam) (*UserUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserCommandServer) Delete(context.Context, *UserUpParam) (*UserUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserCommandServer) mustEmbedUnimplementedUserCommandServer() {}

// UnsafeUserCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCommandServer will
// result in compilation errors.
type UnsafeUserCommandServer interface {
	mustEmbedUnimplementedUserCommandServer()
}

func RegisterUserCommandServer(s grpc.ServiceRegistrar, srv UserCommandServer) {
	s.RegisterService(&UserCommand_ServiceDesc, srv)
}

func _UserCommand_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommandServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCommand_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommandServer).Create(ctx, req.(*UserUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCommand_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommandServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCommand_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommandServer).Update(ctx, req.(*UserUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCommand_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCommandServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCommand_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCommandServer).Delete(ctx, req.(*UserUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCommand_ServiceDesc is the grpc.ServiceDesc for UserCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_protoc.UserCommand",
	HandlerType: (*UserCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserCommand_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserCommand_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserCommand_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_proto/command.proto",
}

const (
	StoreCategoryCommand_Create_FullMethodName = "/go_protoc.StoreCategoryCommand/Create"
	StoreCategoryCommand_Update_FullMethodName = "/go_protoc.StoreCategoryCommand/Update"
	StoreCategoryCommand_Delete_FullMethodName = "/go_protoc.StoreCategoryCommand/Delete"
)

// StoreCategoryCommandClient is the client API for StoreCategoryCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *****************************************
// 飲食店カテゴリと飲食店の更新サービス型の定義
// *****************************************
// 飲食店カテゴリ更新サービス型
type StoreCategoryCommandClient interface {
	// 飲食店カテゴリを追加した結果を返す
	Create(ctx context.Context, in *StoreCategoryUpParam, opts ...grpc.CallOption) (*StoreCategoryUpResult, error)
	// 飲食店カテゴリを更新した結果を返す
	Update(ctx context.Context, in *StoreCategoryUpParam, opts ...grpc.CallOption) (*StoreCategoryUpResult, error)
	// 飲食店カテゴリを削除した結果を返す
	Delete(ctx context.Context, in *StoreCategoryUpParam, opts ...grpc.CallOption) (*StoreCategoryUpResult, error)
}

type storeCategoryCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreCategoryCommandClient(cc grpc.ClientConnInterface) StoreCategoryCommandClient {
	return &storeCategoryCommandClient{cc}
}

func (c *storeCategoryCommandClient) Create(ctx context.Context, in *StoreCategoryUpParam, opts ...grpc.CallOption) (*StoreCategoryUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreCategoryUpResult)
	err := c.cc.Invoke(ctx, StoreCategoryCommand_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeCategoryCommandClient) Update(ctx context.Context, in *StoreCategoryUpParam, opts ...grpc.CallOption) (*StoreCategoryUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreCategoryUpResult)
	err := c.cc.Invoke(ctx, StoreCategoryCommand_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeCategoryCommandClient) Delete(ctx context.Context, in *StoreCategoryUpParam, opts ...grpc.CallOption) (*StoreCategoryUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreCategoryUpResult)
	err := c.cc.Invoke(ctx, StoreCategoryCommand_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreCategoryCommandServer is the server API for StoreCategoryCommand service.
// All implementations must embed UnimplementedStoreCategoryCommandServer
// for forward compatibility
//
// *****************************************
// 飲食店カテゴリと飲食店の更新サービス型の定義
// *****************************************
// 飲食店カテゴリ更新サービス型
type StoreCategoryCommandServer interface {
	// 飲食店カテゴリを追加した結果を返す
	Create(context.Context, *StoreCategoryUpParam) (*StoreCategoryUpResult, error)
	// 飲食店カテゴリを更新した結果を返す
	Update(context.Context, *StoreCategoryUpParam) (*StoreCategoryUpResult, error)
	// 飲食店カテゴリを削除した結果を返す
	Delete(context.Context, *StoreCategoryUpParam) (*StoreCategoryUpResult, error)
	mustEmbedUnimplementedStoreCategoryCommandServer()
}

// UnimplementedStoreCategoryCommandServer must be embedded to have forward compatible implementations.
type UnimplementedStoreCategoryCommandServer struct {
}

func (UnimplementedStoreCategoryCommandServer) Create(context.Context, *StoreCategoryUpParam) (*StoreCategoryUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStoreCategoryCommandServer) Update(context.Context, *StoreCategoryUpParam) (*StoreCategoryUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStoreCategoryCommandServer) Delete(context.Context, *StoreCategoryUpParam) (*StoreCategoryUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStoreCategoryCommandServer) mustEmbedUnimplementedStoreCategoryCommandServer() {}

// UnsafeStoreCategoryCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreCategoryCommandServer will
// result in compilation errors.
type UnsafeStoreCategoryCommandServer interface {
	mustEmbedUnimplementedStoreCategoryCommandServer()
}

func RegisterStoreCategoryCommandServer(s grpc.ServiceRegistrar, srv StoreCategoryCommandServer) {
	s.RegisterService(&StoreCategoryCommand_ServiceDesc, srv)
}

func _StoreCategoryCommand_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCategoryUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreCategoryCommandServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreCategoryCommand_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreCategoryCommandServer).Create(ctx, req.(*StoreCategoryUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreCategoryCommand_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCategoryUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreCategoryCommandServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreCategoryCommand_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreCategoryCommandServer).Update(ctx, req.(*StoreCategoryUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreCategoryCommand_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCategoryUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreCategoryCommandServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreCategoryCommand_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreCategoryCommandServer).Delete(ctx, req.(*StoreCategoryUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreCategoryCommand_ServiceDesc is the grpc.ServiceDesc for StoreCategoryCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreCategoryCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_protoc.StoreCategoryCommand",
	HandlerType: (*StoreCategoryCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StoreCategoryCommand_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StoreCategoryCommand_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StoreCategoryCommand_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_proto/command.proto",
}

const (
	StoreCommand_Create_FullMethodName = "/go_protoc.StoreCommand/Create"
	StoreCommand_Update_FullMethodName = "/go_protoc.StoreCommand/Update"
	StoreCommand_Delete_FullMethodName = "/go_protoc.StoreCommand/Delete"
)

// StoreCommandClient is the client API for StoreCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 飲食店更新サービス型
type StoreCommandClient interface {
	// 飲食店を追加した結果を返す
	Create(ctx context.Context, in *StoreUpParam, opts ...grpc.CallOption) (*StoreUpResult, error)
	// 飲食店を更新した結果を返す
	Update(ctx context.Context, in *StoreUpParam, opts ...grpc.CallOption) (*StoreUpResult, error)
	// 飲食店を削除した結果を返す
	Delete(ctx context.Context, in *StoreUpParam, opts ...grpc.CallOption) (*StoreUpResult, error)
}

type storeCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreCommandClient(cc grpc.ClientConnInterface) StoreCommandClient {
	return &storeCommandClient{cc}
}

func (c *storeCommandClient) Create(ctx context.Context, in *StoreUpParam, opts ...grpc.CallOption) (*StoreUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreUpResult)
	err := c.cc.Invoke(ctx, StoreCommand_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeCommandClient) Update(ctx context.Context, in *StoreUpParam, opts ...grpc.CallOption) (*StoreUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreUpResult)
	err := c.cc.Invoke(ctx, StoreCommand_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeCommandClient) Delete(ctx context.Context, in *StoreUpParam, opts ...grpc.CallOption) (*StoreUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreUpResult)
	err := c.cc.Invoke(ctx, StoreCommand_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreCommandServer is the server API for StoreCommand service.
// All implementations must embed UnimplementedStoreCommandServer
// for forward compatibility
//
// 飲食店更新サービス型
type StoreCommandServer interface {
	// 飲食店を追加した結果を返す
	Create(context.Context, *StoreUpParam) (*StoreUpResult, error)
	// 飲食店を更新した結果を返す
	Update(context.Context, *StoreUpParam) (*StoreUpResult, error)
	// 飲食店を削除した結果を返す
	Delete(context.Context, *StoreUpParam) (*StoreUpResult, error)
	mustEmbedUnimplementedStoreCommandServer()
}

// UnimplementedStoreCommandServer must be embedded to have forward compatible implementations.
type UnimplementedStoreCommandServer struct {
}

func (UnimplementedStoreCommandServer) Create(context.Context, *StoreUpParam) (*StoreUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStoreCommandServer) Update(context.Context, *StoreUpParam) (*StoreUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStoreCommandServer) Delete(context.Context, *StoreUpParam) (*StoreUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStoreCommandServer) mustEmbedUnimplementedStoreCommandServer() {}

// UnsafeStoreCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreCommandServer will
// result in compilation errors.
type UnsafeStoreCommandServer interface {
	mustEmbedUnimplementedStoreCommandServer()
}

func RegisterStoreCommandServer(s grpc.ServiceRegistrar, srv StoreCommandServer) {
	s.RegisterService(&StoreCommand_ServiceDesc, srv)
}

func _StoreCommand_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreCommandServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreCommand_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreCommandServer).Create(ctx, req.(*StoreUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreCommand_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreCommandServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreCommand_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreCommandServer).Update(ctx, req.(*StoreUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreCommand_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreCommandServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreCommand_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreCommandServer).Delete(ctx, req.(*StoreUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreCommand_ServiceDesc is the grpc.ServiceDesc for StoreCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_protoc.StoreCommand",
	HandlerType: (*StoreCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StoreCommand_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StoreCommand_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StoreCommand_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_proto/command.proto",
}
