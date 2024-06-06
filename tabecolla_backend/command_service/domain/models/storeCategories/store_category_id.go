package storeCategories

import (
	"commandservice/errs"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// 飲食店カテゴリ番号を保持する値オブジェクト(UUIDを保持する)
type StoreCategoryUid struct {
	value string
}

// valueフィールドのゲッター
func (ins *StoreCategoryUid) Value() string {
	return ins.value
}

// 同一性検証
func (ins *StoreCategoryUid) Equals(value *StoreCategoryUid) bool {
	if ins == value { // アドレスが同じ?
		return true
	}
	// 値が同じ
	return ins.value == value.Value()
}

// コンストラクタ
func NewStoreCategoryUid(value string) (*StoreCategoryUid, *errs.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("飲食店カテゴリIDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("飲食店カテゴリIDはUUIDの形式でなければなりません。")
	}
	return &StoreCategoryUid{value: value}, nil
}
