package loginUsers

import (
	"context"
	"database/sql"
)

// ユーザをアクセスするリポジトリインターフェイス
type LoginUserRepository interface {
	// 同名のユーザがあるか
	Exists(ctx context.Context, tran *sql.Tx, user *LoginUser) error
	// 新しいユーザを永続化する
	Create(ctx context.Context, tran *sql.Tx, user *LoginUser) error
	// ユーザを変更する
	UpdateById(ctx context.Context, tran *sql.Tx, user *LoginUser) error
	// ユーザを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, user *LoginUser) error
}
