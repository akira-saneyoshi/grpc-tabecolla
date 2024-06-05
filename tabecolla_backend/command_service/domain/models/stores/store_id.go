package stores

import (
	"commandservice/errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// 飲食店番号を保持する値オブジェクト(UUIDを保持する)
type StoreId struct {
	value string
}

// valueフィールドのゲッター
func (ins *StoreId) Value() string {
	return ins.value
}

// 同一性検証
func (ins *StoreId) Equlas(value *StoreId) bool {
	if ins == value { // アドレスが同じ?
		return true
	}
	// 値の比較
	return ins.value == value.Value()
}

// コンストラクタ
func NewStoreId(value string) (*StoreId, *errors.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errors.NewDomainError(fmt.Sprintf("飲食店IDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errors.NewDomainError("飲食店IDはUUIDの形式でなければなりません。")
	}
	return &StoreId{value: value}, nil
}
