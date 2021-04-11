package main

import (
	"reflect"

	"github.com/tskdsb/nonil/random"
)

// Zero returns instance with every field set by type.
func Zero(t reflect.Type) (interface{}, error) {
	if v := random.GetValueByType(t); v != nil {
		return v, nil
	}

	k := t.Kind()
	switch k {
	case reflect.Struct:
		return fillStruct(t)
	case reflect.Slice:
		return fillSlice(t)
	case reflect.Map:
		return fillMap(t)
	}

	return nil, nil
}

func fillStruct(t reflect.Type) (interface{}, error) {
	zeroValue := reflect.New(t).Elem()

	l := zeroValue.NumField()
	for i := 0; i < l; i++ {
		ft := zeroValue.Field(i).Type()
		v, err := Zero(ft)
		if err != nil {
			return nil, err
		}
		zeroValue.Field(i).Set(reflect.ValueOf(v))
	}

	return zeroValue.Interface(), nil
}

func fillSlice(t reflect.Type) (interface{}, error) {
	zeroValue := reflect.MakeSlice(t, random.DefaultSliceLen, random.DefaultSliceLen)

	v, err := Zero(t.Elem())
	if err != nil {
		return nil, err
	}

	for i := 0; i < random.DefaultSliceLen; i++ {
		zeroValue.Index(i).Set(reflect.ValueOf(v))
	}

	return zeroValue.Interface(), nil
}

func fillMap(t reflect.Type) (interface{}, error) {
	zeroValue := reflect.MakeMapWithSize(t, random.DefaultMapLen)

	for i := 0; i < random.DefaultSliceLen; i++ {
		k, err := Zero(t.Key())
		if err != nil {
			return nil, err
		}
		v, err := Zero(t.Elem())
		if err != nil {
			return nil, err
		}
		zeroValue.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	}

	return zeroValue.Interface(), nil
}
