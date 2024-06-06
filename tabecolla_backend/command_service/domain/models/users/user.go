package users

import (
	"commandservice/errs"
)

// ユーザを表すEntity
type User struct {
	user_id       *UserId       // ユーザId
	user_name     *UserName     // ユーザ名
	user_mail     *UserMail     // メールアドレス
	user_password *UserPassword // ユーザパスワード
	user_enable   *UserEnable   // ユーザの有効/無効
}

// ゲッター
func (ins *User) UserId() *UserId {
	return ins.user_id
}
func (ins *User) UserName() *UserName {
	return ins.user_name
}
func (ins *User) UserMail() *UserMail {
	return ins.user_mail
}
func (ins *User) UserPassword() *UserPassword {
	return ins.user_password
}
func (ins *User) UserEnable() *UserEnable {
	return ins.user_enable
}

// 値の変更
func (ins *User) ChangeUserId(user_id *UserId) {
	ins.user_id = user_id
}
func (ins *User) ChangeUserName(user_name *UserName) {
	ins.user_name = user_name
}
func (ins *User) ChangeUserMail(user_mail *UserMail) {
	ins.user_mail = user_mail
}
func (ins *User) ChangeUserPassword(user_password *UserPassword) {
	ins.user_password = user_password
}
func (ins *User) ChangeUserEnable(user_enable *UserEnable) {
	ins.user_enable = user_enable
}

// 同一性検証メソッド
func (ins *User) Equals(obj *User) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.user_id.Equals(obj.UserId())
	return result, nil
}

// コンストラクタ
func NewUser(user_id *UserId, user_name *UserName, user_mail *UserMail, user_password *UserPassword, user_enable *UserEnable) (*User, *errs.DomainError) {
	// ユーザエンティティのインスタンスを生成して返す
	return &User{
		user_id:       user_id,
		user_name:     user_name,
		user_mail:     user_mail,
		user_password: user_password,
		user_enable:   user_enable,
	}, nil
}

// ユーザエンティティの再構築
func BuildUser(user_id *UserId, user_name *UserName, user_mail *UserMail, user_password *UserPassword, user_enable *UserEnable) *User {
	user := User{ // ユーザエンティティのインスタンスを生成して返す
		user_id:       user_id,
		user_name:     user_name,
		user_mail:     user_mail,
		user_password: user_password,
		user_enable:   user_enable,
	}
	return &user
}
