package config

import (
	"fmt"
	"strconv"
)

// GetUint64 returns the uint64 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an uint64, err is a
// non-nil error describing the invalid value.
func GetUint64(b Bucket, k string) (v uint64, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v, err = strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid unsigned 64-bit integer: %w`,
			k,
			err,
		)
	}

	return v, true, nil
}

// GetUint64Default returns the uint64 representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an uint64, it returns an
// error describing the invalid value.
func GetUint64Default(b Bucket, k string, v uint64) (uint64, error) {
	x, ok, err := GetUint64(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetUint64 returns the uint64 representation of the value associated with
// k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an uint64.
func MustGetUint64(b Bucket, k string) (v uint64, ok bool) {
	v, ok, err := GetUint64(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetUint64Default returns the uint64 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an uint64.
func MustGetUint64Default(b Bucket, k string, v uint64) uint64 {
	if x, ok := MustGetUint64(b, k); ok {
		return x
	}

	return v
}
