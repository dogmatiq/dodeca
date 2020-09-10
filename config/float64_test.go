package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsFloat64()", func() {
	It("returns a float64 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat64(b, "<key>")
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsFloat64(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat64(b, "<key>")
		}).To(PanicWith(`expected <key> to be a 64-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsFloat64Default()", func() {
	It("returns a float64 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat64Default(b, "<key>", 50.5)
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsFloat64Default(b, "<key>", 50.5)
		Expect(v).To(BeNumerically("==", 50.5))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat64Default(b, "<key>", 50.5)
		}).To(PanicWith(`expected <key> to be a 64-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsFloat64Between()", func() {
	It("returns a float64 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat64Between(b, "<key>", -100, 100)
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsFloat64Between(b, "<key>", -100, 100)
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsFloat64Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got -120.000000`))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsFloat64Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got 120.000000`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat64Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be a 64-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsFloat64DefaultBetween()", func() {
	It("returns a float64 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat64DefaultBetween(b, "<key>", 50.5, -100, 100)
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsFloat64DefaultBetween(b, "<key>", 50.5, -100, 100)
		Expect(v).To(BeNumerically("==", 50.5))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsFloat64DefaultBetween(b, "<key>", 50.5, -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got -120.000000`))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsFloat64DefaultBetween(b, "<key>", 50.5, -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got 120.000000`))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50.5")}

		Expect(func() {
			AsFloat64DefaultBetween(b, "<key>", -120, -100, 100)
		}).To(PanicWith(`expected the default value for <key> to be between -100.000000 and 100.000000 (inclusive), got -120.000000`))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsFloat64DefaultBetween(b, "<key>", 120, -100, 100)
		}).To(PanicWith(`expected the default value for <key> to be between -100.000000 and 100.000000 (inclusive), got 120.000000`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat64DefaultBetween(b, "<key>", 50.5, -100, 100)
		}).To(PanicWith(`expected <key> to be a 64-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})
