package xoptional_test

import (
	"testing"

	ginkgo "github.com/onsi/ginkgo/v2"
	gomega "github.com/onsi/gomega"
)

type MockObj struct {
	NullableNumber *int
	Number         int
}

type Mock struct {
	NullableString *string
	String         string
	NullableObject *MockObj
	Object         MockObj
}

func TestXoptional(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Xoptional Suite")
}
