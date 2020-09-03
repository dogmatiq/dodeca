package logging

// Wrapper is an interface for loggers that "wrap" some other logger.
type Wrapper interface {
	Logger

	// UnwrapLogger returns the logger wrapped by this logger.
	//
	// If ok is true it means that this logger is wrapping another logger, even
	// if that logger is nil, indicating that DefaultLogger should be used.
	//
	// If ok is false it means that this logger is not wrapping another logger.
	UnwrapLogger() (l Logger, ok bool)
}

// Unwrap returns the "inner-most" logger wrapped by l.
//
// If l is nil, or any of the wrapped loggers a nil, DefaultLogger is returned.
func Unwrap(l Logger) Logger {
	for {
		w, ok := l.(Wrapper)
		if !ok {
			break
		}

		next, ok := w.UnwrapLogger()
		if !ok {
			break
		}

		l = next
	}

	return getLogger(l)
}
