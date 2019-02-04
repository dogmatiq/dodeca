package config

// EachFunc is a function used to visit the key/value pairs in a bucket using
// Bucket.Each().
type EachFunc func(k string, v Value) bool

// Bucket is a container for configuration key/value pairs.
type Bucket interface {
	// Get returns the value associated with the given key.
	//
	// If they key is not defined, it returns an empty value.
	Get(k string) Value

	// GetDefault returns the value associated with the given key.
	//
	// If the key is not defined, it returns a value with the content of v.
	GetDefault(k string, v string) Value

	// Each calls fn for each key/value pair in the bucket.
	//
	// If fn returns false, iteration is stopped.
	//
	// Each returns true if iteration completes fully, or false if fn()
	// returns false.
	Each(fn EachFunc) bool
}
