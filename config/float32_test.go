package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsFloat32()", func() {
	It("returns a float32 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat32(b, "<key>")
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsFloat32(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat32(b, "<key>")
		}).To(PanicWith(`expected <key> to be a 32-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsFloat32Default()", func() {
	It("returns a float32 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat32Default(b, "<key>", 50.5)
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsFloat32Default(b, "<key>", 50.5)
		Expect(v).To(BeNumerically("==", 50.5))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat32Default(b, "<key>", 50.5)
		}).To(PanicWith(`expected <key> to be a 32-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsFloat32Between()", func() {
	It("returns a float32 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat32Between(b, "<key>", -100, 100)
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsFloat32Between(b, "<key>", -100, 100)
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsFloat32Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got -120.000000`))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsFloat32Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got 120.000000`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat32Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be a 32-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsFloat32DefaultBetween()", func() {
	It("returns a float32 value", func() {
		b := Map{"<key>": String("-50.5")}

		v := AsFloat32DefaultBetween(b, "<key>", 50.5, -100, 100)
		Expect(v).To(BeNumerically("==", -50.5))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsFloat32DefaultBetween(b, "<key>", 50.5, -100, 100)
		Expect(v).To(BeNumerically("==", 50.5))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsFloat32DefaultBetween(b, "<key>", 50.5, -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got -120.000000`))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsFloat32DefaultBetween(b, "<key>", 50.5, -100, 100)
		}).To(PanicWith(`expected <key> to be between -100.000000 and 100.000000 (inclusive), got 120.000000`))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50.5")}

		Expect(func() {
			AsFloat32DefaultBetween(b, "<key>", -120, -100, 100)
		}).To(PanicWith(`expected the default value for <key> to be between -100.000000 and 100.000000 (inclusive), got -120.000000`))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsFloat32DefaultBetween(b, "<key>", 120, -100, 100)
		}).To(PanicWith(`expected the default value for <key> to be between -100.000000 and 100.000000 (inclusive), got 120.000000`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsFloat32DefaultBetween(b, "<key>", 50.5, -100, 100)
		}).To(PanicWith(`expected <key> to be a 32-bit floating-point number: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})
