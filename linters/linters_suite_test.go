package linters_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLinters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Linters Suite")
}
