package cerrors_test

import (
	"errors"

	"github.com/kc-workspace/go-lib/xcache/cerrors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NoItemError", func() {
	var (
		successEntries = []TableEntry{
			Entry("No item on <empty> key",
				"", "No item found on data key ''"),
			Entry("No item on test key",
				"test", "No item found on data key 'test'"),
		}

		failureEntries = []TableEntry{
			Entry("unknown error type",
				errors.New("unknown error")),
		}
	)

	DescribeTable("New", func(key string, expected string) {
		var err = cerrors.NewNoItemError(key)
		Expect(err).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Success check", func(key string, expected string) {
		var err = cerrors.NewNoItemError(key)
		var actual, ok = cerrors.IsNoItemError(err)

		Expect(ok).To(BeTrue())
		Expect(actual).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Failure check", func(err error) {
		var actual, ok = cerrors.IsNoItemError(err)

		Expect(ok).To(BeFalse())
		Expect(actual).To(BeNil())
	},
		failureEntries,
	)
})
