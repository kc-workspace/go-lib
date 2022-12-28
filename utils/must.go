package utils

// MustR will panic if error occurred; otherwise, return input value
func MustR[K interface{}](input K, err error) K {
	if err != nil {
		panic(err)
	}
	return input
}

// Must will panic if error occurred.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
