package stores

import (
	"commandservice/errors"
	"fmt"
	"unicode/utf8"
)

// 飲食店の場所を保持する値オブジェクト
type StorePlace struct {
	value string
}

// ゲッター
func (ins *StorePlace) Value() string {
	return ins.value
}

// コンストラクタ
func NewStorePlace(value string) (*StorePlace, *errors.DomainError) {
	const MIN_LENGTH int = 3  // フィールドの最小文字数
	const MAX_LENGTH int = 50 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errors.NewDomainError(fmt.Sprintf("飲食店の場所の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &StorePlace{value: value}, nil
}
