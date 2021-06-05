package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

const size byte = 10

func init() {
	_, exists := os.LookupEnv("DEBUG")
	if exists {
		log.SetLevel(log.InfoLevel)
		log.Info("Debug is enabled")
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	ar := [size]int{5, 4, 3, 8, 9, 0, 6, 2, 7, 1}
	fmt.Println(insertionSort(ar))
}

func insertionSort(input [size]int) [size]int {
	for i := 1; i < len(input); i++ {
		element := input[i]
		for j := i; j > 0; j-- {
			log.Info("i: ", i, " j: ", j, " element: ", element, input)
			if input[j] < input[j-1] {
				input[j-1], input[j] = input[j], input[j-1]
			} else {
				break
			}
		}
	}
	return input
}
