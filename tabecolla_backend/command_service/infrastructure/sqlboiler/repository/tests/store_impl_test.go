package repository_test

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/domain/models/stores"
	"commandservice/errs"
	"commandservice/infrastructure/sqlboiler/repository"
	"context"
	"database/sql"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ = Describe("storeRepositorySQLBoiler構造体", Ordered, Label("StoreRepositoryインターフェイスメソッドのテスト"), func() {
	var store_category *storeCategories.StoreCategory
	var rep stores.StoreRepository
	var ctx context.Context
	var tran *sql.Tx
	// 前処理
	BeforeAll(func() {
		// リポジトリの生成
		rep = repository.NewStoreRepositorySQLBoiler() // Repositoryの生成
		// カテゴリエンティティの生成
		store_category_id, _ := storeCategories.NewStoreCategoryUid("eeee1111-22ff-33cc-44dd-666666eeeeee")
		store_category_name, _ := storeCategories.NewStoreCategoryName("ハンバーガー")
		store_category = storeCategories.BuildStoreCategory(store_category_id, store_category_name)
	})
	// テスト毎の前処理
	BeforeEach(func() {
		ctx = context.Background()       // Contextの生成
		tran, _ = boil.BeginTx(ctx, nil) // トランザクションの開始
	})
	// テスト毎の後処理
	AfterEach(func() {
		tran.Rollback() // トランザクションのロールバック
	})
	// Exists()メソッドのテスト
	Context("同名の飲食店の存在確認結果を返す", Label("Exists"), func() {
		It("存在しない飲食店の場合nilが返る", func() {
			store_name, _ := stores.NewStoreName("McDonald's")
			store := stores.BuildStore(nil, store_name, nil, nil)
			result := rep.Exists(ctx, tran, store)
			Expect(result).To(BeNil())
		})
		It("存在する飲食店の場合、errs.CRUDErrorが返る", func() {
			store_name, _ := stores.NewStoreName("モスバーガー")
			store := stores.BuildStore(nil, store_name, nil, nil)
			result := rep.Exists(ctx, tran, store)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", store_name.Value()))))
		})
	})
	// Create()メソッドのテスト
	Context("新しい飲食店を永続化する", Label("Create"), func() {
		It("飲食店が登録成功し、nilが返る", func() {
			store_name, _ := stores.NewStoreName("McDonald's")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２ 2 番地 東京スカイツリータウン・ソラマチ 4F ウエストヤード")
			store, _ := stores.NewStore(store_name, store_place, store_category)
			result := rep.Create(ctx, tran, store)
			Expect(result).To(BeNil())
		})
		//	productテーブルのobj_idをユニーク設定後テスト!!
		It("obj_idが同じ飲食店カテゴリを追加すると、errs.CRUDErrorが返る", func() {
			store_id, _ := stores.NewStoreId("ffff1111-99cc-bb88-22aa-77ee77ee77ee")
			store_name, _ := stores.NewStoreName("McDonald's")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２ 2 番地 東京スカイツリータウン・ソラマチ 4F ウエストヤード")
			store := stores.BuildStore(store_id, store_name, store_place, store_category)
			result := rep.Create(ctx, tran, store)
			err := errs.NewCRUDError("一意制約違反です。")
			Expect(result).To(Equal(err))
		})
	})
	// UpdateById()メソッドのテスト
	Context("飲食店を変更する", Label("UpdateById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			store_name, _ := stores.NewStoreName("McDonald's")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２ 2 番地 東京スカイツリータウン・ソラマチ 4F ウエストヤード")
			store, _ := stores.NewStore(store_name, store_place, store_category)
			result := rep.UpdateById(ctx, tran, store)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("飲食店番号:%sは存在しないため、更新できませんでした。", store.StoreId().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			store_id, _ := stores.NewStoreId("ffff1111-99cc-bb88-22aa-77ee77ee77ee")
			store_name, _ := stores.NewStoreName("McDonald's")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２ 2 番地 東京スカイツリータウン・ソラマチ 4F ウエストヤード")
			store := stores.BuildStore(store_id, store_name, store_place, store_category)
			result := rep.UpdateById(ctx, tran, store)
			Expect(result).To(BeNil())
		})
	})
	// DeleteById()メソッドのテスト
	Context("飲食店を削除する", Label("DeleteById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			store_name, _ := stores.NewStoreName("McDonald's")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目１−２ 2 番地 東京スカイツリータウン・ソラマチ 4F ウエストヤード")
			store, _ := stores.NewStore(store_name, store_place, nil)
			result := rep.DeleteById(ctx, tran, store)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("飲食店番号:%sは存在しないため、削除できませんでした。", store.StoreId().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 削除対象の飲食店カテゴリを登録する
			store_id, _ := stores.NewStoreId("ffff1111-99cc-bb88-22aa-77ee77ee77ee")
			store_name, _ := stores.NewStoreName("モスバーガー")
			store_place, _ := stores.NewStorePlace("東京都墨田区押上１丁目２３−１ 土居ビル １階")
			store := stores.BuildStore(store_id, store_name, store_place, store_category)
			rep.Create(ctx, tran, store)
			// 登録した飲食店カテゴリを削除する
			result := rep.DeleteById(ctx, tran, store)
			Expect(result).To(BeNil())
		})
	})
})
