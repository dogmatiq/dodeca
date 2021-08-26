package config_test

import (
	. "github.com/dogmatiq/dodeca/config"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable(
	"error types",
	func(message string, err Error) {
		Expect(err.ConfigKey()).To(Equal("<key>"))
		Expect(err.Error()).To(Equal(message))
	},
	Entry(
		"type NotDefined",
		"<key> is not defined",
		NotDefined{
			Key: "<key>",
		},
	),
	Entry(
		"type InvalidValue",
		`<key> has an invalid value ("<value>"): <explanation>`,
		InvalidValue{
			Key:         "<key>",
			Value:       "<value>",
			Explanation: "<explanation>",
		},
	),
	Entry(
		"type InvalidDefaultValue",
		`<key> has an invalid default value ("<value>"): <explanation>`,
		InvalidDefaultValue{
			Key:          "<key>",
			DefaultValue: "<value>",
			Explanation:  "<explanation>",
		},
	),
)
