package linters

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/onsi/gomega"
	"golang.org/x/tools/go/analysis/analysistest"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("ocmlogger analyzer", func() {
	It("reports only expected diagnostics according to // want comments", func() {
		td := analysistest.TestData()
		newPlugin, err := register.GetPlugin("ocmlogger")
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		plugin, err := newPlugin(nil)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		analyzers, err := plugin.BuildAnalyzers()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		analysistest.Run(GinkgoT(), td, analyzers[0], "data")
	})
})
