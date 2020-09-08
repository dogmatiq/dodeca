package config

// EachFunc is a function used to visit the key/value pairs in a bucket using
// Bucket.Each().
type EachFunc func(k string, v Value) bool

// Bucket is a container for configuration key/value pairs.
type Bucket interface {
	// Get returns the value associated with the given key.
	//
	// If they key is not defined, it returns a zero-value.
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

// Map is an in-memory implementation of Bucket.
type Map map[string]Value

// Get returns the value associated with the given key.
//
// If they key is not defined, it returns a zero-value.
func (m Map) Get(k string) Value {
	return m[k]
}

// GetDefault returns the value associated with the given key.
//
// If the key is not defined, it returns a value with the content of v.
func (m Map) GetDefault(k string, v string) Value {
	x := m.Get(k)

	if x.IsZero() {
		return String(v)
	}

	return x
}

// Each calls fn for each key/value pair in the bucket.
//
// If fn returns false, iteration is stopped.
//
// Each returns true if iteration completes fully, or false if fn()
// returns false.
func (m Map) Each(fn EachFunc) bool {
	for k, v := range m {
		if !fn(k, v) {
			return false
		}
	}

	return true
}
