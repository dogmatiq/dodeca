package logging

// DiscardLogger is a logger that produces no output.
type DiscardLogger struct{}

// Log is a no-op.
func (DiscardLogger) Log(f string, v ...interface{}) {}

// LogString is a no-op.
func (DiscardLogger) LogString(s string) {}

// Debug is a no-op.
func (DiscardLogger) Debug(f string, v ...interface{}) {}

// DebugString is a no-op.
func (DiscardLogger) DebugString(s string) {}

// IsDebug always returns false.
func (DiscardLogger) IsDebug() bool {
	return false
}
