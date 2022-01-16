package main

import "testing"

func TestAdd(t *testing.T) {

	got := M2fConvert(3048)
	var want float64 = 10000

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
