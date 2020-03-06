package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewWriter", func() {
	It("returns a writer that produces regular log messages", func() {
		logger := &BufferedLogger{}
		writer := NewWriter(logger)

		m := []byte("<message1>\n<message2>\n<message3>")
		n, err := writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(Equal(20))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message1>",
				IsDebug: false,
			},
			BufferedLogMessage{
				Message: "<message2>",
				IsDebug: false,
			},
		))

		logger.Reset()

		m = []byte("<message4>\n<message5>\n")
		n, err = writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(Equal(30))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message3><message4>",
				IsDebug: false,
			},
			BufferedLogMessage{
				Message: "<message5>",
				IsDebug: false,
			},
		))
	})
})

var _ = Describe("NewDebugWriter", func() {
	It("returns a writer that produces debug log messages", func() {
		logger := &BufferedLogger{CaptureDebug: true}
		writer := NewDebugWriter(logger)

		m := []byte("<message1>\n<message2>\n<message3>")
		n, err := writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(Equal(20))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message1>",
				IsDebug: true,
			},
			BufferedLogMessage{
				Message: "<message2>",
				IsDebug: true,
			},
		))

		logger.Reset()

		m = []byte("<message4>\n<message5>\n")
		n, err = writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(Equal(30))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message3><message4>",
				IsDebug: true,
			},
			BufferedLogMessage{
				Message: "<message5>",
				IsDebug: true,
			},
		))
	})
})
