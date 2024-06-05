package store_categories_test

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StoreCategoryエンティティを構成する値オブジェクト", Ordered, Label("StoreCategoryId構造体の生成"), func() {
	var empty_str *errors.DomainError                      // 空文字列　長さ36に違反する
	var length_over *errors.DomainError                    // 36文字より大きい文字列　長さ36に違反する
	var not_uuid *errors.DomainError                       // UUID以外の文字列を指定する
	var store_category_id *storeCategories.StoreCategoryId // UUID文字列を指定する
	var uid string

	BeforeAll(func() {
		_, empty_str = storeCategories.NewStoreCategoryId("")
		_, length_over = storeCategories.NewStoreCategoryId("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd")
		_, not_uuid = storeCategories.NewStoreCategoryId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		store_category_id, _ = storeCategories.NewStoreCategoryId(id.String())
	})

	Context("文字数の検証", Label("文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errors.NewDomainError("飲食店カテゴリIDの長さは36文字でなければなりません。")))
		})
		It("36文字より大きい文字列の場合、errors.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errors.NewDomainError("飲食店カテゴリIDの長さは36文字でなければなりません。")))
		})
	})

	Context("UUID形式の検証", Label("UUID形式"), func() {
		It("uuid以外の文字列の場合、errors.DomainErrorが返る", func() {
			Expect(not_uuid).To(Equal(errors.NewDomainError("飲食店カテゴリIDはUUIDの形式でなければなりません。")))
		})
		It("36文字のuuid文字列の場合、StoreCategoryIdが返る", func() {
			id, _ := storeCategories.NewStoreCategoryId(uid)
			Expect(store_category_id).To(Equal(id))
		})
	})
})

var _ = Describe("StoreCategoryエンティティを構成する値オブジェクト", Ordered, Label("StoreCategoryName構造体の生成"), func() {
	var empty_str *errors.DomainError                          // 空文字列　長さ2文字以上、20文字以内に違反する
	var length_over *errors.DomainError                        // 20文字より大きい文字列　長さ2文字以上、20文字以内に違反する
	var store_category_name *storeCategories.StoreCategoryName // 20文字以内の文字列を指定する

	BeforeAll(func() {
		_, empty_str = storeCategories.NewStoreCategoryName("")
		_, length_over = storeCategories.NewStoreCategoryName("aaaaaaaaaabbbbbbbbbbc")
		store_category_name, _ = storeCategories.NewStoreCategoryName("イタリアン")
	})

	Context("文字数の検証", Label("無効な文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errors.NewDomainError("飲食店カテゴリ名の長さは2文字以上、20文字以内です。")))
		})
		It("20文字より大きいの場合,errors.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errors.NewDomainError("飲食店カテゴリ名の長さは2文字以上、20文字以内です。")))
		})
	})

	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("2文字以上20文字以内場合,NewStoreCategoryNameが返る", func() {
			Expect(store_category_name.Value()).To(Equal("イタリアン"))
		})
	})
})
