package config

import (
	"fmt"
	"strconv"
)

// GetUint32 returns the uint32 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an uint32, err is a
// non-nil error describing the invalid value.
func GetUint32(b Bucket, k string) (v uint32, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid unsigned 32-bit integer: %w`,
			k,
			err,
		)
	}

	return uint32(v64), true, nil
}

// GetUint32Default returns the uint32 representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an uint32, it returns an
// error describing the invalid value.
func GetUint32Default(b Bucket, k string, v uint32) (uint32, error) {
	x, ok, err := GetUint32(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetUint32 returns the uint32 representation of the value associated with
// k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an uint32.
func MustGetUint32(b Bucket, k string) (v uint32, ok bool) {
	v, ok, err := GetUint32(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetUint32Default returns the uint32 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an uint32.
func MustGetUint32Default(b Bucket, k string, v uint32) uint32 {
	if x, ok := MustGetUint32(b, k); ok {
		return x
	}

	return v
}
