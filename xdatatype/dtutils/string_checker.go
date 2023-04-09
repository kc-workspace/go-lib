package dtutils

import "reflect"

// IsStringOrSimilar will do soft check if
// input can be string or not
func IsStringOrSimilar(input any) bool {
	// For string all type in golang can be convert to string
	return true
}

// IsString will check if input type
// must be string
func IsString(input any) bool {
	var t = reflect.TypeOf(input)
	return t.Kind() == reflect.String
}
