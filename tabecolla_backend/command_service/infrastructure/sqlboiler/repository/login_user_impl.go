package repository

import (
	"commandservice/domain/models/loginUsers"
	"commandservice/errs"
	"commandservice/infrastructure/sqlboiler/handler"
	"commandservice/infrastructure/sqlboiler/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UserRepositoryインターフェイスの実装
type loginUserRepositorySQLBoiler struct{}

// コンストラクタ
func NewLoginUserRepositorySQLBoiler() loginUsers.LoginUserRepository {
	// フック関数の登録
	models.AddLoginUserHook(boil.AfterInsertHook, LoginUserAfterInsertHook)
	models.AddLoginUserHook(boil.AfterUpdateHook, LoginUserAfterUpdateHook)
	models.AddLoginUserHook(boil.AfterDeleteHook, LoginUserAfterDeleteHook)
	return &loginUserRepositorySQLBoiler{}
}

// 同名のユーザIdが存在確認結果を返す
func (rep *loginUserRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, login_user *loginUsers.LoginUser) error {
	// レコードの存在確認条件を作成する
	// TODO:condition変数と同様に、メールアドレス被りの検証も実装する。 44~56行目のfunc関数を移管する。
	condition := models.LoginUserWhere.UserID.EQ(login_user.LoginUserId().Value())
	// レコードの存在を確認するクエリを実行する
	if exists, err := models.LoginUsers(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists { // 同じユーザIdは存在していない?
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", login_user.LoginUserId().Value()))
	}
}

// 同名のメールアドレスが存在確認結果を返す
// func (rep *loginUserRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, login_user *loginUsers.LoginUser) error {
// 	// レコードの存在確認条件を作成する
// 	condition := models.LoginUserWhere.UserMail.EQ(login_user.LoginUserMail().Value())
// 	// レコードの存在を確認するクエリを実行する
// 	if exists, err := models.LoginUsers(condition).Exists(ctx, tran); err != nil {
// 		return handler.DBErrHandler(err)
// 	} else if !exists { // 同じメールアドレスは存在していない?
// 		return nil
// 	} else {
// 		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", login_user.LoginUserMail().Value()))
// 	}
// }

// 新しいユーザを永続化する
func (rep *loginUserRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, login_user *loginUsers.LoginUser) error {
	// SqlBoilerのモデルを生成する
	new_login_user := models.LoginUser{
		ID:    0,
		UserID: login_user.LoginUserId().Value(),
		UserName:  login_user.LoginUserName().Value(),
		UserMail:  login_user.LoginUserMail().Value(),
		UserPassword:  login_user.LoginUserPassword().Value(),
		UserEnable:  true,
	}
	// ユーザを永続化する
	if err := new_login_user.Insert(ctx, tran, boil.Whitelist("user_id", "user_name", "user_mail", "user_password", "user_enable")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// ユーザを変更する
func (rep *loginUserRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, login_user *loginUsers.LoginUser) error {
	// 更新対象を取得する
	up_model, err := models.LoginUsers(qm.Where("user_id = ?", login_user.LoginUserId().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("ユーザID:%sは存在しないため、更新できませんでした。", login_user.LoginUserId().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 取得したモデルの値を変更する
	up_model.UserID = login_user.LoginUserId().Value()
	up_model.UserName = login_user.LoginUserName().Value()
	up_model.UserMail = login_user.LoginUserMail().Value()
	up_model.UserPassword = login_user.LoginUserPassword().Value()
	up_model.UserEnable = login_user.LoginUserEnable().Value()
	// 更新を実行する
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("user_id", "user_name", "user_mail", "user_password", "user_enable")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// ユーザを削除する
func (rep *loginUserRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, login_user *loginUsers.LoginUser) error {
	// 削除対象を取得する
	del_model, err := models.LoginUsers(qm.Where("user_id = ?", login_user.LoginUserId().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("ユーザID:%sは存在しないため、削除できませんでした。",
		login_user.LoginUserId().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 削除を実行する
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 登録処理後に実行されるフック
func LoginUserAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, login_user *models.LoginUser) error {
	log.Printf("ユーザID:%s ユーザ名:%s メールアドレス:%s パスワード:%s ユーザの有効/無効:%sを登録しました。\n",
		login_user.UserID, login_user.UserName, login_user.UserMail, login_user.UserPassword, login_user.UserEnable)
	return nil
}

// 変更処理後に実行されるフック
func LoginUserAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, login_user *models.LoginUser) error {
	log.Printf("ユーザID:%s ユーザ名:%s メールアドレス:%s パスワード:%s ユーザの有効/無効:%sを変更しました。\n",
		login_user.UserID, login_user.UserName, login_user.UserMail, login_user.UserPassword, login_user.UserEnable)
	return nil
}

// 削除処理後に実行されるフック
func LoginUserAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, login_user *models.LoginUser) error {
	log.Printf("ユーザID:%s ユーザ名:%s メールアドレス:%s パスワード:%s ユーザの有効/無効:%sを削除しました。\n",
		login_user.UserID, login_user.UserName, login_user.UserMail, login_user.UserPassword, login_user.UserEnable)
	return nil
}
