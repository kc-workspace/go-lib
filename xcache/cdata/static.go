package cdata

import "time"

// StaticData should implement BaseData[T]
type StaticData[T any] struct {
	key      string
	value    T
	createAt time.Time
	modifyAt time.Time
	expireAt time.Duration
	deleteAt time.Time
}

func (d *StaticData[T]) Key() string {
	return d.key
}

func (d *StaticData[T]) Get() (*T, error) {
	return &d.value, nil
}

func (d *StaticData[T]) Force() T {
	return d.value
}

func (d *StaticData[T]) Update() bool {
	if d.IsDel() {
		return false
	}

	if d.HasExp() && !d.IsExp() {
		return false
	}

	d.modifyAt = time.Now()
	return true
}

func (d *StaticData[T]) IsEmpty() bool {
	return false
}

func (d *StaticData[T]) Del() bool {
	d.deleteAt = time.Now()
	return true
}

func (d *StaticData[T]) IsDel() bool {
	return !d.deleteAt.IsZero()
}

func (d *StaticData[T]) HasExp() bool {
	return d.expireAt > 0
}

func (d *StaticData[T]) IsExp() bool {
	if !d.HasExp() {
		return false
	}

	return d.modifyAt.Add(d.expireAt).After(time.Now())
}

func (d *StaticData[T]) IsErr() bool {
	return false
}
