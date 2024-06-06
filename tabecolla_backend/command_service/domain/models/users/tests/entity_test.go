package users_test

import (
	"commandservice/domain/models/users"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Userエンティティ", Ordered, Label("User構造体の生成"), func() {
	// 前処理
	BeforeAll(func() {
	})
	var _ = Describe("Userエンティティ", Ordered, Label("User構造体の生成"), func() {
		Context("インスタンス生成", Label("Create User"), func() {
			It("新しいユーザのインスタンス生成", Label("NewUser"), func() {
				user_id, _ := users.NewUserId("a-saneyoshi")
				user_name, _ := users.NewUserName("實好聖佳")
				user_mail, _ := users.NewUserMail("saneyoshi@dummymail.com")
				user_password, _ := users.NewUserPassword("Valid123!")
				user_enable, _ := users.NewUserEnable(true)
				user, _ := users.NewUser(user_id, user_name, user_mail, user_password, user_enable)
				Expect(user.UserId().Value()).To(Equal("a-saneyoshi"))
				Expect(user.UserName().Value()).To(Equal("實好聖佳"))
				Expect(user.UserMail().Value()).To(Equal("saneyoshi@dummymail.com"))
				Expect(user.UserPassword().Value()).To(Equal("Valid123!"))
				Expect(user.UserEnable().Value()).To(Equal(true))
			})
			It("ユーザのインスタンスを構築する", Label("BuildUser"), func() {
				user_id, _ := users.NewUserId("a-saneyoshi")
				user_mail, _ := users.NewUserMail("saneyoshi@dummymail.com")
				user_password, _ := users.NewUserPassword("Valid123!")
				user_enable, _ := users.NewUserEnable(true)
				user := users.BuildUser(user_id, nil, user_mail, user_password, user_enable)
				Expect(user.UserId().Value()).To(Equal("a-saneyoshi"))
				Expect(user.UserName()).To(BeNil())
				Expect(user.UserMail().Value()).To(Equal("saneyoshi@dummymail.com"))
				Expect(user.UserPassword().Value()).To(Equal("Valid123!"))
				Expect(user.UserEnable().Value()).To(Equal(true))
			})
		})
	})
})

var _ = Describe("Userエンティティ", Ordered, Label("Userの同一性検証"), func() {
	var user *users.User

	// 前処理
	BeforeAll(func() {
		user_id, _ := users.NewUserId("sa-saneyoshi")
		user_name, _ := users.NewUserName("實好聖佳")
		user_mail, _ := users.NewUserMail("saneyoshi@dummymail.com")
		user_password, _ := users.NewUserPassword("Valid123!")
		user_enable, _ := users.NewUserEnable(true)
		user, _ = users.NewUser(user_id, user_name, user_mail, user_password, user_enable)
	})

	// エラーの検証
	Context("エラーの検証", func() {
		It("比較対象がnil", Label("nil検証"), func() {
			By("nilを指定し,DomainErrorを返すことを検証する")
			_, err := user.Equals(nil)
			Expect(err).To(Equal(errs.NewDomainError("引数でnilが指定されました。")))
		})
	})

	// 比較結果の検証
	Context("比較結果の検証", func() {
		It("異なるユーザID", Label("false検証"), func() {
			user_id, _ := users.NewUserId("a-saneyoshi")
			user_name, _ := users.NewUserName("實好聖佳")
			user_mail, _ := users.NewUserMail("saneyoshi@dummymail.com")
			user_password, _ := users.NewUserPassword("Valid123!")
			user_enable, _ := users.NewUserEnable(true)
			u, _ := users.NewUser(user_id, user_name, user_mail, user_password, user_enable)
			By("異なるUserを指定し,falseを返すことを検証する")
			result, _ := user.Equals(u)
			Expect(result).To(Equal(false))
		})
		It("同一のユーザID", Label("trueの検証"), func() {
			u := users.BuildUser(
				user.UserId(),
				user.UserName(),
				user.UserMail(),
				user.UserPassword(),
				user.UserEnable())
			By("同一のStoreを指定し,trueを返すことを検証する")
			result, _ := user.Equals(u)
			Expect(result).To(Equal(true))
		})
	})
})
