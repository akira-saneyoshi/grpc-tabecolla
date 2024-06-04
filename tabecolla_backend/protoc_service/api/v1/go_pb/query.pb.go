//
//  QueryService用のParam型とResult型を定義したprotoファイル
//

// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: go_proto/query.proto

// パッケージの宣言

package go_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *****************************************
// ユーザ検索用param型とResult型の定義
// *****************************************
//
//	ユーザ検索用Param型
type UserParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ユーザId
}

func (x *UserParam) Reset() {
	*x = UserParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserParam) ProtoMessage() {}

func (x *UserParam) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserParam.ProtoReflect.Descriptor instead.
func (*UserParam) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{0}
}

func (x *UserParam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// ユーザ複数件を返すResult型
type UsersResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users     []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`         // ユーザ複数
	Error     *ApiError              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`         // エラー
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *UsersResult) Reset() {
	*x = UsersResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersResult) ProtoMessage() {}

func (x *UsersResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersResult.ProtoReflect.Descriptor instead.
func (*UsersResult) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{1}
}

func (x *UsersResult) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UsersResult) GetError() *ApiError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UsersResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// ユーザ1件を返すResult型
type UserResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// エラーか検索結果のいずれかを返す
	//
	// Types that are assignable to Result:
	//
	//	*UserResult_User
	//	*UserResult_Error
	Result    isUserResult_Result    `protobuf_oneof:"result"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *UserResult) Reset() {
	*x = UserResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResult) ProtoMessage() {}

func (x *UserResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResult.ProtoReflect.Descriptor instead.
func (*UserResult) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{2}
}

func (m *UserResult) GetResult() isUserResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *UserResult) GetUser() *User {
	if x, ok := x.GetResult().(*UserResult_User); ok {
		return x.User
	}
	return nil
}

func (x *UserResult) GetError() *ApiError {
	if x, ok := x.GetResult().(*UserResult_Error); ok {
		return x.Error
	}
	return nil
}

func (x *UserResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isUserResult_Result interface {
	isUserResult_Result()
}

type UserResult_User struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3,oneof"` // 検索結果
}

type UserResult_Error struct {
	Error *ApiError `protobuf:"bytes,2,opt,name=error,proto3,oneof"` // 検索エラー
}

func (*UserResult_User) isUserResult_Result() {}

func (*UserResult_Error) isUserResult_Result() {}

// *****************************************
// 飲食店カテゴリ検索用param型とResult型の定義
// *****************************************
//
//	飲食店カテゴリ検索用Param型
type StoreCategoryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreCategoryId int32 `protobuf:"varint,1,opt,name=store_category_id,json=storeCategoryId,proto3" json:"store_category_id,omitempty"` // 飲食店カテゴリ番号
}

func (x *StoreCategoryParam) Reset() {
	*x = StoreCategoryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCategoryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCategoryParam) ProtoMessage() {}

func (x *StoreCategoryParam) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCategoryParam.ProtoReflect.Descriptor instead.
func (*StoreCategoryParam) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{3}
}

func (x *StoreCategoryParam) GetStoreCategoryId() int32 {
	if x != nil {
		return x.StoreCategoryId
	}
	return 0
}

// 飲食店カテゴリ複数件を返すResult型
type StoreCategoriesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreCategories []*StoreCategory       `protobuf:"bytes,1,rep,name=store_categories,json=storeCategories,proto3" json:"store_categories,omitempty"` // 飲食店カテゴリ複数
	Error           *ApiError              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                                            // エラー
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                    // タイムスタンプ
}

func (x *StoreCategoriesResult) Reset() {
	*x = StoreCategoriesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCategoriesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCategoriesResult) ProtoMessage() {}

func (x *StoreCategoriesResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCategoriesResult.ProtoReflect.Descriptor instead.
func (*StoreCategoriesResult) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{4}
}

func (x *StoreCategoriesResult) GetStoreCategories() []*StoreCategory {
	if x != nil {
		return x.StoreCategories
	}
	return nil
}

