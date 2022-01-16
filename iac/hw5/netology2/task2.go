package main

import "fmt"

func main() {
	var arr = []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	ind, val := FindMin(arr)
	fmt.Printf("minValue = %d position = %d", val, ind)

}

func FindMin(array []int) (int, int) {
	var min = array[0]
	var ind int = 0
	for i, value := range array {
		if min > value {
			min = value
			ind = i
		}
	}
	return ind, min
}
