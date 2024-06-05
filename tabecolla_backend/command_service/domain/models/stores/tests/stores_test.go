package stores_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 飲食店エンティティ、値オブジェクトテスト用エントリーポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "domain/models/storesパッケージのテスト")
}
