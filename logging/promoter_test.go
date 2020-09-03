package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Promoter", func() {
	var (
		target   *BufferedLogger
		promoter *Promoter
	)

	BeforeEach(func() {
		target = &BufferedLogger{}

		promoter = &Promoter{
			Target: target,
		}
	})

	Describe("func Log()", func() {
		It("forwards to the target", func() {
			promoter.Log("message <%s>", "arg")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "message <arg>",
				IsDebug: false,
			}))
		})
	})

	Describe("func LogString()", func() {
		It("forwards to the target", func() {
			promoter.LogString("<message>")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			}))
		})
	})

	Describe("func Debug()", func() {
		It("forwards a non-debug message to the target", func() {
			promoter.Debug("message <%s>", "arg")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "message <arg>",
				IsDebug: false,
			}))
		})
	})

	Describe("func DebugString()", func() {
		It("forwards a non-debug message to the target", func() {
			promoter.DebugString("<message>")

			Expect(target.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			}))
		})
	})

	Describe("func IsDebug()", func() {
		It("returns true even if the target does not capture debug messages", func() {
			target.CaptureDebug = false
			Expect(promoter.IsDebug()).To(BeTrue())
		})
	})
})
