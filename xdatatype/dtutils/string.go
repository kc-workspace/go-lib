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
		func(
			value reflect.Value,
			kind reflect.Kind,
		) (result string, err error, accept bool) {
			switch kind {
			case reflect.Pointer:
				result, err = StringError(value.Elem().Interface())
				accept = true
			}

			return
		}, func(
			value reflect.Value,
			kind reflect.Kind,
		) (result string, err error, accept bool) {
			switch kind {
			case reflect.Slice,
				reflect.Array:
				var stringArray = make([]string, 0)
				var len = value.Cap()

				for i := 0; i < len; i++ {
					var internal, internalError = StringError(value.Index(i).Interface())
					// Force return error if
					// any array element cannot convert to string
					if internalError != nil {
						err = internalError
						accept = true
						return
					}

					stringArray = append(stringArray, internal)
				}

				result = fmt.Sprintf("[%s]", strings.Join(stringArray, ","))
				accept = true
			}

			return
		})
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
		func(
			value reflect.Value,
			kind reflect.Kind,
		) (result string, err error, accept bool) {
			switch kind {
			case reflect.Pointer:
				result = StringForce(value.Elem().Interface())
				accept = true
			}

			return
		}, func(
			value reflect.Value,
			kind reflect.Kind,
		) (result string, err error, accept bool) {
			switch kind {
			case reflect.Slice,
				reflect.Array:
				var stringArray = make([]string, 0)
				var len = value.Cap()

				for i := 0; i < len; i++ {
					var internal = StringForce(value.Index(i).Interface())
					stringArray = append(stringArray, internal)
				}

				result = fmt.Sprintf("[%s]", strings.Join(stringArray, ","))
				accept = true
			}

			return
		})

	return
}

type underlayCallback func(
	value reflect.Value,
	kind reflect.Kind,
) (result string, err error, accept bool)

func underlayString(
	input any,
	before underlayCallback,
	after underlayCallback,
) (result string, err error) {
	var accept bool
	var value = reflect.ValueOf(input)
	var kind = value.Kind()

	result, err, accept = before(value, kind)
	if accept {
		return result, err
	}

	// Reset error to nil
	err = nil
	switch kind {
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
	default:
		result, err, accept = after(value, kind)
		if accept {
			return result, err
		} else {
			result = ""
			err = dterrors.NewConvertFailError(input, "string")
		}
	}

	return
}
