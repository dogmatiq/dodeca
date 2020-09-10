package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ExampleAsInt() {
	os.Setenv("FOO", "123")

	v := config.AsInt(config.Environment(), "FOO")

	fmt.Printf("the value is %d!\n", v)

	// Output: the value is 123!
}

func ExampleAsIntDefault() {
	os.Setenv("FOO", "")

	v := config.AsIntDefault(config.Environment(), "FOO", -456)

	fmt.Printf("the value is %d!\n", v)

	// Output: the value is -456!
}

var _ = Describe("func AsInt()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-123")}

		v := AsInt(b, "<key>")
		Expect(v).To(BeNumerically("==", -123))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsInt(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsInt(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsIntDefault()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("-123")}

		v := AsIntDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", -123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsIntDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsIntDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be a signed integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsIntP()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("123")}

		v := AsIntP(b, "<key>")
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsIntP(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})

	It("panics if the value is zero", func() {
		b := Map{"<key>": String("0")}

		Expect(func() {
			AsIntP(b, "<key>")
		}).To(PanicWith(`expected <key> to be positive, got 0`))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			AsIntP(b, "<key>")
		}).To(PanicWith(`expected <key> to be positive, got -123`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsIntP(b, "<key>")
		}).To(PanicWith(`expected <key> to be a signed integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

var _ = Describe("func AsIntPDefault()", func() {
	It("returns an int value", func() {
		b := Map{"<key>": String("123")}

		v := AsIntPDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsIntPDefault(b, "<key>", 123)
		Expect(v).To(BeNumerically("==", 123))
	})

	It("panics if the value is zero", func() {
		b := Map{"<key>": String("0")}

		Expect(func() {
			AsIntPDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be positive, got 0`))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			AsIntPDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be positive, got -123`))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsIntPDefault(b, "<key>", 123)
		}).To(PanicWith(`expected <key> to be a signed integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})
