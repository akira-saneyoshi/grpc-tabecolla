# grpc-tabecolla

## 使用技術

- フロントエンド構成

| 種別 | 技術 | バージョン |
| --- | --- | --- |
| 言語 | TypeScript | - |
| ランタイム | Node.js | - |
| フレームワーク | Next.js | - |
| ライブラリ | React | - |

- バックエンド構成(gRPC API)

| 種別 | 技術 | バージョン |
| --- | --- | --- |
| 言語 | Golang | 1.22.3 |
| 言語 | Protobuf | 5.27.0 |
| ライブラリ/パッケージ | go.mod参照 | - |

- infra構成

| 種別 | 技術 | バージョン |
| --- | --- | --- |
| 基盤 | Docker | - |
| 基盤 | Docker for Desktop | 4.30.0 |
| データベース | MySQL | 8.0.32 |


## フロントエンド コーディング規約
- https://kinsta.com/jp/blog/react-best-practices/
- https://typescript-jp.gitbook.io/deep-dive/styleguide

## ローカル環境構築　Docker環境を利用する場合

1. まずは docker が必要です。以下よりダウンロードしてDocker for Desktopのセットアップを完了させてください。もし、 Ubuntu のインストールが要求された場合はあわせて Microsoft Storeより インストールします。

[Docker for Desktop](https://www.docker.com/)　インストーラ

2. Docker Containerを起動する。プロジェクトのルートディレクトリにて実行してください。
```bash
saneyoshi@LAPTOP-IOE8NK0E MINGW64 ~/workspace/lab-coding/grpc-tabecolla (develop)
$ make up
```

### 環境構築時に使用したコマンド集
1. protocセットアップ時

```
go get -u google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. protocを利用したGoコードの生成(protoコンテナで実行)

```
protoc --go_out=./ --go-grpc_out=./ go_proto/*.proto
```

3. Model層におけるGoコードの生成(commandコンテナで実行)

```
sqlboiler mysql -c config/database.toml -o models -p models --no-tests --wipe
```

### protocの共通項目の定義

```
syntax = "proto3"; // ライセンスヘッダ
option go_package = "/api/v1/go_pb"; // 生成コードの格納先パッケージを指定
package go_protoc; // パッケージの宣言
```

### Ginkgo V2を使用したビヘイビア駆動開発用テストフレームワークの実行手順

1. テストコードの実行コマンド

下記は、飲食店モデル

```
/opt/command_service/domain/models/stores/tests # ginkgo -v
```

<details><summary>コンソールの実行ログ例</summary>

```
/opt/command_service/domain/models/stores/tests # ginkgo -v
Running Suite: domain/models/storesパッケージのテスト - /opt/command_service/domain/models/stores/tests
================================================================================================================
Random Seed: 1717514749

Will run 4 of 4 specs
------------------------------
Storeエンティティを構成する値オブジェクト 文字数の検証 空文字列の場合、errs.DomainErrorが返る [StoreId構造体の生成, 文字数]
/opt/command_service/domain/models/stores/tests/value_test.go:29
• [0.000 seconds]
------------------------------
Storeエンティティを構成する値オブジェクト 文字数の検証 36文字より大きい文字列の場合、errs.DomainErrorが返る [StoreId構造体の生成, 文字数]
/opt/command_service/domain/models/stores/tests/value_test.go:32
• [0.000 seconds]
------------------------------
Storeエンティティを構成する値オブジェクト UUID形式の検証 uuid以外の文字列の場合、errs.DomainErrorが返る [StoreId構造体の生成, UUID形式]
/opt/command_service/domain/models/stores/tests/value_test.go:38
• [0.000 seconds]
------------------------------
Storeエンティティを構成する値オブジェクト UUID形式の検証 36文字のuuid文字列の場合、StoreIdが返る [StoreId構造体の生成, UUID形式]
/opt/command_service/domain/models/stores/tests/value_test.go:41
• [0.000 seconds]
------------------------------

Ran 4 of 4 Specs in 0.002 seconds
SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 936.73792ms
Test Suite Passed
```

</details>

### MySQLコンテナのmy.cnfをマウントするときの備忘録

- 概要

コンテナログにて下記のWarningが出ているとき、my.cnfが反映されない問題が起きる。

```
[Warning] World-writable config file '/etc/mysql/conf.d/my.cnf' is ignored.
```

- 解決策

windowsで環境構築している場合、/etc/mysql/conf.d/my.cnfにマウントしたファイルは読み取り専用にしとかなければならない
