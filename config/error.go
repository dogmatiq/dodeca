package config

import "fmt"

// KeyError is an error that indicates a problem with the value associated with
// a specific key.
type KeyError interface {
	error

	// ConfigKey returns the config key that the error relates to.
	ConfigKey() string
}

// NotDefined is an error used as a panic value when a requested key is not
// defined.
type NotDefined struct {
	Key string
}

// ConfigKey returns the config key that the error relates to.
func (e NotDefined) ConfigKey() string {
	return e.Key
}

func (e NotDefined) Error() string {
	return fmt.Sprintf("%s is not defined", e.Key)
}

// InvalidValue is an error used as a panic value when the value associated with
// a key is not well-formed or is otherwise invalid.
type InvalidValue struct {
	Key         string
	Value       string
	Explanation string
}

// ConfigKey returns the config key that the error relates to.
func (e InvalidValue) ConfigKey() string {
	return e.Key
}

func (e InvalidValue) Error() string {
	return fmt.Sprintf(
		"%s has an invalid value (%#v): %s",
		e.Key,
		e.Value,
		e.Explanation,
	)
}

// InvalidDefaultValue is an error used as a panic value when the default value
// associated with a key is not well-formed or is otherwise invalid.
type InvalidDefaultValue struct {
	Key          string
	DefaultValue string
	Explanation  string
}

// ConfigKey returns the config key that the error relates to.
func (e InvalidDefaultValue) ConfigKey() string {
	return e.Key
}

func (e InvalidDefaultValue) Error() string {
	return fmt.Sprintf(
		"%s has an invalid default value (%#v): %s",
		e.Key,
		e.DefaultValue,
		e.Explanation,
	)
}
