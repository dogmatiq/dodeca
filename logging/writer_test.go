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

		m := []byte("<message>")
		n, err := writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(Equal(len(m)))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			},
		))
	})
})

var _ = Describe("NewDebugWriter", func() {
	It("returns a writer that produces debug log messages", func() {
		logger := &BufferedLogger{CaptureDebug: true}
		writer := NewDebugWriter(logger)

		m := []byte("<message>")
		n, err := writer.Write(m)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(Equal(len(m)))

		Expect(logger.Messages()).To(ConsistOf(
			BufferedLogMessage{
				Message: "<message>",
				IsDebug: true,
			},
		))
	})
})
