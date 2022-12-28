package xtemplates

import (
	"text/template"

	"github.com/kc-workspace/go-lib/mapper"
)

func json(m interface{}) (string, error) {
	var a, e = mapper.ToJson(m)
	return string(a), e
}

func indentJson(m interface{}) (string, error) {
	var a, e = mapper.ToFormatJson(m)
	return string(a), e
}

var jsonFuncs template.FuncMap = map[string]interface{}{
	"json":       json,
	"indentJson": indentJson,
}