func (x *StoreCategoriesResult) GetError() *ApiError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StoreCategoriesResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// 飲食店カテゴリ1件を返すResult型
type StoreCategoryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// エラーか検索結果のいずれかを返す
	//
	// Types that are assignable to Result:
	//
	//	*StoreCategoryResult_StoreCategory
	//	*StoreCategoryResult_Error
	Result    isStoreCategoryResult_Result `protobuf_oneof:"result"`
	Timestamp *timestamppb.Timestamp       `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *StoreCategoryResult) Reset() {
	*x = StoreCategoryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCategoryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCategoryResult) ProtoMessage() {}

func (x *StoreCategoryResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCategoryResult.ProtoReflect.Descriptor instead.
func (*StoreCategoryResult) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{5}
}

func (m *StoreCategoryResult) GetResult() isStoreCategoryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *StoreCategoryResult) GetStoreCategory() *StoreCategory {
	if x, ok := x.GetResult().(*StoreCategoryResult_StoreCategory); ok {
		return x.StoreCategory
	}
	return nil
}

func (x *StoreCategoryResult) GetError() *ApiError {
	if x, ok := x.GetResult().(*StoreCategoryResult_Error); ok {
		return x.Error
	}
	return nil
}

func (x *StoreCategoryResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isStoreCategoryResult_Result interface {
	isStoreCategoryResult_Result()
}

type StoreCategoryResult_StoreCategory struct {
	StoreCategory *StoreCategory `protobuf:"bytes,1,opt,name=store_category,json=storeCategory,proto3,oneof"` // 飲食店カテゴリ
}

type StoreCategoryResult_Error struct {
	Error *ApiError `protobuf:"bytes,2,opt,name=error,proto3,oneof"` // エラー
}

func (*StoreCategoryResult_StoreCategory) isStoreCategoryResult_Result() {}

func (*StoreCategoryResult_Error) isStoreCategoryResult_Result() {}

// *****************************************
// 飲食店検索用param型とResult型の定義
// *****************************************
//
//	飲食店検索用Param型
type StoreParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId      string `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`                // 飲食店番号
	StoreKeyword string `protobuf:"bytes,2,opt,name=store_keyword,json=storeKeyword,proto3" json:"store_keyword,omitempty"` // キーワード
}

func (x *StoreParam) Reset() {
	*x = StoreParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreParam) ProtoMessage() {}

func (x *StoreParam) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreParam.ProtoReflect.Descriptor instead.
func (*StoreParam) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{6}
}

func (x *StoreParam) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *StoreParam) GetStoreKeyword() string {
	if x != nil {
		return x.StoreKeyword
	}
	return ""
}

// 飲食店複数件を返すResult型
type StoresResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stores    []*Store               `protobuf:"bytes,1,rep,name=stores,proto3" json:"stores,omitempty"`       // 飲食店複数
	Error     *ApiError              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`         // エラー
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *StoresResult) Reset() {
	*x = StoresResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoresResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoresResult) ProtoMessage() {}

func (x *StoresResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoresResult.ProtoReflect.Descriptor instead.
func (*StoresResult) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{7}
}

func (x *StoresResult) GetStores() []*Store {
	if x != nil {
		return x.Stores
	}
	return nil
}

