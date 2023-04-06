package xdatatype_test

import (
	"fmt"

	"github.com/kc-workspace/go-lib/xdatatype"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestObj struct {
	Name    string
	private bool
}

var _ = Describe("Converter", func() {
	var (
		str    = "hello world"
		nilArr = new([]string)
	)

	DescribeTable("StringWithBool", func(input any, expected string, ok bool) string {
		return fmt.Sprintf(
			"converting '%#v' should return [%s, %t]",
			input,
			expected,
			ok,
		)
	}, func(input any, expected string, expectedOk bool) {
		convert := xdatatype.NewStringConverter(input)
		result, ok := convert.B()

		Expect(ok).To(Equal(expectedOk))
		Expect(result).To(Equal(expected))
	},
		Entry(nil, str, str, true),
		Entry(nil, &str, str, true),
		Entry(nil, nil, "", false),
		Entry(nil, nilArr, "[]", true),
		Entry(nil, []byte("string"), "[115,116,114,105,110,103]", true),
		Entry(nil, []string{"a", "b", "c", "d"}, "[a,b,c,d]", true),
		Entry(nil, []bool{false, true, false}, "[false,true,false]", true),
		Entry(nil, map[string]string{
			"":      str,
			"first": "",
		}, "{\"\":\"hello world\",\"first\":\"\"}", true),
		Entry(nil, []TestObj{{
			Name: str, private: true,
		}, {}}, "[TestObj{\"Name\":\"hello world\"},TestObj{\"Name\":\"\"}]", true),
	)
})
