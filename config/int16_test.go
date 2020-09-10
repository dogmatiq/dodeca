package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsInt16()", func() {
	It("returns an int16 value", func() {
		b := Map{"<key>": String("-50")}

		v := AsInt16(b, "<key>")
		Expect(v).To(BeNumerically("==", -50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt16(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt16(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed 16-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt16Default()", func() {
	It("returns an int16 value", func() {
		b := Map{"<key>": String("-50")}

		v := AsInt16Default(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", -50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsInt16Default(b, "<key>", 50)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt16Default(b, "<key>", 50)
		}).To(PanicWith(`expected <key> to be a signed 16-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt16Between()", func() {
	It("returns an int16 value", func() {
		b := Map{"<key>": String("-50")}

		v := AsInt16Between(b, "<key>", -100, 100)
		Expect(v).To(BeNumerically("==", -50))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt16Between(b, "<key>", -100, 100)
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsInt16Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be between -100 and 100 (inclusive), got -120`))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsInt16Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be between -100 and 100 (inclusive), got 120`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt16Between(b, "<key>", -100, 100)
		}).To(PanicWith(`expected <key> to be a signed 16-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt16DefaultBetween()", func() {
	It("returns an int16 value", func() {
		b := Map{"<key>": String("-50")}

		v := AsInt16DefaultBetween(b, "<key>", 50, -100, 100)
		Expect(v).To(BeNumerically("==", -50))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsInt16DefaultBetween(b, "<key>", 50, -100, 100)
		Expect(v).To(BeNumerically("==", 50))
	})

	It("panics if the value is lower than the minimum", func() {
		b := Map{"<key>": String("-120")}

		Expect(func() {
			AsInt16DefaultBetween(b, "<key>", 50, -100, 100)
		}).To(PanicWith(`expected <key> to be between -100 and 100 (inclusive), got -120`))
	})

	It("panics if the value is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsInt16DefaultBetween(b, "<key>", 50, -100, 100)
		}).To(PanicWith(`expected <key> to be between -100 and 100 (inclusive), got 120`))
	})

	It("panics if the default is lower than the minimum", func() {
		b := Map{"<key>": String("50")}

		Expect(func() {
			AsInt16DefaultBetween(b, "<key>", -120, -100, 100)
		}).To(PanicWith(`expected the default value for <key> to be between -100 and 100 (inclusive), got -120`))
	})

	It("panics if the default is greater than the maximum", func() {
		b := Map{"<key>": String("120")}

		Expect(func() {
			AsInt16DefaultBetween(b, "<key>", 120, -100, 100)
		}).To(PanicWith(`expected the default value for <key> to be between -100 and 100 (inclusive), got 120`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt16DefaultBetween(b, "<key>", 50, -100, 100)
		}).To(PanicWith(`expected <key> to be a signed 16-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})
