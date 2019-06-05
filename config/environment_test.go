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

func ExampleEnvironment_get() {
	os.Setenv("FOO", "<foo>")

	v := config.Environment().Get("FOO")

	fmt.Println(v)

	// Output: <foo>
}

func ExampleEnvironment_getWithUndefinedVariable() {
	os.Setenv("FOO", "")

	v := config.Environment().Get("FOO")

	fmt.Println(v.IsEmpty())

	// Output: true
}

func ExampleEnvironment_getDefault() {
	os.Setenv("FOO", "<foo>")

	v := config.Environment().GetDefault("FOO", "<default>")

	fmt.Println(v)

	// Output: <foo>
}

func ExampleEnvironment_getDefaultWithUndefinedVariable() {
	os.Setenv("FOO", "")

	v := config.Environment().GetDefault("FOO", "<default>")

	fmt.Println(v)

	// Output: <default>
}

func ExampleEnvironment_each() {
	os.Setenv("FOO", "<foo>")
	os.Setenv("BAR", "<bar>")

	config.Environment().Each(
		func(k string, v config.Value) bool {
			if k == "FOO" || k == "BAR" {
				fmt.Printf("%s=%s\n", k, v)
			}

			return true
		},
	)

	// Output: FOO=<foo>
	// BAR=<bar>
}
