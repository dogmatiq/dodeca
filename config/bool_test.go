package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dogmatiq/dodeca/config"
)

func ExampleGetBoolT() {
	os.Setenv("FOO", "false")

	v, err := config.GetBoolT(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

func ExampleGetBoolT_withUndefinedVariable() {
	os.Setenv("FOO", "")

	v, err := config.GetBoolT(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

func ExampleGetBoolF() {
	os.Setenv("FOO", "true")

	v, err := config.GetBoolF(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

func ExampleGetBoolF_withUndefinedVariable() {
	os.Setenv("FOO", "")

	v, err := config.GetBoolF(config.Environment(), "FOO")
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

func ExampleGetBoolDefault() {
	os.Setenv("FOO", "false")

	v, err := config.GetBoolDefault(config.Environment(), "FOO", true)
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

func ExampleGetBoolDefault_withUndefinedVariable() {
	os.Setenv("FOO", "")

	v, err := config.GetBoolDefault(config.Environment(), "FOO", true)
	if err != nil {
		panic(err)
	}

	if v {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

func TestGetBoolDefault_withInvalidValue(t *testing.T) {
	os.Setenv("FOO", "<invalid>")

	_, err := config.GetBoolDefault(config.Environment(), "FOO", true)

	if err == nil || err.Error() != "FOO is not a valid boolean value: <invalid>" {
		t.Fatal("unexpected error, got:", err)
	}
}

func ExampleMustGetBoolT() {
	os.Setenv("FOO", "false")

	if config.MustGetBoolT(config.Environment(), "FOO") {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

func ExampleMustGetBoolT_withUndefinedVariable() {
	os.Setenv("FOO", "")

	if config.MustGetBoolT(config.Environment(), "FOO") {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

func ExampleMustGetBoolF() {
	os.Setenv("FOO", "true")

	if config.MustGetBoolF(config.Environment(), "FOO") {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

func ExampleMustGetBoolF_withUndefinedVariable() {
	os.Setenv("FOO", "")

	if config.MustGetBoolF(config.Environment(), "FOO") {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

func ExampleMustGetBoolDefault() {
	os.Setenv("FOO", "false")

	if config.MustGetBoolDefault(config.Environment(), "FOO", true) {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: false!
}

func ExampleMustGetBoolDefault_withUndefinedVariable() {
	os.Setenv("FOO", "")

	if config.MustGetBoolDefault(config.Environment(), "FOO", true) {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}

	// Output: true!
}

func TestMustGetBoolDefault_withInvalidValue(t *testing.T) {
	defer func() {
		r := recover()

		err, ok := r.(error)
		if !ok {
			panic(r)
		}

		if err.Error() != "FOO is not a valid boolean value: <invalid>" {
			t.Fatal("unexpected error, got:", err)
		}
	}()

	os.Setenv("FOO", "<invalid>")

	config.MustGetBoolDefault(config.Environment(), "FOO", true)
}
