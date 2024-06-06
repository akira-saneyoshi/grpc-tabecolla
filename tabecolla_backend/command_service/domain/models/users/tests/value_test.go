package users_test

import (
	"commandservice/domain/models/users"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Userエンティティを構成する値オブジェクト", Ordered, Label("UserId構造体の生成"), func() {
	var empty_str *errs.DomainError   // 空文字列　長さ20に違反する
	var length_over *errs.DomainError // 20文字より大きい文字列　長さ20に違反する
	var user_id *users.UserId         // 20文字以内の文字列を指定する

	BeforeAll(func() {
		_, empty_str = users.NewUserId("")
		_, length_over = users.NewUserId("aaaaaaaaaabbbbbbbbbbc")
		user_id, _ = users.NewUserId("a-saneyoshi")
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
		It("3文字以上20文字以内場合,UserIdが返る", func() {
			Expect(user_id.Value()).To(Equal("a-saneyoshi"))
		})
	})
})

var _ = Describe("Userエンティティを構成する値オブジェクト", Ordered, Label("UserName構造体の生成"), func() {
	var empty_str *errs.DomainError   // 空文字列　長さ3文字以上、20文字以内に違反する
	var length_over *errs.DomainError // 20文字より大きい文字列　長さ3文字以上、20文字以内に違反する
	var user_name *users.UserName     // 20文字以内の文字列を指定する

	BeforeAll(func() {
		_, empty_str = users.NewUserName("")
		_, length_over = users.NewUserName("aaaaaaaaaabbbbbbbbbbc")
		user_name, _ = users.NewUserName("實好聖佳")
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
		It("3文字以上20文字以内場合,UserNameが返る", func() {
			Expect(user_name.Value()).To(Equal("實好聖佳"))
		})
	})
})

var _ = Describe("Userエンティティを構成する値オブジェクト", Ordered, Label("UserEmail構造体の生成"), func() {
	var empty_str *errs.DomainError      // 空文字列　長さ7文字以上、50文字以内に違反する
	var length_over *errs.DomainError    // 50文字より大きい文字列　長さ7文字以上、50文字以内に違反する
	var invalid_format *errs.DomainError // 無効な形式のメールアドレス
	var invalid_domain *errs.DomainError // ドメイン部分が無効なメールアドレス
	var user_email *users.UserMail       // 有効なメールアドレス

	BeforeAll(func() {
		_, empty_str = users.NewUserMail("")
		_, length_over = users.NewUserMail("aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeee@gmail.com")
		_, invalid_format = users.NewUserMail("invalid-email")
		_, invalid_domain = users.NewUserMail("aaa@.com")
		user_email, _ = users.NewUserMail("saneyoshi@dummymail.com")
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
		It("7文字以上50文字以内の場合,UserEmailが返る", func() {
			Expect(user_email.Value()).To(Equal("saneyoshi@dummymail.com"))
		})
	})
})

var _ = Describe("Userエンティティを構成する値オブジェクト", Ordered, Label("UserPassword構造体の生成"), func() {
	var short_password *errs.DomainError   // パスワードが8文字未満
	var long_password *errs.DomainError    // パスワードが15文字を超える
	var invalid_password *errs.DomainError // パスワードに英数字や記号が含まれていない
	var valid_password *users.UserPassword // 有効なパスワード

	BeforeAll(func() {
		_, short_password = users.NewUserPassword("short")
		_, long_password = users.NewUserPassword("thispasswordiswaytoolong")
		_, invalid_password = users.NewUserPassword("無効なパスワード")
		valid_password, _ = users.NewUserPassword("Valid123!")
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
		It("8文字以上15文字以内で、英数字や記号を含む場合、UserPasswordが返る", func() {
			Expect(valid_password.Value()).To(Equal("Valid123!"))
		})
	})
})

var _ = Describe("Userエンティティを構成する値オブジェクト", Ordered, Label("UserEnable構造体の生成"), func() {
	var user_enabled *users.UserEnable  // 有効なユーザ
	var user_disabled *users.UserEnable // 無効なユーザ

	BeforeAll(func() {
		user_enabled, _ = users.NewUserEnable(true)
		user_disabled, _ = users.NewUserEnable(false)
	})

	Context("ユーザの有効か無効かどうかの検証", Label("ユーザの有効性の状態"), func() {
		It("ユーザが有効の場合、trueが返る", func() {
			Expect(user_enabled.Value()).To(BeTrue())
		})
		It("ユーザが無効の場合、falseが返る", func() {
			Expect(user_disabled.Value()).To(BeFalse())
		})
	})
})
