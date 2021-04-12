package random

import (
	"reflect"
	"time"
)

const (
	DefaultBool = true

	DefaultSliceLen = 1
	DefaultMapLen   = 1
)

func SimpleValue(t reflect.Type) (reflect.Value, bool) {
	v := reflect.New(t).Elem()

	switch t.Kind() {
	case reflect.Bool:
		v.Set(reflect.ValueOf(DefaultBool))
		return v, true
	case reflect.String:
		v.Set(reflect.ValueOf(randomString()))
		return v, true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.Set(reflect.ValueOf(randomInt()))
		return v, true
	case reflect.Float32:
		v.Set(reflect.ValueOf(randomFloat32()))
		return v, true
	case reflect.Float64:
		v.Set(reflect.ValueOf(randomFloat64()))
		return v, true
	}

	switch t {
	case reflect.TypeOf(time.Time{}):
		v.Set(reflect.ValueOf(time.Now()))
		return v, true
	}

	return v, false
}
