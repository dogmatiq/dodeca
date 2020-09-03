package logging

// CallbackLogger is an implementation of Logger that forwards log messages to
// user-supplied callback functions.
type CallbackLogger struct {
	// LogTarget is the target for non-debug messages.
	LogTarget Callback

	// DebugTarget is the target for debug messages.
	// If it is nil no debug logging will occur and IsDebug() returns false.
	DebugTarget Callback
}

// Callback is the function signature for Printf-style callbacks used by
// CallbackLogger.
type Callback func(f string, v ...interface{})

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l *CallbackLogger) Log(f string, v ...interface{}) {
	l.LogTarget(f, v...)
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
// f is the format specifier, as per fmt.Printf(), etc.
func (l *CallbackLogger) Debug(f string, v ...interface{}) {
	if l.DebugTarget != nil {
		l.DebugTarget(f, v...)
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
