package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func GetFloat32()", func() {
	It("returns a positive float value", func() {
		b := Map{"<key>": String("123.45")}

		v, ok, err := GetFloat32(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeNumerically("~", float32(123.45)))
		Expect(ok).To(BeTrue())
	})

	It("returns a negative float value", func() {
		b := Map{"<key>": String("-123.45")}

		v, ok, err := GetFloat32(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeNumerically("~", float32(-123.45)))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok, err := GetFloat32(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ok).To(BeFalse())
	})

	It("returns an error if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, _, err := GetFloat32(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid 32-bit float: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

func ExampleGetFloat32() {
	os.Setenv("FOO", "1.5")

	v, ok, err := config.GetFloat32(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %f!\n", v)
	}

	// Output: the value is 1.500000!
}

var _ = Describe("func GetFloat32Default()", func() {
	It("returns a positive float value", func() {
		b := Map{"<key>": String("123.45")}

		v, err := GetFloat32Default(b, "<key>", 456.78)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeNumerically("~", float32(123.45)))
	})

	It("returns a negative float value", func() {
		b := Map{"<key>": String("-123.45")}

		v, err := GetFloat32Default(b, "<key>", 456.78)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeNumerically("~", float32(-123.45)))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v, err := GetFloat32Default(b, "<key>", 456.78)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeNumerically("~", float32(456.78)))
	})

	It("returns an error if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetFloat32Default(b, "<key>", 456.78)
		Expect(err).To(MatchError(`<key> is not a valid 32-bit float: strconv.ParseFloat: parsing "<invalid>": invalid syntax`))
	})
})

func ExampleGetFloat32Default() {
	os.Setenv("FOO", "")

	v, err := config.GetFloat32Default(config.Environment(), "FOO", 4.5)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the value is %f!\n", v)

	// Output: the value is 4.500000!
}

var _ = Describe("func MustGetFloat32()", func() {
	It("returns a positive float value", func() {
		b := Map{"<key>": String("123.45")}

		v, ok := MustGetFloat32(b, "<key>")
		Expect(v).To(BeNumerically("~", float32(123.45)))
		Expect(ok).To(BeTrue())
	})

	It("returns a negative float value", func() {
		b := Map{"<key>": String("-123.45")}

		v, ok := MustGetFloat32(b, "<key>")
		Expect(v).To(BeNumerically("~", float32(-123.45)))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok := MustGetFloat32(b, "<key>")
		Expect(ok).To(BeFalse())
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetFloat32(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid 32-bit float: strconv.ParseFloat: parsing "<invalid>": invalid syntax`),
		))
	})
})

func ExampleMustGetFloat32() {
	os.Setenv("FOO", "1.5")

	v, ok := config.MustGetFloat32(config.Environment(), "FOO")

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %f!\n", v)
	}

	// Output: the value is 1.500000!
}

var _ = Describe("func MustGetFloat32Default()", func() {
	It("returns a positive float value", func() {
		b := Map{"<key>": String("123.45")}

		v := MustGetFloat32Default(b, "<key>", 456.78)
		Expect(v).To(BeNumerically("~", float32(123.45)))
	})

	It("returns a negative float value", func() {
		b := Map{"<key>": String("-123.45")}

		v := MustGetFloat32Default(b, "<key>", 456.78)
		Expect(v).To(BeNumerically("~", float32(-123.45)))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := MustGetFloat32Default(b, "<key>", 456.78)
		Expect(v).To(BeNumerically("~", float32(456.78)))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetFloat32Default(b, "<key>", 456.78)
		}).To(PanicWith(
			MatchError(`<key> is not a valid 32-bit float: strconv.ParseFloat: parsing "<invalid>": invalid syntax`),
		))
	})
})

func ExampleMustGetFloat32Default() {
	os.Setenv("FOO", "")

	v := config.MustGetFloat32Default(config.Environment(), "FOO", 4.5)

	fmt.Printf("the value is %f!\n", v)

	// Output: the value is 4.500000!
}
