package config_test

import (
	"os"
	"testing"

	. "github.com/dogmatiq/dodeca/config"
)

func TestValue_AsReader_withEmptyValue(t *testing.T) {
	_, err := Value{}.AsReader()

	if !os.IsNotExist(err) {
		t.Fatal("expected a not-exist error")
	}
}

func TestValue_AsPath_withEmptyValue(t *testing.T) {
	_, _, err := Value{}.AsPath()

	if !os.IsNotExist(err) {
		t.Fatal("expected a not-exist error")
	}
}

func TestValue_AsString_withEmptyValue(t *testing.T) {
	v, err := Value{}.AsString()

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if v != "" {
		t.Fatal("unexpected value", v)
	}
}

func TestValue_AsBytes_withEmptyValue(t *testing.T) {
	v, err := Value{}.AsBytes()

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(v) != 0 {
		t.Fatal("unexpected value", v)
	}
}
