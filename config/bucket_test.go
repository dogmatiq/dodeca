package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Map", func() {
	var bucket Map

	BeforeEach(func() {
		bucket = Map{
			"<key-1>": String("<value-1>"),
			"<key-2>": String("<value-2>"),
			"<key-3>": Value{},
		}
	})

	Describe("func Get()", func() {
		It("returns the value associated with the key", func() {
			v := bucket.Get("<key-1>")
			Expect(v).To(Equal(String("<value-1>")))
		})
	})

	Describe("func GetDefault()", func() {
		It("returns the value associated with the key", func() {
			v := bucket.GetDefault("<key-1>", "<default>")
			Expect(v).To(Equal(String("<value-1>")))
		})

		It("returns the default value if the key is undefined", func() {
			v := bucket.GetDefault("<undefined>", "<default>")
			Expect(v).To(Equal(String("<default>")))
		})

		It("returns the default value if the value is the zero-value", func() {
			v := bucket.GetDefault("<key-3>", "<default>")
			Expect(v).To(Equal(String("<default>")))
		})
	})

	Describe("func Each()", func() {
		It("invokes the function for each key/value pair", func() {
			calls := map[string]Value{}

			fn := func(k string, v Value) bool {
				calls[k] = v
				return true
			}

			Expect(bucket.Each(fn)).To(BeTrue())
			Expect(calls).To(BeEquivalentTo(bucket))
		})

		It("stops iterating if the function returns false", func() {
			count := 0

			fn := func(k string, v Value) bool {
				count++
				return false
			}

			Expect(bucket.Each(fn)).To(BeFalse())
			Expect(count).To(Equal(1))
		})
	})
})
