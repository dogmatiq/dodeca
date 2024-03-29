package config_test

import (
	"time"

	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsDuration()", func() {
	It("returns a duration value", func() {
		b := Map{"<key>": String("-50ms")}

		v := AsDuration(b, "<key>")
		Expect(v).To(BeNumerically("==", -50*time.Millisecond))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsDuration(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsDuration(b, "<key>")
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected a duration`,
		}))
	})
})

var _ = Describe("func AsDurationDefault()", func() {
	It("returns a duration value", func() {
		b := Map{"<key>": String("-50ms")}

		v := AsDurationDefault(b, "<key>", 50*time.Millisecond)
		Expect(v).To(BeNumerically("==", -50*time.Millisecond))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsDurationDefault(b, "<key>", 50*time.Millisecond)
		Expect(v).To(BeNumerically("==", 50*time.Millisecond))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsDurationDefault(b, "<key>", 50*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected a duration`,
		}))
	})
})

var _ = Describe("func AsDurationBetween()", func() {
	It("returns a duration value", func() {
		b := Map{"<key>": String("-50ms")}

		v := AsDurationBetween(b, "<key>", -100*time.Millisecond, 100*time.Millisecond)
		Expect(v).To(BeNumerically("==", -50*time.Millisecond))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsDurationBetween(b, "<key>", -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120ms")}

		Expect(func() {
			AsDurationBetween(b, "<key>", -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "-120ms",
			Explanation: `expected a duration between -100ms and 100ms (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120ms")}

		Expect(func() {
			AsDurationBetween(b, "<key>", -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120ms",
			Explanation: `expected a duration between -100ms and 100ms (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsDurationBetween(b, "<key>", -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected a duration`,
		}))
	})
})

var _ = Describe("func AsDurationDefaultBetween()", func() {
	It("returns a duration value", func() {
		b := Map{"<key>": String("-50ms")}

		v := AsDurationDefaultBetween(b, "<key>", 50*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		Expect(v).To(BeNumerically("==", -50*time.Millisecond))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsDurationDefaultBetween(b, "<key>", 50*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		Expect(v).To(BeNumerically("==", 50*time.Millisecond))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120ms")}

		Expect(func() {
			AsDurationDefaultBetween(b, "<key>", 50*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "-120ms",
			Explanation: `expected a duration between -100ms and 100ms (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120ms")}

		Expect(func() {
			AsDurationDefaultBetween(b, "<key>", 50*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120ms",
			Explanation: `expected a duration between -100ms and 100ms (inclusive)`,
		}))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50ms")}

		Expect(func() {
			AsDurationDefaultBetween(b, "<key>", -120*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "-120ms",
			Explanation:  `expected a duration between -100ms and 100ms (inclusive)`,
		}))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120ms")}

		Expect(func() {
			AsDurationDefaultBetween(b, "<key>", 120*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "120ms",
			Explanation:  `expected a duration between -100ms and 100ms (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsDurationDefaultBetween(b, "<key>", 50*time.Millisecond, -100*time.Millisecond, 100*time.Millisecond)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected a duration`,
		}))
	})
})
