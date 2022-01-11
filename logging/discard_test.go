package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type DiscardLogger", func() {
	var logger DiscardLogger

	Describe("func Log()", func() {
		It("does not panic", func() {
			logger.Log("<message>")
		})
	})

	Describe("func LogString()", func() {
		It("does not panic", func() {
			logger.LogString("<message>")
		})
	})

	Describe("func Debug()", func() {
		It("does not panic", func() {
			logger.Debug("<message>")
		})
	})

	Describe("func DebugString()", func() {
		It("does not panic", func() {
			logger.DebugString("<message>")
		})
	})

	Describe("func IsDebug()", func() {
		It("returns false", func() {
			Expect(logger.IsDebug()).To(BeFalse())
		})
	})
})
