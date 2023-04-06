package xdatatype

import "github.com/kc-workspace/go-lib/xdatatype/dtmodels"

func NewStringConverter(input any) dtmodels.Converter[string] {
	return dtmodels.NewString(input)
}
