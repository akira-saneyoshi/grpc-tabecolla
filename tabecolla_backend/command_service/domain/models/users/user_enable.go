package users

import (
	"commandservice/errs"
)

// UserEnable ユーザの有効/無効を保持する値オブジェクト
type UserEnable struct {
	value bool
}

// ゲッター
func (ins *UserEnable) Value() bool {
	return ins.value
}

// コンストラクタ
func NewUserEnable(value bool) (*UserEnable, *errs.DomainError) {
	return &UserEnable{value: value}, nil
}
