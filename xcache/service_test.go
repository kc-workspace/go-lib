package xcache_test

import (
	"github.com/kc-workspace/go-lib/xcache"
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
	It("create new", func() {
		service := xcache.New[string](csetting.Setting{
			AutoUpdate: false,
		})

		Expect(service).NotTo(BeNil())
		Expect(service.Size()).To(Equal(0))
		Expect(service.Has("invalid")).To(BeFalse())
	})

	It("create new", func() {
	})
})
