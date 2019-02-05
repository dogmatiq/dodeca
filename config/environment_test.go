package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
)

func ExampleGetEnv() {
	// Setup the environment such that the EXAMPLE variable contains the path to
	// a configuration file, and the EXAMPLE__DATASOURCE specifies the source
	// type as "file".
	//
	// Outside of an example these environment variables would be set in the
	// operating system shell, and the path would refer to a real configuration
	// file.
	os.Setenv("EXAMPLE", "./testdata/example.json")
	os.Setenv("EXAMPLE__DATASOURCE", "file")

	fmt.Println(
		config.GetEnv("EXAMPLE"),
	)

	// Output: {"example_config": true}
}
