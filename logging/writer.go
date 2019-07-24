package logging

import "io"

// writer is an implementation of io.Writer that writes to a logger.
type writer struct {
	l Logger
	f func(Logger, string)
}

// NewWriter returns an io.Writer that writes content to the given logger.
//
// Each call to Write() produces a separate log message.
func NewWriter(l Logger) io.Writer {
	return &writer{l, LogString}
}

// NewDebugWriter returns an io.Writer that writes content to the given logger
// as debug messages.
//
// Each call to Write() produces a separate log message.
func NewDebugWriter(l Logger) io.Writer {
	return &writer{l, DebugString}
}

func (w *writer) Write(data []byte) (int, error) {
	w.f(w.l, string(data))
	return len(data), nil
}
