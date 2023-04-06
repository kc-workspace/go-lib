package dtutils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/kc-workspace/go-lib/xdatatype/dterrors"
)

// StringError will try to convert input to string with error
func StringError(input any) (string, error) {
	return underlayString(
		input,
		func(value interface{}) (string, error) {
			return StringError(value)
		},
	)
}

// String will try to convert input to string with flag
func String(input any) (result string, ok bool) {
	var underlayResult, underlayError = StringError(input)
	if underlayError != nil {
		return "", false
	}

	return underlayResult, true
}

// StringDefault will try to convert input to string
// otherwise, return <def> string
func StringDefault(input any, def string) string {
	var underlayResult, underlayError = StringError(input)
	if underlayError != nil {
		return def
	}

	return underlayResult
}

// StringForce will force convert input to string
func StringForce(input any) (result string) {
	result, _ = underlayString(
		input,
		func(value interface{}) (string, error) {
			return StringForce(input), nil
		},
	)

	return
}

type underlayCallback func(
	value interface{},
) (result string, err error)

func underlayString(
	input any,
	recurCB underlayCallback,
) (result string, err error) {
	var value = reflect.ValueOf(input)
	var kind = value.Kind()

	switch kind {
	case reflect.Pointer:
		result, err = recurCB(value.Elem().Interface())
	case reflect.String:
		result = value.String()
	case reflect.Bool:
		result = strconv.FormatBool(value.Bool())
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		result = strconv.FormatInt(value.Int(), 10)
	case reflect.Uint,
		reflect.Uintptr,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		result = strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		result = strconv.FormatFloat(value.Float(), 'f', -1, 32)
	case reflect.Float64:
		result = strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.Complex64:
		result = strconv.FormatComplex(value.Complex(), 'f', -1, 64)
	case reflect.Complex128:
		result = strconv.FormatComplex(value.Complex(), 'f', -1, 128)
	case reflect.Map:
		var inJson, inErr = json.Marshal(value.Interface())
		if inErr == nil {
			result = string(inJson)
		} else {
			err = inErr
		}
	case reflect.Struct:
		var inJson, inErr = json.Marshal(value.Interface())
		if inErr == nil {
			result = fmt.Sprintf("%s%s", value.Type().Name(), string(inJson))
		} else {
			err = inErr
		}
	case reflect.Slice,
		reflect.Array:
		var stringArray = make([]string, 0)
		var len = value.Cap()

		for i := 0; i < len; i++ {
			var internal string
			internal, err = recurCB(value.Index(i).Interface())
			if err != nil {
				return
			}

			stringArray = append(stringArray, internal)
		}

		result = fmt.Sprintf("[%s]", strings.Join(stringArray, ","))
	default:
		err = dterrors.NewConvertFailError(input, "string")
	}

	return
}
