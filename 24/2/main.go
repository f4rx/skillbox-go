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
	fmt.Println(sortBubbleReverse(2, 4, 3, 5, 7, 6))
	fmt.Println(sortBubbleReverse(1, 2, 3))
	fmt.Println(sortBubbleReverse(3, 2, 1, 0))
	fmt.Println(sortBubbleReverse())
}

func sortBubbleReverse(input ...int) []int {
	log.Info(input)

	for i := len(input) - 1; i > 0; i-- {
		for j := len(input) - 2; j >= len(input)-1-i; j-- {
			if input[j] < input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
			log.Info("i: ", i, " j: ", j, input)
		}
	}
	return input
}
