package cerrors_test

import (
	"errors"

	"github.com/kc-workspace/go-lib/xcache/cerrors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateFailError", func() {
	var (
		successEntries = []TableEntry{
			Entry("Key empty is exist, force create is needed",
				"", "cannot update key ''"),
			Entry("Key test is exist, force update is needed",
				"test", "cannot update key 'test'"),
		}

		failureEntries = []TableEntry{
			Entry("unknown error type",
				errors.New("unknown error")),
		}
	)

	DescribeTable("New", func(key string, expected string) {
		var err = cerrors.NewUpdateFailError(key)
		Expect(err).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Success check", func(key string, expected string) {
		var err = cerrors.NewUpdateFailError(key)
		var actual, ok = cerrors.IsUpdateFailError(err)

		Expect(ok).To(BeTrue())
		Expect(actual).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Failure check", func(err error) {
		var actual, ok = cerrors.IsUpdateFailError(err)

		Expect(ok).To(BeFalse())
		Expect(actual).To(BeNil())
	},
		failureEntries,
	)
})
