-- テーブル作成
DROP DATABASE IF EXISTS new_tabecolla;
CREATE DATABASE tabecolla;

-- データベースを選択
USE tabecolla;

-- APIユーザ作成
CREATE USER "tabecolla_repl"@"%" IDENTIFIED WITH mysql_native_password BY 'tabecolla_repl_pass';
GRANT ALL PRIVILEGES ON tabecolla.* TO "tabecolla_repl"@"%";

-- レプリケーション権限の作成
GRANT REPLICATION SLAVE ON *.* TO "tabecolla_repl"@"%";

-- I/Oスレッドを停止
STOP REPLICA IO_THREAD FOR CHANNEL '';

-- レプリケーションソースの変更
CHANGE REPLICATION SOURCE TO 
    SOURCE_HOST='commandservice_db',
    SOURCE_PORT=3306,
    SOURCE_USER='tabecolla_repl',
    SOURCE_PASSWORD='tabecolla_repl_pass',
    SOURCE_LOG_FILE='binlog.000003',
    SOURCE_LOG_POS=157;

-- I/Oスレッドを再開
START REPLICA;