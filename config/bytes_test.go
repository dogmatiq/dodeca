package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})
})

var _ = Describe("func AsBytesDefault()", func() {
	It("returns a byte-slice value", func() {
		b := Map{"<key>": String("<value>")}

		v := AsBytesDefault(b, "<key>", []byte("<default>"))
		Expect(v).To(Equal([]byte("<value>")))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsBytesDefault(b, "<key>", []byte("<default>"))
		Expect(v).To(Equal([]byte("<default>")))
	})
})
