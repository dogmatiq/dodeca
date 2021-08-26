package config_test

import (
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

var _ = Describe("func AsBool()", func() {
	DescribeTable(
		"it returns a boolean value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := AsBool(b, "<key>")
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsBool(b, "<key>")
		}).To(PanicWith(NotDefined{Key: "<key>"}))
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsBool(b, "<key>")
		}).To(PanicWith(`expected <key> to be a boolean ("true", "false", "yes", "no", "on" or "off"), got "<invalid>"`))
	})
})

var _ = Describe("func AsBoolF()", func() {
	DescribeTable(
		"it returns a boolean value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := AsBoolF(b, "<key>")
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns true if the key is not defined", func() {
		b := Map{}

		v := AsBoolT(b, "<key>")
		Expect(v).To(BeTrue())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsBoolT(b, "<key>")
		}).To(PanicWith(`expected <key> to be a boolean ("true", "false", "yes", "no", "on" or "off"), got "<invalid>"`))
	})
})

var _ = Describe("func AsBoolF()", func() {
	DescribeTable(
		"it returns a boolean value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := AsBoolF(b, "<key>")
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns false if the key is not defined", func() {
		b := Map{}

		v := AsBoolF(b, "<key>")
		Expect(v).To(BeFalse())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsBoolF(b, "<key>")
		}).To(PanicWith(`expected <key> to be a boolean ("true", "false", "yes", "no", "on" or "off"), got "<invalid>"`))
	})
})

var _ = Describe("func AsBoolDefault()", func() {
	DescribeTable(
		"it returns a boolean value",
		func(str string, expect bool) {
			b := Map{"<key>": String(str)}

			v := AsBoolDefault(b, "<key>", true)
			Expect(v).To(Equal(expect))
		},
		boolTableEntries...,
	)

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsBoolDefault(b, "<key>", true)
		Expect(v).To(BeTrue())
	})

	It("panics if the value is not one of the accepted values", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			AsBoolDefault(b, "<key>", true)
		}).To(PanicWith(`expected <key> to be a boolean ("true", "false", "yes", "no", "on" or "off"), got "<invalid>"`))
	})
})
