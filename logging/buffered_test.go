package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type BufferedLogger", func() {
	var (
		logger *BufferedLogger
	)

	BeforeEach(func() {
		logger = &BufferedLogger{}
	})

	Context("when debug logging is enabled", func() {
		BeforeEach(func() {
			logger.CaptureDebug = true
		})

		Describe("func Log", func() {
			It("stores a formatted message", func() {
				logger.Log("message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("func LogString", func() {
			It("stores a message", func() {
				logger.LogString("<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})

			It("does not substitute placeholders", func() {
				logger.LogString("<%s>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<%s>",
					IsDebug: false,
				}))
			})
		})

		Describe("func Debug", func() {
			It("stores a formatted message", func() {
				logger.Debug("message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: true,
				}))
			})
		})

		Describe("func DebugString", func() {
			It("stores a message", func() {
				logger.DebugString("<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: true,
				}))
			})

			It("does not substitute placeholders", func() {
				logger.DebugString("<%s>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<%s>",
					IsDebug: true,
				}))
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
			It("stores a formatted message", func() {
				logger.Log("message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("func LogString", func() {
			It("stores a message", func() {
				logger.LogString("<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})

			It("does not substitute placeholders", func() {
				logger.LogString("<%s>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<%s>",
					IsDebug: false,
				}))
			})
		})

		Describe("func Debug", func() {
			It("does not produce any output", func() {
				logger.Debug("message <%s>", "arg")

				Expect(logger.Messages()).To(BeEmpty())
			})
		})

		Describe("func DebugString", func() {
			It("does not produce any output", func() {
				logger.DebugString("<message>")

				Expect(logger.Messages()).To(BeEmpty())
			})
		})

		Describe("func IsDebug", func() {
			It("returns false", func() {
				Expect(logger.IsDebug()).To(BeFalse())
			})
		})
	})

	Describe("Reset", func() {
		It("clears messages from the logger", func() {
			logger.LogString("<message>")
			logger.Reset()

			Expect(logger.Messages()).To(BeEmpty())
		})
	})

	Describe("TakeMessages", func() {
		It("returns and clears messages from the logger", func() {
			logger.LogString("<message>")

			m := logger.TakeMessages()
			Expect(m).To(HaveLen(1))
			Expect(logger.Messages()).To(BeEmpty())
		})
	})

	Describe("Flush", func() {
		dest := &BufferedLogger{
			CaptureDebug: true,
		}

		BeforeEach(func() {
			logger.CaptureDebug = true
			dest.Reset()
		})

		It("logs captures messages to the destination logger", func() {
			logger.LogString("<message>")
			logger.FlushTo(dest)

			Expect(logger.Messages()).To(BeEmpty())
			Expect(dest.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			}))
		})

		It("logs debug messages as debug messages", func() {
			logger.DebugString("<message>")
			logger.FlushTo(dest)

			Expect(logger.Messages()).To(BeEmpty())
			Expect(dest.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: true,
			}))
		})
	})
})

var _ = Describe("type BufferedLogMessage", func() {
	Describe("string String", func() {
		It("returns the log message", func() {
			m := BufferedLogMessage{Message: "<message>"}
			Expect(m.String()).To(Equal("<message>"))
		})
	})
})
