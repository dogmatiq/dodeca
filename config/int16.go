package config

import (
	"fmt"
	"strconv"
)

// GetInt16 returns the int16 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an int16, err is a
// non-nil error describing the invalid value.
func GetInt16(b Bucket, k string) (v int16, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v64, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid signed 16-bit integer: %w`,
			k,
			err,
		)
	}

	return int16(v64), true, nil
}

// GetInt16Default returns the int16 representation of the value associated with
// k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an int16, it returns an
// error describing the invalid value.
func GetInt16Default(b Bucket, k string, v int16) (int16, error) {
	x, ok, err := GetInt16(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetInt16 returns the int16 representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an int16.
func MustGetInt16(b Bucket, k string) (v int16, ok bool) {
	v, ok, err := GetInt16(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetInt16Default returns the int16 representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an int16.
func MustGetInt16Default(b Bucket, k string, v int16) int16 {
	if x, ok := MustGetInt16(b, k); ok {
		return x
	}

	return v
}
