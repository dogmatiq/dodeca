package config

import (
	"fmt"
	"net/url"
)

// AsURL returns the url.URL representation of the value associated with k or
// panics if unable to do so.
func AsURL(b Bucket, k string) *url.URL {
	if v, ok := tryAsURL(b, k); ok {
		return v
	}

	panic(NotDefined{k})
}

// AsURLDefault returns the url.URL representation of the value associated with
// k, or the default value v if k is undefined.
func AsURLDefault(b Bucket, k, v string) *url.URL {
	if v, ok := tryAsURL(b, k); ok {
		return v
	}

	u, err := url.Parse(v)
	if err != nil {
		panic(InvalidDefaultValue{
			k,
			v,
			fmt.Sprintf(
				`expected a URL (%s)`,
				err.(*url.Error).Unwrap(),
			),
		})
	}

	return u
}

func tryAsURL(
	b Bucket,
	k string,
) (*url.URL, bool) {
	x := b.Get(k)

	if x.IsZero() {
		return nil, false
	}

	s := mustAsString(k, x)
	v, err := url.Parse(s)
	if err != nil {
		panic(InvalidValue{
			k,
			s,
			fmt.Sprintf(
				`expected a URL (%s)`,
				err.(*url.Error).Unwrap(),
			),
		})
	}

	return v, true
}
