package config

import (
	"bytes"
	"fmt"
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

// AsBytes returns the byte-slice representation of the value associated with k
// or panics if unable to do so.
func AsBytes(b Bucket, k string) []byte {
	if v, ok := asBytes(b, k); ok {
		return v
	}

	panic(fmt.Sprintf("%s is not defined", k))
}

// AsBytesDefault returns the byte-slice representation of the value associated
// with k, or the default value v if k is undefined.
func AsBytesDefault(b Bucket, k string, v []byte) []byte {
	if x, ok := asBytes(b, k); ok {
		return x
	}

	return v
}

func asBytes(b Bucket, k string) ([]byte, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return nil, false
	}

	s, err := x.AsBytes()
	if err != nil {
		panic(fmt.Sprintf("%s can not be converted to a byte-slice: %s", k, err))
	}

	return s, true
}
