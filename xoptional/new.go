package xoptional

func New[T any](t *T) Optional[T] {
	return Optional[T]{
		value: t,
	}
}

func Empty[T any]() Optional[T] {
	return New[T](nil)
}

func Next[T, N any](t Optional[T], fn func(t T) *N) Optional[N] {
	if t.Empty() {
		return Empty[N]()
	}

	return New(fn(t.Get()))
}
