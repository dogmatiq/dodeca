package config

import (
	"fmt"
	"time"
)

// GetDuration returns the time.Duration representation of the value associated
// with k.
//
// Durations are specified using the syntax supported by time.ParseDuration.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value cannot be parsed as a duration, err is a
// non-nil error describing the invalid value.
func GetDuration(b Bucket, k string) (v time.Duration, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return 0, false, err
	}

	v, err = time.ParseDuration(s)
	if err != nil {
		return 0, false, fmt.Errorf(
			`%s is not a valid duration: %w`,
			k,
			err,
		)
	}

	return v, true, nil
}

// GetDurationDefault returns the time.Duration representation of the value
// associated with k, or the default value v if k is undefined.
//
// Durations are specified using the syntax supported by time.ParseDuration.
//
// If k is defined but its value cannot be parsed as a duration, it returns an
// error describing the invalid value.
func GetDurationDefault(b Bucket, k string, v time.Duration) (time.Duration, error) {
	x, ok, err := GetDuration(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetDuration returns the time.Duration representation of the value
// associated with k.
//
// Durations are specified using the syntax supported by time.ParseDuration.
//
// If k is undefined, ok is false and err is nil.
//
// It panics if k is defined but its value cannot be parsed as a duration.
func MustGetDuration(b Bucket, k string) (time.Duration, bool) {
	v, ok, err := GetDuration(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetDurationDefault returns the time.Duration representation of the value
// associated with k, or the default value v if k is undefined.
//
// Durations are specified using the syntax supported by time.ParseDuration.
//
// It panics if k is defined but its value cannot be parsed as a duration.
func MustGetDurationDefault(b Bucket, k string, v time.Duration) time.Duration {
	if x, ok := MustGetDuration(b, k); ok {
		return x
	}

	return v
}
