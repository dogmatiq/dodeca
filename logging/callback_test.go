package logging_test

import (
	"bytes"
	"log"

	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type CallbackLogger", func() {
	var (
		buffer bytes.Buffer
		target *log.Logger = log.New(&buffer, "", 0)
		logger *CallbackLogger
	)

	BeforeEach(func() {
		buffer.Reset()

		logger = &CallbackLogger{
			LogTarget: target.Printf,
		}
	})

	Context("when debug logging is enabled", func() {
		BeforeEach(func() {
			logger.DebugTarget = target.Printf
		})

		Describe("func Log", func() {
			It("writes a formatted message to the target", func() {
				logger.Log("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("func LogString", func() {
			It("writes a message to the target", func() {
				logger.LogString("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})

			It("does not substitute placeholders", func() {
				logger.LogString("<%s>")

				Expect(buffer.String()).To(Equal("<%s>\n"))
			})
		})

		Describe("func Debug", func() {
			It("writes a formatted message to the target", func() {
				logger.Debug("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("func DebugString", func() {
			It("writes a message to the target", func() {
				logger.DebugString("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})

			It("does not substitute placeholders", func() {
				logger.DebugString("<%s>")

				Expect(buffer.String()).To(Equal("<%s>\n"))
			})
		})

		Describe("func IsDebug", func() {
			It("returns true", func() {
				Expect(logger.IsDebug()).To(BeTrue())
			})
		})
	})

	Context("when debug logging is disabled", func() {
		Describe("func Log", func() {
			It("writes a formatted message to the target", func() {
				logger.Log("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("func LogString", func() {
			It("writes a message to the target", func() {
				logger.LogString("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})
		})

		Describe("func Debug", func() {
			It("does not produce any output", func() {
				logger.Debug("message <%s>", "arg")

				Expect(buffer.String()).To(Equal(""))
			})
		})

		Describe("func DebugString", func() {
			It("does not produce any output", func() {
				logger.DebugString("<message>")

				Expect(buffer.String()).To(Equal(""))
			})
		})

		Describe("func IsDebug", func() {
			It("returns false", func() {
				Expect(logger.IsDebug()).To(BeFalse())
			})
		})
	})
})
