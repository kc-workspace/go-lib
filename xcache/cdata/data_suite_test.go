package cdata_test

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
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

func TestXcacheData(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Xcache.Data Suite")
}
