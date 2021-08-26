package config

import "fmt"

// NotDefined is an error used as a panic value when a requested key is not
// defined.
type NotDefined struct {
	Key string
}

func (e NotDefined) Error() string {
	return fmt.Sprintf("%s is not defined", e.Key)
}
