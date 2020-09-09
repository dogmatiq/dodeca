package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func GetUint()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, ok, err := GetUint(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(123))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok, err := GetUint(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ok).To(BeFalse())
	})

	It("returns an error if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, _, err := GetUint(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`))
	})

	It("returns an error if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		_, _, err := GetUint(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "-123": invalid syntax`))
	})
})

func ExampleGetUint() {
	os.Setenv("FOO", "123")

	v, ok, err := config.GetUint(config.Environment(), "FOO")
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

var _ = Describe("func GetUintDefault()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, err := GetUintDefault(b, "<key>", 10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v, err := GetUintDefault(b, "<key>", 10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(10))
	})

	It("returns an error if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetUintDefault(b, "<key>", 10)
		Expect(err).To(MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`))
	})

	It("returns an error if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		_, err := GetUintDefault(b, "<key>", 10)
		Expect(err).To(MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "-123": invalid syntax`))
	})
})

func ExampleGetUintDefault() {
	os.Setenv("FOO", "123")

	v, err := config.GetUintDefault(config.Environment(), "FOO", 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the value is %d!\n", v)

	// Output: the value is 123!
}

var _ = Describe("func MustGetUint()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, ok := MustGetUint(b, "<key>")
		Expect(v).To(BeEquivalentTo(123))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok := MustGetUint(b, "<key>")
		Expect(ok).To(BeFalse())
	})

	It("panics if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetUint(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`),
		))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			MustGetUint(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "-123": invalid syntax`),
		))
	})
})

func ExampleMustGetUint() {
	os.Setenv("FOO", "123")

	v, ok := config.MustGetUint(config.Environment(), "FOO")

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %d!\n", v)
	}

	// Output: the value is 123!
}

var _ = Describe("func MustGetUintDefault()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v := MustGetUintDefault(b, "<key>", 10)
		Expect(v).To(BeEquivalentTo(123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := MustGetUintDefault(b, "<key>", 10)
		Expect(v).To(BeEquivalentTo(10))
	})

	It("panics if the value can not be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetUintDefault(b, "<key>", 10)
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`),
		))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			MustGetUintDefault(b, "<key>", 10)
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned integer: strconv.ParseUint: parsing "-123": invalid syntax`),
		))
	})
})

func ExampleMustGetUintDefault() {
	os.Setenv("FOO", "123")

	v := config.MustGetUintDefault(config.Environment(), "FOO", 10)

	fmt.Printf("the value is %d!\n", v)

	// Output: the value is 123!
}
