package dtchecker

import "github.com/kc-workspace/go-lib/xdatatype/dtconverter"

type Checker[T any] interface {
	// Check will perform soft check
	// if input convertable to checker type or not
	Check(input any) bool

	// StrictCheck will stricter check
	// if input must be checker type only
	StrictCheck(input any) bool

	// Get converter of current checker
	Converter() dtconverter.Converter[T]
}
