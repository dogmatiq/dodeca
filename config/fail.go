package config

import (
	"io"
)

// fail returns a configuration value that returns an error whenever it is
// consumed by any method.
func fail(err error) Value {
	return Value{failSource{err: err}}
}

// failSource is an implementation of the source interface that always returns
// an error.
type failSource struct {
	err error
}

func (s failSource) AsReader() (io.ReadCloser, error) {
	return nil, s.err
}

func (s failSource) AsPath() (string, io.Closer, error) {
	return "", nil, s.err
}

func (s failSource) AsString() (string, error) {
	return "", s.err
}

func (s failSource) AsBytes() ([]byte, error) {
	return nil, s.err
}
