package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const size = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(7) //nolint:gosec
	}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Printf("Введите число для поиска: ")
	var searchNumer int
	fmt.Scan(&searchNumer)

	resultIndex := findIndex(arr, searchNumer)
	fmt.Println(resultIndex, arr[resultIndex])
}

func findIndex(arr []int, searchNumer int) int {
	left := 0
	right := len(arr)
	middle := len(arr) / 2

	if searchNumer < arr[0] || searchNumer > arr[len(arr)-1] {
		return -1
	}

	for {
		// fmt.Println(left, right, middle)
		if left == middle || middle == right { //nolint:gocritic
			return -1
		} else if arr[middle] == searchNumer {
			return checkLeftIndex(arr, middle, searchNumer)
		} else if arr[middle] > searchNumer {
			right = middle - 1
			middle = right / 2
		} else if arr[middle] < searchNumer {
			left = middle + 1
			middle = (middle + right) / 2
		}
	}
}

func checkLeftIndex(arr []int, position, value int) int {
	for {
		if position == 0 { //nolint:gocritic
			return position
		} else if arr[position-1] == value {
			position--
		} else {
			return position
		}
	}
}
