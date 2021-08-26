package config_test

import (
	"fmt"

	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsInt()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-50")}

		v := AsInt(b, "<key>")
		Expect(v).To(BeNumerically("==", -50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt(b, "<key>")
		}).To(PanicWith(InvalidValue{
			Key:   "<key>",
			Value: "<invalid>",
			Explanation: fmt.Sprintf(
				`expected an integer between %d and %d (inclusive)`,
				MinInt,
				MaxInt,
			),
		}))
	})
})

var _ = Describe("func AsIntDefault()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-50")}

		v := AsIntDefault(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", -50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsIntDefault(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsIntDefault(b, "<key>", 50)
		}).To(PanicWith(InvalidValue{
			Key:   "<key>",
			Value: "<invalid>",
			Explanation: fmt.Sprintf(
				`expected an integer between %d and %d (inclusive)`,
				MinInt,
				MaxInt,
			),
		}))
	})
})

var _ = Describe("func AsIntBetween()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-50")}

		v := AsIntBetween(b, "<key>", -100, 100)
		Expect(v).To(BeNumerically("==", -50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsIntBetween(b, "<key>", -100, 100)
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsIntBetween(b, "<key>", -100, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "-120",
			Explanation: `expected an integer between -100 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsIntBetween(b, "<key>", -100, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between -100 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsIntBetween(b, "<key>", -100, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between -100 and 100 (inclusive)`,
		}))
	})
})

var _ = Describe("func AsIntDefaultBetween()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-50")}

		v := AsIntDefaultBetween(b, "<key>", 50, -100, 100)
		Expect(v).To(BeNumerically("==", -50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsIntDefaultBetween(b, "<key>", 50, -100, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsIntDefaultBetween(b, "<key>", 50, -100, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "-120",
			Explanation: `expected an integer between -100 and 100 (inclusive)`,
		}))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsIntDefaultBetween(b, "<key>", 50, -100, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "120",
			Explanation: `expected an integer between -100 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50")}

		Expect(func() {
			AsIntDefaultBetween(b, "<key>", -120, -100, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "-120",
			Explanation:  `expected an integer between -100 and 100 (inclusive)`,
		}))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsIntDefaultBetween(b, "<key>", 120, -100, 100)
		}).To(PanicWith(InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "120",
			Explanation:  `expected an integer between -100 and 100 (inclusive)`,
		}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsIntDefaultBetween(b, "<key>", 50, -100, 100)
		}).To(PanicWith(InvalidValue{
			Key:         "<key>",
			Value:       "<invalid>",
			Explanation: `expected an integer between -100 and 100 (inclusive)`,
		}))
	})
})
