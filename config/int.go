package config

import (
	"fmt"
	"strconv"
)

// GetInt64 returns the int64 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an int64, err is a
// non-nil error describing the invalid value.
func GetInt64(b Bucket, k string) (v int64, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid signed 64-bit integer: %w`,
			k,
			err,
		)
	}

	return v, true, nil
}

// MustGetInt64 returns the int64 representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an int64.
func MustGetInt64(b Bucket, k string) (v int64, ok bool) {
	v, ok, err := GetInt64(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}
