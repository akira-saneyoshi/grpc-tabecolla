package repository_test

import (
	"commandservice/domain/models/storeCategories"
	"commandservice/errs"
	"commandservice/infrastructure/sqlboiler/repository"
	"context"
	"database/sql"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ = Describe("storeCategoryRepositorySQLBoiler構造体", Ordered, Label("StoreCategoryRepositoryインターフェイスメソッドのテスト"), func() {
	var rep storeCategories.StoreCategoryRepository
	var ctx context.Context
	var tran *sql.Tx
	// 前処理
	BeforeAll(func() {
		// リポジトリの生成
		rep = repository.NewStoreCategoryRepositorySQLBoiler()
	})

	// テスト毎の前処理
	BeforeEach(func() {
		ctx = context.Background()       // Contextの生成
		tran, _ = boil.BeginTx(ctx, nil) // トランザクションの開始
	})

	// テスト毎の後処理
	AfterEach(func() {
		tran.Rollback() // トランザクションのロールバックc
	})

	// Exists()メソッドのテスト
	Context("同名の飲食店カテゴリが存在確認結果を返す", Label("Exists"), func() {
		It("存在しない飲食店カテゴリの場合nilが返る", func() {
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category, _ := storeCategories.NewStoreCategory(store_category_name)
			result := rep.Exists(ctx, tran, store_category)
			Expect(result).To(BeNil())
		})
		It("存在する飲食店カテゴリ名の場合、errs.CRUDErrorが返る", func() {
			store_category_name, _ := storeCategories.NewStoreCategoryName("和食")
			store_category, _ := storeCategories.NewStoreCategory(store_category_name)
			result := rep.Exists(ctx, tran, store_category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", store_category.StoreCategoryName().Value()))))
		})
	})
	// Create()メソッドのテスト
	Context("新しい飲食店カテゴリを永続化する", Label("Create"), func() {
		It("飲食店カテゴリが登録成功し、nilが返る", func() {
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category, _ := storeCategories.NewStoreCategory(store_category_name)
			result := rep.Create(ctx, tran, store_category)
			Expect(result).To(BeNil())
		})
		It("obj_idが同じ飲食店カテゴリを追加すると、errs.CRUDErrorが返る", func() {
			store_category_id, _ := storeCategories.NewStoreCategoryUid("aaaa1111-22bb-33cc-44dd-555555eeeeee")
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category := storeCategories.BuildStoreCategory(store_category_id, store_category_name)
			result := rep.Create(ctx, tran, store_category)
			Expect(result).To(Equal(errs.NewCRUDError("一意制約違反です。")))
		})
	})
	// UpdateById()メソッドのテスト
	Context("飲食店カテゴリを変更する", Label("UpdateById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			store_category_id, _ := storeCategories.NewStoreCategoryUid("b1524011-b6af-417e-8bf2-f449dd58b5c1")
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category := storeCategories.BuildStoreCategory(store_category_id, store_category_name)
			result := rep.UpdateById(ctx, tran, store_category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("飲食店カテゴリ番号:%sは存在しないため、更新できませんでした。", store_category.StoreCategoryUid().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			store_category_id, _ := storeCategories.NewStoreCategoryUid("bbbb1111-22aa-dd33-cc44-eeeeee666666")
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category := storeCategories.BuildStoreCategory(store_category_id, store_category_name)
			result := rep.UpdateById(ctx, tran, store_category)
			Expect(result).To(BeNil())
		})
	})
	// DeleteById()メソッドのテスト
	Context("飲食店カテゴリを削除する", Label("DeleteById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			store_category_id, _ := storeCategories.NewStoreCategoryUid("b1524011-b6af-417e-8bf2-f449dd58b5c1")
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category := storeCategories.BuildStoreCategory(store_category_id, store_category_name)
			result := rep.DeleteById(ctx, tran, store_category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("飲食店カテゴリ番号:%sは存在しないため、削除できませんでした。",
				store_category.StoreCategoryUid().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 削除対象のカテゴリを登録する
			store_category_name, _ := storeCategories.NewStoreCategoryName("焼き鳥")
			store_category, _ := storeCategories.NewStoreCategory(store_category_name)
			rep.Create(ctx, tran, store_category)
			// 登録したカテゴリを削除する
			result := rep.DeleteById(ctx, tran, store_category)
			Expect(result).To(BeNil())
		})
	})
})
