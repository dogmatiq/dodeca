package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
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
		It("writes each line as a separate log message", func() {
			m := []byte("<message1>\n<message2>\n<message3>")
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

			m = []byte("<message4>\n<message5>\n")
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
		})
	})

	Describe("func Close()", func() {
		It("writes the remaining buffered content when Close() is called", func() {
			m := []byte("<message1>\n<message2>\n<message3>")
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

			err = writer.Close()
			Expect(err).ShouldNot(HaveOccurred())

			Expect(logger.Messages()).To(ConsistOf(
				BufferedLogMessage{
					Message: "<message3>",
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
