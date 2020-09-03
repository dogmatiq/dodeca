package logging

import (
	"fmt"
	"sync"
)

// BufferedLogger is an implementation of Logger that buffers log messages in
// memory.
type BufferedLogger struct {
	// CaptureDebug controls whether debug messages should be stored.
	CaptureDebug bool

	m        sync.RWMutex
	messages []BufferedLogMessage
}

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l *BufferedLogger) Log(f string, v ...interface{}) {
	l.LogString(
		fmt.Sprintf(f, v...),
	)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l *BufferedLogger) LogString(s string) {
	l.m.Lock()
	defer l.m.Unlock()

	l.messages = append(
		l.messages,
		BufferedLogMessage{s, false},
	)
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l *BufferedLogger) Debug(f string, v ...interface{}) {
	if l.CaptureDebug {
		l.DebugString(
			fmt.Sprintf(f, v...),
		)
	}
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l *BufferedLogger) DebugString(s string) {
	if l.CaptureDebug {
		l.m.Lock()
		defer l.m.Unlock()

		l.messages = append(
			l.messages,
			BufferedLogMessage{s, true},
		)
	}
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l *BufferedLogger) IsDebug() bool {
	return l.CaptureDebug
}

// Reset removes all buffered log messages.
func (l *BufferedLogger) Reset() {
	l.m.Lock()
	defer l.m.Unlock()

	l.messages = nil
}

// Messages returns the messages that have been logged.
func (l *BufferedLogger) Messages() []BufferedLogMessage {
	l.m.RLock()
	defer l.m.RUnlock()

	return append(
		[]BufferedLogMessage(nil),
		l.messages...,
	)
}

// TakeMessages returns the messages that have been logged and resets the logger
// in a single operation.
func (l *BufferedLogger) TakeMessages() []BufferedLogMessage {
	l.m.Lock()
	defer l.m.Unlock()

	m := l.messages
	l.messages = nil

	return m
}

// FlushTo logs the buffered messages to dest, and resets the logger.
func (l *BufferedLogger) FlushTo(dest Logger) {
	for _, m := range l.TakeMessages() {
		if m.IsDebug {
			dest.DebugString(m.Message)
		} else {
			dest.LogString(m.Message)
		}
	}
}

// BufferedLogMessage is a log message stored by a BufferedLogger
type BufferedLogMessage struct {
	Message string
	IsDebug bool
}

func (m BufferedLogMessage) String() string {
	return m.Message
}
