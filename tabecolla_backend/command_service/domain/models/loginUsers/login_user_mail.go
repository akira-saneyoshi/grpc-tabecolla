package loginUsers

import (
	"commandservice/errs"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ユーザのメールアドレスを保持する値オブジェクト
type LoginUserMail struct {
	value string
}

// メールアドレスの正規表現パターン
var emailRegex = regexp.MustCompile(`^[^@]+@[^@]+$`)

// ゲッター
func (ins *LoginUserMail) Value() string {
	return ins.value
}

// 同一性検証
func (ins *LoginUserMail) Equals(value *LoginUserMail) bool {
	if ins == value {
		return true
	}
	// 値の比較
	return ins.value == value.Value()
}

// コンストラクタ
func NewLoginUserMail(value string) (*LoginUserMail, *errs.DomainError) {
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
	if len(parts) != 2 || parts[1] == "" || parts[1][0] == '.' {
		errMsg := "ユーザのメールアドレスのドメイン部分が '@' で始まっています。"
		return nil, errs.NewDomainError(errMsg)
	}

	return &LoginUserMail{value: value}, nil
}
