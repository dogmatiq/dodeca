package config

import (
	"fmt"
	"strconv"
)

// GetInt8 returns the int8 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an int8, err is a
// non-nil error describing the invalid value.
func GetInt8(b Bucket, k string) (v int8, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v64, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid signed 8-bit integer: %w`,
			k,
			err,
		)
	}

	return int8(v64), true, nil
}

// GetInt8Default returns the int8 representation of the value associated with
// k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an int8, it returns an
// error describing the invalid value.
func GetInt8Default(b Bucket, k string, v int8) (int8, error) {
	x, ok, err := GetInt8(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetInt8 returns the int8 representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an int8.
func MustGetInt8(b Bucket, k string) (v int8, ok bool) {
	v, ok, err := GetInt8(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetInt8Default returns the int8 representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an int8.
func MustGetInt8Default(b Bucket, k string, v int8) int8 {
	if x, ok := MustGetInt8(b, k); ok {
		return x
	}

	return v
}
