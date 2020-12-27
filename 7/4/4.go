package main

import (
	"fmt"
)

func main() {

	unhappyChainSize := 0
	lastHappyNumber := 0
	for i := 100_000; i <= 999_999; i++ {
		if isHappyTiket(i) {
			if unhappyChainSize < i-lastHappyNumber && lastHappyNumber != 0 {
				unhappyChainSize = i - lastHappyNumber
			}
			lastHappyNumber = i
		}
	}

	fmt.Println(unhappyChainSize)
}

func isHappyTiket(number int) bool {

	a1 := number % 1_000_000 / 100_000
	a2 := number % 100_000 / 10_000
	a3 := number % 10_000 / 1000
	a4 := number % 1_000 / 100
	a5 := number % 100 / 10
	a6 := number % 10 / 1

	return a1+a2+a3 == a4+a5+a6
}
