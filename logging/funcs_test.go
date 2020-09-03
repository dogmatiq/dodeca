package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logging Functions", func() {
	var (
		logger          BufferedLogger
		originalDefault Logger
	)

	BeforeEach(func() {
		logger.CaptureDebug = true
		logger.Reset()
	})

	When("passed a nil logger", func() {
		BeforeSuite(func() {
			originalDefault = DefaultLogger
			DefaultLogger = &logger
		})

		AfterSuite(func() {
			DefaultLogger = originalDefault
		})

		Describe("func Log()", func() {
			It("forwards to the default logger", func() {
				Log(nil, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("func LogString()", func() {
			It("forwards to the default logger", func() {
				LogString(nil, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})
		})

		Describe("func Debug()", func() {
			It("forwards to the default logger", func() {
				Debug(nil, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: true,
				}))
			})
		})

		Describe("func DebugString()", func() {
			It("forwards to the default logger", func() {
				DebugString(nil, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: true,
				}))
			})
		})

		Describe("func IsDebug()", func() {
			It("forwards to the default logger", func() {
				Expect(IsDebug(nil)).To(BeTrue())
				logger.CaptureDebug = false
				Expect(IsDebug(nil)).To(BeFalse())
			})
		})
	})

	When("passed a non-nil logger", func() {
		Describe("func Log()", func() {
			It("forwards to the given logger", func() {
				Log(&logger, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("func LogString()", func() {
			It("forwards to the given logger", func() {
				LogString(&logger, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})
		})

		Describe("func Debug()", func() {
			It("forwards to the given logger", func() {
				Debug(&logger, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: true,
				}))
			})
		})

		Describe("func DebugString()", func() {
			It("forwards to the given logger", func() {
				DebugString(&logger, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: true,
				}))
			})
		})

		Describe("func IsDebug()", func() {
			It("forwards to the given logger", func() {
				Expect(IsDebug(&logger)).To(BeTrue())
				logger.CaptureDebug = false
				Expect(IsDebug(&logger)).To(BeFalse())
			})
		})
	})
})
