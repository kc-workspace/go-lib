package xoptional_test

import (
	"github.com/kc-workspace/go-lib/xoptional"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Optional.New", func() {
	Context("String", func() {
		value := "test"
		optional := xoptional.New(&value)

		It("should return false on Empty() method", func() {
			Expect(optional.Empty()).To(BeFalse())
		})
		It("should return true on Present() method", func() {
			Expect(optional.Present()).To(BeTrue())
		})

		It("should return value on Get()", func() {
			Expect(optional.Get()).To(Equal(value))
		})

		It("should ignore default on OrElse()", func() {
			Expect(optional.OrElse("")).To(Equal(value))
		})
	})

	Context("Object", func() {
		value := &Mock{
			String: "test",
		}
		optional := xoptional.New(value)

		It("should return false on Empty() method", func() {
			Expect(optional.Empty()).To(BeFalse())
		})
		It("should return true on Present() method", func() {
			Expect(optional.Present()).To(BeTrue())
		})

		It("should return value on Get()", func() {
			Expect(optional.Get().String).To(Equal(value.String))
		})

		It("should ignore default on OrElse()", func() {
			Expect(optional.OrElse(Mock{}).String).To(Equal(value.String))
		})

		Context("with Next()", func() {
			nilObj := xoptional.Next(optional, func(t Mock) *MockObj {
				return t.NullableObject
			})

			It("should no panic when current point is nil", func() {
				Expect(nilObj.Empty()).To(BeTrue())

				number := xoptional.Next(nilObj, func(t MockObj) *int {
					return t.NullableNumber
				})

				Expect(number.Raw()).To(BeNil())
			})

			obj := xoptional.Next(optional, func(t Mock) *MockObj {
				return &t.Object
			})

			It("should get value from nested object", func() {
				number := xoptional.Next(obj, func(t MockObj) *int {
					return &t.Number
				})
				Expect(number.Present()).To(BeTrue())
				Expect(number.Get()).To(Equal(0))
			})
		})
	})
})
