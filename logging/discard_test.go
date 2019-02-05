package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type DiscardLogger", func() {
	Describe("func IsDebug", func() {
		It("returns false", func() {
			Expect(DiscardLogger{}.IsDebug()).To(BeFalse())
		})
	})
})
