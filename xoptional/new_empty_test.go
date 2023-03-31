package xoptional_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kc-workspace/go-lib/xoptional"
)

var _ = Describe("Empty function", func() {
	var optional = xoptional.Empty[string]()

	It("should return true on Empty() method", func() {
		Expect(optional.Empty()).To(Equal(true))
	})
	It("should return false on Present() method", func() {
		Expect(optional.Present()).To(Equal(false))
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
