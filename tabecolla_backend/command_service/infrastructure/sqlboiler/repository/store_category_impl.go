package repository

import (
	"commandservice/domain/models/storeCategories"
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

// CategoryRepositoryインターフェイスの実装
type storeCategoryRepositorySQLBoiler struct{}

// コンストラクタ
func NewStoreCategoryRepositorySQLBoiler() storeCategories.StoreCategoryRepository {
	// フック関数の登録
	models.AddStoreCategoryHook(boil.AfterInsertHook, StoreCategoryAfterInsertHook)
	models.AddStoreCategoryHook(boil.AfterUpdateHook, StoreCategoryAfterUpdateHook)
	models.AddStoreCategoryHook(boil.AfterDeleteHook, StoreCategoryAfterDeleteHook)
	return &storeCategoryRepositorySQLBoiler{}
}

// 同名の飲食店カテゴリが存在確認結果を返す
func (rep *storeCategoryRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, store_category *storeCategories.StoreCategory) error {
	// レコードの存在確認条件を作成する
	condition := models.StoreCategoryWhere.Name.EQ(store_category.StoreCategoryName().Value())
	// レコードの存在を確認するクエリを実行する
	if exists, err := models.StoreCategories(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists { // 同じ名称の飲食店カテゴリは存在していない?
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", store_category.StoreCategoryName().Value()))
	}
}

// 新しい飲食店カテゴリを永続化する
func (rep *storeCategoryRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, store_category *storeCategories.StoreCategory) error {
	// SqlBoilerのモデルを生成する
	new_store_category := models.StoreCategory{
		ID:    0,
		ObjID: store_category.StoreCategoryUid().Value(),
		Name:  store_category.StoreCategoryName().Value(),
	}
	// 飲食店カテゴリを永続化する
	if err := new_store_category.Insert(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 飲食店カテゴリを変更する
func (rep *storeCategoryRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, store_category *storeCategories.StoreCategory) error {
	// 更新対象を取得する
	up_model, err := models.StoreCategories(qm.Where("obj_id = ?", store_category.StoreCategoryUid().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("飲食店カテゴリ番号:%sは存在しないため、更新できませんでした。", store_category.StoreCategoryUid().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 取得したモデルの値を変更する
	up_model.ObjID = store_category.StoreCategoryUid().Value()
	up_model.Name = store_category.StoreCategoryName().Value()
	// 更新を実行する
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 飲食店カテゴリを削除する
func (rep *storeCategoryRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, store_category *storeCategories.StoreCategory) error {
	// 削除対象を取得する
	del_model, err := models.StoreCategories(qm.Where("obj_id = ?", store_category.StoreCategoryUid().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("飲食店カテゴリ番号:%sは存在しないため、削除できませんでした。",
		store_category.StoreCategoryUid().Value()))
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
func StoreCategoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, store_category *models.StoreCategory) error {
	log.Printf("飲食店カテゴリID:%s 飲食店カテゴリ名:%sを登録しました。\n", store_category.ObjID, store_category.Name)
	return nil
}

// 変更処理後に実行されるフック
func StoreCategoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, store_category *models.StoreCategory) error {
	log.Printf("飲食店カテゴリID:%s 飲食店カテゴリ名:%sを変更しました。\n", store_category.ObjID, store_category.Name)
	return nil
}

// 削除処理後に実行されるフック
func StoreCategoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, store_category *models.StoreCategory) error {
	log.Printf("飲食店カテゴリID:%s 飲食店カテゴリ名:%sを削除しました。\n", store_category.ObjID, store_category.Name)
	return nil
}
