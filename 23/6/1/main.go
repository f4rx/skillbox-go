package main

import (
	"fmt"
)

const size = 10

func main() {
	a := [size]int{-2, -1, 0, 1, 2, 4, 5, 9, 10, 11}

	a1, a2 := parseArray(a)
	fmt.Printf("Массив с чётными данными:\t%v\n", a1)
	fmt.Printf("Массив с нечётными данными:\t%v\n", a2)
}

func parseArray(a [size]int) ([]int, []int) {
	evenArray := make([]int, 0, size)
	oddArray := make([]int, 0, size)

	for _, i := range a {
		if i%2 == 0 {
			evenArray = append(evenArray, i)
		} else {
			oddArray = append(oddArray, i)
		}
	}
	return evenArray, oddArray
}
