package config_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dogmatiq/dodeca/config"
)

// This example demonstrates how to specify a configuration value as a file and
// consume it using an io.ReadCloser.
func ExampleValue_file_as_reader() {
	// Setup the environment such that the EXAMPLE variable contains the path to
	// a configuration file, and the EXAMPLE__DATASOURCE specifies the source
	// type as "file".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "./testdata/example.json")
	os.Setenv("EXAMPLE__DATASOURCE", "file")

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

	// Output: {"example_config": true}
}

// This example demonstrates how to specify a configuration value as a file and
// consume it as a path to that file.
func ExampleValue_file_as_path() {
	// Setup the environment such that the EXAMPLE variable contains the path to
	// a configuration file, and the EXAMPLE__DATASOURCE specifies the source
	// type as "file".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "./testdata/example.json")
	os.Setenv("EXAMPLE__DATASOURCE", "file")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Obtain the path to the configuration file.
	//
	// Because the configuration value is specified as a file, Path() returns
	// the path to the original file. If the configuration value were specified
	// by some other means, this would be a path to a temporary file.
	p, c, err := value.AsPath()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Print the path.
	fmt.Println(p)

	// Output: ./testdata/example.json
}

// This example demonstrates how to specify a configuration value as a file and
// consume it as string.
func ExampleValue_file_as_string() {
	// Setup the environment such that the EXAMPLE variable contains the path to
	// a configuration file, and the EXAMPLE__DATASOURCE specifies the source
	// type as "file".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "./testdata/example.json")
	os.Setenv("EXAMPLE__DATASOURCE", "file")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Get the configuration value as a string.
	str, err := value.AsString()
	if err != nil {
		panic(err)
	}

	// Print the value.
	fmt.Println(str)

	// Output: {"example_config": true}
}

// This example demonstrates how to specify a configuration value as a file and
// consume it as a byte-slice.
func ExampleValue_file_as_bytes() {
	// Setup the environment such that the EXAMPLE variable contains the path to
	// a configuration file, and the EXAMPLE__DATASOURCE specifies the source
	// type as "file".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "./testdata/example.json")
	os.Setenv("EXAMPLE__DATASOURCE", "file")

	// Get the configuration value from the Environment bucket.
	value := config.Environment().Get("EXAMPLE")

	// Get the configuration value as a byte-slice.
	buf, err := value.AsBytes()
	if err != nil {
		panic(err)
	}

	// Print the value as a string.
	fmt.Println(string(buf))

	// Output: {"example_config": true}
}
