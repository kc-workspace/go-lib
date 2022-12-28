package xtemplates

import (
	"bytes"
	"text/template"

	"github.com/kc-workspace/go-lib/random"
)

func New(name string) *template.Template {
	return template.New(name).
		Funcs(stringFuncs).
		Funcs(numberFuncs).
		Funcs(jsonFuncs).
		Funcs(durationFuncs).
		Funcs(envFuncs).
		Option("missingkey=error")
}

func File(name string, path string) (*template.Template, error) {
	return New(name).ParseFiles(path)
}

func Buffer(content string, data interface{}, target *bytes.Buffer) error {
	var rand = random.New().FixedAlphaNumericString(7)
	var tpl, err = New(rand).Parse(content)
	if err != nil {
		return err
	}

	return tpl.Execute(target, data)
}

func Text(content string, data interface{}) (string, error) {
	var target bytes.Buffer
	var err = Buffer(content, data, &target)
	return target.String(), err
}
