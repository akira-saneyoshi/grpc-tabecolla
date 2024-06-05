package storeCategories

import (
	"commandservice/errors"

	"github.com/google/uuid"
)

// 飲食店カテゴリを表すEntity
type StoreCategory struct {
	store_category_id   *StoreCategoryId   // 飲食店カテゴリ番号
	store_category_name *StoreCategoryName // 飲食店カテゴリ名
}

// ゲッター
func (ins *StoreCategory) StoreCategoryId() *StoreCategoryId {
	return ins.store_category_id
}

func (ins *StoreCategory) StoreCategoryName() *StoreCategoryName {
	return ins.store_category_name
}

// 値の変更
func (ins *StoreCategory) ChangeStoreCategoryName(store_category_name *StoreCategoryName) {
	ins.store_category_name = store_category_name
}

// 同一性検証
func (ins *StoreCategory) Equals(obj *StoreCategory) (bool, *errors.DomainError) {
	if obj == nil {
		return false, errors.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.store_category_id.Equals(obj.StoreCategoryId())
	return result, nil
}

// コンストラクタ
func NewStoreCategory(store_category_name *StoreCategoryName) (*StoreCategory, *errors.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errors.NewDomainError(err.Error())
	} else {
		if store_category_id, err := NewStoreCategoryId(uid.String()); err != nil {
			return nil, errors.NewDomainError(err.Error())
		} else {
			return &StoreCategory{
				store_category_id:   store_category_id,
				store_category_name: store_category_name,
			}, nil
		}
	}
}

// StoreCategoryエンティティを再構築する
func BuildStoreCategory(store_category_id *StoreCategoryId, store_category_name *StoreCategoryName) *StoreCategory {
	return &StoreCategory{
		store_category_id:   store_category_id,
		store_category_name: store_category_name,
	}
}
