package main

import "fmt"

func main() {
	var startInterval int = 1
	var EndInterval int = 100
	var divisor int = 3

	result := FindAliquot(startInterval, EndInterval, divisor)
	fmt.Println(result)
}

func FindAliquot(start int, end int, divisor int) []int {
	var aliquot []int
	for i := start; i <= end; i++ {
		if i%divisor == 0 {
			aliquot = append(aliquot, i)
		}
	}
	return aliquot
}
