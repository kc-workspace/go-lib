package cerrors_test

import (
	"errors"

	"github.com/kc-workspace/go-lib/xcache/cerrors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NoDataError", func() {
	var (
		successEntries = []TableEntry{
			Entry("No data on <empty> key",
				"", "No data found on key ''"),
			Entry("No data on test key",
				"test", "No data found on key 'test'"),
		}

		failureEntries = []TableEntry{
			Entry("unknown error type",
				errors.New("unknown error")),
		}
	)

	DescribeTable("New", func(key string, expected string) {
		var err = cerrors.NewNoDataError(key)
		Expect(err).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Success check", func(key string, expected string) {
		var err = cerrors.NewNoDataError(key)
		var actual, ok = cerrors.IsNoDataError(err)

		Expect(ok).To(BeTrue())
		Expect(actual).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Failure check", func(err error) {
		var actual, ok = cerrors.IsNoDataError(err)

		Expect(ok).To(BeFalse())
		Expect(actual).To(BeNil())
	},
		failureEntries,
	)
})
