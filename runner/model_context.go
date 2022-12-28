package runner

import "github.com/kc-workspace/go-lib/logger"

type Context struct {
	Disabled bool
	Logger   *logger.Logger

	input  interface{}
	output interface{}
}

func (c *Context) In(i interface{}) {
	c.input = i
}

func (c *Context) Input() interface{} {
	return c.input
}

func (c *Context) Out(o interface{}) {
	c.output = o
}

func (c *Context) Output() interface{} {
	return c.output
}

func NewContext(name string) *Context {
	return &Context{
		Disabled: false,
		Logger:   logger.Get("runner", name),
	}
}
