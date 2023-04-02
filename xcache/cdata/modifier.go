package cdata

// Modifier is function to modify data to latest data
type Modifier[T any] func(prev *T) (*T, error)
