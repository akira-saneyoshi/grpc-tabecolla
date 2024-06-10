-- テーブル作成
DROP DATABASE IF EXISTS new_tabecolla;
CREATE DATABASE tabecolla;

-- データベースを選択
USE tabecolla;

-- APIユーザ作成
CREATE USER "tabecolla_repl"@"%" IDENTIFIED BY 'tabecolla_repl_pass';
GRANT ALL PRIVILEGES ON tabecolla.* TO "tabecolla_repl"@"%";

-- レプリケーション権限の作成
GRANT REPLICATION SLAVE ON *.* TO "tabecolla_repl"@"%";

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
  , `place` VARCHAR(250) not null comment '飲食店の場所'
  , `store_category_id` VARCHAR(36) not null comment '飲食店カテゴリID'
  , constraint `stores_PKC` primary key (`id`)
  , unique key `stores_IDX1` (`obj_id`)
  , foreign key `store_categories_FK` (`store_category_id`) references `store_categories` (`obj_id`)
) comment '飲食店';

INSERT INTO store_categories (obj_id,name) VALUES ('aaaa1111-22bb-33cc-44dd-555555eeeeee','和食');
INSERT INTO store_categories (obj_id,name) VALUES ('bbbb1111-22aa-dd33-cc44-eeeeee666666','洋食');
INSERT INTO store_categories (obj_id,name) VALUES ('cccc1111-22dd-33cc-44dd-555555eeeeee','和食(鍋)');
INSERT INTO store_categories (obj_id,name) VALUES ('dddd1111-cc22-dd33-cc44-ffffff666666','イタリアン');
INSERT INTO store_categories (obj_id,name) VALUES ('eeee1111-22ff-33cc-44dd-666666eeeeee','ハンバーガー');

INSERT INTO stores (obj_id,name,place,store_category_id) VALUES ('ffff1111-99cc-bb88-22aa-77ee77ee77ee','モスバーガー','東京都墨田区押上１丁目２３−１ 土居ビル １階','eeee1111-22ff-33cc-44dd-666666eeeeee');
INSERT INTO stores (obj_id,name,place,store_category_id) VALUES ('1111dddd-4567-ffff-abcd-5555fda34fcb','蟻月','東京都墨田区押上１丁目１−２ 東京スカイツリータウン・ソラマチ ３１Ｆ','cccc1111-22dd-33cc-44dd-555555eeeeee');
INSERT INTO stores (obj_id,name,place,store_category_id) VALUES ('ee22ee44-fcba-76dc-cd78-333fcbacd377','Issare shu cielo','東京都墨田区押上1-1-2 東京ソラマチ 30F','dddd1111-cc22-dd33-cc44-ffffff666666');
INSERT INTO stores (obj_id,name,place,store_category_id) VALUES ('1432daca-cfbd-cd55-cbb7-acc367b75v78','RIGOLETTO ROTISSERIE AND WINE','東京都墨田区押上１丁目１−２','dddd1111-cc22-dd33-cc44-ffffff666666');
