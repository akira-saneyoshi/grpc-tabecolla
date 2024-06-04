package stores

import (
	"commandservice/errors"
	"fmt"
	"unicode/utf8"
)

// 飲食店名を保持する値オブジェクト
type StoreName struct {
	value string
}

// ゲッター
func (ins *StoreName) Value() string {
	return ins.value
}

// コンストラクタ
func NewStoreName(value string) (*StoreName, *errs.DomainError) {
	const MIN_LENGTH int = 5  // フィールドの最小文字数
	const MAX_LENGTH int = 30 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("飲食店名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &StoreName{value: value}, nil
}
