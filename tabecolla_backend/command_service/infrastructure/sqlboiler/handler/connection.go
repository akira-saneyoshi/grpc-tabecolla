package handler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// データベース接続情報
type DBConfig struct {
	DbName string `toml:"dbName"` //	データベース名
	Host   string `toml:"host"`   //	ホスト名
	Port   int64  `toml:"port"`   //	ポート番号
	User   string `toml:"user"`   //	ユーザー名
	Pass   string `toml:"pass"`   //	パスワード
}

// database.tomlから接続情報を取得してDbConfig型で返す
func tomlRead() (*DBConfig, error) {
	path := os.Getenv("DATABASE_TOML_PATH")
	if path == "" {
		path = "infrastructure/sqlboiler/config/database.toml"
	}
	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}
	config := m["mysql"]
	return &config, nil
}

// データベース接続
func DBConnect() error {

	config, err := tomlRead()
	if err != nil {
		return DBErrHandler(err)
	}

	rdbms := "mysql"
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.DbName)

	conn, err := sql.Open(rdbms, connect_str)
	if err != nil {
		return DBErrHandler(err)
	}
	if err = conn.Ping(); err != nil {
		return DBErrHandler(err)
	}

	MAX_IDLE_CONNS := 10                   // 初期接続数
	MAX_OPEN_CONNS := 100                  // 最大接続数
	CONN_MAX_LIFETIME := 300 * time.Second // 最大生存期間

	// コネクションプールの設定
	conn.SetMaxIdleConns(MAX_IDLE_CONNS)       // 初期接続数
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)       // 最大接続数
	conn.SetConnMaxLifetime(CONN_MAX_LIFETIME) // 最大利用生存期間

	boil.SetDB(conn)      // グローバルコネクション設定
	boil.DebugMode = true // デバッグモードに設定 生成されたSQLを出力する
	return nil
}
