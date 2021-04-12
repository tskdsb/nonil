package main

import (
	"reflect"

	"github.com/tskdsb/nonil/random"
)

// Zero returns instance with every field set by type.
func Zero(t reflect.Type) interface{} {
	isPtr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isPtr = true
	}

	var value reflect.Value
	if v, has := random.SimpleValue(t); has {
		value = v
	} else {
		k := t.Kind()
		switch k {
		case reflect.Struct:
			value = fillStruct(t)
		case reflect.Slice:
			value = fillSlice(t)
		case reflect.Map:
			value = fillMap(t)
		}
	}

	if isPtr {
		value = value.Addr()
	}

	return value.Interface()
}

func fillStruct(t reflect.Type) reflect.Value {
	zeroValue := reflect.New(t).Elem()

	l := zeroValue.NumField()
	for i := 0; i < l; i++ {
		ft := zeroValue.Field(i).Type()
		v := Zero(ft)
		zeroValue.Field(i).Set(reflect.ValueOf(v))
	}

	return zeroValue
}

func fillSlice(t reflect.Type) reflect.Value {
	zeroValue := reflect.MakeSlice(t, random.DefaultSliceLen, random.DefaultSliceLen)

	v := Zero(t.Elem())
	for i := 0; i < random.DefaultSliceLen; i++ {
		zeroValue.Index(i).Set(reflect.ValueOf(v))
	}

	return zeroValue
}

func fillMap(t reflect.Type) reflect.Value {
	zeroValue := reflect.MakeMapWithSize(t, random.DefaultMapLen)

	for i := 0; i < random.DefaultSliceLen; i++ {
		k := Zero(t.Key())
		v := Zero(t.Elem())
		zeroValue.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	}

	return zeroValue
}
