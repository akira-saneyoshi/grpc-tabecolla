//
// CommandService用のParam型とResult型を定義したprotoファイル
//

// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: go_proto/command.proto

// パッケージの宣言

package go_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// 更新の種類
type CRUD int32

const (
	CRUD_UNKNOWN CRUD = 0 // 不明
	CRUD_INSERT  CRUD = 1 // 追加
	CRUD_UPDATE  CRUD = 2 // 変更
	CRUD_DELETE  CRUD = 3 // 削除
)

// Enum value maps for CRUD.
var (
	CRUD_name = map[int32]string{
		0: "UNKNOWN",
		1: "INSERT",
		2: "UPDATE",
		3: "DELETE",
	}
	CRUD_value = map[string]int32{
		"UNKNOWN": 0,
		"INSERT":  1,
		"UPDATE":  2,
		"DELETE":  3,
	}
)

func (x CRUD) Enum() *CRUD {
	p := new(CRUD)
	*p = x
	return p
}

func (x CRUD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CRUD) Descriptor() protoreflect.EnumDescriptor {
	return file_go_proto_command_proto_enumTypes[0].Descriptor()
}

func (CRUD) Type() protoreflect.EnumType {
	return &file_go_proto_command_proto_enumTypes[0]
}

func (x CRUD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CRUD.Descriptor instead.
func (CRUD) EnumDescriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{0}
}

// *****************************************
// ユーザ更新用param型とResult型の定義
// *****************************************
// ユーザ更新Param型
type UserUpParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crud         CRUD   `protobuf:"varint,1,opt,name=crud,proto3,enum=go_protoc.CRUD" json:"crud,omitempty"`                // 更新の種類
	UserId       string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                   // ユーザId
	UserName     string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`             // ユーザ名
	UserMail     string `protobuf:"bytes,4,opt,name=user_mail,json=userMail,proto3" json:"user_mail,omitempty"`             // メールアドレス
	UserPassword string `protobuf:"bytes,5,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"` // ユーザパスワード
	UserEnable   bool   `protobuf:"varint,6,opt,name=user_enable,json=userEnable,proto3" json:"user_enable,omitempty"`      // ユーザの有効/無効
}

func (x *UserUpParam) Reset() {
	*x = UserUpParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpParam) ProtoMessage() {}

func (x *UserUpParam) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpParam.ProtoReflect.Descriptor instead.
func (*UserUpParam) Descriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{0}
}

func (x *UserUpParam) GetCrud() CRUD {
	if x != nil {
		return x.Crud
	}
	return CRUD_UNKNOWN
}

func (x *UserUpParam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserUpParam) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserUpParam) GetUserMail() string {
	if x != nil {
		return x.UserMail
	}
	return ""
}

func (x *UserUpParam) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

func (x *UserUpParam) GetUserEnable() bool {
	if x != nil {
		return x.UserEnable
	}
	return false
}

// ユーザ更新Result型
type UserUpResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`           // 更新結果
	Error     *ApiError              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`         // 更新エラー
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *UserUpResult) Reset() {
	*x = UserUpResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpResult) ProtoMessage() {}

func (x *UserUpResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpResult.ProtoReflect.Descriptor instead.
func (*UserUpResult) Descriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{1}
}

func (x *UserUpResult) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserUpResult) GetError() *ApiError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UserUpResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// *****************************************
// 飲食店カテゴリ更新用param型とResult型の定義
// *****************************************
// 飲食店カテゴリ更新Param型
type StoreCategoryUpParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crud              CRUD   `protobuf:"varint,1,opt,name=crud,proto3,enum=go_protoc.CRUD" json:"crud,omitempty"`                                 // 更新の種類
	StoreCategoryUid  string `protobuf:"bytes,2,opt,name=store_category_uid,json=storeCategoryUid,proto3" json:"store_category_uid,omitempty"`    // 飲食店カテゴリ番号
	StoreCategoryName string `protobuf:"bytes,3,opt,name=store_category_name,json=storeCategoryName,proto3" json:"store_category_name,omitempty"` // 飲食店カテゴリ名
}

func (x *StoreCategoryUpParam) Reset() {
	*x = StoreCategoryUpParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCategoryUpParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCategoryUpParam) ProtoMessage() {}

func (x *StoreCategoryUpParam) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCategoryUpParam.ProtoReflect.Descriptor instead.
func (*StoreCategoryUpParam) Descriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{2}
}

func (x *StoreCategoryUpParam) GetCrud() CRUD {
	if x != nil {
		return x.Crud
	}
	return CRUD_UNKNOWN
}

func (x *StoreCategoryUpParam) GetStoreCategoryUid() string {
	if x != nil {
		return x.StoreCategoryUid
	}
	return ""
}

func (x *StoreCategoryUpParam) GetStoreCategoryName() string {
	if x != nil {
		return x.StoreCategoryName
	}
	return ""
}

// 飲食店更新Param型
type StoreUpParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crud             CRUD   `protobuf:"varint,1,opt,name=crud,proto3,enum=go_protoc.CRUD" json:"crud,omitempty"`          // 更新の種類
	StoreId          string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`          // 飲食店番号
	StoreName        string `protobuf:"bytes,3,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`    // 飲食店名
	StorePlace       string `protobuf:"bytes,4,opt,name=store_place,json=storePlace,proto3" json:"store_place,omitempty"` // 飲食店の場所
	StoreCategoryUid string `protobuf:"bytes,5,opt,name=storeCategoryUid,proto3" json:"storeCategoryUid,omitempty"`       // 飲食店カテゴリ番号
}

