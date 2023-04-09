package dterrors_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kc-workspace/go-lib/xdatatype/dterrors"
)

type TestObject struct {
	Name string
}

var _ = Describe("ConvertFailError", func() {
	var (
		successEntries = []TableEntry{
			Entry("Empty input to string",
				"", "string",
				"Cannot convert '' to string"),
			Entry("nil input to string",
				nil, "string",
				"Cannot convert '<nil>' to string"),
			Entry("object input to int",
				TestObject{Name: "hello"}, "int",
				"Cannot convert '{hello}' to int"),
		}

		failureEntries = []TableEntry{
			Entry("unknown error type",
				errors.New("unknown error")),
		}
	)

	DescribeTable("New", func(input any, outType string, expected string) {
		var err = dterrors.NewConvertFailError(input, outType)
		Expect(err).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Success check", func(input any, outType string, expected string) {
		var err = dterrors.NewConvertFailError(input, outType)
		var actual, ok = dterrors.IsConvertFailError(err)

		Expect(ok).To(BeTrue())
		Expect(actual).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Failure check", func(err error) {
		var actual, ok = dterrors.IsConvertFailError(err)

		Expect(ok).To(BeFalse())
		Expect(actual).To(BeNil())
	},
		failureEntries,
	)
})
