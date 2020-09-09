package config_test

import (
	"bytes"
	"os"
	"testing"

	. "github.com/dogmatiq/dodeca/config"
)

func TestValue_AsReader_withZeroValue(t *testing.T) {
	_, err := Value{}.AsReader()

	if !os.IsNotExist(err) {
		t.Fatal("expected a not-exist error")
	}
}

func TestValue_AsPath_withZeroValue(t *testing.T) {
	_, _, err := Value{}.AsPath()

	if !os.IsNotExist(err) {
		t.Fatal("expected a not-exist error")
	}
}

func TestValue_AsString_withZeroValue(t *testing.T) {
	_, err := Value{}.AsString()

	if err == nil || err.Error() != "cannot represent a zero-value as a string" {
		t.Fatal("unexpected error, got:", err)
	}
}

func TestValue_AsBytes_withZeroValue(t *testing.T) {
	_, err := Value{}.AsBytes()

	if err == nil || err.Error() != "cannot represent a zero-value as a byte-slice" {
		t.Fatal("unexpected error, got:", err)
	}
}

func TestValue_String(t *testing.T) {
	v := String("<value>").String()

	if v != "<value>" {
		t.Fatalf("unexpected value: %s", v)
	}
}

func TestValue_String_withZeroValue(t *testing.T) {
	defer func() {
		p := recover()

		switch e := p.(type) {
		case error:
			if e.Error() != "cannot represent a zero-value as a string" {
				t.Fatalf("expected panic did not occur, got: %s", p)
			}
		default:
			t.Fatal("expected panic did not occur")
		}
	}()

	_ = Value{}.String()
}

func TestValue_Bytes(t *testing.T) {
	v := String("<value>").Bytes()

	if !bytes.Equal(v, []byte("<value>")) {
		t.Fatalf("unexpected value: %s", v)
	}
}

func TestValue_Bytes_withZeroValue(t *testing.T) {
	defer func() {
		p := recover()

		switch e := p.(type) {
		case error:
			if e.Error() != "cannot represent a zero-value as a byte-slice" {
				t.Fatalf("expected panic did not occur, got: %s", p)
			}
		default:
			t.Fatal("expected panic did not occur")
		}
	}()

	_ = Value{}.Bytes()
}
