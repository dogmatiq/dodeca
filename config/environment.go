package config

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Environment is a Bucket that produces configuration values from the operating
// system's environment variables.
//
// For any given environment variable K, the environment variable K__DATASOURCE
// indicates how the content of K should be interpreted.
//
// If K__DATASOURCE is:
//
// - empty, undefined or the value "string:plain", then the content of K is
// treated as a standard environment variable
//
// - the value "string:hex", then the content of K is treated as a binary value,
// encoded as a hexadecimal string
//
// - the value "string:base64", then the content of K is treated as a binary
// value, encoded as a standard base-64 string
//
// - the value "file", then the content of K is treated as a path to a file
// containing the value
var Environment Bucket = environment{}

type environment struct{}

// Get returns the value associated with the given key.
//
// If they key is not defined, it returns an empty value.
func (environment) Get(k string) Value {
	if strings.HasSuffix(k, suffix) {
		// never return the value of "source type" variables, they are meta-data
		// about configuration values, not configuration values themselves.
		return Value{}
	}

	return getenv(k)
}

// GetDefault returns the value associated with the given key.
//
// If the key is not defined, it returns a value with the content of v.
func (environment) GetDefault(k string, v string) Value {
	panic("<ni>")
}

// Each calls fn for each key/value pair in the bucket.
//
// If fn returns false, iteration is stopped.
//
// Each returns true if iteration completes fully, or false if fn()
// returns false.
func (environment) Each(fn EachFunc) bool {
	for _, str := range os.Environ() {
		pair := strings.SplitN(str, "=", 2)

		k := pair[0]

		if strings.HasSuffix(k, suffix) {
			continue
		}

		var v Value

		if len(pair) == 2 {
			v = getenv(k)
		}

		if !fn(k, v) {
			return false
		}
	}

	return true
}

// suffix is the suffix used to identify environment variables that specify the
// "source type" of the environment variable without this suffix.
const suffix = "__DATASOURCE"

const (
	sourceStringPlain  = "string:plain"
	sourceStringHex    = "string:hex"
	sourceStringBase64 = "string:base64"
	sourceFile         = "file"
)

// getenv returns the Value for the environment variable named k.
func getenv(k string) Value {
	raw := os.Getenv(k)

	if raw == "" {
		return Value{}
	}

	src := os.Getenv(k + suffix)

	switch src {
	case "", sourceStringPlain:
		return String(raw)

	case sourceStringHex:
		buf, err := hex.DecodeString(raw)
		if err != nil {
			return fail(err)
		}
		return Bytes(buf)

	case sourceStringBase64:
		buf, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return fail(err)
		}
		return Bytes(buf)

	case sourceFile:
		return File(raw)

	default:
		return fail(
			fmt.Errorf("unrecognised environment variable source type: %s", src),
		)
	}
}
