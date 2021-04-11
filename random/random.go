package random

import (
	"reflect"
	"time"
)

const (
	DefaultBool  = true
	DefaultFloat = 3.14

	DefaultSliceLen = 1
	DefaultMapLen   = 1
)

func GetValueByType(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Bool:
		return DefaultBool
	case reflect.String:
		return randomString()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return randomInt()
	case reflect.Float32, reflect.Float64:
		return DefaultFloat
	}

	switch t {
	case reflect.TypeOf(time.Time{}):
		return time.Now()
	}

	return nil
}
