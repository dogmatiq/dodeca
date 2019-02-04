package config

import (
	"io"
	"os"
)

// Value is a configuration value.
type Value struct {
	src source
}

// IsEmpty returns true if this value is empty or undefined.
func (v *Value) IsEmpty() bool {
	return v.src == nil
}

// AsReader returns an io.ReadCloser that produces the configuration value.
//
// If v is the zero-value, it returns an os.ErrNotExist error.
func (v Value) AsReader() (io.ReadCloser, error) {
	if v.src == nil {
		return nil, os.ErrNotExist
	}

	return v.src.AsReader()
}

// AsPath returns the path to a real file on disk that contains the
// configuration value.
//
// If the configuration value was originally specified as a file, this will be
// the path to the original file. Otherwise, the path may be to a temporary
// file.
//
// This method should be used when some existing code requires a path to a file.
// Otherwise, it is preferable to use AsReader(), AsString() or AsBytes().
//
// It returns an io.Closer that must be closed when the file is no longer
// needed, regardless of how the configuration was specified.
//
// If v is the zero-value, it returns an os.ErrNotExist error.
func (v Value) AsPath() (string, io.Closer, error) {
	if v.src == nil {
		return "", nil, os.ErrNotExist
	}

	return v.src.AsPath()
}

// AsString returns the configuration value as a string.
//
// If v is the zero-value, it returns an empty string.
func (v Value) AsString() (string, error) {
	if v.src == nil {
		return "", nil
	}

	return v.src.AsString()
}

// AsBytes returns the configuration value as a byte-slice.
//
// If v is the zero-value, it returns a nil slice.
func (v Value) AsBytes() ([]byte, error) {
	if v.src == nil {
		return nil, nil
	}

	return v.src.AsBytes()
}

// String returns the value as a string, or panics if unable to do so.
//
// If v is the zero-value, it returns an empty string.
func (v Value) String() string {
	s, err := v.AsString()
	if err != nil {
		panic(err)
	}

	return s
}
