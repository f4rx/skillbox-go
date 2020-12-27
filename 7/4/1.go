package main

import (
	"fmt"
	"math"
)

func main() {
	counter := 0

	MainLoop:
	for i := 100_000; i <= 999_999; i++ {
		// Не понимаю зачем тут 2 вложенных цикла, всё решается достаточно просто и так
		/*
		a1 := i % 1_000_000 / 100_000
		a2 := i % 100_000 / 10_000
		a3 := i % 10_000 / 1000
		a4 := i % 1_000 / 100
		a5 := i % 100 / 10
		a6 := i % 10 / 1

		if a1 == a6 && a2 == a5 && a3 == a4 {
			counter++
		}
		*/

		var a [6]int

		for j := 0; j < 6; j ++{
			a[j] = i % int(math.Pow10(6 - j)) / int(math.Pow10(5 - j))
		}

		for j := 0; j <3; j++ {
			if a[j] != a[5-j] {
				continue MainLoop
			}
		}

		counter++
	}
	fmt.Println(counter)
}
