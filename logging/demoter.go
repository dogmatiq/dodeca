package logging

// Demoter is an implementation of Logger that forwards all messages to a target
// logger as debug messages. Thus, it "demotes" non-debug messages to the debug
// level.
type Demoter struct {
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
func (l *Demoter) Log(f string, v ...interface{}) {
	l.Debug(f, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l *Demoter) LogString(s string) {
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
func (l *Demoter) Debug(f string, v ...interface{}) {
	Debug(l.Target, f, v...)
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l *Demoter) DebugString(s string) {
	DebugString(l.Target, s)
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l *Demoter) IsDebug() bool {
	return l.Target.IsDebug()
}
