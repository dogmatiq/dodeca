package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func GetUint64()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, ok, err := GetUint64(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(123))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok, err := GetUint64(b, "<key>")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ok).To(BeFalse())
	})

	It("returns an error if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, _, err := GetUint64(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`))
	})

	It("returns an error if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		_, _, err := GetUint64(b, "<key>")
		Expect(err).To(MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "-123": invalid syntax`))
	})
})

var _ = Describe("func GetUint64Default()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, err := GetUint64Default(b, "<key>", 10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v, err := GetUint64Default(b, "<key>", 10)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(v).To(BeEquivalentTo(10))
	})

	It("returns an error if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		_, err := GetUint64Default(b, "<key>", 10)
		Expect(err).To(MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`))
	})

	It("returns an error if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		_, err := GetUint64Default(b, "<key>", 10)
		Expect(err).To(MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "-123": invalid syntax`))
	})
})

var _ = Describe("func MustGetUint64()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v, ok := MustGetUint64(b, "<key>")
		Expect(v).To(BeEquivalentTo(123))
		Expect(ok).To(BeTrue())
	})

	It("sets ok to false if the key is not defined", func() {
		b := Map{}

		_, ok := MustGetUint64(b, "<key>")
		Expect(ok).To(BeFalse())
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetUint64(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`),
		))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			MustGetUint64(b, "<key>")
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "-123": invalid syntax`),
		))
	})
})

var _ = Describe("func MustGetUint64Default()", func() {
	It("returns a positive integer value", func() {
		b := Map{"<key>": String("123")}

		v := MustGetUint64Default(b, "<key>", 10)
		Expect(v).To(BeEquivalentTo(123))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := MustGetUint64Default(b, "<key>", 10)
		Expect(v).To(BeEquivalentTo(10))
	})

	It("panics if the value cannot be parsed", func() {
		b := Map{"<key>": String("<invalid>")}

		Expect(func() {
			MustGetUint64Default(b, "<key>", 10)
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "<invalid>": invalid syntax`),
		))
	})

	It("panics if the value is negative", func() {
		b := Map{"<key>": String("-123")}

		Expect(func() {
			MustGetUint64Default(b, "<key>", 10)
		}).To(PanicWith(
			MatchError(`<key> is not a valid unsigned 64-bit integer: strconv.ParseUint: parsing "-123": invalid syntax`),
		))
	})
})
