package storeCategories

import (
	"context"
	"database/sql"
)

// 飲食店カテゴリをアクセスするリポジトリインターフェイス
type StoreCategoryRepository interface {
	// 同名の飲食店カテゴリが存在確認結果を返す
	Exists(ctx context.Context, tran *sql.Tx, storeCategory *StoreCategory) error
	// 新しい飲食店カテゴリを永続化する
	Create(ctx context.Context, tran *sql.Tx, storeCategory *StoreCategory) error
	// 飲食店カテゴリを変更する
	UpdateById(ctx context.Context, tran *sql.Tx, storeCategory *StoreCategory) error
	// 飲食店カテゴリを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, storeCategory *StoreCategory) error
}
