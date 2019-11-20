package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type StandardLogger", func() {
	var (
		target *testingTMock
		logger *TestLogger
	)

	BeforeEach(func() {
		target = &testingTMock{}

		logger = &TestLogger{
			Target: target,
		}
	})

	Context("when debug logging is enabled", func() {
		Describe("func Log", func() {
			It("writes a formatted message to the target", func() {
				logger.Log("message <%s>", "arg")

				Expect(target.call).To(Equal([]callRecord{{
					format: "message <%s>",
					args:   []interface{}{"arg"},
				}}))
			})
		})

		Describe("func LogString", func() {
			It("writes a message to the target", func() {
				logger.LogString("<message>")

				Expect(target.call).To(Equal([]callRecord{{
					format: "<message>",
				}}))
			})
		})

		Describe("func Debug", func() {
			It("writes a formatted message to the target", func() {
				logger.Debug("message <%s>", "arg")

				Expect(target.call).To(Equal([]callRecord{{
					format: "message <%s>",
					args:   []interface{}{"arg"},
				}}))
			})
		})

		Describe("func DebugString", func() {
			It("writes a message to the target", func() {
				logger.DebugString("<message>")

				Expect(target.call).To(Equal([]callRecord{{
					format: "<message>",
				}}))
			})
		})

		Describe("func IsDebug", func() {
			It("returns true", func() {
				Expect(logger.IsDebug()).To(BeTrue())
			})
		})
	})
})

type testingTMock struct {
	call []callRecord
}

type callRecord struct {
	format string
	args   []interface{}
}

func (t *testingTMock) Logf(format string, args ...interface{}) {
	t.call = append(t.call, callRecord{
		format: format,
		args:   args,
	})
}
