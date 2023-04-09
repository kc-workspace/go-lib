package xerror_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// xer := xerror.Cast(err)
// xer := xerror.Builder()
//          .Category(0, "General")
//          .Key(5, "Unknown")
//          .Message("{key.name} ({key.id}) something went wrong [category={category.id}]")

// Traditional error message supported by golang
// xer.Error()

var _ = Describe("Error", func() {
	It("test", func() {
		Expect("").To(Equal(""))
	})
})
