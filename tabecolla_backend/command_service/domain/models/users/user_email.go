package users

import (
	"commandservice/errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// ユーザのメールアドレスを保持する値オブジェクト
type UserEmail struct {
	value string
}

// メールアドレスの正規表現パターン
var emailRegex = regexp.MustCompile(`^[^@]+@[^@]+$`)

// ゲッター
func (ins *UserEmail) Value() string {
	return ins.value
}

// コンストラクタ
func NewUserEmail(value string) (*UserEmail, *errs.DomainError) {
	const MIN_LENGTH int = 7  // フィールドの最小文字数
	const MAX_LENGTH int = 50 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)

	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("ユーザのメールアドレスの長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}

	// メールアドレスの形式チェック
	if !emailRegex.MatchString(value) {
		errMsg := "ユーザのメールアドレスの形式が正しくありません。"
		return nil, errs.NewDomainError(errMsg)
	}

	// ドメイン部分が '@' で始まるかどうかのチェック
	parts := strings.Split(value, "@")
	if len(parts) != 2 || parts[1] == "" || parts[1][0] == '@' {
		errMsg := "ユーザのメールアドレスのドメイン部分が '@' で始まっています。"
		return nil, errs.NewDomainError(errMsg)
	}

	return &UserEmail{value: value}, nil
}