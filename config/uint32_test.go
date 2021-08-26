package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsUint32()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint32(b, "<key>")
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsUint32(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint32(b, "<key>")
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 0 and 4294967295 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUint32Default()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint32Default(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsUint32Default(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint32Default(b, "<key>", 50)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 0 and 4294967295 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUint32Between()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint32Between(b, "<key>", 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsUint32Between(b, "<key>", 10, 100)
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("5")}

		Expect(func() {
			AsUint32Between(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "5",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUint32Between(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint32Between(b, "<key>", 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsUint32DefaultBetween()", func() {
	It("returns a uint value", func() {
		b := Map{"<key>": String("50")}

		v := AsUint32DefaultBetween(b, "<key>", 50, 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsUint32DefaultBetween(b, "<key>", 50, 10, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("5")}

		Expect(func() {
			AsUint32DefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "5",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUint32DefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50")}

		Expect(func() {
			AsUint32DefaultBetween(b, "<key>", 5, 10, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "5",
			Explanation:  `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsUint32DefaultBetween(b, "<key>", 120, 10, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "120",
			Explanation:  `expected an integer between 10 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsUint32DefaultBetween(b, "<key>", 50, 10, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between 10 and 100 (inclusive)`,
		}))
	})
})
