package config

import "sync/atomic"

// closer is an implementation of io.Closer that calls fn() at most once.
type closer struct {
	fn   func() error
	done int32
}

func (c *closer) Close() error {
	if atomic.SwapInt32(&c.done, 1) == 0 {
		return c.fn()
	}

	return nil
}
