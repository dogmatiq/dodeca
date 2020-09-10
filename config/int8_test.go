package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsInt8()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-123")}

		v := AsInt8(b, "<key>")
		Expect(v).To(BeNumerically("==", -123))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt8(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt8(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed 8-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt8Default()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-123")}

		v := AsInt8Default(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", -123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsInt8Default(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt8Default(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be a signed 8-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt8P()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("123")}

		v := AsInt8P(b, "<key>")
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt8P(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value is zero", func() {
		b := Map{"<key>": String("0")}

		Expect(func() {
			AsInt8P(b, "<key>")
		}).To(PanicWith(`expected <key> to be positive, got 0`))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			AsInt8P(b, "<key>")
		}).To(PanicWith(`expected <key> to be positive, got -123`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt8P(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed 8-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt8PDefault()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("123")}

		v := AsInt8PDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsInt8PDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the value is zero", func() {
		b := Map{"<key>": String("0")}

		Expect(func() {
			AsInt8PDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be positive, got 0`))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			AsInt8PDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be positive, got -123`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt8PDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be a signed 8-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})
