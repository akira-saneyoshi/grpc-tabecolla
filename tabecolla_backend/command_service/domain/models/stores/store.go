package stores

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/errs"

	"github.com/google/uuid"
)

// 飲食店を表すEntity
type Store struct {
	store_id       *StoreId                       // 飲食店Id
	store_name     *StoreName                     // 飲食店名
	store_place    *StorePlace                    // 飲食店の場所
	store_category *storeCategories.StoreCategory // 飲食店カテゴリ
}

// ゲッター
func (ins *Store) StoreId() *StoreId {
	return ins.store_id
}
func (ins *Store) StoreName() *StoreName {
	return ins.store_name
}
func (ins *Store) StorePlace() *StorePlace {
	return ins.store_place
}
func (ins *Store) StoreCategory() *storeCategories.StoreCategory {
	return ins.store_category
}

// 値の変更
func (ins *Store) ChangeProductName(store_name *StoreName) {
	ins.store_name = store_name
}
func (ins *Store) ChangeProductPrice(store_place *StorePlace) {
	ins.store_place = store_place
}
func (ins *Store) ChangeCategory(store_category *storeCategories.StoreCategory) {
	ins.store_category = store_category
}

// 同一性検証メソッド
func (ins *Store) Equals(obj *Store) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.store_id.Equals(obj.StoreId())
	return result, nil
}

// コンストラクタ
func NewStore(store_name *StoreName, store_place *StorePlace, store_category *storeCategories.StoreCategory) (*Store, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	} else {
		// 飲食店ID用値オブジェクトを生成する
		if store_id, err := NewStoreId(uid.String()); err != nil {
			return nil, err
		} else {
			// 飲食店エンティティのインスタンスを生成して返す
			return &Store{
				store_id:       store_id,
				store_name:     store_name,
				store_place:    store_place,
				store_category: store_category,
			}, nil
		}
	}
}

// 飲食店エンティティの再構築
func BuildStore(store_id *StoreId, store_name *StoreName, store_place *StorePlace, store_category *storeCategories.StoreCategory) *Store {
	store := Store{ // 飲食店エンティティのインスタンスを生成して返す
		store_id:       store_id,
		store_name:     store_name,
		store_place:    store_place,
		store_category: store_category,
	}
	return &store
}
