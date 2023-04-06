package cdata

import "github.com/kc-workspace/go-lib/xcache/cerrors"

// EmptyData should implement BaseData[T]
type EmptyData[T any] struct {
	key string
}

func (d *EmptyData[T]) Key() string {
	return d.key
}

func (d *EmptyData[T]) Get() (*T, error) {
	return nil, cerrors.NewNoDataError(d.key)
}

func (d *EmptyData[T]) Force() T {
	panic(cerrors.NewNoDataError(d.key))
}

func (d *EmptyData[T]) Update() bool {
	return false
}

func (d *EmptyData[T]) Del() bool {
	return false
}

func (d *EmptyData[T]) IsEmpty() bool {
	return true
}

func (d *EmptyData[T]) IsDel() bool {
	return false
}

func (d *EmptyData[T]) IsErr() bool {
	return true
}

func (d *EmptyData[T]) IsExp() bool {
	return false
}

func (d *EmptyData[T]) HasExp() bool {
	return false
}
