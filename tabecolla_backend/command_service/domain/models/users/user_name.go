package users

import (
	"commandservice/errors"
	"fmt"
	"unicode/utf8"
)

// ユーザ名を保持する値オブジェクト
type UserName struct {
	value string
}

// ゲッター
func (ins *UserName) Value() string {
	return ins.value
}

// コンストラクタ
func NewUserName(value string) (*UserName, *errors.DomainError) {
	const MIN_LENGTH int = 3  // フィールドの最小文字数
	const MAX_LENGTH int = 20 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errors.NewDomainError(fmt.Sprintf("ユーザ名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &UserName{value: value}, nil
}
