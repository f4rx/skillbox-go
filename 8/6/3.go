package main

import (
	"fmt"
)

func main() {
	var canChange bool
	income := []int{5, 5, 5, 10, 20}
	fmt.Printf("Входящая последовательность %v\n", income)
	canChange = lemonadeChange(income)
	fmt.Println(canChange)
	fmt.Println()

	income = []int{10, 10}
	fmt.Printf("Входящая последовательность %v\n", income)
	canChange = lemonadeChange(income)
	fmt.Println(canChange)
	fmt.Println()

	income = []int{5, 5, 10, 10, 20}
	fmt.Printf("Входящая последовательность %v\n", income)
	canChange = lemonadeChange(income)
	fmt.Println(canChange)
}

func lemonadeChange(bills []int) bool {
	var bank map[int]int
	bank = make(map[int]int)
	bank[5] = 0
	bank[10] = 0
	bank[20] = 0
	canChange := true
	lastBill := 0

	for _, v := range bills {
		// fmt.Printf("%v %v\n", i, v)
		bank[v]++
		lastBill = v

		if v == 10 {
			if bank[5] > 0 {
				bank[5]--
			} else {
				canChange = false
				break
			}
		} else if v == 20 {
			if bank[10] >= 1 && bank[5] >= 1 {
				bank[10]--
				bank[5]--
			} else if bank[5] >= 3 {
				bank[5]--
			} else {
				canChange = false
				break
			}
		}
	}

	if canChange {
		fmt.Printf("Купюры в конце %v\n", bank)
	} else {
		fmt.Printf("Предалагаемая курюра: %d\n", lastBill)
		fmt.Printf("Купюры в кассе %v\n", bank)
	}

	return canChange
}
