package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsUint16()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint16(b, "<key>")
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsUint16(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint16(b, "<key>")
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 0 and 65535 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUint16Default()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint16Default(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsUint16Default(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint16Default(b, "<key>", 50)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 0 and 65535 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUint16Between()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint16Between(b, "<key>", 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsUint16Between(b, "<key>", 10, 100)
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("5")}

		Expect(func() {
			AsUint16Between(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "5",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUint16Between(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint16Between(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUint16DefaultBetween()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint16DefaultBetween(b, "<key>", 50, 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsUint16DefaultBetween(b, "<key>", 50, 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("5")}

		Expect(func() {
			AsUint16DefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "5",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUint16DefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50")}

		Expect(func() {
			AsUint16DefaultBetween(b, "<key>", 5, 10, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "5",
			Explanation:  `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUint16DefaultBetween(b, "<key>", 120, 10, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "120",
			Explanation:  `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint16DefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})
})
