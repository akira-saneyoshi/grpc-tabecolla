package storeCategories

import (
	"commandservice/errors"
	"fmt"
	"unicode/utf8"
)

// カテゴリ名を保持する値オブジェクト
type StoreCategoryName struct {
	value string
}

// valueフィールドのゲッター
func (ins *StoreCategoryName) Value() string {
	return ins.value
}

// コンストラクタ
func NewStoreCategoryName(value string) (*StoreCategoryName, *errors.DomainError) {
	const MIN_LENGTH int = 2  // フィールドの最小文字数
	const MAX_LENGTH int = 20 // フィールドの最大文字数
	// 引数の文字数チェック
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errors.NewDomainError(fmt.Sprintf("飲食店カテゴリ名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &StoreCategoryName{value: value}, nil
}
