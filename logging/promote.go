package logging

// Promote returns a Logger that forwards all messages to a target logger as
// non-debug messages. Thus, it "promotes" debug messages to the non-debug
// level.
//
// If the target is nil, DefaultLogger is used.
func Promote(target Logger) Logger {
	return promoter{target}
}

// promoter is an implementation of Logger that forwards all messages to a
// target logger as non-debug messages. Thus, it "promotes" debug messages to
// the non-debug level.
type promoter struct {
	// Target is the logger to which messages are forwarded.
	Target Logger
}

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l promoter) Log(f string, v ...interface{}) {
	Log(l.Target, f, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l promoter) LogString(s string) {
	LogString(l.Target, s)
}

// Debug writes an APPLICATION (non-debug) log message formatted according to a
// format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l promoter) Debug(f string, v ...interface{}) {
	l.Log(f, v...)
}

// DebugString writes a pre-formatted APPLICATION (non-debug) log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l promoter) DebugString(s string) {
	l.LogString(s)
}

// IsDebug returns true if this logger will perform debug logging.
//
// Because debug messages are "promoted" to the non-debug level, this
// implementation always returns true.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l promoter) IsDebug() bool {
	return true
}

// UnwrapLogger returns the logger wrapped by this logger.
//
// If ok is true it means that this logger is wrapping another logger, even
// if that logger is nil, indicating that DefaultLogger should be used.
//
// If ok is false it means that this logger is not wrapping another logger.
func (l promoter) UnwrapLogger() (Logger, bool) {
	return l.Target, true
}
