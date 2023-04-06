package xcache

import (
	"github.com/kc-workspace/go-lib/xcache/cdata"
	"github.com/kc-workspace/go-lib/xcache/csetting"
)

func New[T any](setting csetting.Setting) Service[T] {
	return Service[T]{
		values:  make(map[string]cdata.BaseData[T]),
		setting: setting,
	}
}

func NewAny(setting csetting.Setting) Service[any] {
	return Service[any]{
		values:  make(map[string]cdata.BaseData[any]),
		setting: setting,
	}
}
