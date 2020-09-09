package config

// GetUint8 returns the uint8 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value cannot be parsed as an uint8, err is a non-nil
// error describing the invalid value.
func GetUint8(b Bucket, k string) (v uint8, ok bool, err error) {
	v64, ok, err := getUint(b, k, 8)
	return uint8(v64), ok, err
}

// GetUint8Default returns the uint8 representation of the value associated with
// k, or the default value v if k is undefined.
//
// If k is defined but its value cannot be parsed as an uint8, it returns an
// error describing the invalid value.
func GetUint8Default(b Bucket, k string, v uint8) (uint8, error) {
	x, ok, err := GetUint8(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetUint8 returns the uint8 representation of the value associated with k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value cannot be parsed as an uint8.
func MustGetUint8(b Bucket, k string) (v uint8, ok bool) {
	v, ok, err := GetUint8(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetUint8Default returns the uint8 representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value cannot be parsed as an uint8.
func MustGetUint8Default(b Bucket, k string, v uint8) uint8 {
	if x, ok := MustGetUint8(b, k); ok {
		return x
	}

	return v
}