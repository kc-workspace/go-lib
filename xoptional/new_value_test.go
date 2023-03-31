package xoptional_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kc-workspace/go-lib/xoptional"
)

var _ = Describe("Optional.New", func() {
	Context("String", func() {
		var value = "test"
		var optional = xoptional.New(&value)

		It("should return false on Empty() method", func() {
			Expect(optional.Empty()).To(Equal(false))
		})
		It("should return true on Present() method", func() {
			Expect(optional.Present()).To(Equal(true))
		})

		It("should return value on Get()", func() {
			Expect(optional.Get()).To(Equal(value))
		})

		It("should ignore default on OrElse()", func() {
			Expect(optional.OrElse("")).To(Equal(value))
		})
	})

	Context("Object", func() {
		var value = &Mock{
			String: "test",
		}
		var optional = xoptional.New(value)

		It("should return false on Empty() method", func() {
			Expect(optional.Empty()).To(Equal(false))
		})
		It("should return true on Present() method", func() {
			Expect(optional.Present()).To(Equal(true))
		})

		It("should return value on Get()", func() {
			Expect(optional.Get().String).To(Equal(value.String))
		})

		It("should ignore default on OrElse()", func() {
			Expect(optional.OrElse(Mock{}).String).To(Equal(value.String))
		})

		Context("with Next()", func() {
			var nilObj = xoptional.Next(optional, func(t Mock) *MockObj {
				return t.NullableObject
			})

			It("should no panic when current point is nil", func() {
				Expect(nilObj.Empty()).To(Equal(true))

				var number = xoptional.Next(nilObj, func(t MockObj) *int {
					return t.NullableNumber
				})

				Expect(number.Raw()).To(BeNil())
			})

			var obj = xoptional.Next(optional, func(t Mock) *MockObj {
				return &t.Object
			})

			It("should get value from nested object", func() {
				var number = xoptional.Next(obj, func(t MockObj) *int {
					return &t.Number
				})
				Expect(number.Present()).To(Equal(true))
				Expect(number.Get()).To(Equal(0))
			})
		})
	})
})
