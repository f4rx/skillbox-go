package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	arraySize    = 12
	maxRandomNum = 7
)

func init() { //nolint:gochecknoinits
	_, exists := os.LookupEnv("DEBUG")
	if exists {
		log.SetLevel(log.InfoLevel)
		log.Info("Debug is enabled")
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, arraySize)
	for i := range arr {
		arr[i] = rand.Intn(maxRandomNum) //nolint:gosec
	}

	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Printf("Введите число для поиска: ")
	var searchNumer int
	fmt.Scan(&searchNumer)

	resultIndex := findIndex(arr, searchNumer)
	if resultIndex >= 0 {
		fmt.Println(resultIndex, arr[resultIndex])
	} else {
		fmt.Println("Элемент не найден")
	}
}

func findIndex(arr []int, searchNumer int) int {
	left := 0
	right := len(arr)
	middle := len(arr) / 2

	if searchNumer < arr[0] || searchNumer > arr[len(arr)-1] {
		return -1
	}

	for {
		log.Info(left, right, middle)
		if left == middle && middle == right || left == right { //nolint:gocritic
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
	if position == 0 {
		return position
	} else if arr[position-1] == value {
		return checkLeftIndex(arr, position-1, value)
	}

	return position
}
