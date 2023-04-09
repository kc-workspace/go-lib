package xmap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// mapper = xmap.Empty()
// mapper = xmap.New(make(map[string]any))
// mapper = xmap.New(map[string]string{"test":"hello world"})

// mapper.Set("test", "true")
// mapper.Set("message.string", "hello world")

var _ = Describe("Mapper", func() {
	It("test", func() {
		Expect("").To(Equal(""))
	})
})
