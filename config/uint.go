package config

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// MaxUint is the maximum value that can be expressed using the uint type.
	MaxUint = 1<<uintBitSize - 1 // 1<<32 - 1 or 1<<64 - 1

	uintBitSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
)

// AsUint returns the uint representation of the value associated with k or
// panics if unable to do so.
func AsUint(b Bucket, k string) uint {
	return uint(asUint(b, k, 0, 0, MaxUint))
}

// AsUintDefault returns the uint representation of the value associated with k,
// or the default value v if k is undefined.
func AsUintDefault(b Bucket, k string, v uint) uint {
	return uint(asUintDefault(b, k, 0, uint64(v), 0, MaxUint))
}

// AsUintBetween returns the uint representation of the value associated with k
// or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsUintBetween(b Bucket, k string, min, max int) uint {
	return uint(asUint(b, k, 0, uint64(min), uint64(max)))
}

// AsUintDefaultBetween returns the uint representation of the value associated
// with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsUintDefaultBetween(b Bucket, k string, v, min, max int) uint {
	return uint(asUintDefault(b, k, 0, uint64(v), uint64(min), uint64(max)))
}

// AsUint8 returns the uint8 representation of the value associated with k or
// panics if unable to do so.
func AsUint8(b Bucket, k string) uint8 {
	return uint8(asUint(b, k, 8, 0, math.MaxUint8))
}

// AsUint8Default returns the uint8 representation of the value associated with
// k, or the default value v if k is undefined.
func AsUint8Default(b Bucket, k string, v uint8) uint8 {
	return uint8(asUintDefault(b, k, 8, uint64(v), 0, math.MaxUint8))
}

// AsUint8Between returns the uint8 representation of the value associated with
// k or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsUint8Between(b Bucket, k string, min, max int8) uint8 {
	return uint8(asUint(b, k, 8, uint64(min), uint64(max)))
}

// AsUint8DefaultBetween returns the uint8 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsUint8DefaultBetween(b Bucket, k string, v, min, max int8) uint8 {
	return uint8(asUintDefault(b, k, 8, uint64(v), uint64(min), uint64(max)))
}

// AsUint16 returns the uint16 representation of the value associated with k or
// panics if unable to do so.
func AsUint16(b Bucket, k string) uint16 {
	return uint16(asUint(b, k, 16, 0, math.MaxUint16))
}

// AsUint16Default returns the uint16 representation of the value associated
// with k, or the default value v if k is undefined.
func AsUint16Default(b Bucket, k string, v uint16) uint16 {
	return uint16(asUintDefault(b, k, 16, uint64(v), 0, math.MaxUint16))
}

// AsUint16Between returns the uint16 representation of the value associated
// with k or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsUint16Between(b Bucket, k string, min, max int16) uint16 {
	return uint16(asUint(b, k, 16, uint64(min), uint64(max)))
}

// AsUint16DefaultBetween returns the uint16 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsUint16DefaultBetween(b Bucket, k string, v, min, max int16) uint16 {
	return uint16(asUintDefault(b, k, 16, uint64(v), uint64(min), uint64(max)))
}

// AsUint32 returns the uint32 representation of the value associated with k or
// panics if unable to do so.
func AsUint32(b Bucket, k string) uint32 {
	return uint32(asUint(b, k, 32, 0, math.MaxUint32))
}

// AsUint32Default returns the uint32 representation of the value associated
// with k, or the default value v if k is undefined.
func AsUint32Default(b Bucket, k string, v uint32) uint32 {
	return uint32(asUintDefault(b, k, 32, uint64(v), 0, math.MaxUint32))
}

// AsUint32Between returns the uint32 representation of the value associated
// with k or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsUint32Between(b Bucket, k string, min, max int32) uint32 {
	return uint32(asUint(b, k, 32, uint64(min), uint64(max)))
}

// AsUint32DefaultBetween returns the uint32 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsUint32DefaultBetween(b Bucket, k string, v, min, max int32) uint32 {
	return uint32(asUintDefault(b, k, 32, uint64(v), uint64(min), uint64(max)))
}

// AsUint64 returns the uint64 representation of the value associated with k or
// panics if unable to do so.
func AsUint64(b Bucket, k string) uint64 {
	return asUint(b, k, 64, 0, math.MaxUint64)
}

// AsUint64Default returns the uint64 representation of the value associated
// with k, or the default value v if k is undefined.
func AsUint64Default(b Bucket, k string, v uint64) uint64 {
	return asUintDefault(b, k, 64, v, 0, math.MaxUint64)
}

// AsUint64Between returns the uint64 representation of the value associated
// with k or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsUint64Between(b Bucket, k string, min, max uint64) uint64 {
	return asUint(b, k, 64, min, max)
}

// AsUint64DefaultBetween returns the uint64 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsUint64DefaultBetween(b Bucket, k string, v, min, max uint64) uint64 {
	return asUintDefault(b, k, 64, v, min, max)
}

func tryAsUint(
	b Bucket,
	k string,
	bitSize int,
	min, max uint64,
) (uint64, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false
	}

	v, err := strconv.ParseUint(
		mustAsString(k, x),
		10,
		bitSize,
	)
	if err != nil {
		if bitSize == 0 {
			panic(fmt.Sprintf(
				`expected %s to be an unsigned integer: %s`,
				k,
				err,
			))
		}

		panic(fmt.Sprintf(
			`expected %s to be an unsigned %d-bit integer: %s`,
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

func asUint(
	b Bucket,
	k string,
	bitSize int,
	min, max uint64,
) uint64 {
	if v, ok := tryAsUint(b, k, bitSize, min, max); ok {
		return v
	}

	panic(fmt.Sprintf("%s is not defined", k))
}

func asUintDefault(
	b Bucket,
	k string,
	bitSize int,
	d, min, max uint64,
) uint64 {
	if min > d || d > max {
		panic(fmt.Sprintf(
			`expected the default value for %s to be between %d and %d (inclusive), got %d`,
			k,
			min,
			max,
			d,
		))
	}

	if v, ok := tryAsUint(b, k, bitSize, min, max); ok {
		return v
	}

	return d
}
