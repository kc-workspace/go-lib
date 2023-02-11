package utils

// CloneArray will create copy of input array
func CloneArray[T any](a []T, extra ...T) []T {
	var base = make([]T, 0)

	base = append(base, a...)
	base = append(base, extra...)

	return base
}
