package logging

import (
	"log"
	"os"
)

// StandardLogger is an implementation of Logger that uses Go's standard logger
// package.
type StandardLogger struct {
	// Target is the standard Go logger used to write messages. If it is nil,
	// logs are sent to STDOUT as per the 12-factor logging recomendations. Note
	// that this differs to Go's default logger, which writes to STDERR.
	Target *log.Logger

	// CaptureDebug controls whether debug messages should be written to the
	// target logger.
	CaptureDebug bool
}

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l *StandardLogger) Log(f string, v ...interface{}) {
	l.target().Printf(f, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l *StandardLogger) LogString(s string) {
	l.target().Println(s)
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l *StandardLogger) Debug(f string, v ...interface{}) {
	if l.CaptureDebug {
		l.target().Printf(f, v...)
	}
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l *StandardLogger) DebugString(s string) {
	if l.CaptureDebug {
		l.target().Println(s)
	}
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l *StandardLogger) IsDebug() bool {
	return l.CaptureDebug
}

func (l *StandardLogger) target() *log.Logger {
	if l.Target == nil {
		return defaultTarget
	}

	return l.Target
}

var defaultTarget = log.New(os.Stdout, "", 0)
