package xcache_test

import (
	"github.com/kc-workspace/go-lib/xcache"
	"github.com/kc-workspace/go-lib/xcache/cdata"
	"github.com/kc-workspace/go-lib/xcache/csetting"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// service.Set("<key>", "<value>")          // boolean
// service.SetData("<key>", &xcache.Data{}) // boolean
// service.Update("<key>")                  // -> boolean
// service.Get("<key>")                     // -> "<value>"
// service.GetData("<key>")                 // -> *xcache.Data
// service.Has("<key>")                     // -> boolean
// service.IsExp("<key>")                   // -> boolean
// service.IsDel("<key>")                   // -> boolean
// service.Del("<key>")                     // -> boolean
// service.Remove("<key>")                  // -> boolean

var _ = Describe("Service", func() {
	Describe("Typed", func() {
		It("New", func() {
			service := xcache.New[string](csetting.Setting{
				AutoUpdate: false,
			})

			Expect(service).NotTo(BeNil())
			Expect(service.Size()).To(Equal(0))
			Expect(service.Has("invalid")).To(BeFalse())
			Expect(service.Get("invalid")).Error().
				Should(HaveOccurred())
		})

		It("Static Set", func() {
			service := xcache.New[string](csetting.Setting{
				AutoUpdate: false,
			})

			Expect(service.Set("test", "string")).To(Succeed())
			Expect(service.Size()).To(Equal(1))
			Expect(service.Get("test")).To(HaveValue(Equal("string")))
		})

		It("Data Set", func() {
			service := xcache.New[string](csetting.Setting{
				AutoUpdate: false,
			})

			Expect(service.SetData(cdata.NewStatic("test", "string"))).
				To(Succeed())

			Expect(service.Size()).To(Equal(1))
			Expect(service.Get("test")).To(HaveValue(Equal("string")))
		})
	})
})
