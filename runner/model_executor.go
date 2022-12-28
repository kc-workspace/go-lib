package runner

type Validator func(context *Context) error

type Executor func(context *Context) error
