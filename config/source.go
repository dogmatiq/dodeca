package config

import "io"

// source is an interface for the underlying implementation of a Value.
type source interface {
	AsReader() (io.ReadCloser, error)
	AsPath() (string, io.Closer, error)
	AsString() (string, error)
	AsBytes() ([]byte, error)
}
