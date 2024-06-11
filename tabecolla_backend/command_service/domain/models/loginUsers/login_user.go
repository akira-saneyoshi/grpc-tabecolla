package loginUsers

import (
	"commandservice/errs"
)

// ユーザを表すEntity
type LoginUser struct {
	login_user_id       *LoginUserId       // ユーザId
	login_user_name     *LoginUserName     // ユーザ名
	login_user_mail     *LoginUserMail     // メールアドレス
	login_user_password *LoginUserPassword // ユーザパスワード
	login_user_enable   *LoginUserEnable   // ユーザの有効/無効
}

// ゲッター
func (ins *LoginUser) LoginUserId() *LoginUserId {
	return ins.login_user_id
}
func (ins *LoginUser) LoginUserName() *LoginUserName {
	return ins.login_user_name
}
func (ins *LoginUser) LoginUserMail() *LoginUserMail {
	return ins.login_user_mail
}
func (ins *LoginUser) LoginUserPassword() *LoginUserPassword {
	return ins.login_user_password
}
func (ins *LoginUser) LoginUserEnable() *LoginUserEnable {
	return ins.login_user_enable
}

// 値の変更
func (ins *LoginUser) ChangeLoginUserId(login_user_id *LoginUserId) {
	ins.login_user_id = login_user_id
}
func (ins *LoginUser) ChangeLoginUserName(login_user_name *LoginUserName) {
	ins.login_user_name = login_user_name
}
func (ins *LoginUser) ChangeLoginUserMail(login_user_mail *LoginUserMail) {
	ins.login_user_mail = login_user_mail
}
func (ins *LoginUser) ChangeLoginUserPassword(login_user_password *LoginUserPassword) {
	ins.login_user_password = login_user_password
}
func (ins *LoginUser) ChangeLoginUserEnable(login_user_enable *LoginUserEnable) {
	ins.login_user_enable = login_user_enable
}

// 同一性検証メソッド
func (ins *LoginUser) Equals(obj *LoginUser) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.login_user_id.Equals(obj.LoginUserId())
	return result, nil
}

// コンストラクタ
func NewLoginUser(login_user_id *LoginUserId, login_user_name *LoginUserName, login_user_mail *LoginUserMail, login_user_password *LoginUserPassword, login_user_enable *LoginUserEnable) (*LoginUser, *errs.DomainError) {
	// ユーザエンティティのインスタンスを生成して返す
	return &LoginUser{
		login_user_id:       login_user_id,
		login_user_name:     login_user_name,
		login_user_mail:     login_user_mail,
		login_user_password: login_user_password,
		login_user_enable:   login_user_enable,
	}, nil
}

// ユーザエンティティの再構築
func BuildLoginUser(login_user_id *LoginUserId, login_user_name *LoginUserName, login_user_mail *LoginUserMail, login_user_password *LoginUserPassword, login_user_enable *LoginUserEnable) *LoginUser {
	return &LoginUser{ // ユーザエンティティのインスタンスを生成して返す
		login_user_id:       login_user_id,
		login_user_name:     login_user_name,
		login_user_mail:     login_user_mail,
		login_user_password: login_user_password,
		login_user_enable:   login_user_enable,
	}
}
