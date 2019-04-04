package nonil

import (
	"testing"
)

type exampleType struct {
	i *int
	s []int
	m map[string]interface{}
}

func TestZero(t *testing.T) {
	e1 := exampleType{}
	e2, err := Zero(e1)
	if err != nil {
		t.Logf("%s\n", err)
		return
	}

	e3, ok := e2.(exampleType)
	if !ok {
		t.Fatalf("type changed after Zero: %T", e3)
	}

	if e3.i == nil {
		t.Fatalf("failed on Pointer")
	}
	if e3.s == nil {
		t.Fatalf("failed on Slice")
	}
	if e3.m == nil {
		t.Fatalf("failed on Map")
	}
}
