package config

// GetUint16 returns the uint16 representation of the value associated with k.
//
// If k is undefined, ok is false and err is nil.
//
// If k is defined but its value can not be parsed as an uint16, err is a
// non-nil error describing the invalid value.
func GetUint16(b Bucket, k string) (v uint16, ok bool, err error) {
	v64, ok, err := getUint(b, k, 16)
	return uint16(v64), ok, err
}

// GetUint16Default returns the uint16 representation of the value associated
// with k, or the default value v if k is undefined.
//
// If k is defined but its value can not be parsed as an uint16, it returns an
// error describing the invalid value.
func GetUint16Default(b Bucket, k string, v uint16) (uint16, error) {
	x, ok, err := GetUint16(b, k)
	if err != nil {
		return 0, err
	}

	if ok {
		return x, nil
	}

	return v, nil
}

// MustGetUint16 returns the uint16 representation of the value associated with
// k.
//
// If k is undefined, ok is false.
//
// It panics if k is defined but its value can not be parsed as an uint16.
func MustGetUint16(b Bucket, k string) (v uint16, ok bool) {
	v, ok, err := GetUint16(b, k)
	if err != nil {
		panic(err)
	}

	return v, ok
}

// MustGetUint16Default returns the uint16 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if k is defined but its value can not be parsed as an uint16.
func MustGetUint16Default(b Bucket, k string, v uint16) uint16 {
	if x, ok := MustGetUint16(b, k); ok {
		return x
	}

	return v
}
