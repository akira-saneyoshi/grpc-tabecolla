package store_categories_test

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StoreCategoryエンティティ", Ordered, Label("StoreCategory構造体の生成"), func() {
	Context("インスタンス生成", Label("Create StoreCategory"), func() {
		It("新しい飲食店カテゴリを生成する", Label("NewStoreCategory"), func() {
			store_category_name, _ := storeCategories.NewStoreCategoryName("イタリアン")
			storeCategory, _ := storeCategories.NewStoreCategory(store_category_name)
			Expect(storeCategory.StoreCategoryId()).ToNot(BeNil())
			Expect(storeCategory.StoreCategoryName().Value()).To(Equal("イタリアン"))
		})
		It("飲食店カテゴリを再構築する", Label("BuildStoreCategory"), func() {
			store_category_id, _ := storeCategories.NewStoreCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			store_category_name, _ := storeCategories.NewStoreCategoryName("イタリアン")
			storeCategory := storeCategories.BuildStoreCategory(store_category_id, store_category_name)
			Expect(storeCategory.StoreCategoryId().Value()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
			Expect(storeCategory.StoreCategoryName().Value()).To(Equal("イタリアン"))
		})
	})
})

var _ = Describe("StoreCategoryエンティティ", Ordered, Label("StoreCategoryの同一性検証"), func() {
	It("比較対象がnil", Label("nil検証"), func() {
		By("インスタンスの生成")
		store_category := storeCategories.BuildStoreCategory(nil, nil)
		By("Equals()にnilを指定")
		_, err := store_category.Equals(nil)
		By("errors.DomainErrorを評価する")
		Expect(err).To(Equal(errors.NewDomainError("引数でnilが指定されました。")))
	})
	It("異なる飲食店カテゴリID", Label("false検証"), func() {
		By("2つインスタンスの生成")
		store_category_name, _ := storeCategories.NewStoreCategoryName("イタリアン")
		store_category_a, _ := storeCategories.NewStoreCategory(store_category_name)
		store_category_name, _ = storeCategories.NewStoreCategoryName("イタリアン")
		store_category_b, _ := storeCategories.NewStoreCategory(store_category_name)
		By("Equals()にstore_category_bを指定")
		result, _ := store_category_a.Equals(store_category_b)
		By("falseを評価する")
		Expect(result).To(Equal(false))
	})
	It("同一の飲食店カテゴリID", Label("trueの検証"), func() {
		By("2つインスタンスの生成")
		store_category_name, _ := storeCategories.NewStoreCategoryName("イタリアン")
		store_category_a, _ := storeCategories.NewStoreCategory(store_category_name)

		store_category_b := storeCategories.BuildStoreCategory(store_category_a.StoreCategoryId(), store_category_a.StoreCategoryName())
		By("Equals()にstore_category_bを指定")
		result, _ := store_category_a.Equals(store_category_b)
		By("trueを評価する")
		Expect(result).To(Equal(true))
	})
})
