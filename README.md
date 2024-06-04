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

### protocの共通項目の定義

```
syntax = "proto3"; // ライセンスヘッダ
option go_package = "/api/v1/go_pb"; // 生成コードの格納先パッケージを指定
package go_protoc; // パッケージの宣言
```
