package config_test

import (
	"fmt"

	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsUint()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint(b, "<key>")
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsUint(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint(b, "<key>")
		}).To(PanicWith(InvalidValue{
			Key:   "<key>",
			Value: "<invalid>",
			Explanation: fmt.Sprintf(
				`expected an integer between 0 and %d (inclusive)`,
				uint(MaxUint),
			),
		}))
	})
})

var _ = Describe("func AsUintDefault()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUintDefault(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsUintDefault(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUintDefault(b, "<key>", 50)
		}).To(PanicWith(InvalidValue{
			Key:   "<key>",
			Value: "<invalid>",
			Explanation: fmt.Sprintf(
				`expected an integer between 0 and %d (inclusive)`,
				uint(MaxUint),
			),
		}))
	})
})

var _ = Describe("func AsUintBetween()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUintBetween(b, "<key>", 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsUintBetween(b, "<key>", 10, 100)
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("5")}

		Expect(func() {
			AsUintBetween(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "5",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUintBetween(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUintBetween(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUintDefaultBetween()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUintDefaultBetween(b, "<key>", 50, 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsUintDefaultBetween(b, "<key>", 50, 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("5")}

		Expect(func() {
			AsUintDefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "5",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUintDefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50")}

		Expect(func() {
			AsUintDefaultBetween(b, "<key>", 5, 10, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "5",
			Explanation:  `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUintDefaultBetween(b, "<key>", 120, 10, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "120",
			Explanation:  `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUintDefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})
})
