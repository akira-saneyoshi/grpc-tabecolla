package sqlboiler

import (
	"commandservice/infra/sqlboiler/repository"

	"go.uber.org/fx"
)

// SQLBoilerを利用したRepositoryの依存定義
var RepDepend = fx.Options(
	fx.Provide(
		// Repositoryインターフェイス実装のコンストラクタを指定
		repository.NewStoreCategoryRepositorySQLBoiler, // 飲食店カテゴリ用Reposititory
		repository.NewStoreRepositorySQLBoiler,  // 飲食店用Repository
		repository.NewLoginUserRepositorySQLBoiler,  // ユーザ用Repository
	),
)
