package config

import (
	"fmt"
	"strconv"
)

// AsInt returns the int representation of the value associated with k or panics
// if unable to do so.
func AsInt(b Bucket, k string) int {
	return int(asInt(b, k, 0, false))
}

// AsIntDefault returns the int representation of the value associated with k,
// or the default value v if k is undefined.
func AsIntDefault(b Bucket, k string, v int) int {
	return int(asIntDefault(b, k, 0, int64(v), false))
}

// AsIntP returns the int representation of the value associated with k or panics
// if unable to do so.
//
// It panics if the value is not zero or negative.
func AsIntP(b Bucket, k string) int {
	return int(asInt(b, k, 0, true))
}

// AsIntPDefault returns the int representation of the value associated with k,
// or the default value v if k is undefined.
//
// It panics if the value is not zero or negative.
func AsIntPDefault(b Bucket, k string, v int) int {
	return int(asIntDefault(b, k, 0, int64(v), true))
}

// AsInt8 returns the int8 representation of the value associated with k or
// panics if unable to do so.
func AsInt8(b Bucket, k string) int8 {
	return int8(asInt(b, k, 8, false))
}

// AsInt8Default returns the int8 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt8Default(b Bucket, k string, v int8) int8 {
	return int8(asIntDefault(b, k, 8, int64(v), false))
}

// AsInt8P returns the int8 representation of the value associated with k or
// panics if unable to do so.
//
// It panics if the value is not zero or negative.
func AsInt8P(b Bucket, k string) int8 {
	return int8(asInt(b, k, 8, true))
}

// AsInt8PDefault returns the int8 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not zero or negative.
func AsInt8PDefault(b Bucket, k string, v int8) int8 {
	return int8(asIntDefault(b, k, 8, int64(v), true))
}

// AsInt16 returns the int16 representation of the value associated with k or
// panics if unable to do so.
func AsInt16(b Bucket, k string) int16 {
	return int16(asInt(b, k, 16, false))
}

// AsInt16Default returns the int16 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt16Default(b Bucket, k string, v int16) int16 {
	return int16(asIntDefault(b, k, 16, int64(v), false))
}

// AsInt16P returns the int16 representation of the value associated with k or
// panics if unable to do so.
//
// It panics if the value is not zero or negative.
func AsInt16P(b Bucket, k string) int16 {
	return int16(asInt(b, k, 16, true))
}

// AsInt16PDefault returns the int16 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not zero or negative.
func AsInt16PDefault(b Bucket, k string, v int16) int16 {
	return int16(asIntDefault(b, k, 16, int64(v), true))
}

// AsInt32 returns the int32 representation of the value associated with k or
// panics if unable to do so.
func AsInt32(b Bucket, k string) int32 {
	return int32(asInt(b, k, 32, false))
}

// AsInt32Default returns the int32 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt32Default(b Bucket, k string, v int32) int32 {
	return int32(asIntDefault(b, k, 32, int64(v), false))
}

// AsInt32P returns the int32 representation of the value associated with k or
// panics if unable to do so.
//
// It panics if the value is not zero or negative.
func AsInt32P(b Bucket, k string) int32 {
	return int32(asInt(b, k, 32, true))
}

// AsInt32PDefault returns the int32 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not zero or negative.
func AsInt32PDefault(b Bucket, k string, v int32) int32 {
	return int32(asIntDefault(b, k, 32, int64(v), true))
}

// AsInt64 returns the int64 representation of the value associated with k or
// panics if unable to do so.
func AsInt64(b Bucket, k string) int64 {
	return asInt(b, k, 64, false)
}

// AsInt64Default returns the int64 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt64Default(b Bucket, k string, v int64) int64 {
	return asIntDefault(b, k, 64, v, false)
}

// AsInt64P returns the int64 representation of the value associated with k or
// panics if unable to do so.
//
// It panics if the value is not zero or negative.
func AsInt64P(b Bucket, k string) int64 {
	return asInt(b, k, 64, true)
}

// AsInt64PDefault returns the int64 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not zero or negative.
func AsInt64PDefault(b Bucket, k string, v int64) int64 {
	return asIntDefault(b, k, 64, v, true)
}

func tryAsInt(
	b Bucket,
	k string,
	bitSize int,
	positiveOnly bool,
) (int64, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false
	}

	s, err := x.AsString()
	if err != nil {
		panic(fmt.Sprintf("cannot read %s: %s", k, err))
	}

	v, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		if bitSize == 0 {
			panic(fmt.Sprintf(
				`expected %s to be a signed integer: %s`,
				k,
				err,
			))
		}

		panic(fmt.Sprintf(
			`expected %s to be a signed %d-bit integer: %s`,
			k,
			bitSize,
			err,
		))
	}

	if positiveOnly && v <= 0 {
		panic(fmt.Sprintf(
			`expected %s to be positive, got %d`,
			k,
			v,
		))
	}

	return v, true
}

func asInt(
	b Bucket,
	k string,
	bitSize int,
	positiveOnly bool,
) int64 {
	if v, ok := tryAsInt(b, k, bitSize, positiveOnly); ok {
		return v
	}

	panic(fmt.Sprintf("%s is not defined", k))
}

func asIntDefault(
	b Bucket,
	k string,
	bitSize int,
	d int64,
	positiveOnly bool,
) int64 {
	if v, ok := tryAsInt(b, k, bitSize, positiveOnly); ok {
		return v
	}

	return d
}
