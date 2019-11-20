package logging

// CallbackLogger is an implementation of Logger that forwards calls to Log and
// Debug to a delegate callback. CallbackLogger MUST have LogTarget set.
type CallbackLogger struct {
	// LogTarget is the delegate used for log messages.
	LogTarget Printf
	// DebugTarget is the delegate used for debug messages. When this
	// field isn't set, no debug logging will occur and IsDebug() will
	// return false
	DebugTarget Printf
}

// Printf is the function signature for fmt.Printf() and alike functions.
type Printf func(fmt string, v ...interface{})

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// fmt is the format specifier, as per fmt.Printf(), etc.
func (l *CallbackLogger) Log(fmt string, v ...interface{}) {
	l.LogTarget(fmt, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l *CallbackLogger) LogString(s string) {
	l.LogTarget("%s", s)
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// fmt is the format specifier, as per fmt.Printf(), etc.
func (l *CallbackLogger) Debug(fmt string, v ...interface{}) {
	if l.DebugTarget != nil {
		l.DebugTarget(fmt, v...)
	}
}

// DebugString writes a pre-formatted debug log message.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l *CallbackLogger) DebugString(s string) {
	if l.DebugTarget != nil {
		l.DebugTarget("%s", s)
	}
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l *CallbackLogger) IsDebug() bool {
	return l.DebugTarget != nil
}
