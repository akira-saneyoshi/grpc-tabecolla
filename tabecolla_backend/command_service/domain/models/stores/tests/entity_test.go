package stores_test

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/domain/models/stores"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storeエンティティ", Ordered, Label("Store構造体の生成"), func() {
	// 前処理
	BeforeAll(func() {
	})
	var _ = Describe("Storeエンティティ", Ordered, Label("Store構造体の生成"), func() {
		Context("インスタンス生成", Label("Create Store"), func() {
			It("新しい飲食店のインスタンス生成", Label("NewStore"), func() {
				store_name, _ := stores.NewStoreName("グランメゾン東京")
				store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２")
				store, _ := stores.NewStore(store_name, store_place, nil)
				Expect(store.StoreId().Value()).ToNot(BeNil())
				Expect(store.StoreName().Value()).To(Equal("グランメゾン東京"))
				Expect(store.StorePlace().Value()).To(Equal("東京都墨田区押上１丁目１−２"))
				Expect(store.StoreCategory()).To(BeNil())
			})
			It("飲食店のインスタンスを構築する", Label("BuildStore"), func() {
				store_id, _ := stores.NewStoreId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				store_name, _ := stores.NewStoreName("グランメゾン東京")
				store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２")
				store := stores.BuildStore(store_id, store_name, store_place, nil)
				Expect(store.StoreId().Value()).To(Equal("ac413f22-0cf1-490a-9635-7e9ca810e544"))
				Expect(store.StoreName().Value()).To(Equal("グランメゾン東京"))
				Expect(store.StorePlace().Value()).To(Equal("東京都墨田区押上１丁目１−２"))
				Expect(store.StoreCategory()).To(BeNil())
			})
		})
	})
})

var _ = Describe("Storeエンティティ", Ordered, Label("Storeの同一性検証"), func() {
	var store_category *storeCategories.StoreCategory
	var store *stores.Store

	// 前処理
	BeforeAll(func() {
		store_category_name, _ := storeCategories.NewStoreCategoryName("フレンチ")
		store_category, _ = storeCategories.NewStoreCategory(store_category_name)
		store_name, _ := stores.NewStoreName("グランメゾン東京")
		store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２")
		store, _ = stores.NewStore(store_name, store_place, store_category)
	})

	// エラーの検証
	Context("エラーの検証", func() {
		It("比較対象がnil", Label("nil検証"), func() {
			By("nilを指定し,DomainErrorを返すことを検証する")
			_, err := store.Equals(nil)
			Expect(err).To(Equal(errs.NewDomainError("引数でnilが指定されました。")))
		})
	})

	// 比較結果の検証
	Context("比較結果の検証", func() {
		It("異なる飲食店ID", Label("false検証"), func() {
			store_name, _ := stores.NewStoreName("グランメゾン東京")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２")
			s, _ := stores.NewStore(store_name, store_place, store_category)
			By("異なるStoreを指定し,falseを返すことを検証する")
			result, _ := store.Equals(s)
			Expect(result).To(Equal(false))
		})
		It("同一の飲食店ID", Label("trueの検証"), func() {
			s := stores.BuildStore(
				store.StoreId(),
				store.StoreName(),
				store.StorePlace(), store_category)
			By("同一のStoreを指定し,trueを返すことを検証する")
			result, _ := store.Equals(s)
			Expect(result).To(Equal(true))
		})
	})
})
