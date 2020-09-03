package logging

// Demote returns a logger that forwards all messages to the target logger as
// debug messages. Thus, it "demotes" non-debug messages to the debug level.
func Demote(target Logger) Logger {
	return demoter{target}
}

// demoter is an implementation of Logger that forwards all messages to a target
// logger as debug messages. Thus, it "demotes" non-debug messages to the debug
// level.
type demoter struct {
	// Target is the logger to which messages are forwarded.
	Target Logger
}

// Log writes a DEBUG log message formatted according to a format specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l demoter) Log(f string, v ...interface{}) {
	l.Debug(f, v...)
}

// LogString writes a pre-formatted DEBUG log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l demoter) LogString(s string) {
	l.DebugString(s)
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l demoter) Debug(f string, v ...interface{}) {
	Debug(l.Target, f, v...)
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l demoter) DebugString(s string) {
	DebugString(l.Target, s)
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l demoter) IsDebug() bool {
	return l.Target.IsDebug()
}

// UnwrapLogger returns the logger wrapped by this logger.
//
// If ok is true it means that this logger is wrapping another logger, even
// if that logger is nil, indicating that DefaultLogger should be used.
//
// If ok is false it means that this logger is not wrapping another logger.
func (l demoter) UnwrapLogger() (Logger, bool) {
	return l.Target, true
}
