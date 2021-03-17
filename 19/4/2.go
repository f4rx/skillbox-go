package main

import (
	"fmt"
)

func main() {
	bubbleSort([6]int{9, 10, 5, 8, 3, 7})
	bubbleSort([6]int{5, 1, 4, 2, 8, 7})
	bubbleSort([6]int{5, 4, 0, 0, -5, -7})
	bubbleSort([6]int{1, 2, 3, 4, 5, 6})
	bubbleSort([6]int{6, 5, 4, 3, 2, 1})
}

func bubbleSort(a [6]int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			// fmt.Println(a)
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				continue
			}
		}
	}
	fmt.Println(a)
}
