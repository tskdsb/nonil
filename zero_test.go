package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type (
	exampleType struct {
		String string
		Int    int
		Struct exampleStruct

		Array []int
		Map   map[string]string
		// Pointer *int
	}

	exampleStruct struct {
		Name string
	}
)

func TestZero(t *testing.T) {
	e1 := exampleType{}
	e2, err := Zero(reflect.TypeOf(e1))
	if err != nil {
		t.Logf("%s\n", err)
		return
	}

	e3, ok := e2.(exampleType)
	if !ok {
		t.Fatalf("type changed after Zero: %T", e3)
	}
	bytes, err := json.MarshalIndent(e3, "", "  ")
	fmt.Printf("%s\n", bytes)

	// if e3.pointer == nil {
	// 	t.Fatalf("failed on Pointer")
	// }
	// if e3.array == nil {
	// 	t.Fatalf("failed on Slice")
	// }
	// if e3.m == nil {
	// 	t.Fatalf("failed on Map")
	// }
}
