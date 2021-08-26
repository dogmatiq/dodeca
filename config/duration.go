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

	s := mustAsString(k, x)
	v, err := time.ParseDuration(s)
	if err != nil {
		panic(InvalidValue{
			k,
			s,
			`expected a duration`,
		})
	}

	if min > v || v > max {
		panic(InvalidValue{
			k,
			s,
			fmt.Sprintf(
				`expected a duration between %s and %s (inclusive)`,
				min,
				max,
			),
		})
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

	panic(NotDefined{k})
}

func asDurationDefault(
	b Bucket,
	k string,
	d, min, max time.Duration,
) time.Duration {
	if min > d || d > max {
		panic(InvalidDefaultValue{
			k,
			d,
			fmt.Sprintf(
				`expected a duration between %s and %s (inclusive)`,
				min,
				max,
			),
		})
	}

	if v, ok := tryAsDuration(b, k, min, max); ok {
		return v
	}

	return d
}
