//
// クライアントとサーバ間で利用するデータ格納型を定義したprotoファイル
//

// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: go_proto/models.proto

// パッケージの宣言

package go_pb

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

// ユーザ型の定義
type LoginUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginUserId       string `protobuf:"bytes,1,opt,name=login_user_id,json=loginUserId,proto3" json:"login_user_id,omitempty"`                   // ユーザId
	LoginUserName     string `protobuf:"bytes,2,opt,name=login_user_name,json=loginUserName,proto3" json:"login_user_name,omitempty"`             // ユーザ名
	LoginUserMail     string `protobuf:"bytes,3,opt,name=login_user_mail,json=loginUserMail,proto3" json:"login_user_mail,omitempty"`             // メールアドレス
	LoginUserPassword string `protobuf:"bytes,4,opt,name=login_user_password,json=loginUserPassword,proto3" json:"login_user_password,omitempty"` // ユーザパスワード
	LoginUserEnable   bool   `protobuf:"varint,5,opt,name=login_user_enable,json=loginUserEnable,proto3" json:"login_user_enable,omitempty"`      // ユーザの有効/無効
}

func (x *LoginUser) Reset() {
	*x = LoginUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUser) ProtoMessage() {}

func (x *LoginUser) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUser.ProtoReflect.Descriptor instead.
func (*LoginUser) Descriptor() ([]byte, []int) {
	return file_go_proto_models_proto_rawDescGZIP(), []int{0}
}

func (x *LoginUser) GetLoginUserId() string {
	if x != nil {
		return x.LoginUserId
	}
	return ""
}

func (x *LoginUser) GetLoginUserName() string {
	if x != nil {
		return x.LoginUserName
	}
	return ""
}

func (x *LoginUser) GetLoginUserMail() string {
	if x != nil {
		return x.LoginUserMail
	}
	return ""
}

func (x *LoginUser) GetLoginUserPassword() string {
	if x != nil {
		return x.LoginUserPassword
	}
	return ""
}

func (x *LoginUser) GetLoginUserEnable() bool {
	if x != nil {
		return x.LoginUserEnable
	}
	return false
}

// 飲食店カテゴリ型の定義
type StoreCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreCategoryUid  string `protobuf:"bytes,1,opt,name=store_category_uid,json=storeCategoryUid,proto3" json:"store_category_uid,omitempty"`    // 飲食店カテゴリ番号
	StoreCategoryName string `protobuf:"bytes,2,opt,name=store_category_name,json=storeCategoryName,proto3" json:"store_category_name,omitempty"` // 飲食店カテゴリ名
}

func (x *StoreCategory) Reset() {
	*x = StoreCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCategory) ProtoMessage() {}

func (x *StoreCategory) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCategory.ProtoReflect.Descriptor instead.
func (*StoreCategory) Descriptor() ([]byte, []int) {
	return file_go_proto_models_proto_rawDescGZIP(), []int{1}
}

func (x *StoreCategory) GetStoreCategoryUid() string {
	if x != nil {
		return x.StoreCategoryUid
	}
	return ""
}

func (x *StoreCategory) GetStoreCategoryName() string {
	if x != nil {
		return x.StoreCategoryName
	}
	return ""
}

// 飲食店型の定義
type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId       string         `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`                         // 飲食店Id
	StoreName     string         `protobuf:"bytes,2,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`                   // 飲食店名
	StorePlace    string         `protobuf:"bytes,3,opt,name=store_place,json=storePlace,proto3" json:"store_place,omitempty"`                // 飲食店の場所
	StoreCategory *StoreCategory `protobuf:"bytes,4,opt,name=store_category,json=storeCategory,proto3,oneof" json:"store_category,omitempty"` // 飲食店カテゴリ
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_proto_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_go_proto_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_go_proto_models_proto_rawDescGZIP(), []int{2}
}

func (x *Store) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Store) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Store) GetStorePlace() string {
	if x != nil {
		return x.StorePlace
	}
	return ""
}

func (x *Store) GetStoreCategory() *StoreCategory {
	if x != nil {
		return x.StoreCategory
	}
	return nil
}

var File_go_proto_models_proto protoreflect.FileDescriptor

var file_go_proto_models_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x22, 0xdb, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x6d, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xbb, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x0f, 0x5a,
	0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_proto_models_proto_rawDescOnce sync.Once
	file_go_proto_models_proto_rawDescData = file_go_proto_models_proto_rawDesc
)

func file_go_proto_models_proto_rawDescGZIP() []byte {
	file_go_proto_models_proto_rawDescOnce.Do(func() {
		file_go_proto_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_proto_models_proto_rawDescData)
	})
	return file_go_proto_models_proto_rawDescData
}

var file_go_proto_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_proto_models_proto_goTypes = []interface{}{
	(*LoginUser)(nil),     // 0: go_protoc.LoginUser
	(*StoreCategory)(nil), // 1: go_protoc.StoreCategory
	(*Store)(nil),         // 2: go_protoc.Store
}
var file_go_proto_models_proto_depIdxs = []int32{
	1, // 0: go_protoc.Store.store_category:type_name -> go_protoc.StoreCategory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_proto_models_proto_init() }
func file_go_proto_models_proto_init() {
	if File_go_proto_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_proto_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUser); i {
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
		file_go_proto_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCategory); i {
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
		file_go_proto_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
	file_go_proto_models_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_proto_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_proto_models_proto_goTypes,
		DependencyIndexes: file_go_proto_models_proto_depIdxs,
		MessageInfos:      file_go_proto_models_proto_msgTypes,
	}.Build()
	File_go_proto_models_proto = out.File
	file_go_proto_models_proto_rawDesc = nil
	file_go_proto_models_proto_goTypes = nil
	file_go_proto_models_proto_depIdxs = nil
}
