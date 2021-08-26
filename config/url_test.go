package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func AsURL()", func() {
	It("returns a URL value", func() {
		b := Map{"<key>": String("http://www.example.com/path")}

		v := AsURL(b, "<key>")
		Expect(v.String()).To(Equal("http://www.example.com/path"))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsURL(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String(":")}

		Expect(func() {
			AsURL(b, "<key>")
		}).To(PanicWith(`expected <key> to be a URL: parse ":": missing protocol scheme`))
	})
})

var _ = Describe("func AsURLDefault()", func() {
	It("returns a URL value", func() {
		b := Map{"<key>": String("http://www.example.com/path")}

		v := AsURLDefault(b, "<key>", "http://www.example.org/default")
		Expect(v.String()).To(Equal("http://www.example.com/path"))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsURLDefault(b, "<key>", "http://www.example.org/default")
		Expect(v.String()).To(Equal("http://www.example.org/default"))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String(":")}

		Expect(func() {
			AsURLDefault(b, "<key>", "http://www.example.org/default")
		}).To(PanicWith(`expected <key> to be a URL: parse ":": missing protocol scheme`))
	})

	It("panics if the default value cannot be parsed", func() {
		b := Map{}

		Expect(func() {
			AsURLDefault(b, "<key>", ":")
		}).To(PanicWith(`expected the default value for <key> to be a URL: parse ":": missing protocol scheme`))
	})
})
