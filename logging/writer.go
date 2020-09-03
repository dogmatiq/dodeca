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

	m   sync.Mutex
	buf strings.Builder
}

// NewWriter returns an io.WriteCloser that writes content to the given logger.
//
// Each line-feed-terminated line in the data passed to Write() is logged as a
// separate message. Any unterminated line is buffered until a line-feed is
// encountered in a future call to Write(), or the writer is closed.
func NewWriter(l Logger) io.WriteCloser {
	return &writer{l: l, f: LogString}
}

// NewDebugWriter returns an io.Writer that writes content to the given logger
// as debug messages.
//
// Each line-feed-terminated line in the data passed to Write() is logged as a
// separate message. Any unterminated line is buffered until a line-feed is
// encountered in a future call to Write(), or the writer is closed.
func NewDebugWriter(l Logger) io.WriteCloser {
	return &writer{l: l, f: DebugString}
}

func (w *writer) Write(data []byte) (int, error) {
	n := len(data)

	w.m.Lock()
	defer w.m.Unlock()

	i := bytes.IndexRune(data, '\n')
	for i != -1 {
		w.buf.Write(data[:i])
		w.f(w.l, w.buf.String())
		w.buf.Reset()

		data = data[i+1:]

		i = bytes.IndexRune(data, '\n')
	}

	// Write the remaining data, if any, into the temporary string buffer
	w.buf.Write(data)

	return n, nil
}

func (w *writer) Close() error {
	w.m.Lock()
	defer w.m.Unlock()

	if w.buf.Len() > 0 {
		w.f(w.l, w.buf.String())
		w.buf.Reset()
	}

	return nil
}
