package logging

import (
	"fmt"
	"strings"
)

// Prefix returns a logger that prefixes all messages with a fixed string.
//
// f is the format specifier for the prefix, as per fmt.Printf(), etc.
//
// If the target is nil, DefaultLogger is used.
func Prefix(target Logger, f string, v ...interface{}) Logger {
	p := fmt.Sprintf(f, v...)

	return prefixer{
		target,
		p,
		strings.ReplaceAll(p, "%", "%%"),
	}
}

// prefixer is an implementation of Logger that prefixes all message with a
// fixed string.
type prefixer struct {
	// Target is the logger to which messages are forwarded.
	Target Logger

	// StringPrefix is the string to prefix to all messages.
	StringPrefix string

	// FormatSpecifierPrefix is the prefix string escaped for use as a format
	// specifier.
	FormatSpecifierPrefix string
}

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l prefixer) Log(f string, v ...interface{}) {
	Log(
		l.Target,
		l.FormatSpecifierPrefix+f,
		v...,
	)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l prefixer) LogString(s string) {
	LogString(
		l.Target,
		l.StringPrefix+s,
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
func (l prefixer) Debug(f string, v ...interface{}) {
	Debug(
		l.Target,
		l.FormatSpecifierPrefix+f,
		v...,
	)
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l prefixer) DebugString(s string) {
	DebugString(
		l.Target,
		l.StringPrefix+s,
	)
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l prefixer) IsDebug() bool {
	return IsDebug(l.Target)
}

// UnwrapLogger returns the logger wrapped by this logger.
//
// If ok is true it means that this logger is wrapping another logger, even
// if that logger is nil, indicating that DefaultLogger should be used.
//
// If ok is false it means that this logger is not wrapping another logger.
func (l prefixer) UnwrapLogger() (Logger, bool) {
	return l.Target, true
}
