package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ExampleAsBytes() {
	os.Setenv("FOO", "<value>")

	v := config.AsBytes(config.Environment(), "FOO")

	fmt.Printf("the value is %s!\n", v)

	// Output: the value is <value>!
}

func ExampleAsBytesDefault() {
	os.Setenv("FOO", "")

	v := config.AsBytesDefault(config.Environment(), "FOO", []byte("<default>"))

	fmt.Printf("the value is %s!\n", v)

	// Output: the value is <default>!
}

var _ = Describe("func AsBytes()", func() {
	It("returns a byte-slice value", func() {
		b := Map{"<key>": String("<value>")}

		v := AsBytes(b, "<key>")
		Expect(v).To(Equal([]byte("<value>")))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsBytes(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})
})

var _ = Describe("func AsBytesDefault()", func() {
	It("returns a byte-slice value", func() {
		b := Map{"<key>": String("<value>")}

		v := AsBytesDefault(b, "<key>", []byte("<default>"))
		Expect(v).To(Equal([]byte("<value>")))
	})

	It("returns the default value key is not defined", func() {
		b := Map{}

		v := AsBytesDefault(b, "<key>", []byte("<default>"))
		Expect(v).To(Equal([]byte("<default>")))
	})
})
