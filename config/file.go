package config

import (
	"io"
	"io/ioutil"
	"os"
)

// File returns a configuration value that is specified as a path to a file.
func File(p string) Value {
	return Value{&fileSource{path: p}}
}

// stringSource is an implementation of the source interface for configuration
// values that are specified as a path to a file.
type fileSource struct {
	path string
}

func (s *fileSource) AsReader() (io.ReadCloser, error) {
	return os.Open(s.path)
}

func (s *fileSource) AsPath() (string, io.Closer, error) {
	return s.path, &closer{done: 1}, nil
}

func (s *fileSource) AsString() (string, error) {
	buf, err := ioutil.ReadFile(s.path)
	return string(buf), err
}

func (s *fileSource) AsBytes() ([]byte, error) {
	return ioutil.ReadFile(s.path)
}
