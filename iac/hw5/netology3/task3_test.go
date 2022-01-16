package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	got := FindAliquot(2, 10, 3)
	want := []int{3, 6, 9}

	gotLen := len(got)
	gotWant := len(want)

	if gotLen != gotWant {
		t.Errorf("got len %d, wanted len %d", gotLen, gotWant)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("arrays not equal")
	}
}
