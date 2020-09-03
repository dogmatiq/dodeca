package logging

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
//
// If l is nil, DefaultLogger is used.
func Log(l Logger, f string, v ...interface{}) {
	getLogger(l).Log(f, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// If l is nil, DefaultLogger is used.
func LogString(l Logger, s string) {
	getLogger(l).LogString(s)
}

// Debug writes a debug log message formatted according to a format
// specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
//
// If l is nil, DefaultLogger is used.
func Debug(l Logger, f string, v ...interface{}) {
	getLogger(l).Debug(f, v...)
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// If l is nil, DefaultLogger is used.
func DebugString(l Logger, s string) {
	getLogger(l).DebugString(s)
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
//
// If l is nil, DefaultLogger is used.
func IsDebug(l Logger) bool {
	return getLogger(l).IsDebug()
}

// getLogger returns l, or DefaultLogger if l is nil.
func getLogger(l Logger) Logger {
	if l != nil {
		return l
	}

	return DefaultLogger
}
