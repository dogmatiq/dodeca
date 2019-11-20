package logging

// TestLogger is an implementation of Logger that uses Go's standard logger
// package.
type TestLogger struct {
	// Target is a TestingT where all debug and log messages are written to
	Target TestingT
}

// TestingT is used to abstract over *testing.T and *testing.B
type TestingT interface {
	// Logf is used to log to the underlying *testing.T or *testing.B
	Logf(format string, args ...interface{})
}

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// fmt is the format specifier, as per fmt.Printf(), etc.
func (l *TestLogger) Log(fmt string, v ...interface{}) {
	l.Target.Logf(fmt, v...)
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (l *TestLogger) LogString(s string) {
	l.Target.Logf(s)
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// fmt is the format specifier, as per fmt.Printf(), etc.
func (l *TestLogger) Debug(fmt string, v ...interface{}) {
	l.Target.Logf(fmt, v...)
}

// DebugString writes a pre-formatted debug log message.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (l *TestLogger) DebugString(s string) {
	l.Target.Logf(s)
}

// IsDebug returns true
func (l *TestLogger) IsDebug() bool {
	return true
}
