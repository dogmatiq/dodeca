package config

import (
	"fmt"
	"math"
	"time"
)

// AsDuration returns the time.Duration representation of the value associated
// with k or panics if unable to do so.
//
// Durations are specified using the syntax supported by time.ParseDuration.
func AsDuration(b Bucket, k string) time.Duration {
	return asDuration(b, k, math.MinInt64, math.MaxInt64)
}

// AsDurationDefault returns the time.Duration representation of the value
// associated with k, or the default value v if k is undefined.
//
// Durations are specified using the syntax supported by time.ParseDuration.
func AsDurationDefault(b Bucket, k string, v time.Duration) time.Duration {
	return asDurationDefault(b, k, v, math.MinInt64, math.MaxInt64)
}

// AsDurationBetween returns the time.Duration representation of the value
// associated with k or panics if unable to do so.
//
// Durations are specified using the syntax supported by time.ParseDuration.
//
// It panics if the value is not between min and max (inclusive).
func AsDurationBetween(b Bucket, k string, min, max time.Duration) time.Duration {
	return asDuration(b, k, min, max)
}

// AsDurationDefaultBetween returns the time.Duration representation of the
// value associated with k, or the default value v if k is undefined.
//
// Durations are specified using the syntax supported by time.ParseDuration.
//
// It panics if the value is not between min and max (inclusive).
func AsDurationDefaultBetween(b Bucket, k string, v, min, max time.Duration) time.Duration {
	return asDurationDefault(b, k, v, min, max)
}

func tryAsDuration(
	b Bucket,
	k string,
	min, max time.Duration,
) (time.Duration, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return 0, false
	}

	v, err := time.ParseDuration(mustAsString(k, x))
	if err != nil {
		panic(fmt.Sprintf(
			`expected %s to be a duration: %s`,
			k,
			err,
		))
	}

	if min > v || v > max {
		panic(fmt.Sprintf(
			`expected %s to be between %s and %s (inclusive), got %s`,
			k,
			min,
			max,
			v,
		))
	}

	return v, true
}

func asDuration(
	b Bucket,
	k string,
	min, max time.Duration,
) time.Duration {
	if v, ok := tryAsDuration(b, k, min, max); ok {
		return v
	}

	panic(fmt.Sprintf("%s is not defined", k))
}

func asDurationDefault(
	b Bucket,
	k string,
	d, min, max time.Duration,
) time.Duration {
	if min > d || d > max {
		panic(fmt.Sprintf(
			`expected the default value for %s to be between %s and %s (inclusive), got %s`,
			k,
			min,
			max,
			d,
		))
	}

	if v, ok := tryAsDuration(b, k, min, max); ok {
		return v
	}

	return d
}
