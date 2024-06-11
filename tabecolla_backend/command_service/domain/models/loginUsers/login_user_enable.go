package loginUsers

import (
	"commandservice/errs"
)

// LoginUserEnable ユーザの有効/無効を保持する値オブジェクト
type LoginUserEnable struct {
	value bool
}

// ゲッター
func (ins *LoginUserEnable) Value() bool {
	return ins.value
}

// コンストラクタ
func NewLoginUserEnable(value bool) (*LoginUserEnable, *errs.DomainError) {
	return &LoginUserEnable{value: value}, nil
}
