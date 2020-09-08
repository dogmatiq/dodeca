package config

import (
	"fmt"
	"strings"
)

// GetBoolT returns the value associated with the given key as a boolean.
//
// If the value is "true", "yes", or "on", it returns true. If the value is
// "false", "no", or "off", it returns false. These string comparisons are
// case-insensitive. It returns an error if any other value is used.
//
// If they key is not defined, it returns true.
func GetBoolT(b Bucket, k string) (bool, error) {
	return GetBoolDefault(b, k, true)
}

// GetBoolF returns the value associated with the given key as a boolean.
//
// If the value is "true", "yes", or "on", it returns true. If the value is
// "false", "no", or "off", it returns false. These string comparisons are
// case-insensitive. It returns an error if any other value is used.
//
// If they key is not defined, it returns false.
func GetBoolF(b Bucket, k string) (bool, error) {
	return GetBoolDefault(b, k, false)
}

// GetBoolDefault returns the value associated with the given key as a
// boolean.
//
// If the value is "true", "yes", or "on", it returns true. If the value is
// "false", "no", or "off", it returns false. These string comparisons are
// case-insensitive. It returns an error if any other value is used.
//
// If they key is not defined, it returns v.
func GetBoolDefault(b Bucket, k string, v bool) (bool, error) {
	x := b.Get(k)

	if x.IsZero() {
		return v, nil
	}

	s, err := x.AsString()
	if err != nil {
		return false, err
	}

	switch strings.ToLower(s) {
	case "true", "yes", "on":
		return true, nil
	case "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("%s is not a valid boolean value: %s", k, s)
	}
}

// MustGetBoolT returns the value associated with the given key as a boolean.
//
// If the value is "true", "yes", or "on", it returns true. If the value is
// "false", "no", or "off", it returns false. These string comparisons are
// case-insensitive. It panics if any other value is used.
//
// If they key is not defined, it returns true.
func MustGetBoolT(b Bucket, k string) bool {
	return MustGetBoolDefault(b, k, true)
}

// MustGetBoolF returns the value associated with the given key as a boolean.
//
// If the value is "true", "yes", or "on", it returns true. If the value is
// "false", "no", or "off", it returns false. These string comparisons are
// case-insensitive. It panics if any other value is used.
//
// If they key is not defined, it returns false.
func MustGetBoolF(b Bucket, k string) bool {
	return MustGetBoolDefault(b, k, false)
}

// MustGetBoolDefault returns the value associated with the given key as a
// boolean.
//
// If the value is "true", "yes", or "on", it returns true. If the value is
// "false", "no", or "off", it returns false. These string comparisons are
// case-insensitive. It panics if any other value is used.
//
// If they key is not defined, it returns v.
func MustGetBoolDefault(b Bucket, k string, v bool) bool {
	v, err := GetBoolDefault(b, k, v)
	if err != nil {
		panic(err)
	}

	return v
}