func (x *StoresResult) GetError() *ApiError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StoresResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// 飲食店1件を返すResult型
type StoreResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// エラーか検索結果のいずれかを返す
	//
	// Types that are assignable to Result:
	//
	//	*StoreResult_Store
	//	*StoreResult_Error
	Result    isStoreResult_Result   `protobuf_oneof:"result"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *StoreResult) Reset() {
	*x = StoreResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_query_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResult) ProtoMessage() {}

func (x *StoreResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_query_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResult.ProtoReflect.Descriptor instead.
func (*StoreResult) Descriptor() ([]byte, []int) {
	return file_go_proto_query_proto_rawDescGZIP(), []int{8}
}

func (m *StoreResult) GetResult() isStoreResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *StoreResult) GetStore() *Store {
	if x, ok := x.GetResult().(*StoreResult_Store); ok {
		return x.Store
	}
	return nil
}

func (x *StoreResult) GetError() *ApiError {
	if x, ok := x.GetResult().(*StoreResult_Error); ok {
		return x.Error
	}
	return nil
}

func (x *StoreResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isStoreResult_Result interface {
	isStoreResult_Result()
}

type StoreResult_Store struct {
	Store *Store `protobuf:"bytes,1,opt,name=store,proto3,oneof"` // 検索結果
}

type StoreResult_Error struct {
	Error *ApiError `protobuf:"bytes,2,opt,name=error,proto3,oneof"` // 検索エラー
}

func (*StoreResult_Store) isStoreResult_Result() {}

func (*StoreResult_Error) isStoreResult_Result() {}

var File_go_proto_query_proto protoreflect.FileDescriptor

var file_go_proto_query_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x1a, 0x14, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa4,
	0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x43, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xc9, 0x01, 0x0a, 0x13,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x41, 0x70, 0x69, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x41, 0x70, 0x69, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x9d, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xf3, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x38, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x42, 0x79, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_proto_query_proto_rawDescOnce sync.Once
	file_go_proto_query_proto_rawDescData = file_go_proto_query_proto_rawDesc
)

func file_go_proto_query_proto_rawDescGZIP() []byte {
	file_go_proto_query_proto_rawDescOnce.Do(func() {
		file_go_proto_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_proto_query_proto_rawDescData)
	})
	return file_go_proto_query_proto_rawDescData
}

var file_go_proto_query_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_go_proto_query_proto_goTypes = []interface{}{
	(*UserParam)(nil),             // 0: go_protoc.UserParam
	(*UsersResult)(nil),           // 1: go_protoc.UsersResult
	(*UserResult)(nil),            // 2: go_protoc.UserResult
	(*StoreCategoryParam)(nil),    // 3: go_protoc.StoreCategoryParam
	(*StoreCategoriesResult)(nil), // 4: go_protoc.StoreCategoriesResult
	(*StoreCategoryResult)(nil),   // 5: go_protoc.StoreCategoryResult
	(*StoreParam)(nil),            // 6: go_protoc.StoreParam
	(*StoresResult)(nil),          // 7: go_protoc.StoresResult
	(*StoreResult)(nil),           // 8: go_protoc.StoreResult
	(*User)(nil),                  // 9: go_protoc.User
	(*ApiError)(nil),              // 10: go_protoc.ApiError
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*StoreCategory)(nil),         // 12: go_protoc.StoreCategory
	(*Store)(nil),                 // 13: go_protoc.Store
	(*emptypb.Empty)(nil),         // 14: google.protobuf.Empty
}
var file_go_proto_query_proto_depIdxs = []int32{
	9,  // 0: go_protoc.UsersResult.users:type_name -> go_protoc.User
	10, // 1: go_protoc.UsersResult.error:type_name -> go_protoc.ApiError
	11, // 2: go_protoc.UsersResult.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 3: go_protoc.UserResult.user:type_name -> go_protoc.User
	10, // 4: go_protoc.UserResult.error:type_name -> go_protoc.ApiError
	11, // 5: go_protoc.UserResult.timestamp:type_name -> google.protobuf.Timestamp
	12, // 6: go_protoc.StoreCategoriesResult.store_categories:type_name -> go_protoc.StoreCategory
	10, // 7: go_protoc.StoreCategoriesResult.error:type_name -> go_protoc.ApiError
	11, // 8: go_protoc.StoreCategoriesResult.timestamp:type_name -> google.protobuf.Timestamp
	12, // 9: go_protoc.StoreCategoryResult.store_category:type_name -> go_protoc.StoreCategory
	10, // 10: go_protoc.StoreCategoryResult.error:type_name -> go_protoc.ApiError
	11, // 11: go_protoc.StoreCategoryResult.timestamp:type_name -> google.protobuf.Timestamp
	13, // 12: go_protoc.StoresResult.stores:type_name -> go_protoc.Store
	10, // 13: go_protoc.StoresResult.error:type_name -> go_protoc.ApiError
	11, // 14: go_protoc.StoresResult.timestamp:type_name -> google.protobuf.Timestamp
	13, // 15: go_protoc.StoreResult.store:type_name -> go_protoc.Store
	10, // 16: go_protoc.StoreResult.error:type_name -> go_protoc.ApiError
	11, // 17: go_protoc.StoreResult.timestamp:type_name -> google.protobuf.Timestamp
	14, // 18: go_protoc.StoreCategoryQuery.List:input_type -> google.protobuf.Empty
	3,  // 19: go_protoc.StoreCategoryQuery.ById:input_type -> go_protoc.StoreCategoryParam
	14, // 20: go_protoc.StoreQuery.ListStream:input_type -> google.protobuf.Empty
	14, // 21: go_protoc.StoreQuery.List:input_type -> google.protobuf.Empty
	6,  // 22: go_protoc.StoreQuery.ById:input_type -> go_protoc.StoreParam
	6,  // 23: go_protoc.StoreQuery.ByKeyword:input_type -> go_protoc.StoreParam
	4,  // 24: go_protoc.StoreCategoryQuery.List:output_type -> go_protoc.StoreCategoriesResult
	5,  // 25: go_protoc.StoreCategoryQuery.ById:output_type -> go_protoc.StoreCategoryResult
	13, // 26: go_protoc.StoreQuery.ListStream:output_type -> go_protoc.Store
	7,  // 27: go_protoc.StoreQuery.List:output_type -> go_protoc.StoresResult
	8,  // 28: go_protoc.StoreQuery.ById:output_type -> go_protoc.StoreResult
	7,  // 29: go_protoc.StoreQuery.ByKeyword:output_type -> go_protoc.StoresResult
	24, // [24:30] is the sub-list for method output_type
	18, // [18:24] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_go_proto_query_proto_init() }
func file_go_proto_query_proto_init() {
	if File_go_proto_query_proto != nil {
		return
	}
	file_go_proto_error_proto_init()
	file_go_proto_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_proto_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserParam); i {
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
		file_go_proto_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersResult); i {
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
		file_go_proto_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResult); i {
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
		file_go_proto_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCategoryParam); i {
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
		file_go_proto_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCategoriesResult); i {
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
		file_go_proto_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCategoryResult); i {
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
		file_go_proto_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreParam); i {
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
		file_go_proto_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoresResult); i {
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
		file_go_proto_query_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResult); i {
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
	file_go_proto_query_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UserResult_User)(nil),
		(*UserResult_Error)(nil),
	}
	file_go_proto_query_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*StoreCategoryResult_StoreCategory)(nil),
		(*StoreCategoryResult_Error)(nil),
	}
	file_go_proto_query_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*StoreResult_Store)(nil),
		(*StoreResult_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_proto_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_go_proto_query_proto_goTypes,
		DependencyIndexes: file_go_proto_query_proto_depIdxs,
		MessageInfos:      file_go_proto_query_proto_msgTypes,
	}.Build()
	File_go_proto_query_proto = out.File
	file_go_proto_query_proto_rawDesc = nil
	file_go_proto_query_proto_goTypes = nil
	file_go_proto_query_proto_depIdxs = nil
}