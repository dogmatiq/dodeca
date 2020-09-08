package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func GetInt64()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, ok, err := GetInt64(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(123))
		Expect(ok).To(BeTrue())
	})

	It("returns a negative integer value", func() {
		b := Map{"<key>": String("-123")}

		v, ok, err := GetInt64(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(-123))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok, err := GetInt64(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ok).To(BeFalse())
	})

	It("returns an error if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, _, err := GetInt64(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

func ExampleGetInt64() {
	os.Setenv("FOO", "123")

	v, ok, err := config.GetInt64(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %d!\n", v)
	}

	// Output: the value is 123!
}

var _ = Describe("func GetInt64Default()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, err := GetInt64Default(b, "<key>", -10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(123))
	})

	It("returns a negative integer value", func() {
		b := Map{"<key>": String("-123")}

		v, err := GetInt64Default(b, "<key>", -10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(-123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v, err := GetInt64Default(b, "<key>", -10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(-10))
	})

	It("returns an error if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetInt64Default(b, "<key>", -10)
		Expect(err).To(MatchError(`<key> is not a valid signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`))
	})
})

func ExampleGetInt64Default() {
	os.Setenv("FOO", "123")

	v, err := config.GetInt64Default(config.Environment(), "FOO", -10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the value is %d!\n", v)

	// Output: the value is 123!
}

var _ = Describe("func MustGetInt64()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, ok := MustGetInt64(b, "<key>")
		Expect(v).To(BeEquivalentTo(123))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok := MustGetInt64(b, "<key>")
		Expect(ok).To(BeFalse())
	})

	It("panics if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetInt64(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`),
		))
	})
})

func ExampleMustGetInt64() {
	os.Setenv("FOO", "123")

	v, ok := config.MustGetInt64(config.Environment(), "FOO")

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %d!\n", v)
	}

	// Output: the value is 123!
}

var _ = Describe("func MustGetInt64Default()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v := MustGetInt64Default(b, "<key>", -10)
		Expect(v).To(BeEquivalentTo(123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := MustGetInt64Default(b, "<key>", -10)
		Expect(v).To(BeEquivalentTo(-10))
	})

	It("panics if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetInt64Default(b, "<key>", -10)
		}).To(PanicWith(
			MatchError(`<key> is not a valid signed 64-bit integer: strconv.ParseInt: parsing "<invalid>": invalid syntax`),
		))
	})
})

func ExampleMustGetInt64Default() {
	os.Setenv("FOO", "123")

	v := config.MustGetInt64Default(config.Environment(), "FOO", -10)

	fmt.Printf("the value is %d!\n", v)

	// Output: the value is 123!
}
