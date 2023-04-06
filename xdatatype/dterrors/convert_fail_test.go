package dterrors_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kc-workspace/go-lib/xdatatype/dterrors"
)

type TestObject struct {
	Name string
}

var _ = DescribeTable("ConvertFail", func(input any, outType string, expected string) string {
	return fmt.Sprintf(
		"ConvertFail('%v', '%s') should return '%s'",
		input,
		outType,
		expected,
	)
}, func(input any, outType string, expected string) {
	var err = dterrors.NewConvertFailError(input, outType)
	Expect(err).To(MatchError(expected))
},
	Entry(nil, "", "string", "Cannot convert '' to string"),
	Entry(nil, nil, "string", "Cannot convert '<nil>' to string"),
	Entry(nil, TestObject{Name: "hello"}, "int", "Cannot convert '{hello}' to int"),
)
