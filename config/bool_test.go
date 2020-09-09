package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var boolTableEntries = []TableEntry{
	Entry("truthy: true", "true", true),
	Entry("truthy: yes", "yes", true),
	Entry("truthy: on", "on", true),

	Entry("falsey: false", "false", false),
	Entry("falsey: no", "no", false),
	Entry("falsey: off", "off", false),
}

var _ = Describe("func GetBool()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v, ok, err := GetBool(b, "<key>")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(expect))
			Expect(ok).To(BeTrue())
		},
		boolTableEntries...,
	)

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok, err := GetBool(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ok).To(BeFalse())
	})

	It("returns an error if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		_, _, err := GetBool(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`))
	})
})

func ExampleGetBool() {
	os.Setenv("FOO", "true")

	v, ok, err := config.GetBool(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("key is not defined!")
	} else if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

var _ = Describe("func GetBoolT()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v, err := GetBoolT(b, "<key>")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns true if the key is not defined", func() {
		b := Map{}

		v, err := GetBoolT(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeTrue())
	})

	It("returns an error if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetBoolT(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`))
	})
})

func ExampleGetBoolT() {
	os.Setenv("FOO", "false")

	v, err := config.GetBoolT(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

var _ = Describe("func GetBoolF()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v, err := GetBoolF(b, "<key>")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns false if the key is not defined", func() {
		b := Map{}

		v, err := GetBoolF(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeFalse())
	})

	It("returns an error if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetBoolF(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`))
	})
})

func ExampleGetBoolF() {
	os.Setenv("FOO", "true")

	v, err := config.GetBoolF(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

var _ = Describe("func GetBoolDefault()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v, err := GetBoolDefault(b, "<key>", true)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v, err := GetBoolDefault(b, "<key>", true)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeTrue())
	})

	It("returns an error if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetBoolDefault(b, "<key>", true)
		Expect(err).To(MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`))
	})
})

func ExampleGetBoolDefault() {
	os.Setenv("FOO", "")

	v, err := config.GetBoolDefault(config.Environment(), "FOO", true)
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

var _ = Describe("func MustGetBool()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v, ok := MustGetBool(b, "<key>")
			Expect(v).To(Equal(expect))
			Expect(ok).To(BeTrue())
		},
		boolTableEntries...,
	)

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok := MustGetBool(b, "<key>")
		Expect(ok).To(BeFalse())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetBool(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`),
		))
	})
})

func ExampleMustGetBool() {
	os.Setenv("FOO", "false")

	v, ok := config.MustGetBool(config.Environment(), "FOO")

	if !ok {
		fmt.Println("key is not defined!")
	} else if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

var _ = Describe("func MustGetBoolT()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := MustGetBoolT(b, "<key>")
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns true if the key is not defined", func() {
		b := Map{}

		v := MustGetBoolT(b, "<key>")
		Expect(v).To(BeTrue())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetBoolT(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`),
		))
	})
})

func ExampleMustGetBoolT() {
	os.Setenv("FOO", "false")

	if config.MustGetBoolT(config.Environment(), "FOO") {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

var _ = Describe("func MustGetBoolF()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := MustGetBoolF(b, "<key>")
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns false if the key is not defined", func() {
		b := Map{}

		v := MustGetBoolF(b, "<key>")
		Expect(v).To(BeFalse())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetBoolF(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`),
		))
	})
})

func ExampleMustGetBoolF() {
	os.Setenv("FOO", "true")

	if config.MustGetBoolF(config.Environment(), "FOO") {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

var _ = Describe("func MustGetBoolDefault()", func() {
	DescribeTable(
		"it returns the expected value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := MustGetBoolDefault(b, "<key>", true)
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := MustGetBoolDefault(b, "<key>", true)
		Expect(v).To(BeTrue())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetBoolDefault(b, "<key>", true)
		}).To(PanicWith(
			MatchError(`<key> is not a boolean, got "<invalid>", expected "true", "false", "yes", "no", "on" or "off"`),
		))
	})
})

func ExampleMustGetBoolDefault() {
	os.Setenv("FOO", "")

	if config.MustGetBoolDefault(config.Environment(), "FOO", true) {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}
