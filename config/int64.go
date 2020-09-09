package config

// GetInt64 returns the int64 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value cannot be parsed as an int64, err is a non-nil
// error describing the invalid value.
func GetInt64(b Bucket, k string) (v int64, ok bool, err error) {
	return getInt(b, k, 64)
}

// GetInt64Default returns the int64 representation of the value associated with
// k, or the default value v if k is undefined.
//
// If k is defined but its value cannot be parsed as an int64, it returns an
// error describing the invalid value.
func GetInt64Default(b Bucket, k string, v int64) (int64, error) {
	x, ok, err := GetInt64(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetInt64 returns the int64 representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value cannot be parsed as an int64.
func MustGetInt64(b Bucket, k string) (v int64, ok bool) {
	v, ok, err := GetInt64(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetInt64Default returns the int64 representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value cannot be parsed as an int64.
func MustGetInt64Default(b Bucket, k string, v int64) int64 {
	if x, ok := MustGetInt64(b, k); ok {
		return x
	}

	return v
}
