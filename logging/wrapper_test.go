package logging_test

import (
	. "github.com/dogmatiq/dodeca/logging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type wrapper struct {
	Logger
	Ok bool
}

func (w wrapper) UnwrapLogger() (Logger, bool) {
	return w.Logger, w.Ok
}

var _ = Describe("func Unwrap()", func() {
	It("returns the argument if it does not implement Wrapper", func() {
		logger := DiscardLogger{}
		Expect(Unwrap(logger)).To(BeIdenticalTo(logger))
	})

	It("returns the argument if it is a Wrapper that is not currently wrapping", func() {
		logger := wrapper{}
		Expect(Unwrap(logger)).To(BeIdenticalTo(logger))
	})

	It("returns the wrapped logger", func() {
		next := DiscardLogger{}
		logger := wrapper{
			Logger: next,
			Ok:     true,
		}
		Expect(Unwrap(logger)).To(BeIdenticalTo(next))
	})

	It("unwraps multiple levels", func() {
		next := DiscardLogger{}
		logger := wrapper{
			Logger: wrapper{
				Logger: next,
				Ok:     true,
			},
			Ok: true,
		}
		Expect(Unwrap(logger)).To(BeIdenticalTo(next))
	})

	It("returns DefaultLogger if the argument is nil", func() {
		Expect(Unwrap(nil)).To(BeIdenticalTo(DefaultLogger))
	})

	It("returns DefaultLogger if the wrapped logger is nil", func() {
		logger := wrapper{
			Ok: true,
		}
		Expect(Unwrap(logger)).To(BeIdenticalTo(DefaultLogger))
	})
})
