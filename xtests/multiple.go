package xtests

import "testing"

type Assertions struct {
	T *testing.T
}

func (a *Assertions) New() *Assertion {
	return &Assertion{
		enabled:  true,
		name:     "",
		desc:     "",
		actual:   nil,
		expected: nil,
		err:      nil,
		T:        a.T,
	}
}

func (a *Assertions) NewName(name string) *Assertion {
	return &Assertion{
		enabled:  true,
		name:     name,
		desc:     "",
		actual:   nil,
		expected: nil,
		err:      nil,
		T:        a.T,
	}
}

func (a *Assertions) NewCase(c TestCase) bool {
	var _a = a.NewName(c.Name)
	return _a.
		WithActual(c.Actual).
		WithExpected(c.Expected).
		Must(c.Checker...)
}
