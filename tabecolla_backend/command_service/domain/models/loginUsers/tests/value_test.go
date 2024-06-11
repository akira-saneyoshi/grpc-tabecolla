package loginUsers_test

import (
	"commandservice/domain/models/loginUsers"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginUserエンティティを構成する値オブジェクト", Ordered, Label("LoginUserId構造体の生成"), func() {
	var empty_str *errs.DomainError   // 空文字列　長さ20に違反する
	var length_over *errs.DomainError // 20文字より大きい文字列　長さ20に違反する
	var login_user_id *loginUsers.LoginUserId         // 20文字以内の文字列を指定する

	BeforeAll(func() {
		_, empty_str = loginUsers.NewLoginUserId("")
		_, length_over = loginUsers.NewLoginUserId("aaaaaaaaaabbbbbbbbbbc")
		login_user_id, _ = loginUsers.NewLoginUserId("a-saneyoshi")
	})

	Context("文字数の検証", Label("文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("ユーザIDの長さは3文字以上、20文字以内です。")))
		})
		It("20文字より大きい文字列の場合、errors.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("ユーザIDの長さは3文字以上、20文字以内です。")))
		})
	})

	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("3文字以上20文字以内場合, LoginUserIdが返る", func() {
			Expect(login_user_id.Value()).To(Equal("a-saneyoshi"))
		})
	})
})

var _ = Describe("LoginUserエンティティを構成する値オブジェクト", Ordered, Label("LoginUserName構造体の生成"), func() {
	var empty_str *errs.DomainError   // 空文字列　長さ3文字以上、20文字以内に違反する
	var length_over *errs.DomainError // 20文字より大きい文字列　長さ3文字以上、20文字以内に違反する
	var login_user_name *loginUsers.LoginUserName     // 20文字以内の文字列を指定する

	BeforeAll(func() {
		_, empty_str = loginUsers.NewLoginUserName("")
		_, length_over = loginUsers.NewLoginUserName("aaaaaaaaaabbbbbbbbbbc")
		login_user_name, _ = loginUsers.NewLoginUserName("實好聖佳")
	})

	Context("文字数の検証", Label("無効な文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("ユーザ名の長さは3文字以上、20文字以内です。")))
		})
		It("20文字より大きいの場合,errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("ユーザ名の長さは3文字以上、20文字以内です。")))
		})
	})

	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("3文字以上20文字以内場合, LoginUserNameが返る", func() {
			Expect(login_user_name.Value()).To(Equal("實好聖佳"))
		})
	})
})

var _ = Describe("LoginUserエンティティを構成する値オブジェクト", Ordered, Label("LoginUserMail構造体の生成"), func() {
	var empty_str *errs.DomainError      // 空文字列　長さ7文字以上、50文字以内に違反する
	var length_over *errs.DomainError    // 50文字より大きい文字列　長さ7文字以上、50文字以内に違反する
	var invalid_format *errs.DomainError // 無効な形式のメールアドレス
	var invalid_domain *errs.DomainError // ドメイン部分が無効なメールアドレス
	var login_user_email *loginUsers.LoginUserMail       // 有効なメールアドレス

	BeforeAll(func() {
		_, empty_str = loginUsers.NewLoginUserMail("")
		_, length_over = loginUsers.NewLoginUserMail("aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeee@gmail.com")
		_, invalid_format = loginUsers.NewLoginUserMail("invalid-email")
		_, invalid_domain = loginUsers.NewLoginUserMail("aaa@.com")
		login_user_email, _ = loginUsers.NewLoginUserMail("saneyoshi@dummymail.com")
	})

	Context("文字数の検証", Label("無効な文字数"), func() {
		It("空文字列の場合、errors.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("ユーザのメールアドレスの長さは7文字以上、50文字以内です。")))
		})
		It("50文字より大きい場合,errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("ユーザのメールアドレスの長さは7文字以上、50文字以内です。")))
		})
	})

	Context("形式の検証", Label("無効な形式"), func() {
		It("無効な形式の場合、errors.DomainErrorが返る", func() {
			Expect(invalid_format).To(Equal(errs.NewDomainError("ユーザのメールアドレスの形式が正しくありません。")))
		})
	})

	Context("ドメイン部分の検証", Label("無効なドメイン"), func() {
		It("ドメイン部分が '.' で始まっている場合、errors.DomainErrorが返る", func() {
			Expect(invalid_domain).To(Equal(errs.NewDomainError("ユーザのメールアドレスのドメイン部分が '@' で始まっています。")))
		})
	})

	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("7文字以上50文字以内の場合, LoginUserMailが返る", func() {
			Expect(login_user_email.Value()).To(Equal("saneyoshi@dummymail.com"))
		})
	})
})

var _ = Describe("LoginUserエンティティを構成する値オブジェクト", Ordered, Label("LoginUserPassword構造体の生成"), func() {
	var short_password *errs.DomainError   // パスワードが8文字未満
	var long_password *errs.DomainError    // パスワードが15文字を超える
	var invalid_password *errs.DomainError // パスワードに英数字や記号が含まれていない
	var valid_password *loginUsers.LoginUserPassword // 有効なパスワード

	BeforeAll(func() {
		_, short_password = loginUsers.NewLoginUserPassword("short")
		_, long_password = loginUsers.NewLoginUserPassword("thispasswordiswaytoolong")
		_, invalid_password = loginUsers.NewLoginUserPassword("無効なパスワード")
		valid_password, _ = loginUsers.NewLoginUserPassword("Valid123!")
	})

	Context("パスワードの長さの検証", Label("無効な長さ"), func() {
		It("8文字未満の場合、errors.DomainErrorが返る", func() {
			Expect(short_password).To(Equal(errs.NewDomainError("ユーザのパスワードの長さは8文字以上、15文字以内です。")))
		})
		It("15文字を超える場合、errors.DomainErrorが返る", func() {
			Expect(long_password).To(Equal(errs.NewDomainError("ユーザのパスワードの長さは8文字以上、15文字以内です。")))
		})
	})

	Context("パスワードの内容の検証", Label("無効な内容"), func() {
		It("英数字や記号を含まない場合、errors.DomainErrorが返る", func() {
			Expect(invalid_password).To(Equal(errs.NewDomainError("パスワードには半角の英数字や記号を含める必要があります。")))
		})
	})

	Context("有効なパスワードの検証", Label("有効なパスワード"), func() {
		It("8文字以上15文字以内で、英数字や記号を含む場合、LoginUserPasswordが返る", func() {
			Expect(valid_password.Value()).To(Equal("Valid123!"))
		})
	})
})

var _ = Describe("LoginUserエンティティを構成する値オブジェクト", Ordered, Label("LoginUserEnable構造体の生成"), func() {
	var login_user_enabled *loginUsers.LoginUserEnable  // 有効なユーザ
	var login_user_disabled *loginUsers.LoginUserEnable // 無効なユーザ

	BeforeAll(func() {
		login_user_enabled, _ = loginUsers.NewLoginUserEnable(true)
		login_user_disabled, _ = loginUsers.NewLoginUserEnable(false)
	})

	Context("ユーザの有効か無効かどうかの検証", Label("ユーザの有効性の状態"), func() {
		It("ユーザが有効の場合、trueが返る", func() {
			Expect(login_user_enabled.Value()).To(BeTrue())
		})
		It("ユーザが無効の場合、falseが返る", func() {
			Expect(login_user_disabled.Value()).To(BeFalse())
		})
	})
})
