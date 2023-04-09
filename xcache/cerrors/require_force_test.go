package cerrors_test

import (
	"errors"

	"github.com/kc-workspace/go-lib/xcache/cerrors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RequireForceError", func() {
	var (
		successEntries = []TableEntry{
			Entry("Key empty is exist, force create is needed",
				"create", "", "require force create on key ''"),
			Entry("Key test is exist, force update is needed",
				"update", "test", "require force update on key 'test'"),
		}

		failureEntries = []TableEntry{
			Entry("unknown error type",
				errors.New("unknown error")),
		}
	)

	DescribeTable("New", func(action, key string, expected string) {
		var err = cerrors.NewRequireForceError(action, key)
		Expect(err).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Success check", func(action, key string, expected string) {
		var err = cerrors.NewRequireForceError(action, key)
		var actual, ok = cerrors.IsRequireForceError(err)

		Expect(ok).To(BeTrue())
		Expect(actual).To(MatchError(expected))
	},
		successEntries,
	)

	DescribeTable("Failure check", func(err error) {
		var actual, ok = cerrors.IsRequireForceError(err)

		Expect(ok).To(BeFalse())
		Expect(actual).To(BeNil())
	},
		failureEntries,
	)
})
