package cdata_test

import (
	"github.com/kc-workspace/go-lib/xcache/cdata"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StaticData", func() {
	Describe("create", func() {
		It("string data", func() {
			var data = cdata.NewStatic("key", "string")

			Expect(data).ToNot(BeNil())
			Expect(data.Force()).To(Equal("string"))
		})

		It("object data", func() {
			var data = cdata.NewStatic("key", Mock{})

			Expect(data).ToNot(BeNil())
			Expect(data.Force()).To(Equal(Mock{}))
		})

		It("nil data", func() {
			var data = cdata.NewStatic[*Mock]("key", nil)

			Expect(data).ToNot(BeNil())
			Expect(data.Force()).To(BeNil())
		})
	})
})
