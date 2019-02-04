package config

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Bytes returns a configuration value that is specified as a byte-slice.
func Bytes(v []byte) Value {
	return Value{&bytesSource{value: v}}
}

// bytesSource is an implementation of the source interface for configuration
// values that are specified as a byte-slice.
type bytesSource struct {
	value []byte
	temp  tempfile
}

func (s *bytesSource) AsReader() (io.ReadCloser, error) {
	return ioutil.NopCloser(
		bytes.NewBuffer(s.value),
	), nil
}

func (s *bytesSource) AsPath() (string, io.Closer, error) {
	fn := func(w io.Writer) error {
		_, err := w.Write(s.value)
		return err
	}

	path, err := s.temp.addRef(fn)
	if err != nil {
		return "", nil, err
	}

	return path, &closer{fn: s.temp.decRef}, nil
}

func (s *bytesSource) AsString() (string, error) {
	return string(s.value), nil
}

func (s *bytesSource) AsBytes() ([]byte, error) {
	return s.value, nil
}
