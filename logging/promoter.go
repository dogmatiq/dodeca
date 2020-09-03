package logging

// Promoter is an implementation of Logger that forwards all messages to a
// target logger as non-debug messages. Thus, it "promotes" debug messages to
// the non-debug level.
type Promoter struct {
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
func (l *Promoter) Log(f string, v ...interface{}) {
	Log(l.Target, f, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l *Promoter) LogString(s string) {
	LogString(l.Target, s)
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (l *Promoter) Debug(f string, v ...interface{}) {
	l.Log(f, v...)
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l *Promoter) DebugString(s string) {
	l.LogString(s)
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (l *Promoter) IsDebug() bool {
	return true
}
