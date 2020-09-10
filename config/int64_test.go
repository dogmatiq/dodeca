package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsInt64()", func() {
	It("returns an int64 value", func() {
		b := Map{"<key>": String("-123")}

		v := AsInt64(b, "<key>")
		Expect(v).To(BeNumerically("==", -123))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt64(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt64(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt64Default()", func() {
	It("returns an int64 value", func() {
		b := Map{"<key>": String("-123")}

		v := AsInt64Default(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", -123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsInt64Default(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt64Default(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be a signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt64P()", func() {
	It("returns an int64 value", func() {
		b := Map{"<key>": String("123")}

		v := AsInt64P(b, "<key>")
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt64P(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value is zero", func() {
		b := Map{"<key>": String("0")}

		Expect(func() {
			AsInt64P(b, "<key>")
		}).To(PanicWith(`expected <key> to be positive, got 0`))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			AsInt64P(b, "<key>")
		}).To(PanicWith(`expected <key> to be positive, got -123`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt64P(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsInt64PDefault()", func() {
	It("returns an int64 value", func() {
		b := Map{"<key>": String("123")}

		v := AsInt64PDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsInt64PDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the value is zero", func() {
		b := Map{"<key>": String("0")}

		Expect(func() {
			AsInt64PDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be positive, got 0`))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			AsInt64PDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be positive, got -123`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt64PDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be a signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})
