package cdata

import "time"

func NewStatic[T any](key string, value T) BaseData[T] {
	return NewStaticExp(key, value, -1)
}

func NewStaticExp[T any](key string, value T, expire time.Duration) BaseData[T] {
	var now = time.Now()
	return &StaticData[T]{
		key:      key,
		value:    value,
		createAt: now,
		modifyAt: now,
		expireAt: expire,
	}
}

func NewDynamic[T any](key string, modifier Modifier[T]) BaseData[T] {
	return NewDynamicExp(key, modifier, -1)
}

func NewDynamicExp[T any](key string, modifier Modifier[T], expire time.Duration) BaseData[T] {
	var now = time.Now()
	return &DynamicData[T]{
		key:      key,
		value:    nil,
		modifier: modifier,
		createAt: now,
		modifyAt: now,
		expireAt: expire,
	}
}
