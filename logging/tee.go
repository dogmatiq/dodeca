package logging

// Tee returns a logger that forwards all messages to multiple target loggers.
//
// It panics if no targets are provided, ensuring that log messages are not lost
// due to misconfiguration.
//
// The target list is de-duplicated, ensuring that log messages are not
// repeatedly sent to the same logger. Two loggers are considered equal if they
// compare the same using a regular shallow interface comparison via ==.
//
// Although the returned logger "wraps" the target loggers, it does not
// implement Wrapper, as that interface does not support multiple loggers.
func Tee(targets ...Logger) Logger {
	if len(targets) == 0 {
		panic("at least one target logger must be provided")
	}

	d := duplicator{}

next:
	for _, t := range targets {
		for _, x := range d.Targets {
			if x == t {
				continue next
			}
		}

		d.Targets = append(d.Targets, t)

		if t.IsDebug() {
			d.CaptureDebug = true
		}
	}

	return d
}

// duplicator is a logger that forwards log messages to multiple loggers.
type duplicator struct {
	Targets      []Logger
	CaptureDebug bool
}

// Log writes an application log message formatted according to a format
// specifier.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (d duplicator) Log(f string, v ...interface{}) {
	for _, t := range d.Targets {
		Log(t, f, v...)
	}
}

// LogString writes a pre-formatted application log message.
//
// It should be used for messages that are intended for people responsible for
// operating the application, such as the end-user or operations staff.
func (d duplicator) LogString(s string) {
	for _, t := range d.Targets {
		LogString(t, s)
	}
}

// Debug writes a debug log message formatted according to a format specifier.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
//
// f is the format specifier, as per fmt.Printf(), etc.
func (d duplicator) Debug(f string, v ...interface{}) {
	for _, t := range d.Targets {
		Debug(t, f, v...)
	}
}

// DebugString writes a pre-formatted debug log message.
//
// If IsDebug() returns false, no logging is performed.
//
// It should be used for messages that are intended for the software developers
// that maintain the application.
func (d duplicator) DebugString(s string) {
	for _, t := range d.Targets {
		DebugString(t, s)
	}
}

// IsDebug returns true if this logger will perform debug logging.
//
// Generally the application should just call Debug() or DebugString() without
// calling IsDebug(), however it can be used to check if debug logging is
// necessary before executing expensive code that is only used to obtain debug
// information.
func (d duplicator) IsDebug() bool {
	return d.CaptureDebug
}
