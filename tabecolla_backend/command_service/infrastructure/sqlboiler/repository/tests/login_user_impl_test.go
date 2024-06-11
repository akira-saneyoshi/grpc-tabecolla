package repository_test

import (
	"commandservice/domain/models/loginUsers"
	"commandservice/errs"
	"commandservice/infrastructure/sqlboiler/repository"
	"context"
	"database/sql"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ = Describe("loginUserRepositorySQLBoiler構造体", Ordered, Label("LoginUserRepositoryインターフェイスメソッドのテスト"), func() {
	var rep loginUsers.LoginUserRepository
	var ctx context.Context
	var tran *sql.Tx
	// 前処理
	BeforeAll(func() {
		// リポジトリの生成
		rep = repository.NewLoginUserRepositorySQLBoiler()
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
	Context("同名のユーザが存在確認結果を返す", Label("Exists"), func() {
		It("存在しないユーザIDの場合nilが返る", func() {
			login_user_id, _ := loginUsers.NewLoginUserId("yamada")
			login_user_name, _ := loginUsers.NewLoginUserName("山田桃太郎")
			login_user_mail, _ := loginUsers.NewLoginUserMail("yamada@dummy.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			login_user := loginUsers.BuildLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			result := rep.Exists(ctx, tran, login_user)
			Expect(result).To(BeNil())
		})
		It("存在するユーザIDの場合、errs.CRUDErrorが返る", func() {
			login_user_id, _ := loginUsers.NewLoginUserId("tanaka")
			login_user := loginUsers.BuildLoginUser(login_user_id, nil, nil, nil, nil)
			result := rep.Exists(ctx, tran, login_user)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", login_user_id.Value()))))
		})
		// TODO: 別の層でエラーロジックの実装可否を検討する。
		// It("存在しないメールアドレスの場合nilが返る", func() {
		// 	login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummy.com")
		// 	login_user, _ := loginUsers.NewLoginUser(nil, nil, login_user_mail, nil, nil)
		// 	result := rep.Exists(ctx, tran, login_user)
		// 	Expect(result).To(BeNil())
		// })
		// It("存在するメールアドレスの場合、errs.CRUDErrorが返る", func() {
		// 	login_user_mail, _ := loginUsers.NewLoginUserMail("tanaka@dummy.com")
		// 	login_user, _ := loginUsers.NewLoginUser(login_user_mail)
		// 	result := rep.Exists(ctx, tran, login_user)
		// 	Expect(result).To(Equal(errs.NewCRUDError(
		// 		fmt.Sprintf("%sは既に登録されています。", login_user.NewLoginUserMail().Value()))))
		// })
	})
	// Create()メソッドのテスト
	Context("新しいユーザを永続化する", Label("Create"), func() {
		It("ユーザが登録成功し、nilが返る", func() {
			login_user_id, _ := loginUsers.NewLoginUserId("saneyoshi")
			login_user_name, _ := loginUsers.NewLoginUserName("實好聖佳")
			login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummy.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			login_user, _ := loginUsers.NewLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			result := rep.Create(ctx, tran, login_user)
			Expect(result).To(BeNil())
		})
		// TODO: 別の層でエラーロジックの実装可否を検討する。
		// It("user_idが同じユーザを追加すると、errs.CRUDErrorが返る", func() {
		// 	login_user_id, _ := loginUsers.NewLoginUserId("saneyoshi")
		// 	login_user_name, _ := loginUsers.NewLoginUserName("實好聖佳")
		// 	login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummy.com")
		// 	login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
		// 	login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
		// 	login_user := loginUsers.BuildLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
		// 	result := rep.Create(ctx, tran, login_user)
		// 	Expect(result).To(Equal(errs.NewCRUDError("一意制約違反です。")))
		// })
		// It("user_mailが同じユーザを追加すると、errs.CRUDErrorが返る", func() {
		// 	login_user_id, _ := loginUsers.NewLoginUserId("tanaka")
		// 	login_user_name, _ := loginUsers.NewLoginUserName("田中角栄")
		// 	login_user_mail, _ := loginUsers.NewLoginUserMail("tanaka@dummy.com")
		// 	login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
		// 	login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
		// 	login_user, _ := loginUsers.NewLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
		// 	result := rep.Create(ctx, tran, login_user)
		// 	Expect(result).To(Equal(errs.NewCRUDError("一意制約違反です。")))
		// })
	})
	// UpdateById()メソッドのテスト
	Context("ユーザを変更する", Label("UpdateById"), func() {
		It("存在しないuser_idを指定すると、errs.CRUDErrorが返る", func() {
			login_user_id, _ := loginUsers.NewLoginUserId("yamada")
			login_user_name, _ := loginUsers.NewLoginUserName("山田桃太郎")
			login_user_mail, _ := loginUsers.NewLoginUserMail("yamada@dummy.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			login_user := loginUsers.BuildLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			result := rep.UpdateById(ctx, tran, login_user)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("ユーザID:%sは存在しないため、更新できませんでした。", login_user.LoginUserId().Value()))))
		})
		It("存在するuser_idを指定すると、nilが返る", func() {
			login_user_id, _ := loginUsers.NewLoginUserId("tanaka")
			login_user_name, _ := loginUsers.NewLoginUserName("田中角栄")
			login_user_mail, _ := loginUsers.NewLoginUserMail("tanaka@dummy.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			login_user := loginUsers.BuildLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			result := rep.UpdateById(ctx, tran, login_user)
			Expect(result).To(BeNil())
		})
	})
	// DeleteById()メソッドのテスト
	Context("ユーザを削除する", Label("DeleteById"), func() {
		It("存在しないuser_idを指定すると、errs.CRUDErrorが返る", func() {
			login_user_id, _ := loginUsers.NewLoginUserId("yamada")
			login_user_name, _ := loginUsers.NewLoginUserName("山田桃太郎")
			login_user_mail, _ := loginUsers.NewLoginUserMail("yamada@dummy.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			login_user := loginUsers.BuildLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			result := rep.DeleteById(ctx, tran, login_user)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("ユーザID:%sは存在しないため、削除できませんでした。",
					login_user.LoginUserId().Value()))))
		})
		It("存在するuser_idを指定すると、nilが返る", func() {
			// 削除対象のユーザを登録する
			login_user_id, _ := loginUsers.NewLoginUserId("tanaka")
			login_user_name, _ := loginUsers.NewLoginUserName("田中角栄")
			login_user_mail, _ := loginUsers.NewLoginUserMail("tanaka@dummy.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			login_user, _ := loginUsers.NewLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			rep.Create(ctx, tran, login_user)
			// 登録したユーザを削除する
			result := rep.DeleteById(ctx, tran, login_user)
			Expect(result).To(BeNil())
		})
	})
})
