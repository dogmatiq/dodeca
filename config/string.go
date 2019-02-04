package config

import (
	"io"
	"io/ioutil"
	"strings"
)

// String returns a configuration value that is specified as a string.
func String(v string) Value {
	return Value{&stringSource{value: v}}
}

// stringSource is an implementation of the source interface for configuration
// values that are specified as a string.
type stringSource struct {
	value string
	temp  tempfile
}

func (s *stringSource) AsReader() (io.ReadCloser, error) {
	return ioutil.NopCloser(
		strings.NewReader(s.value),
	), nil
}

func (s *stringSource) AsPath() (string, io.Closer, error) {
	fn := func(w io.Writer) error {
		_, err := io.WriteString(w, s.value)
		return err
	}

	path, err := s.temp.addRef(fn)
	if err != nil {
		return "", nil, err
	}

	return path, &closer{fn: s.temp.decRef}, nil
}

func (s *stringSource) AsString() (string, error) {
	return s.value, nil
}

func (s *stringSource) AsBytes() ([]byte, error) {
	return []byte(s.value), nil
}
