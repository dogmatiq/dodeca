package config

// GetInt32 returns the int32 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an int32, err is a non-nil
// error describing the invalid value.
func GetInt32(b Bucket, k string) (v int32, ok bool, err error) {
	v64, ok, err := getInt(b, k, 32)
	return int32(v64), ok, err
}

// GetInt32Default returns the int32 representation of the value associated with
// k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an int32, it returns an
// error describing the invalid value.
func GetInt32Default(b Bucket, k string, v int32) (int32, error) {
	x, ok, err := GetInt32(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetInt32 returns the int32 representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an int32.
func MustGetInt32(b Bucket, k string) (v int32, ok bool) {
	v, ok, err := GetInt32(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetInt32Default returns the int32 representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an int32.
func MustGetInt32Default(b Bucket, k string, v int32) int32 {
	if x, ok := MustGetInt32(b, k); ok {
		return x
	}

	return v
}
