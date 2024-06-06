package users

import (
	"commandservice/errs"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// ユーザのパスワードを保持する値オブジェクト
type UserPassword struct {
	value string
}

// ゲッター
func (ins *UserPassword) Value() string {
	return ins.value
}

// コンストラクタ
func NewUserPassword(value string) (*UserPassword, *errs.DomainError) {
	const MIN_LENGTH int = 8  // フィールドの最小文字数
	const MAX_LENGTH int = 15 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("ユーザのパスワードの長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}

	// 英数字と記号が含まれているかをチェック
	var validPassword = regexp.MustCompile(`^[a-zA-Z0-9!-/:-@¥[-` + "`" + `{-~]+$`)
	if !validPassword.MatchString(value) {
		return nil, errs.NewDomainError("パスワードには半角の英数字や記号を含める必要があります。")
	}

	return &UserPassword{value: value}, nil
}
