package dtconverter

import "github.com/kc-workspace/go-lib/xdatatype/dtutils"

func NewString() Converter[string] {
	return StringConverter{}
}

type StringConverter struct {
}

func (c StringConverter) WithBool(input any) (string, bool) {
	return dtutils.String(input)
}

func (c StringConverter) B(input any) (string, bool) {
	return c.WithBool(input)
}

func (c StringConverter) WithError(input any) (string, error) {
	return dtutils.StringError(input)
}

func (c StringConverter) E(input any) (string, error) {
	return c.WithError(input)
}

func (c StringConverter) WithDefault(input any, def string) string {
	return dtutils.StringDefault(input, def)
}

func (c StringConverter) D(input any, def string) string {
	return c.WithDefault(input, def)
}

func (c StringConverter) Force(input any) string {
	return dtutils.StringForce(input)
}

func (c StringConverter) F(input any) string {
	return c.Force(input)
}

func (c StringConverter) WithPointerBool(input any) (*string, bool) {
	var result, ok = c.WithBool(input)
	if !ok {
		return nil, ok
	}

	return &result, ok
}

func (c StringConverter) PB(input any) (*string, bool) {
	return c.WithPointerBool(input)
}

func (c StringConverter) WithPointerError(input any) (*string, error) {
	var result, err = c.WithError(input)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c StringConverter) PE(input any) (*string, error) {
	return c.WithPointerError(input)
}

func (c StringConverter) WithPointerDefault(input any, def string) *string {
	var result = c.WithDefault(input, def)
	return &result
}

func (c StringConverter) PD(input any, def string) *string {
	return c.WithPointerDefault(input, def)
}
