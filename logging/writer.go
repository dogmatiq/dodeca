package logging

import (
	"bytes"
	"strings"
	"sync"
)

// StreamWriter is an adaptor that presents a Logger as an io.WriteCloser.
//
// Each LF (line-feed) terminated line of text written via Write() is logged as
// a separate message. Any call to write with text that does not end in LF is
// buffered until a LF is encountered in a subsequent call to Write(), or
// Close() is called.
type StreamWriter struct {
	// Target is the logger that receives the log messages.
	Target Logger

	m   sync.Mutex
	buf strings.Builder
}

func (w *StreamWriter) Write(data []byte) (int, error) {
	n := len(data)

	w.m.Lock()
	defer w.m.Unlock()

	i := bytes.IndexRune(data, '\n')
	for i != -1 {
		w.buf.Write(data[:i])
		LogString(w.Target, w.buf.String())
		w.buf.Reset()

		data = data[i+1:]

		i = bytes.IndexRune(data, '\n')
	}

	// Write the remaining data, if any, into the temporary string buffer
	w.buf.Write(data)

	return n, nil
}

// Close closes the writer, producing a log message from any remaining buffered
// text.
func (w *StreamWriter) Close() error {
	w.m.Lock()
	defer w.m.Unlock()

	if w.buf.Len() > 0 {
		LogString(w.Target, w.buf.String())
		w.buf.Reset()
	}

	return nil
}

// LineWriter is an adaptor that presents a Logger as an io.Writer.
//
// For each call to Write() the data is forwarded to the logger as a separate
// message.
type LineWriter struct {
	// Target is the logger that receives the log messages.
	Target Logger
}

func (w *LineWriter) Write(data []byte) (int, error) {
	LogString(w.Target, string(data))
	return len(data), nil
}
