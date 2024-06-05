package storeCategories

import (
	"commandservice/errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// 飲食店カテゴリ番号を保持する値オブジェクト(UUIDを保持する)
type StoreCategoryId struct {
	value string
}

// valueフィールドのゲッター
func (ins *StoreCategoryId) Value() string {
	return ins.value
}

// 同一性検証
func (ins *StoreCategoryId) Equals(value *StoreCategoryId) bool {
	if ins == value { // アドレスが同じ?
		return true
	}
	// 値が同じ
	return ins.value == value.Value()
}

// コンストラクタ
func NewStoreCategoryId(value string) (*StoreCategoryId, *errors.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errors.NewDomainError(fmt.Sprintf("飲食店カテゴリIDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errors.NewDomainError("飲食店カテゴリIDはUUIDの形式でなければなりません。")
	}
	return &StoreCategoryId{value: value}, nil
}
