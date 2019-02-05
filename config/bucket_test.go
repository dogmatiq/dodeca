package config_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/dodeca/config"
)

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
