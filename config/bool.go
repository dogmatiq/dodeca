package config

import (
	"strings"
)

// AsBool returns the boolean representation of the value associated with k or
// panics if unable to do so.
//
// It returns true if the value "true", "yes" or "on", or false if the value is
// "false", "no" or "off".
func AsBool(b Bucket, k string) bool {
	if v, ok := asBool(b, k); ok {
		return v
	}

	panic(NotDefined{k})
}

// AsBoolT returns the boolean representation of the value associated with k, or
// true if k is undefined.
//
// It returns true if the value "true", "yes" or "on", or false if the value is
// "false", "no" or "off".
func AsBoolT(b Bucket, k string) bool {
	return AsBoolDefault(b, k, true)
}

// AsBoolF returns the boolean representation of the value associated with k, or
// true if k is undefined.
//
// It returns true if the value "true", "yes" or "on", or false if the value is
// "false", "no" or "off".
func AsBoolF(b Bucket, k string) bool {
	return AsBoolDefault(b, k, false)
}

// AsBoolDefault returns the boolean representation of the value associated with
// k, or the default value v if k is undefined.
//
// It returns true if the value "true", "yes" or "on", or false if the value is
// "false", "no" or "off".
func AsBoolDefault(b Bucket, k string, v bool) bool {
	if v, ok := asBool(b, k); ok {
		return v
	}

	return v
}

func asBool(b Bucket, k string) (bool, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return false, false
	}

	s := mustAsString(k, x)

	switch strings.ToLower(s) {
	case "true", "yes", "on":
		return true, true
	case "false", "no", "off":
		return false, true
	default:
		panic(InvalidValue{
			k,
			s,
			`expected a boolean ("true", "false", "yes", "no", "on" or "off")`,
		})
	}
}
