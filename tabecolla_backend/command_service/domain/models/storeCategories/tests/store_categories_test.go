package store_categories_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 飲食店カテゴリエンティティ、値オブジェクトテスト用エントリーポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "domain/models/storeCategoriesパッケージのテスト")
}
