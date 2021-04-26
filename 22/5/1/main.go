package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := [size]int{}
	for i := range arr {
		arr[i] = rand.Intn(10) //nolint:gosec
	}
	fmt.Println(arr)

	fmt.Printf("Введите число для поиска: ")
	var searchNumer int
	fmt.Scan(&searchNumer)

	result := findIndex(arr, searchNumer)
	fmt.Printf("Справа от числа %d - %d чисел\n", searchNumer, result)
}

func findIndex(arr [size]int, searchNumer int) int {
	for i, v := range arr {
		if v == searchNumer {
			return len(arr) - i - 1
		}
	}
	return 0
}
