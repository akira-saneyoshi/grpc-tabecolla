package store_categories_test

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/errs"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StoreCategoryエンティティを構成する値オブジェクト", Ordered, Label("StoreCategoryUid構造体の生成"), func() {
	var empty_str *errs.DomainError                          // 空文字列　長さ36に違反する
	var length_over *errs.DomainError                        // 36文字より大きい文字列　長さ36に違反する
	var not_uuid *errs.DomainError                           // UUID以外の文字列を指定する
	var store_category_uid *storeCategories.StoreCategoryUid // UUID文字列を指定する
	var uid string

	BeforeAll(func() {
		_, empty_str = storeCategories.NewStoreCategoryUid("")
		_, length_over = storeCategories.NewStoreCategoryUid("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd")
		_, not_uuid = storeCategories.NewStoreCategoryUid("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		store_category_uid, _ = storeCategories.NewStoreCategoryUid(id.String())
	})

	Context("文字数の検証", Label("文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("飲食店カテゴリIDの長さは36文字でなければなりません。")))
		})
		It("36文字より大きい文字列の場合、errors.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("飲食店カテゴリIDの長さは36文字でなければなりません。")))
		})
	})

	Context("UUID形式の検証", Label("UUID形式"), func() {
		It("uuid以外の文字列の場合、errors.DomainErrorが返る", func() {
			Expect(not_uuid).To(Equal(errs.NewDomainError("飲食店カテゴリIDはUUIDの形式でなければなりません。")))
		})
		It("36文字のuuid文字列の場合、StoreCategoryUidが返る", func() {
			id, _ := storeCategories.NewStoreCategoryUid(uid)
			Expect(store_category_uid).To(Equal(id))
		})
	})
})

var _ = Describe("StoreCategoryエンティティを構成する値オブジェクト", Ordered, Label("StoreCategoryName構造体の生成"), func() {
	var empty_str *errs.DomainError                            // 空文字列　長さ2文字以上、20文字以内に違反する
	var length_over *errs.DomainError                          // 20文字より大きい文字列　長さ2文字以上、20文字以内に違反する
	var store_category_name *storeCategories.StoreCategoryName // 20文字以内の文字列を指定する

	BeforeAll(func() {
		_, empty_str = storeCategories.NewStoreCategoryName("")
		_, length_over = storeCategories.NewStoreCategoryName("aaaaaaaaaabbbbbbbbbbc")
		store_category_name, _ = storeCategories.NewStoreCategoryName("イタリアン")
	})

	Context("文字数の検証", Label("無効な文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("飲食店カテゴリ名の長さは2文字以上、20文字以内です。")))
		})
		It("20文字より大きいの場合,errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("飲食店カテゴリ名の長さは2文字以上、20文字以内です。")))
		})
	})

	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("2文字以上20文字以内場合,NewStoreCategoryNameが返る", func() {
			Expect(store_category_name.Value()).To(Equal("イタリアン"))
		})
	})
})
