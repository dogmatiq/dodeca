package config

import (
	"fmt"
	"strconv"
)

// GetUint returns the uint representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an uint, err is a
// non-nil error describing the invalid value.
func GetUint(b Bucket, k string) (v uint, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v64, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid unsigned integer: %w`,
			k,
			err,
		)
	}

	return uint(v64), true, nil
}

// GetUintDefault returns the uint representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an uint, it returns an
// error describing the invalid value.
func GetUintDefault(b Bucket, k string, v uint) (uint, error) {
	x, ok, err := GetUint(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetUint returns the uint representation of the value associated with
// k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an uint.
func MustGetUint(b Bucket, k string) (v uint, ok bool) {
	v, ok, err := GetUint(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetUintDefault returns the uint representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an uint.
func MustGetUintDefault(b Bucket, k string, v uint) uint {
	if x, ok := MustGetUint(b, k); ok {
		return x
	}

	return v
}
