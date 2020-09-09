package config_test

import (
	"fmt"
	"os"
	"time"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func GetDuration()", func() {
	It("returns a positive duration value", func() {
		b := Map{"<key>": String("123ms")}

		v, ok, err := GetDuration(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(Equal(123 * time.Millisecond))
		Expect(ok).To(BeTrue())
	})

	It("returns a negative duration value", func() {
		b := Map{"<key>": String("-123ms")}

		v, ok, err := GetDuration(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(Equal(-123 * time.Millisecond))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok, err := GetDuration(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ok).To(BeFalse())
	})

	It("returns an error if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, _, err := GetDuration(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid duration: time: invalid duration "<invalid>"`))
	})
})

func ExampleGetDuration() {
	os.Setenv("FOO", "100ms")

	v, ok, err := config.GetDuration(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %s!\n", v)
	}

	// Output: the value is 100ms!
}

var _ = Describe("func GetDurationDefault()", func() {
	It("returns a positive duration value", func() {
		b := Map{"<key>": String("123ms")}

		v, err := GetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(Equal(123 * time.Millisecond))
	})

	It("returns a negative duration value", func() {
		b := Map{"<key>": String("-123ms")}

		v, err := GetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(Equal(-123 * time.Millisecond))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v, err := GetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(Equal(456 * time.Millisecond))
	})

	It("returns an error if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(err).To(MatchError(`<key> is not a valid duration: time: invalid duration "<invalid>"`))
	})
})

func ExampleGetDurationDefault() {
	os.Setenv("FOO", "")

	v, err := config.GetDurationDefault(config.Environment(), "FOO", 200*time.Millisecond)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the value is %s!\n", v)

	// Output: the value is 200ms!
}

var _ = Describe("func MustGetDuration()", func() {
	It("returns a positive duration value", func() {
		b := Map{"<key>": String("123ms")}

		v, ok := MustGetDuration(b, "<key>")
		Expect(v).To(Equal(123 * time.Millisecond))
		Expect(ok).To(BeTrue())
	})

	It("returns a negative duration value", func() {
		b := Map{"<key>": String("-123ms")}

		v, ok := MustGetDuration(b, "<key>")
		Expect(v).To(Equal(-123 * time.Millisecond))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok := MustGetDuration(b, "<key>")
		Expect(ok).To(BeFalse())
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetDuration(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid duration: time: invalid duration "<invalid>"`),
		))
	})
})

func ExampleMustGetDuration() {
	os.Setenv("FOO", "100ms")

	v, ok := config.MustGetDuration(config.Environment(), "FOO")

	if !ok {
		fmt.Println("key is not defined!")
	} else {
		fmt.Printf("the value is %s!\n", v)
	}

	// Output: the value is 100ms!
}

var _ = Describe("func MustGetDurationDefault()", func() {
	It("returns a positive duration value", func() {
		b := Map{"<key>": String("123ms")}

		v := MustGetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(v).To(Equal(123 * time.Millisecond))
	})

	It("returns a negative duration value", func() {
		b := Map{"<key>": String("-123ms")}

		v := MustGetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(v).To(Equal(-123 * time.Millisecond))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := MustGetDurationDefault(b, "<key>", 456*time.Millisecond)
		Expect(v).To(Equal(456 * time.Millisecond))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetDurationDefault(b, "<key>", 456*time.Millisecond)
		}).To(PanicWith(
			MatchError(`<key> is not a valid duration: time: invalid duration "<invalid>"`),
		))
	})
})

func ExampleMustGetDurationDefault() {
	os.Setenv("FOO", "")

	v := config.MustGetDurationDefault(config.Environment(), "FOO", 200*time.Millisecond)

	fmt.Printf("the value is %s!\n", v)

	// Output: the value is 200ms!
}
