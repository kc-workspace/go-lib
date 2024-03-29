package xdatatype

import (
	"github.com/kc-workspace/go-lib/xdatatype/dtchecker"
	"github.com/kc-workspace/go-lib/xdatatype/dtconverter"
)

var (
	StringChecker   = dtchecker.NewString()
	StringConverter = dtconverter.NewString()
)
