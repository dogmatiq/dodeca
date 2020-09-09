package config

// GetFloat32 returns the float32 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as a float32, err is a
// non-nil error describing the invalid value.
func GetFloat32(b Bucket, k string) (v float32, ok bool, err error) {
	v64, ok, err := getFloat(b, k, 32)
	return float32(v64), ok, err
}

// GetFloat32Default returns the float32 representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as a float32, it returns an
// error describing the invalid value.
func GetFloat32Default(b Bucket, k string, v float32) (float32, error) {
	x, ok, err := GetFloat32(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetFloat32 returns the float32 representation of the value associated
// with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as a float32.
func MustGetFloat32(b Bucket, k string) (v float32, ok bool) {
	v, ok, err := GetFloat32(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetFloat32Default returns the float32 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as a float32.
func MustGetFloat32Default(b Bucket, k string, v float32) float32 {
	if x, ok := MustGetFloat32(b, k); ok {
		return x
	}

	return v
}
