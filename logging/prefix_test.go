package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Prefix()", func() {
	var (
		target *BufferedLogger
		logger Logger
	)

	BeforeEach(func() {
		target = &BufferedLogger{
			CaptureDebug: true,
		}

		logger = Prefix(target, "%s %% ", "<prefix>")
	})

	Describe("func Log()", func() {
		It("forwards to the target with the prefix", func() {
			logger.Log("message <%s>", "arg")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<prefix> % message <arg>",
				IsDebug: false,
			}))
		})
	})

	Describe("func LogString()", func() {
		It("forwards to the target with the prefix", func() {
			logger.LogString("<message>")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<prefix> % <message>",
				IsDebug: false,
			}))
		})
	})

	Describe("func Debug()", func() {
		It("forwards to the target with the prefix", func() {
			logger.Debug("message <%s>", "arg")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<prefix> % message <arg>",
				IsDebug: true,
			}))
		})
	})

	Describe("func DebugString()", func() {
		It("fforwards to the target with the prefix", func() {
			logger.DebugString("<message>")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<prefix> % <message>",
				IsDebug: true,
			}))
		})
	})

	Describe("func IsDebug()", func() {
		It("returns true if the target captures debug messages", func() {
			Expect(logger.IsDebug()).To(BeTrue())
		})

		It("returns false if the target does not capture debug messages", func() {
			target.CaptureDebug = false
			Expect(logger.IsDebug()).To(BeFalse())
		})
	})

	Describe("func UnwrapLogger()", func() {
		It("returns the target", func() {
			l := Unwrap(logger)
			Expect(l).To(BeIdenticalTo(target))
		})
	})
})
