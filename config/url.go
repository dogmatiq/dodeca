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
		panic(fmt.Sprintf(
			`expected the default value for %s to be a URL: %s`,
			k,
			err,
		))
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

	v, err := url.Parse(mustAsString(k, x))
	if err != nil {
		panic(fmt.Sprintf(
			`expected %s to be a URL: %s`,
			k,
			err,
		))
	}

	return v, true
}
