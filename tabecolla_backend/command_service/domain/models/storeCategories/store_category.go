package storeCategories

import (
	"commandservice/errs"

	"github.com/google/uuid"
)

// 飲食店カテゴリを表すEntity
type StoreCategory struct {
	store_category_uid  *StoreCategoryUid  // 飲食店カテゴリ番号
	store_category_name *StoreCategoryName // 飲食店カテゴリ名
}

// ゲッター
func (ins *StoreCategory) StoreCategoryUid() *StoreCategoryUid {
	return ins.store_category_uid
}

func (ins *StoreCategory) StoreCategoryName() *StoreCategoryName {
	return ins.store_category_name
}

// 値の変更
func (ins *StoreCategory) ChangeStoreCategoryName(store_category_name *StoreCategoryName) {
	ins.store_category_name = store_category_name
}

// 同一性検証
func (ins *StoreCategory) Equals(obj *StoreCategory) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.store_category_uid.Equals(obj.StoreCategoryUid())
	return result, nil
}

// コンストラクタ
func NewStoreCategory(store_category_name *StoreCategoryName) (*StoreCategory, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	} else {
		if store_category_uid, err := NewStoreCategoryUid(uid.String()); err != nil {
			return nil, errs.NewDomainError(err.Error())
		} else {
			return &StoreCategory{
				store_category_uid:  store_category_uid,
				store_category_name: store_category_name,
			}, nil
		}
	}
}

// StoreCategoryエンティティを再構築する
func BuildStoreCategory(store_category_uid *StoreCategoryUid, store_category_name *StoreCategoryName) *StoreCategory {
	return &StoreCategory{
		store_category_uid:  store_category_uid,
		store_category_name: store_category_name,
	}
}
