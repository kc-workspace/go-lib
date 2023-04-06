package dtmodels

import "github.com/kc-workspace/go-lib/xdatatype/dtutils"

func NewString(input any) Converter[string] {
	return StringConverter{
		raw: input,
	}
}

type StringConverter struct {
	raw any
}

func (c StringConverter) WithBool() (string, bool) {
	return dtutils.String(c.raw)
}

func (c StringConverter) B() (string, bool) {
	return c.WithBool()
}

func (c StringConverter) WithError() (string, error) {
	return dtutils.StringError(c.raw)
}

func (c StringConverter) E() (string, error) {
	return c.WithError()
}

func (c StringConverter) WithDefault(def string) string {
	return dtutils.StringDefault(c.raw, def)
}

func (c StringConverter) D(def string) string {
	return c.WithDefault(def)
}

func (c StringConverter) Force() string {
	return dtutils.StringForce(c.raw)
}

func (c StringConverter) F() string {
	return c.Force()
}

func (c StringConverter) WithPointerBool() (*string, bool) {
	var result, ok = c.WithBool()
	if !ok {
		return nil, ok
	}

	return &result, ok
}

func (c StringConverter) PB() (*string, bool) {
	return c.WithPointerBool()
}

func (c StringConverter) WithPointerError() (*string, error) {
	var result, err = c.WithError()
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c StringConverter) PE() (*string, error) {
	return c.WithPointerError()
}

func (c StringConverter) WithPointerDefault(def string) *string {
	var result = c.WithDefault(def)
	return &result
}

func (c StringConverter) PD(def string) *string {
	return c.WithPointerDefault(def)
}
