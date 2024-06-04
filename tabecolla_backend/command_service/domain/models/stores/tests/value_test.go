package stores_test

import (
	"commandservice/domain/models/stores"
	"commandservice/errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storeエンティティを構成する値オブジェクト", Ordered, Label("StoreId構造体の生成"), func() {
	var empty_str *errors.DomainError   // 空文字列　長さ36に違反する
	var length_over *errors.DomainError // 36文字より大きい文字列　長さ36に違反する
	var not_uuid *errors.DomainError    // UUID以外の文字列を指定する
	var store_id *stores.StoreId        // UUID文字列を指定する
	var uid string

	BeforeAll(func() {
		_, empty_str = stores.NewStoreId("")
		_, length_over = stores.NewStoreId("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd")
		_, not_uuid = stores.NewStoreId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		store_id, _ = stores.NewStoreId(id.String())
	})

	Context("文字数の検証", Label("文字数"), func() {
		It("空文字列の場合、errs.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errors.NewDomainError("飲食店IDの長さは36文字でなければなりません。")))
		})
		It("36文字より大きい文字列の場合、errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errors.NewDomainError("飲食店IDの長さは36文字でなければなりません。")))
		})
	})

	Context("UUID形式の検証", Label("UUID形式"), func() {
		It("uuid以外の文字列の場合、errs.DomainErrorが返る", func() {
			Expect(not_uuid).To(Equal(errors.NewDomainError("飲食店IDはUUIDの形式でなければなりません。")))
		})
		It("36文字のuuid文字列の場合、StoreIdが返る", func() {
			id, _ := stores.NewStoreId(uid)
			Expect(store_id).To(Equal(id))
		})
	})
})
