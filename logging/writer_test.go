package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("type StreamWriter", func() {
	var (
		logger *BufferedLogger
		writer *StreamWriter
	)

	BeforeEach(func() {
		logger = &BufferedLogger{}
		writer = &StreamWriter{Target: logger}
	})

	Describe("func Write()", func() {
		DescribeTable(
			"writes each line of text as a separate log message",
			func(sep string) {
				m := []byte("<message1>" + sep + "<message2>" + sep + "<message3>")
				n, err := writer.Write(m)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(n).To(BeNumerically("==", len(m)))

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

				// Note: no new line between messages 3 and 4.
				m = []byte("<message4>" + sep + "<message5>" + sep)
				n, err = writer.Write(m)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(n).To(BeNumerically("==", len(m)))

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
			},
			Entry("LF", "\n"),
			Entry("CR", "\r"),
			Entry("CRLF", "\r\n"),
		)

		It("handles CRLF split across separate calls to Write()", func() {
			_, err := writer.Write([]byte("<message1>\r"))
			Expect(err).ShouldNot(HaveOccurred())

			_, err = writer.Write([]byte("\n<message2>\r\n"))
			Expect(err).ShouldNot(HaveOccurred())

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
		})
	})

	Describe("func Close()", func() {
		It("writes unterminated lines when Close() is called", func() {
			m := []byte("<message>")
			n, err := writer.Write(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(n).To(BeNumerically("==", len(m)))

			Expect(logger.Messages()).To(BeEmpty())

			err = writer.Close()
			Expect(err).ShouldNot(HaveOccurred())

			Expect(logger.Messages()).To(ConsistOf(
				BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				},
			))
		})
	})
})

var _ = Describe("type LineWriter", func() {
	var (
		logger *BufferedLogger
		writer *LineWriter
	)

	BeforeEach(func() {
		logger = &BufferedLogger{}
		writer = &LineWriter{Target: logger}
	})

	It("writes a single log message", func() {
		m := []byte("<message>")
		n, err := writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(BeNumerically("==", len(m)))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			},
		))
	})
})