func (x *StoreUpParam) Reset() {
	*x = StoreUpParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreUpParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreUpParam) ProtoMessage() {}

func (x *StoreUpParam) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreUpParam.ProtoReflect.Descriptor instead.
func (*StoreUpParam) Descriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{3}
}

func (x *StoreUpParam) GetCrud() CRUD {
	if x != nil {
		return x.Crud
	}
	return CRUD_UNKNOWN
}

func (x *StoreUpParam) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *StoreUpParam) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *StoreUpParam) GetStorePlace() string {
	if x != nil {
		return x.StorePlace
	}
	return ""
}

func (x *StoreUpParam) GetStoreCategoryUid() string {
	if x != nil {
		return x.StoreCategoryUid
	}
	return ""
}

// 飲食店カテゴリ更新Result型
type StoreCategoryUpResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreCategory *StoreCategory         `protobuf:"bytes,1,opt,name=storeCategory,proto3" json:"storeCategory,omitempty"` // 更新結果
	Error         *ApiError              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                 // 更新エラー
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`         // タイムスタンプ
}

func (x *StoreCategoryUpResult) Reset() {
	*x = StoreCategoryUpResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCategoryUpResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCategoryUpResult) ProtoMessage() {}

func (x *StoreCategoryUpResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCategoryUpResult.ProtoReflect.Descriptor instead.
func (*StoreCategoryUpResult) Descriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{4}
}

func (x *StoreCategoryUpResult) GetStoreCategory() *StoreCategory {
	if x != nil {
		return x.StoreCategory
	}
	return nil
}

func (x *StoreCategoryUpResult) GetError() *ApiError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StoreCategoryUpResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// 飲食店更新Result型
type StoreUpResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Store     *Store                 `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`         // 更新結果
	Error     *ApiError              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`         // 更新エラー
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // タイムスタンプ
}

func (x *StoreUpResult) Reset() {
	*x = StoreUpResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreUpResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreUpResult) ProtoMessage() {}

func (x *StoreUpResult) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreUpResult.ProtoReflect.Descriptor instead.
func (*StoreUpResult) Descriptor() ([]byte, []int) {
	return file_go_proto_command_proto_rawDescGZIP(), []int{5}
}

func (x *StoreUpResult) GetStore() *Store {
	if x != nil {
		return x.Store
	}
	return nil
}

func (x *StoreUpResult) GetError() *ApiError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StoreUpResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_go_proto_command_proto protoreflect.FileDescriptor

var file_go_proto_command_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x1a, 0x14, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x72, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x43, 0x52, 0x55, 0x44,
	0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x72, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x43, 0x52,
	0x55, 0x44, 0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x72, 0x75, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x55, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3e, 0x0a,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0d,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x41, 0x70, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2a, 0x37, 0x0a, 0x04, 0x43, 0x52, 0x55, 0x44, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x32, 0xbe, 0x01, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xfd, 0x01, 0x0a, 0x14,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x20, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x20, 0x2e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4b,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xc5, 0x01, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x18, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6f, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_proto_command_proto_rawDescOnce sync.Once
	file_go_proto_command_proto_rawDescData = file_go_proto_command_proto_rawDesc
)

func file_go_proto_command_proto_rawDescGZIP() []byte {
	file_go_proto_command_proto_rawDescOnce.Do(func() {
		file_go_proto_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_proto_command_proto_rawDescData)
	})
	return file_go_proto_command_proto_rawDescData
}

