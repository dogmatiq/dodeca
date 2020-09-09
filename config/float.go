package config

import (
	"fmt"
	"strconv"
)

// GetFloat64 returns the float64 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as a float64, err is a
// non-nil error describing the invalid value.
func GetFloat64(b Bucket, k string) (v float64, ok bool, err error) {
	return getFloat(b, k, 64)
}

// GetFloat64Default returns the float64 representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as a float64, it returns an
// error describing the invalid value.
func GetFloat64Default(b Bucket, k string, v float64) (float64, error) {
	x, ok, err := GetFloat64(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetFloat64 returns the float64 representation of the value associated
// with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as a float64.
func MustGetFloat64(b Bucket, k string) (v float64, ok bool) {
	v, ok, err := GetFloat64(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetFloat64Default returns the float64 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as a float64.
func MustGetFloat64Default(b Bucket, k string, v float64) float64 {
	if x, ok := MustGetFloat64(b, k); ok {
		return x
	}

	return v
}

// getFloat returns the floating-point representation of the value associated
// with k.
func getFloat(b Bucket, k string, bitSize int) (float64, bool, error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid %d-bit float: %w`,
			k,
			bitSize,
			err,
		)
	}

	return v, true, nil
}
