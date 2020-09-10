package config_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dogmatiq/dodeca/config"
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This example demonstrates how to specify a configuration value as a string
// and consume it using an io.ReadCloser.
func ExampleValue_AsReader_specifiedAsString() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a string, and the EXAMPLE__DATASOURCE specifies
	// the source type as "string:plain".
	//
	// This is the default source type. EXAMPLE__DATASOURCE can be left
	// undefined and the behavior would be the same.
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "<the configuration value>")
	os.Setenv("EXAMPLE__DATASOURCE", "string:plain")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Consume the value as an io.ReadCloser.
	r, err := value.AsReader()
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// Read the entire value into a buffer.
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	// Print the buffer.
	fmt.Println(string(buf))

	// Output: <the configuration value>
}

// This example demonstrates how to specify a configuration value as a string
// and consume it as a path to a temporary file containing the string's content.
func ExampleValue_AsPath_specifiedAsString() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a string, and the EXAMPLE__DATASOURCE specifies
	// the source type as "string:plain".
	//
	// This is the default source type. EXAMPLE__DATASOURCE can be left
	// undefined and the behavior would be the same.
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "<the configuration value>")
	os.Setenv("EXAMPLE__DATASOURCE", "string:plain")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Obtain the path to the configuration file.
	//
	// Because the configuration value is specified as a string, Path() returns
	// the path to a temporary file.
	p, c, err := value.AsPath()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Read the temporary file into a buffer.
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	// Print the value as a string.
	fmt.Println(string(buf))

	// Output: <the configuration value>
}

// This example demonstrates how to specify a configuration value as a string
// and consume it as string.
func ExampleValue_AsString_specifiedAsString() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a string, and the EXAMPLE__DATASOURCE specifies
	// the source type as "string:plain".
	//
	// This is the default source type. EXAMPLE__DATASOURCE can be left
	// undefined and the behavior would be the same.
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "<the configuration value>")
	os.Setenv("EXAMPLE__DATASOURCE", "string:plain")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Get the configuration value as a string.
	str, err := value.AsString()
	if err != nil {
		panic(err)
	}

	// Print the value.
	fmt.Println(str)

	// Output: <the configuration value>
}

// This example demonstrates how to specify a configuration value as a string
// and consume it as a byte-slice.
func ExampleValue_AsBytes_specifiedAsString() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a string, and the EXAMPLE__DATASOURCE specifies
	// the source type as "string:plain".
	//
	// This is the default source type. EXAMPLE__DATASOURCE can be left
	// undefined and the behavior would be the same.
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "<the configuration value>")
	os.Setenv("EXAMPLE__DATASOURCE", "string:plain")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Get the configuration value as a byte-slice.
	buf, err := value.AsBytes()
	if err != nil {
		panic(err)
	}

	// Print the value as a string.
	fmt.Println(string(buf))

	// Output: <the configuration value>
}

var _ = Describe("func AsString()", func() {
	It("returns a string value", func() {
		b := Map{"<key>": String("<value>")}

		v := AsString(b, "<key>")
		Expect(v).To(Equal("<value>"))
	})

	It("panics if the key is not defined", func() {
		b := Map{}

		Expect(func() {
			AsString(b, "<key>")
		}).To(PanicWith(`<key> is not defined`))
	})
})

var _ = Describe("func AsStringDefault()", func() {
	It("returns a string value", func() {
		b := Map{"<key>": String("<value>")}

		v := AsStringDefault(b, "<key>", "<default>")
		Expect(v).To(Equal("<value>"))
	})

	It("returns the default value if the key is not defined", func() {
		b := Map{}

		v := AsStringDefault(b, "<key>", "<default>")
		Expect(v).To(Equal("<default>"))
	})
})
