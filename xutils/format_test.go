package xutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// xutils.Nprintf("hello {name}", xmap.Empty().Set("name", "world"))

var _ = Describe("Format", func() {
	It("test", func() {
		Expect("").To(Equal(""))
	})
})
