package config

import (
	"fmt"
	"math"
	"strconv"
)

// AsFloat32 returns the float32 representation of the value associated with k
// or panics if unable to do so.
func AsFloat32(b Bucket, k string) float32 {
	return float32(asFloat(b, k, 32, -math.MaxFloat32, math.MaxFloat32))
}

// AsFloat32Default returns the float32 representation of the value associated
// with k, or the default value v if k is undefined.
func AsFloat32Default(b Bucket, k string, v float32) float32 {
	return float32(asFloatDefault(b, k, 32, float64(v), -math.MaxFloat32, math.MaxFloat32))
}

// AsFloat32Between returns the float32 representation of the value associated
// with k or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsFloat32Between(b Bucket, k string, min, max float32) float32 {
	return float32(asFloat(b, k, 32, float64(min), float64(max)))
}

// AsFloat32DefaultBetween returns the float32 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsFloat32DefaultBetween(b Bucket, k string, v, min, max float32) float32 {
	return float32(asFloatDefault(b, k, 32, float64(v), float64(min), float64(max)))
}

// AsFloat64 returns the float64 representation of the value associated with k
// or panics if unable to do so.
func AsFloat64(b Bucket, k string) float64 {
	return asFloat(b, k, 64, -math.MaxFloat64, math.MaxFloat64)
}

// AsFloat64Default returns the float64 representation of the value associated
// with k, or the default value v if k is undefined.
func AsFloat64Default(b Bucket, k string, v float64) float64 {
	return asFloatDefault(b, k, 64, v, -math.MaxFloat64, math.MaxFloat64)
}

// AsFloat64Between returns the float64 representation of the value associated
// with k or panics if unable to do so.
//
// It panics if the value is not between min and max (inclusive).
func AsFloat64Between(b Bucket, k string, min, max float64) float64 {
	return asFloat(b, k, 64, min, max)
}

// AsFloat64DefaultBetween returns the float64 representation of the value
// associated with k, or the default value v if k is undefined.
//
// It panics if the value is not between min and max (inclusive).
func AsFloat64DefaultBetween(b Bucket, k string, v, min, max float64) float64 {
	return asFloatDefault(b, k, 64, v, min, max)
}

func tryAsFloat(
	b Bucket,
	k string,
	bitSize int,
	min, max float64,
) (float64, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false
	}

	s := mustAsString(k, x)
	v, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		panic(InvalidValue{
			k,
			s,
			fmt.Sprintf(`expected a %d-bit floating-point number`, bitSize),
		})
	}

	if min > v || v > max {
		panic(InvalidValue{
			k,
			s,
			fmt.Sprintf(
				`expected a number between %f and %f (inclusive)`,
				min,
				max,
			),
		})
	}

	return v, true
}

func asFloat(
	b Bucket,
	k string,
	bitSize int,
	min, max float64,
) float64 {
	if v, ok := tryAsFloat(b, k, bitSize, min, max); ok {
		return v
	}

	panic(NotDefined{k})
}

func asFloatDefault(
	b Bucket,
	k string,
	bitSize int,
	d, min, max float64,
) float64 {
	if min > d || d > max {
		panic(InvalidDefaultValue{
			k,
			fmt.Sprintf(`%f`, d),
			fmt.Sprintf(
				`expected a number between %f and %f (inclusive)`,
				min,
				max,
			),
		})
	}

	if v, ok := tryAsFloat(b, k, bitSize, min, max); ok {
		return v
	}

	return d
}
