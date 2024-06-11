package loginUsers

import (
	"commandservice/errs"
	"fmt"
	"unicode/utf8"
)

// ユーザIDを保持する値オブジェクト
type LoginUserId struct {
	value string
}

// ゲッター
func (ins *LoginUserId) Value() string {
	return ins.value
}

// 同一性検証
func (ins *LoginUserId) Equals(value *LoginUserId) bool {
	if ins == value {
		return true
	}
	// 値の比較
	return ins.value == value.Value()
}

// コンストラクタ
func NewLoginUserId(value string) (*LoginUserId, *errs.DomainError) {
	const MIN_LENGTH int = 3  // フィールドの最小文字数
	const MAX_LENGTH int = 20 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("ユーザIDの長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &LoginUserId{value: value}, nil
}
