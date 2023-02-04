package mapper

import (
	"github.com/kc-workspace/go-lib/datatype"
)

// Mapper is mapping of interface where key is string.
type Mapper map[string]interface{}

// Get data as interface and ok to check if key exist
func (m Mapper) Z(key string) (interface{}, bool) {
	d, ok := m[key]
	return d, ok
}

// Get data as interface or error if key is not exist
func (m Mapper) Ze(key string) (interface{}, error) {
	if d, ok := m.Z(key); ok {
		return d, nil
	}
	return nil, missError(m, key)
}

// Get data as interface or default value if key not exist
func (m Mapper) Zo(key string, i interface{}) interface{} {
	if d, ok := m.Z(key); ok {
		return d
	}
	return i
}

// Get data as interface or nil if key not exist
func (m Mapper) Zi(key string) interface{} {
	return m.Zo(key, nil)
}

// Get data as interface or panic if key not exist
func (m Mapper) Zr(key string) interface{} {
	d, err := m.Ze(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as Mapper and ok to check if key exist and parsable
func (m Mapper) M(key string) (Mapper, bool) {
	return ToMapper(m[key])
}

// Get data as Mapper or error if key not exist or cannot parse
func (m Mapper) Me(key string) (Mapper, error) {
	if d, ok := m.M(key); ok {
		return d, nil
	}

	return New(), convertError(key, m[key], "map[string]interface{}")
}

// Get data as Mapper or default value if key not exist or cannot parse
func (m Mapper) Mo(key string, i Mapper) Mapper {
	if d, ok := m.M(key); ok {
		return d
	}
	return i
}

// Get data as Mapper or empty map if key not exist or cannot parse
func (m Mapper) Mi(key string) Mapper {
	return m.Mo(key, New())
}

// Get data as Mapper or panic if key not exist or cannot parse
func (m Mapper) Mr(key string) Mapper {
	d, err := m.Me(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as Array and ok to check if key exist or parsable
func (m Mapper) A(key string) ([]interface{}, bool) {
	return datatype.ToArray(m[key])
}

// Get data as Array or error if key not exist or cannot parse
func (m Mapper) Ae(key string) ([]interface{}, error) {
	if d, ok := m.A(key); ok {
		return d, nil
	}

	return make([]interface{}, 0), convertError(key, m[key], "array")
}

// Get data as Array or default value if key not exist or cannot parse
func (m Mapper) Ao(key string, i []interface{}) []interface{} {
	if d, ok := m.A(key); ok {
		return d
	}
	return i
}

// Get data as Array or empty array if key not exist or cannot parse
func (m Mapper) Ai(key string) []interface{} {
	return m.Ao(key, make([]interface{}, 0))
}

// Get data as Array or panic if key not exist or cannot parse
func (m Mapper) Ar(key string) []interface{} {
	d, err := m.Ae(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as String and ok to check if key exist or parsable
func (m Mapper) S(key string) (string, bool) {
	return datatype.ToString(m[key])
}

// Get data as String or error if key not exist or cannot parse
func (m Mapper) Se(key string) (string, error) {
	if d, ok := m.S(key); ok {
		return d, nil
	}

	return "", convertError(key, m[key], "string")
}

// Get data as String or default value if key not exist or cannot parse
func (m Mapper) So(key string, i string) string {
	if d, ok := m.S(key); ok {
		return d
	}
	return i
}

// Get data as String or empty string if key not exist or cannot parse
func (m Mapper) Si(key string) string {
	return m.So(key, "")
}

// Get data as String or panic if key not exist or cannot parse
func (m Mapper) Sr(key string) string {
	d, err := m.Se(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as Float64 and ok to check if key exist or parsable
func (m Mapper) F(key string) (float64, bool) {
	return datatype.ToFloat(m[key])
}

// Get data as Float64 or error if key not exist or cannot parse
func (m Mapper) Fe(key string) (float64, error) {
	if d, ok := m.F(key); ok {
		return d, nil
	}

	return 0, convertError(key, m[key], "float64")
}

// Get data as Float64 or default value if key not exist or cannot parse
func (m Mapper) Fo(key string, i float64) float64 {
	if d, ok := m.F(key); ok {
		return d
	}
	return i
}

// Get data as Float64 or 0 if key not exist or cannot parse
func (m Mapper) Fi(key string) float64 {
	return m.Fo(key, 0)
}

// Get data as Float64 or panic if key not exist or cannot parse
func (m Mapper) Fr(key string) float64 {
	d, err := m.Fe(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as Int64 and ok to check if key exist or parsable
func (m Mapper) I(key string) (int64, bool) {
	return datatype.ToInt(m[key])
}

// Get data as Int64 or error if key not exist or cannot parse
func (m Mapper) Ie(key string) (int64, error) {
	if d, ok := m.I(key); ok {
		return d, nil
	}

	return 0, convertError(key, m[key], "int64")
}

// Get data as Int64 or default value if key not exist or cannot parse
func (m Mapper) Io(key string, i int64) int64 {
	if d, ok := m.I(key); ok {
		return d
	}
	return i
}

// Get data as Int64 or 0 if key not exist or cannot parse
func (m Mapper) Ii(key string) int64 {
	return m.Io(key, 0)
}

// Get data as Int64 or panic if key not exist or cannot parse
func (m Mapper) Ir(key string) int64 {
	d, err := m.Ie(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as Number (either int or float) and ok to check if key exist or parsable
func (m Mapper) N(key string) (float64, bool) {
	if f, ok := datatype.ToFloat(m[key]); ok {
		return f, true
	}
	if i, ok := datatype.ToInt(m[key]); ok {
		return float64(i), true
	}
	return -1, false
}

// Get data as Number (either int or float) or error if key not exist or cannot parse
func (m Mapper) Ne(key string) (float64, error) {
	if d, ok := m.N(key); ok {
		return d, nil
	}

	return 0, convertError(key, m[key], "number")
}

// Get data as Number (either int or float) or default value if key not exist or cannot parse
func (m Mapper) No(key string, i float64) float64 {
	if d, ok := m.N(key); ok {
		return d
	}
	return i
}

// Get data as Number (either int or float) or -1 if key not exist or cannot parse
func (m Mapper) Ni(key string) float64 {
	return m.No(key, -1)
}

// Get data as Number (either int or float) or panic if key not exist or cannot parse
func (m Mapper) Nr(key string) float64 {
	d, err := m.Ne(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Get data as Boolean and ok to check if key exist or parsable
func (m Mapper) B(key string) (bool, bool) {
	return datatype.ToBool(m[key])
}

// Get data as Boolean or error if key not exist or cannot parse
func (m Mapper) Be(key string) (bool, error) {
	if d, ok := m.B(key); ok {
		return d, nil
	}

	return false, convertError(key, m[key], "bool")
}

// Get data as Boolean or default value if key not exist or cannot parse
func (m Mapper) Bo(key string, i bool) bool {
	if d, ok := m.B(key); ok {
		return d
	}
	return i
}

// Get data as Boolean or 'false' if key not exist or cannot parse
func (m Mapper) Bi(key string) bool {
	return m.Bo(key, false)
}

// Get data as Boolean or panic if key not exist or cannot parse
func (m Mapper) Br(key string) bool {
	d, err := m.Be(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Convert current mapper to struct
func (m Mapper) Struct(target interface{}) error {
	return ToStruct(m, target)
}

// Get map size, only root map had been calculated
func (m Mapper) Size() int {
	return len(m)
}

// return true if current mapper is empty
func (m Mapper) IsEmpty() bool {
	return m.Size() == 0
}

// Check key (dot-notation key) in mapper, this use same time as Get()
func (m Mapper) Has(key string) bool {
	var d, err = Get(m, key)
	return err == nil && d != nil
}

// Set key (dot-notation key) to current mapper.
// If you would like to set new use Copy() first
func (m Mapper) Set(key string, value interface{}) Mapper {
	Set(m, key, value)
	return m
}

// Get value from key (dot-notation key)
func (m Mapper) Get(key string) (interface{}, error) {
	return Get(m, key)
}

// Get first valid value from keys (dot-notation key)
func (m Mapper) Gets(keys ...string) (interface{}, error) {
	return Gets(m, keys...)
}

// ForEach key-value in current map.
// This will perform deep-loop
func (m Mapper) ForEach(fn ForEachFn) {
	ForEach(m, fn)
}

// Copy current mapper to new mapper
func (m Mapper) Copy() Mapper {
	return Copy(m)
}

// Keys will return key as dot notation on every map level
func (m Mapper) Keys() []string {
	var keys = make([]string, 0)
	m.ForEach(func(key string, value interface{}) {
		keys = append(keys, key)
	})
	return keys
}

// get current mapper as json string
func (m Mapper) String() string {
	var json, err = ToJson(m)
	if err != nil {
		return err.Error()
	}
	return string(json)
}