var file_go_proto_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_proto_command_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_proto_command_proto_goTypes = []interface{}{
	(CRUD)(0),                     // 0: go_protoc.CRUD
	(*UserUpParam)(nil),           // 1: go_protoc.UserUpParam
	(*UserUpResult)(nil),          // 2: go_protoc.UserUpResult
	(*StoreCategoryUpParam)(nil),  // 3: go_protoc.StoreCategoryUpParam
	(*StoreUpParam)(nil),          // 4: go_protoc.StoreUpParam
	(*StoreCategoryUpResult)(nil), // 5: go_protoc.StoreCategoryUpResult
	(*StoreUpResult)(nil),         // 6: go_protoc.StoreUpResult
	(*User)(nil),                  // 7: go_protoc.User
	(*ApiError)(nil),              // 8: go_protoc.ApiError
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*StoreCategory)(nil),         // 10: go_protoc.StoreCategory
	(*Store)(nil),                 // 11: go_protoc.Store
}
var file_go_proto_command_proto_depIdxs = []int32{
	0,  // 0: go_protoc.UserUpParam.crud:type_name -> go_protoc.CRUD
	7,  // 1: go_protoc.UserUpResult.user:type_name -> go_protoc.User
	8,  // 2: go_protoc.UserUpResult.error:type_name -> go_protoc.ApiError
	9,  // 3: go_protoc.UserUpResult.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 4: go_protoc.StoreCategoryUpParam.crud:type_name -> go_protoc.CRUD
	0,  // 5: go_protoc.StoreUpParam.crud:type_name -> go_protoc.CRUD
	10, // 6: go_protoc.StoreCategoryUpResult.storeCategory:type_name -> go_protoc.StoreCategory
	8,  // 7: go_protoc.StoreCategoryUpResult.error:type_name -> go_protoc.ApiError
	9,  // 8: go_protoc.StoreCategoryUpResult.timestamp:type_name -> google.protobuf.Timestamp
	11, // 9: go_protoc.StoreUpResult.store:type_name -> go_protoc.Store
	8,  // 10: go_protoc.StoreUpResult.error:type_name -> go_protoc.ApiError
	9,  // 11: go_protoc.StoreUpResult.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 12: go_protoc.UserCommand.Create:input_type -> go_protoc.UserUpParam
	1,  // 13: go_protoc.UserCommand.Update:input_type -> go_protoc.UserUpParam
	1,  // 14: go_protoc.UserCommand.Delete:input_type -> go_protoc.UserUpParam
	3,  // 15: go_protoc.StoreCategoryCommand.Create:input_type -> go_protoc.StoreCategoryUpParam
	3,  // 16: go_protoc.StoreCategoryCommand.Update:input_type -> go_protoc.StoreCategoryUpParam
	3,  // 17: go_protoc.StoreCategoryCommand.Delete:input_type -> go_protoc.StoreCategoryUpParam
	4,  // 18: go_protoc.StoreCommand.Create:input_type -> go_protoc.StoreUpParam
	4,  // 19: go_protoc.StoreCommand.Update:input_type -> go_protoc.StoreUpParam
	4,  // 20: go_protoc.StoreCommand.Delete:input_type -> go_protoc.StoreUpParam
	2,  // 21: go_protoc.UserCommand.Create:output_type -> go_protoc.UserUpResult
	2,  // 22: go_protoc.UserCommand.Update:output_type -> go_protoc.UserUpResult
	2,  // 23: go_protoc.UserCommand.Delete:output_type -> go_protoc.UserUpResult
	5,  // 24: go_protoc.StoreCategoryCommand.Create:output_type -> go_protoc.StoreCategoryUpResult
	5,  // 25: go_protoc.StoreCategoryCommand.Update:output_type -> go_protoc.StoreCategoryUpResult
	5,  // 26: go_protoc.StoreCategoryCommand.Delete:output_type -> go_protoc.StoreCategoryUpResult
	6,  // 27: go_protoc.StoreCommand.Create:output_type -> go_protoc.StoreUpResult
	6,  // 28: go_protoc.StoreCommand.Update:output_type -> go_protoc.StoreUpResult
	6,  // 29: go_protoc.StoreCommand.Delete:output_type -> go_protoc.StoreUpResult
	21, // [21:30] is the sub-list for method output_type
	12, // [12:21] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_go_proto_command_proto_init() }
func file_go_proto_command_proto_init() {
	if File_go_proto_command_proto != nil {
		return
	}
	file_go_proto_error_proto_init()
	file_go_proto_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_proto_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpParam); i {
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
		file_go_proto_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpResult); i {
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
		file_go_proto_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCategoryUpParam); i {
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
		file_go_proto_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreUpParam); i {
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
		file_go_proto_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCategoryUpResult); i {
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
		file_go_proto_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreUpResult); i {
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
			RawDescriptor: file_go_proto_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_go_proto_command_proto_goTypes,
		DependencyIndexes: file_go_proto_command_proto_depIdxs,
		EnumInfos:         file_go_proto_command_proto_enumTypes,
		MessageInfos:      file_go_proto_command_proto_msgTypes,
	}.Build()
	File_go_proto_command_proto = out.File
	file_go_proto_command_proto_rawDesc = nil
	file_go_proto_command_proto_goTypes = nil
	file_go_proto_command_proto_depIdxs = nil
}
