package cdata

import "time"

// DynamicData should implement BaseData[T]
type DynamicData[T any] struct {
	key      string
	value    *T
	err      error
	modifier Modifier[T]
	createAt time.Time
	modifyAt time.Time
	expireAt time.Duration
	deleteAt time.Time
}

func (d *DynamicData[T]) Key() string {
	return d.key
}

func (d *DynamicData[T]) Get() (*T, error) {
	return d.value, d.err
}

func (d *DynamicData[T]) Force() T {
	return *d.value
}

func (d *DynamicData[T]) Update() bool {
	if d.IsDel() {
		return false
	}

	if d.HasExp() && !d.IsExp() {
		return false
	}

	var value, err = d.modifier(d.value)
	if err != nil {
		d.err = err
		return false
	}

	d.value = value
	d.err = nil
	d.modifyAt = time.Now()
	return true
}

func (d *DynamicData[T]) Del() bool {
	d.deleteAt = time.Now()
	return true
}

func (d *DynamicData[T]) IsDel() bool {
	return !d.deleteAt.IsZero()
}

func (d *DynamicData[T]) HasExp() bool {
	return d.expireAt > 0
}

func (d *DynamicData[T]) IsExp() bool {
	if !d.HasExp() {
		return false
	}

	return d.modifyAt.Add(d.expireAt).After(time.Now())
}

func (d *DynamicData[T]) IsErr() bool {
	return d.err != nil
}
