package config

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// MinInt is the minimum value that can be expressed using the int type.
	MinInt = -MaxInt - 1 // -1 << 31 or -1 << 63

	// MaxInt is the maximum value that can be expressed using the int type.
	MaxInt = 1<<(uintBitSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
)

// AsInt returns the int representation of the value associated with k or panics
// if unable to do so.
func AsInt(b Bucket, k string) int {
	return int(asInt(b, k, 0, MinInt, MaxInt))
}

// AsIntDefault returns the int representation of the value associated with k,
// or the default value v if k is undefined.
func AsIntDefault(b Bucket, k string, v int) int {
	return int(asIntDefault(b, k, 0, int64(v), MinInt, MaxInt))
}

// AsIntBetween returns the int representation of the value associated with k or
// panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsIntBetween(b Bucket, k string, min, max int) int {
	return int(asInt(b, k, 0, int64(min), int64(max)))
}

// AsIntDefaultBetween returns the int representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsIntDefaultBetween(b Bucket, k string, v, min, max int) int {
	return int(asIntDefault(b, k, 0, int64(v), int64(min), int64(max)))
}

// AsInt8 returns the int8 representation of the value associated with k or
// panics if unable to do so.
func AsInt8(b Bucket, k string) int8 {
	return int8(asInt(b, k, 8, math.MinInt8, math.MaxInt8))
}

// AsInt8Default returns the int8 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt8Default(b Bucket, k string, v int8) int8 {
	return int8(asIntDefault(b, k, 8, int64(v), math.MinInt8, math.MaxInt8))
}

// AsInt8Between returns the int8 representation of the value associated with k
// or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsInt8Between(b Bucket, k string, min, max int8) int8 {
	return int8(asInt(b, k, 8, int64(min), int64(max)))
}

// AsInt8DefaultBetween returns the int8 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsInt8DefaultBetween(b Bucket, k string, v, min, max int8) int8 {
	return int8(asIntDefault(b, k, 8, int64(v), int64(min), int64(max)))
}

// AsInt16 returns the int16 representation of the value associated with k or
// panics if unable to do so.
func AsInt16(b Bucket, k string) int16 {
	return int16(asInt(b, k, 16, math.MinInt16, math.MaxInt16))
}

// AsInt16Default returns the int16 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt16Default(b Bucket, k string, v int16) int16 {
	return int16(asIntDefault(b, k, 16, int64(v), math.MinInt16, math.MaxInt16))
}

// AsInt16Between returns the int16 representation of the value associated with k
// or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsInt16Between(b Bucket, k string, min, max int16) int16 {
	return int16(asInt(b, k, 16, int64(min), int64(max)))
}

// AsInt16DefaultBetween returns the int16 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsInt16DefaultBetween(b Bucket, k string, v, min, max int16) int16 {
	return int16(asIntDefault(b, k, 16, int64(v), int64(min), int64(max)))
}

// AsInt32 returns the int32 representation of the value associated with k or
// panics if unable to do so.
func AsInt32(b Bucket, k string) int32 {
	return int32(asInt(b, k, 32, math.MinInt32, math.MaxInt32))
}

// AsInt32Default returns the int32 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt32Default(b Bucket, k string, v int32) int32 {
	return int32(asIntDefault(b, k, 32, int64(v), math.MinInt32, math.MaxInt32))
}

// AsInt32Between returns the int32 representation of the value associated with k
// or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsInt32Between(b Bucket, k string, min, max int32) int32 {
	return int32(asInt(b, k, 32, int64(min), int64(max)))
}

// AsInt32DefaultBetween returns the int32 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsInt32DefaultBetween(b Bucket, k string, v, min, max int32) int32 {
	return int32(asIntDefault(b, k, 32, int64(v), int64(min), int64(max)))
}

// AsInt64 returns the int64 representation of the value associated with k or
// panics if unable to do so.
func AsInt64(b Bucket, k string) int64 {
	return asInt(b, k, 64, math.MinInt64, math.MaxInt64)
}

// AsInt64Default returns the int64 representation of the value associated with k,
// or the default value v if k is undefined.
func AsInt64Default(b Bucket, k string, v int64) int64 {
	return asIntDefault(b, k, 64, v, math.MinInt64, math.MaxInt64)
}

// AsInt64Between returns the int64 representation of the value associated with k
// or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsInt64Between(b Bucket, k string, min, max int64) int64 {
	return asInt(b, k, 64, min, max)
}

// AsInt64DefaultBetween returns the int64 representation of the value associated with
// k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsInt64DefaultBetween(b Bucket, k string, v, min, max int64) int64 {
	return asIntDefault(b, k, 64, v, min, max)
}

func tryAsInt(
	b Bucket,
	k string,
	bitSize int,
	min, max int64,
) (int64, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false
	}

	v, err := strconv.ParseInt(
		mustAsString(k, x),
		10,
		bitSize,
	)
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

	if min > v || v > max {
		panic(fmt.Sprintf(
			`expected %s to be between %d and %d (inclusive), got %d`,
			k,
			min,
			max,
			v,
		))
	}

	return v, true
}

func asInt(
	b Bucket,
	k string,
	bitSize int,
	min, max int64,
) int64 {
	if v, ok := tryAsInt(b, k, bitSize, min, max); ok {
		return v
	}

	panic(NotDefined{k})
}

func asIntDefault(
	b Bucket,
	k string,
	bitSize int,
	d, min, max int64,
) int64 {
	if min > d || d > max {
		panic(fmt.Sprintf(
			`expected the default value for %s to be between %d and %d (inclusive), got %d`,
			k,
			min,
			max,
			d,
		))
	}

	if v, ok := tryAsInt(b, k, bitSize, min, max); ok {
		return v
	}

	return d
}
