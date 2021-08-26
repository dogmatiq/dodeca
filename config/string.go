package config

import (
	"fmt"
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

// AsString returns the string representation of the value associated with k or
// panics if unable to do so.
func AsString(b Bucket, k string) string {
	if v, ok := asString(b, k); ok {
		return v
	}

	panic(NotDefined{k})
}

// AsStringDefault returns the string representation of the value associated
// with k, or the default value v if k is undefined.
func AsStringDefault(b Bucket, k string, v string) string {
	if x, ok := asString(b, k); ok {
		return x
	}

	return v
}

func asString(b Bucket, k string) (string, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return "", false
	}

	return mustAsString(k, x), true
}

func mustAsString(k string, v Value) string {
	s, err := v.AsString()
	if err != nil {
		panic(fmt.Sprintf("cannot read %s: %s", k, err))
	}

	return s
}
