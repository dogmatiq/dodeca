package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
)

func ExampleBucket_Get() {
	os.Setenv("FOO", "<foo>")
	os.Setenv("BAR", "<bar>")

	v := config.Environment().Get("FOO")

	fmt.Println(v)

	// Output: <foo>
}

func ExampleBucket_Get_undefinedVariable() {
	os.Setenv("FOO", "")

	v := config.Environment().Get("FOO")

	fmt.Println(v.IsEmpty())

	// Output: true
}

func ExampleBucket_GetDefault() {
	os.Setenv("FOO", "<foo>")
	os.Setenv("BAR", "<bar>")

	v := config.Environment().GetDefault("FOO", "<default>")

	fmt.Println(v)

	// Output: <foo>
}

func ExampleBucket_GetDefault_undefinedVariablealue() {
	os.Setenv("FOO", "")

	v := config.Environment().GetDefault("FOO", "<default>")

	fmt.Println(v)

	// Output: <default>
}

func ExampleBucket_Each() {
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
