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
	fmt.Printf("%#v\n", sortBubbleReverse())

	_, exists := os.LookupEnv("RESERCH")
	if exists {
		reserchSlice()
	}

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

func reserchSlice() {
	var a []int = nil
	var a0 []int = make([]int, 0)
	var a1 []int


	fmt.Println(a == nil)  // true
	fmt.Println(a0 == nil) // false
	fmt.Println(len(a)) // 0
	fmt.Println(len(a0)) // 0


	fmt.Printf("a:  %#v\n", a)  // []int(nil)
	fmt.Printf("a0: %#v\n", a0) // []int{}
	fmt.Printf("a1: %#v\n", a1) // []int{}
	fmt.Printf("%#v\n", [...]int{}) // [0]int{}
	fmt.Printf("%v\n", [...]int{}) // []
}
