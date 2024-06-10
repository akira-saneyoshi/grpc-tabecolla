package repository

import (
	"commandservice/domain/models/stores"
	"commandservice/errs"
	"commandservice/infrastructure/sqlboiler/handler"
	"commandservice/infrastructure/sqlboiler/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// StoreRepositoryインターフェイスの実装
type storeRepositorySQLBoiler struct{}

// コンストラクタ
func NewStoreRepositorySQLBoiler() stores.StoreRepository {
	// フック関数の登録
	models.AddStoreHook(boil.AfterInsertHook, StoreAfterInsertHook)
	models.AddStoreHook(boil.AfterUpdateHook, StoreAfterUpdateHook)
	models.AddStoreHook(boil.AfterDeleteHook, StoreAfterDeleteHook)
	return &storeRepositorySQLBoiler{}
}

// 同名の飲食店の存在確認結果を返す
func (rep *storeRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, store *stores.Store) error {
	// レコードの存在確認条件を作成する
	condition := models.StoreWhere.Name.EQ(store.StoreName().Value())
	// レコードの存在を確認するクエリを実行する
	exists, err := models.Stores(condition).Exists(ctx, tran)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	if !exists { // 同じ名称の飲食店は存在していない?
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", store.StoreName().Value()))
	}
}

// 新しい飲食店を永続化する
func (rep *storeRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, store *stores.Store) error {
	// SqlBoilerのモデルを生成する
	new_store := models.Store{
		ID:         0,
		ObjID:      store.StoreId().Value(),
		Name:       store.StoreName().Value(),
		Place:      store.StorePlace().Value(),
		StoreCategoryID: store.StoreCategory().StoreCategoryUid().Value(),
	}
	// 飲食店を永続化する
	if err := new_store.Insert(ctx, tran, boil.Whitelist("obj_id", "name", "place", "store_category_id")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 飲食店を変更する
func (rep *storeRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, store *stores.Store) error {
	// 更新対象を取得する
	up_model, err := models.Stores(qm.Where("obj_id = ?", store.StoreId().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("飲食店番号:%sは存在しないため、更新できませんでした。", store.StoreId().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 取得したモデルの値を変更する
	up_model.Name = store.StoreName().Value()
	up_model.Place = store.StorePlace().Value()
	// 更新を実行する
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name", "place")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 飲食店を削除する
func (rep *storeRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, store *stores.Store) error {
	// 削除対象を取得する
	del_model, err := models.Stores(qm.Where("obj_id = ?", store.StoreId().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("飲食店番号:%sは存在しないため、削除できませんでした。", store.StoreId().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 削除を実行する
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 登録処理後に実行されるフック
func StoreAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, store *models.Store) error {
	log.Printf("飲食店ID:%s 飲食店名:%s 単価:%d 飲食店カテゴリ番号: %s を登録しました。\n",
		store.ObjID, store.Name, store.Place, store.StoreCategoryID)
	return nil
}

// 変更処理後に実行されるフック
func StoreAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, store *models.Store) error {
	log.Printf("飲食店ID:%s 飲食店名:%s 単価:%d 飲食店カテゴリ番号: %s を変更しました。\n",
		store.ObjID, store.Name, store.Place, store.StoreCategoryID)
	return nil
}

// 削除処理後に実行されるフック
func StoreAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, store *models.Store) error {
	log.Printf("飲食店ID:%s 飲食店名:%s 単価:%d 飲食店カテゴリ番号: %s を削除しました。\n",
		store.ObjID, store.Name, store.Place, store.StoreCategoryID)
	return nil
}
