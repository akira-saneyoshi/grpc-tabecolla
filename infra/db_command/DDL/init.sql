-- テーブル作成
DROP DATABASE IF EXISTS new_tabecolla;
CREATE DATABASE tabecolla;

-- データベースを選択
USE tabecolla;

-- APIユーザ作成
CREATE USER "tabecolla_repl"@"172.%" IDENTIFIED BY 'tabecolla_repl_pass';
GRANT ALL PRIVILEGES ON tabecolla.* TO "tabecolla_repl"@"172.%";
GRANT ALL PRIVILEGES ON tabecolla.* TO "tabecolla_repl"@"172.%";

-- レプリケーション権限の作成
GRANT REPLICATION SLAVE ON *.* TO "tabecolla_repl"@"172.%";

-- ログインユーザテーブル作成
-- * BackupToTempTable
drop table if exists `login_users` cascade;

-- * RestoreFromTempTable
create table `login_users` (
  `id` INT not null AUTO_INCREMENT comment 'ID'
  , `user_id` VARCHAR(30) not null comment 'ユーザID文字列'
  , `user_mail` VARCHAR(254) not null comment 'メールアドレス'
  , `user_password` VARCHAR(60) not null comment 'パスワード'
  , `user_name` VARCHAR(50) not null comment 'ユーザ名'
  , `user_enable` BOOLEAN not null comment '有効フラグ'
  , constraint `login_users_PKC` primary key (`id`)
  , unique key `stores_IDX1` (`id`)
) comment 'ログインユーザ';

-- 飲食店カテゴリ
-- * BackupToTempTable
drop table if exists `store_categories` cascade;

-- * RestoreFromTempTable
create table `store_categories` (
  `id` INT not null AUTO_INCREMENT comment 'ID'
  , `obj_id` VARCHAR(36) not null comment '飲食店カテゴリID文字列'
  , `name` VARCHAR(30) not null comment '飲食店カテゴリ名'
  , constraint `store_categories_PKC` primary key (`id`)
  , unique key `store_categories_IDX1` (`obj_id`)
) comment '飲食店カテゴリ';

-- 飲食店テーブル作成
-- * BackupToTempTable
drop table if exists `stores` cascade;

-- * RestoreFromTempTable
create table `stores` (
  `id` INT not null AUTO_INCREMENT comment 'ID'
  , `obj_id` VARCHAR(36) not null comment '飲食店ID文字列'
  , `name` VARCHAR(30) not null comment '飲食店名:C_TX_STORE'
  , `place` VARCHAR(100) not null comment '飲食店の場所'
  , `store_category_id` VARCHAR(36) not null comment '飲食店カテゴリID'
  , constraint `stores_PKC` primary key (`id`)
  , unique key `stores_IDX1` (`obj_id`)
  , foreign key `store_categories_FK` (`store_category_id`) references `store_categories` (`obj_id`)
) comment '飲食店';
