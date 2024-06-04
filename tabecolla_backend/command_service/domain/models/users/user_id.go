package users

import (
	"commandservice/errors"
	"fmt"
	"unicode/utf8"
)

// ユーザIDを保持する値オブジェクト
type UserId struct {
	value string
}

// ゲッター
func (ins *UserId) Value() string {
	return ins.value
}

// コンストラクタ
func NewUserId(value string) (*UserId, *errs.DomainError) {
	const MIN_LENGTH int = 3  // フィールドの最小文字数
	const MAX_LENGTH int = 20 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("ユーザIDの長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &UserId{value: value}, nil
}
