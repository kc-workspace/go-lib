package xcache

import (
	"github.com/kc-workspace/go-lib/xcache/cdata"
	"github.com/kc-workspace/go-lib/xcache/cerrors"
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

func (s Service[T]) Set(key string, value T) error {
	return s.SetData(cdata.NewStatic(key, value))
}

func (s Service[T]) SetData(data cdata.BaseData[T]) error {
	var key = data.Key()

	if s.Has(key) {
		return cerrors.NewRequireForceError("create", key)
	}

	if s.setting.AutoUpdate && !data.Update() {
		return cerrors.NewUpdateFailError(key)
	}

	s.values[key] = data
	return nil
}

func (s Service[T]) Get(key string) (*T, error) {
	//nolint:wrapcheck
	return s.GetData(key).Get()
}

func (s Service[T]) GetData(key string) cdata.BaseData[T] {
	if !s.Has(key) {
		return cdata.NewEmpty[T](key)
	}

	return s.values[key]
}
