package users

import (
	"commandservice/errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ユーザのメールアドレスを保持する値オブジェクト
type UserMail struct {
	value string
}

// メールアドレスの正規表現パターン
var emailRegex = regexp.MustCompile(`^[^@]+@[^@]+$`)

// ゲッター
func (ins *UserMail) Value() string {
	return ins.value
}

// コンストラクタ
func NewUserMail(value string) (*UserMail, *errors.DomainError) {
	const MIN_LENGTH int = 7  // フィールドの最小文字数
	const MAX_LENGTH int = 50 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)

	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errors.NewDomainError(fmt.Sprintf("ユーザのメールアドレスの長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}

	// メールアドレスの形式チェック
	if !emailRegex.MatchString(value) {
		errMsg := "ユーザのメールアドレスの形式が正しくありません。"
		return nil, errors.NewDomainError(errMsg)
	}

	// ドメイン部分が '@' で始まるかどうかのチェック
	parts := strings.Split(value, "@")
	if len(parts) != 2 || parts[1] == "" || parts[1][0] == '.' {
		errMsg := "ユーザのメールアドレスのドメイン部分が '@' で始まっています。"
		return nil, errors.NewDomainError(errMsg)
	}

	return &UserMail{value: value}, nil
}
