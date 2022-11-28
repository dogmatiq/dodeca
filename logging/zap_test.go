package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

var _ = Describe("func Zap()", func() {
	var (
		target *zap.Logger
		logs   *observer.ObservedLogs
		logger Logger
	)

	BeforeEach(func() {
		var core zapcore.Core
		core, logs = observer.New(zapcore.DebugLevel)

		target = zap.New(core)
		logger = Zap(target)
	})

	Describe("func Log()", func() {
		It("forwards to the target", func() {
			logger.Log("message <%s>", "arg")

			Expect(logs.AllUntimed()).To(ConsistOf(
				observer.LoggedEntry{
					Entry: zapcore.Entry{
						Level:   zapcore.InfoLevel,
						Message: "message <arg>",
					},
					Context: []zapcore.Field{},
				},
			))
		})
	})

	Describe("func LogString()", func() {
		It("forwards to the target", func() {
			logger.LogString("<message>")

			Expect(logs.AllUntimed()).To(ConsistOf(
				observer.LoggedEntry{
					Entry: zapcore.Entry{
						Level:   zapcore.InfoLevel,
						Message: "<message>",
					},
					Context: []zapcore.Field{},
				},
			))
		})
	})

	Describe("func Debug()", func() {
		It("forwards a non-debug message to the target", func() {
			logger.Debug("message <%s>", "arg")

			Expect(logs.AllUntimed()).To(ConsistOf(
				observer.LoggedEntry{
					Entry: zapcore.Entry{
						Level:   zapcore.DebugLevel,
						Message: "message <arg>",
					},
					Context: []zapcore.Field{},
				},
			))
		})
	})

	Describe("func DebugString()", func() {
		It("forwards a non-debug message to the target", func() {
			logger.DebugString("<message>")

			Expect(logs.AllUntimed()).To(ConsistOf(
				observer.LoggedEntry{
					Entry: zapcore.Entry{
						Level:   zapcore.DebugLevel,
						Message: "<message>",
					},
					Context: []zapcore.Field{},
				},
			))
		})
	})

	Describe("func IsDebug()", func() {
		It("returns true even if the target captures debug messages", func() {
			Expect(logger.IsDebug()).To(BeTrue())
		})

		It("returns false even if the target does not capture debug messages", func() {
			var core zapcore.Core
			core, logs = observer.New(zapcore.InfoLevel)

			target = zap.New(core)
			logger = Zap(target)

			Expect(logger.IsDebug()).To(BeFalse())
		})
	})
})
