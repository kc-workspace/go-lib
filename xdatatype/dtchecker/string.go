package dtchecker

import (
	"github.com/kc-workspace/go-lib/xdatatype/dtconverter"
	"github.com/kc-workspace/go-lib/xdatatype/dtutils"
)

func NewString() Checker[string] {
	return StringChecker{}
}

type StringChecker struct {
}

func (c StringChecker) Check(input any) bool {
	return dtutils.IsStringOrSimilar(input)
}

func (c StringChecker) StrictCheck(input any) bool {
	return dtutils.IsString(input)
}

func (c StringChecker) Converter() dtconverter.Converter[string] {
	return dtconverter.NewString()
}
