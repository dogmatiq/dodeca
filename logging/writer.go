package logging

import (
	"bytes"
	"io"
	"strings"
	"sync"
)

// writer is an implementation of io.Writer that writes to a logger.
type writer struct {
	l Logger
	f func(Logger, string)

	mu  sync.Mutex
	buf strings.Builder
}

// NewWriter returns an io.Writer that writes content to the given logger.
//
// The string passed to Write() is sliced by end-of-line (EOL) characters and
// each substring in between the EOL characters is written as a separate log
// message.
//
// If the entire string or its last section does not terminate with an EOL
// character, it is buffered inside the writer. The writer keeps buffering the
// string until it detects an EOL character and produces a log message with the
// buffered string content before the EOL character.
func NewWriter(l Logger) io.Writer {
	return &writer{l: l, f: LogString}
}

// NewDebugWriter returns an io.Writer that writes content to the given logger
// as debug messages.
//
// The string passed to Write() is sliced by end-of-line (EOL) characters and
// each substring in between the EOL characters is written as a separate log
// message.
//
// If the entire string or its last section does not terminate with an EOL
// character, it is buffered inside the writer. The writer keeps buffering the
// string until it detects an EOL character and produces a log message with the
// buffered string content before the EOL character.
func NewDebugWriter(l Logger) io.Writer {
	return &writer{l: l, f: DebugString}
}

func (w *writer) Write(data []byte) (int, error) {
	var n int

	w.mu.Lock()
	defer w.mu.Unlock()

	i := bytes.IndexRune(data, '\n')
	for i != -1 {
		w.buf.Write(data[:i])
		w.f(w.l, w.buf.String())
		n += w.buf.Len()
		w.buf.Reset()

		data = data[i+1:]

		i = bytes.IndexRune(data, '\n')
	}

	// Write the remaining data, if any, into the temporary string buffer
	w.buf.Write(data)

	return n, nil
}
