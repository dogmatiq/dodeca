package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Tee()", func() {
	var (
		targetA, targetB, targetC *BufferedLogger
		logger                    Logger
	)

	BeforeEach(func() {
		targetA = &BufferedLogger{
			CaptureDebug: true,
		}

		targetB = &BufferedLogger{
			CaptureDebug: true,
		}

		targetC = &BufferedLogger{
			CaptureDebug: false,
		}

		logger = Tee(targetA, targetB, targetC)
	})

	It("removes duplicate loggers", func() {
		logger = Tee(targetA, targetA)
		logger.Log("<message>")

		Expect(targetA.Messages()).To(HaveLen(1))
	})

	It("panics if no target are provided", func() {
		Expect(func() {
			Tee()
		}).To(PanicWith("at least one target logger must be provided"))
	})

	Describe("func Log()", func() {
		It("forwards to the all targets", func() {
			logger.Log("message <%s>", "arg")

			x := BufferedLogMessage{
				Message: "message <arg>",
				IsDebug: false,
			}

			Expect(targetA.Messages()).To(ConsistOf(x))
			Expect(targetB.Messages()).To(ConsistOf(x))
			Expect(targetC.Messages()).To(ConsistOf(x))
		})
	})

	Describe("func LogString()", func() {
		It("forwards to all targets", func() {
			logger.LogString("<message>")

			x := BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			}

			Expect(targetA.Messages()).To(ConsistOf(x))
			Expect(targetB.Messages()).To(ConsistOf(x))
			Expect(targetC.Messages()).To(ConsistOf(x))
		})
	})

	Describe("func Debug()", func() {
		It("forwards to all debug targets", func() {
			logger.Debug("message <%s>", "arg")

			x := BufferedLogMessage{
				Message: "message <arg>",
				IsDebug: true,
			}

			Expect(targetA.Messages()).To(ConsistOf(x))
			Expect(targetB.Messages()).To(ConsistOf(x))
			Expect(targetC.Messages()).To(BeEmpty())
		})
	})

	Describe("func DebugString()", func() {
		It("forwards to all targets", func() {
			logger.DebugString("<message>")

			x := BufferedLogMessage{
				Message: "<message>",
				IsDebug: true,
			}

			Expect(targetA.Messages()).To(ConsistOf(x))
			Expect(targetB.Messages()).To(ConsistOf(x))
			Expect(targetC.Messages()).To(BeEmpty())
		})
	})

	Describe("func IsDebug()", func() {
		It("returns true if any target captures debug messages", func() {
			Expect(logger.IsDebug()).To(BeTrue())
		})

		It("returns false if none of the targets capture debug messages", func() {
			logger = Tee(targetC)
			Expect(logger.IsDebug()).To(BeFalse())
		})
	})
})
