package main

import "testing"

func TestAdd(t *testing.T) {
	var arr = []int{0, 1, 2, 3, 4, 5}
	gotInd, gotVal := FindMin(arr)
	var wantInd int = 0
	var wantVal int = 0

	if gotInd != wantInd || gotVal != wantVal {
		t.Errorf("error test, function return value = %d position = %d", gotVal, gotInd)
	}
}
