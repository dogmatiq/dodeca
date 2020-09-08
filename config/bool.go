package config

import (
	"fmt"
	"strings"
)

// GetBool returns the boolean representation of the value associated with k.
//
// If the value is "true", "yes" or "on", v and ok are both true.
//
// If the value is "false", "no" or "off", v is false and ok is true.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value is not listed above, err is a non-nil error
// describing the invalid value.
func GetBool(b Bucket, k string) (v bool, ok bool, err error) {
	x := b.Get(k)

	if x.IsZero() {
		return false, false, nil
	}

	s, err := x.AsString()
	if err != nil {
		return false, false, err
	}

	switch strings.ToLower(s) {
	case "true", "yes", "on":
		return true, true, nil
	case "false", "no", "off":
		return false, true, nil
	default:
		return false, false, fmt.Errorf(
			`%s is not a boolean, got %#v, expected "true", "false", "yes", "no", "on" or "off"`,
			k,
			s,
		)
	}
}

// GetBoolT returns the boolean representation of the value associated with k,
// or true if k is undefined.
//
// If the value is "true", "yes" or "on", it returns true.
//
// If the value is "false", "no" or "off", it returns false.
//
// If k is defined but its value is not listed above, it returns an error
// describing the invalid value.
func GetBoolT(b Bucket, k string) (bool, error) {
	return GetBoolDefault(b, k, true)
}

// GetBoolF returns the boolean representation of the value associated with k,
// or false if k is undefined.
//
// If the value is "true", "yes" or "on", it returns true.
//
// If the value is "false", "no" or "off", it returns false.
//
// If k is defined but its value is not listed above, it returns an error
// describing the invalid value.
func GetBoolF(b Bucket, k string) (bool, error) {
	return GetBoolDefault(b, k, false)
}

// GetBoolDefault returns boolean representation of the value associated with k,
// or the default value v if k is undefined.
//
// If the value is "true", "yes" or "on", it returns true.
//
// If the value is "false", "no" or "off", it returns false.
//
// If k is defined but its value is not listed above, it returns an error
// describing the invalid value.
func GetBoolDefault(b Bucket, k string, v bool) (bool, error) {
	x, ok, err := GetBool(b, k)
	if err != nil {
		return false, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetBool returns the boolean representation of the value associated with
// k or panics if unable to do so.
//
// If the value is "true", "yes" or "on", v and ok are both true.
//
// If the value is "false", "no" or "off", v is false and ok is true.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value is not listed above.
func MustGetBool(b Bucket, k string) (v bool, ok bool) {
	v, ok, err := GetBool(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetBoolT returns the boolean representation of the value associated with
// k, or true if k is undefined.
//
// If the value is "true", "yes" or "on", it returns true.
//
// If the value is "false", "no" or "off", it returns false.
//
// It panics if k is defined but its value is not listed above.
func MustGetBoolT(b Bucket, k string) bool {
	return MustGetBoolDefault(b, k, true)
}

// MustGetBoolF returns the boolean representation of the value associated with
// k, or false if k is undefined.
//
// If the value is "true", "yes" or "on", it returns true.
//
// If the value is "false", "no" or "off", it returns false.
//
// It panics if k is defined but its value is not listed above.
func MustGetBoolF(b Bucket, k string) bool {
	return MustGetBoolDefault(b, k, false)
}

// MustGetBoolDefault returns boolean representation of the value associated
// with k, or the default value v if k is undefined.
//
// If the value is "true", "yes" or "on", it returns true.
//
// If the value is "false", "no" or "off", it returns false.
//
// It panics if k is defined but its value is not listed above.
func MustGetBoolDefault(b Bucket, k string, v bool) bool {
	if x, ok := MustGetBool(b, k); ok {
		return x
	}

	return v
}
