package xcache

import (
	"github.com/kc-workspace/go-lib/xcache/cdata"
	"github.com/kc-workspace/go-lib/xcache/csetting"
)

type Service[T any] struct {
	values  map[string]cdata.BaseData[T]
	setting csetting.Setting
}

func (s Service[T]) Size() int {
	return len(s.values)
}

func (s Service[T]) Has(key string) bool {
	_, ok := s.values[key]
	return ok
}

func (s Service[T]) Set(key string, value T) bool {
	return s.SetData(key, cdata.NewStatic(key, value))
}

func (s Service[T]) SetData(key string, data cdata.BaseData[T]) bool {
	if s.Has(key) {
		return false
	}
	if s.setting.AutoUpdate && !data.Update() {
		return false
	}

	s.values[key] = data
	return true
}
