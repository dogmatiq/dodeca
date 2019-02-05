package config_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dogmatiq/dodeca/config"
)

// This example demonstrates how to specify a configuration value as a
// base-64 encoded string and consume it using an io.ReadCloser.
func ExampleValue_base64AsReader() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a hexadecimal string, and the EXAMPLE__DATASOURCE
	// specifies the source type as "string:base64".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", base64.StdEncoding.EncodeToString([]byte("<the configuration value>")))
	os.Setenv("EXAMPLE__DATASOURCE", "string:base64")

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

// This example demonstrates how to specify a configuration value as a
// base-64 encoded string and consume it as a path to a temporary file containing
// the string's content.
func ExampleValue_base64AsPath() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a hexadecimal string, and the EXAMPLE__DATASOURCE
	// specifies the source type as "string:base64".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", base64.StdEncoding.EncodeToString([]byte("<the configuration value>")))
	os.Setenv("EXAMPLE__DATASOURCE", "string:base64")

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

// This example demonstrates how to specify a configuration value as a
// base-64 encoded string and consume it as string.
func ExampleValue_base64AsString() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a hexadecimal string, and the EXAMPLE__DATASOURCE
	// specifies the source type as "string:base64".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", base64.StdEncoding.EncodeToString([]byte("<the configuration value>")))
	os.Setenv("EXAMPLE__DATASOURCE", "string:base64")

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

// This example demonstrates how to specify a configuration value as a
// base-64 encoded string and consume it as a byte-slice.
func ExampleValue_base64AsBytes() {
	// Setup the environment such that the EXAMPLE variable contains the
	// configuration value as a hexadecimal string, and the EXAMPLE__DATASOURCE
	// specifies the source type as "string:base64".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", base64.StdEncoding.EncodeToString([]byte("<the configuration value>")))
	os.Setenv("EXAMPLE__DATASOURCE", "string:base64")

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
