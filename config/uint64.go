package config

// GetUint64 returns the uint64 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value cannot be parsed as an uint64, err is a
// non-nil error describing the invalid value.
func GetUint64(b Bucket, k string) (v uint64, ok bool, err error) {
	v, ok, err = getUint(b, k, 64)
	return v, ok, err
}

// GetUint64Default returns the uint64 representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value cannot be parsed as an uint64, it returns an
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
// It panics if k is defined but its value cannot be parsed as an uint64.
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
// It panics if k is defined but its value cannot be parsed as an uint64.
func MustGetUint64Default(b Bucket, k string, v uint64) uint64 {
	if x, ok := MustGetUint64(b, k); ok {
		return x
	}

	return v
}
