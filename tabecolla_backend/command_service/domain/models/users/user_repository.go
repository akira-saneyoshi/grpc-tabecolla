package users

import (
	"context"
	"database/sql"
)

// ユーザをアクセスするリポジトリインターフェイス
type UserRepository interface {
	// 同名のユーザがあるか
	Exists(ctx context.Context, tran *sql.Tx, store *User) error
	// 新しいユーザを永続化する
	Create(ctx context.Context, tran *sql.Tx, store *User) error
	// ユーザを変更する
	UpdateById(ctx context.Context, tran *sql.Tx, store *User) error
	// ユーザを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, store *User) error
}
