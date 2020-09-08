package config

import (
	"fmt"
	"strconv"
)

// GetInt returns the int representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an int, err is a
// non-nil error describing the invalid value.
func GetInt(b Bucket, k string) (v int, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v64, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid signed integer: %w`,
			k,
			err,
		)
	}

	return int(v64), true, nil
}

// GetIntDefault returns the int representation of the value associated with
// k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an int, it returns an
// error describing the invalid value.
func GetIntDefault(b Bucket, k string, v int) (int, error) {
	x, ok, err := GetInt(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetInt returns the int representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an int.
func MustGetInt(b Bucket, k string) (v int, ok bool) {
	v, ok, err := GetInt(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetIntDefault returns the int representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an int.
func MustGetIntDefault(b Bucket, k string, v int) int {
	if x, ok := MustGetInt(b, k); ok {
		return x
	}

	return v
}
