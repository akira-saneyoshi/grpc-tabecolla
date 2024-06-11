package loginUsers_test

import (
	"commandservice/domain/models/loginUsers"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginUserエンティティ", Ordered, Label("LoginUser構造体の生成"), func() {
	// 前処理
	BeforeAll(func() {
	})
	var _ = Describe("LoginUserエンティティ", Ordered, Label("LoginUser構造体の生成"), func() {
		Context("インスタンス生成", Label("Create Login User"), func() {
			It("新しいユーザのインスタンス生成", Label("NewLoginUser"), func() {
				login_user_id, _ := loginUsers.NewLoginUserId("a-saneyoshi")
				login_user_name, _ := loginUsers.NewLoginUserName("實好聖佳")
				login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummymail.com")
				login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
				login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
				login_user, _ := loginUsers.NewLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
				Expect(login_user.LoginUserId().Value()).To(Equal("a-saneyoshi"))
				Expect(login_user.LoginUserName().Value()).To(Equal("實好聖佳"))
				Expect(login_user.LoginUserMail().Value()).To(Equal("saneyoshi@dummymail.com"))
				Expect(login_user.LoginUserPassword().Value()).To(Equal("Valid123!"))
				Expect(login_user.LoginUserEnable().Value()).To(Equal(true))
			})
			It("ユーザのインスタンスを構築する", Label("BuildUser"), func() {
				login_user_id, _ := loginUsers.NewLoginUserId("a-saneyoshi")
				login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummymail.com")
				login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
				login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
				login_user := loginUsers.BuildLoginUser(login_user_id, nil, login_user_mail, login_user_password, login_user_enable)
				Expect(login_user.LoginUserId().Value()).To(Equal("a-saneyoshi"))
				Expect(login_user.LoginUserName()).To(BeNil())
				Expect(login_user.LoginUserMail().Value()).To(Equal("saneyoshi@dummymail.com"))
				Expect(login_user.LoginUserPassword().Value()).To(Equal("Valid123!"))
				Expect(login_user.LoginUserEnable().Value()).To(Equal(true))
			})
		})
	})
})

var _ = Describe("LoginUserエンティティ", Ordered, Label("LoginUserの同一性検証"), func() {
	var login_user *loginUsers.LoginUser

	// 前処理
	BeforeAll(func() {
		login_user_id, _ := loginUsers.NewLoginUserId("sa-saneyoshi")
		login_user_name, _ := loginUsers.NewLoginUserName("實好聖佳")
		login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummymail.com")
		login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
		login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
		login_user, _ = loginUsers.NewLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
	})

	// エラーの検証
	Context("エラーの検証", func() {
		It("比較対象がnil", Label("nil検証"), func() {
			By("nilを指定し,DomainErrorを返すことを検証する")
			_, err := login_user.Equals(nil)
			Expect(err).To(Equal(errs.NewDomainError("引数でnilが指定されました。")))
		})
	})

	// 比較結果の検証
	Context("比較結果の検証", func() {
		It("異なるユーザID", Label("false検証"), func() {
			login_user_id, _ := loginUsers.NewLoginUserId("a-saneyoshi")
			login_user_name, _ := loginUsers.NewLoginUserName("實好聖佳")
			login_user_mail, _ := loginUsers.NewLoginUserMail("saneyoshi@dummymail.com")
			login_user_password, _ := loginUsers.NewLoginUserPassword("Valid123!")
			login_user_enable, _ := loginUsers.NewLoginUserEnable(true)
			u, _ := loginUsers.NewLoginUser(login_user_id, login_user_name, login_user_mail, login_user_password, login_user_enable)
			By("異なるLoginUserを指定し,falseを返すことを検証する")
			result, _ := login_user.Equals(u)
			Expect(result).To(Equal(false))
		})
		It("同一のユーザID", Label("trueの検証"), func() {
			u := loginUsers.BuildLoginUser(
				login_user.LoginUserId(),
				login_user.LoginUserName(),
				login_user.LoginUserMail(),
				login_user.LoginUserPassword(),
				login_user.LoginUserEnable())
			By("同一のLoginUserを指定し,trueを返すことを検証する")
			result, _ := login_user.Equals(u)
			Expect(result).To(Equal(true))
		})
	})
})
