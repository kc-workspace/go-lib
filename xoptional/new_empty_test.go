package xoptional_test

import (
	"github.com/kc-workspace/go-lib/xoptional"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Empty function", func() {
	optional := xoptional.Empty[string]()

	It("should return true on Empty() method", func() {
		Expect(optional.Empty()).To(BeTrue())
	})
	It("should return false on Present() method", func() {
		Expect(optional.Present()).To(BeFalse())
	})

	It("should panic when try to Get()", func() {
		Expect(func() {
			optional.Get()
		}).To(Panic())
	})

	It("should use default value on OrElse()", func() {
		Expect(optional.OrElse("")).To(Equal(""))
	})
})
