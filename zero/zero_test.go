package zero

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
		Float  float64
		Struct exampleStruct

		Array   []int
		Map     map[string]string
		Pointer *int
	}

	exampleStruct struct {
		Name string
	}
)

func TestZero(t *testing.T) {
	e1 := exampleType{}
	e2 := Zero(reflect.TypeOf(e1))
	e3, ok := e2.(exampleType)
	if !ok {
		t.Fatalf("type changed after Zero: %T", e3)
	}
	bytes, err := json.MarshalIndent(e3, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
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
