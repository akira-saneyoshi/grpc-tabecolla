package users

import (
	"commandservice/errors"

	"github.com/google/uuid"
)

// 飲食店を表すEntity
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
func (ins *User) Equals(obj *User) (bool, *errors.DomainError) {
	if obj == nil {
		return false, errors.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.user_id.Equlas(obj.UserId())
	return result, nil
}

// コンストラクタ
func NewUser(user_name *UserName, user_mail *UserMail, user_password *UserPassword, user_enable *UserEnable) (*User, *errors.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errors.NewDomainError(err.Error())
	} else {
		// ユーザID用値オブジェクトを生成する
		if user_id, err := NeeUserId(uid.String()); err != nil {
			return nil, err
		} else {
			// ユーザエンティティのインスタンスを生成して返す
			return &User{
				user_id:       user_id,
				user_name:     user_name,
				user_mail:     user_mail,
				user_password: user_password,
				user_enable:   user_enable,
			}, nil
		}
	}
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
