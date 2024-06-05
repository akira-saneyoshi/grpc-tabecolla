package stores

import (
	"context"
	"database/sql"
)

// 飲食店をアクセスするリポジトリインターフェイス
type StoreRepository interface {
	// 同名の飲食店があるか
	Exists(ctx context.Context, tran *sql.Tx, store *Store) error
	// 新しい飲食店を永続化する
	Create(ctx context.Context, tran *sql.Tx, store *Store) error
	// 飲食店を変更する
	UpdateById(ctx context.Context, tran *sql.Tx, store *Store) error
	// 飲食店を削除する
	DeleteById(ctx context.Context, tran *sql.Tx, store *Store) error
}
