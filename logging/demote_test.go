package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Demote()", func() {
	var (
		target *BufferedLogger
		logger Logger
	)

	BeforeEach(func() {
		target = &BufferedLogger{
			CaptureDebug: true,
		}

		logger = Demote(target)
	})

	Describe("func Log()", func() {
		It("forwards a debug message to the target", func() {
			logger.Log("message <%s>", "arg")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "message <arg>",
				IsDebug: true,
			}))
		})
	})

	Describe("func LogString()", func() {
		It("forwards a debug message to the target", func() {
			logger.LogString("<message>")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: true,
			}))
		})
	})

	Describe("func Debug()", func() {
		It("forwards to the target", func() {
			logger.Debug("message <%s>", "arg")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "message <arg>",
				IsDebug: true,
			}))
		})
	})

	Describe("func DebugString()", func() {
		It("forwards to the target", func() {
			logger.DebugString("<message>")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
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
})
