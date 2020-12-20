package main

import (
	"fmt"
)

func main() {

	totalAmount := 4000000
	entranceCount := 10
	appartmentPerEntranceCount := 40

	fmt.Printf("Сумма, указанная в квитанции: %d\n", totalAmount)
	fmt.Printf("Подъездов в доме:  %d\n", entranceCount)
	fmt.Printf("Квартир в каждом подъезде: %d\n", appartmentPerEntranceCount)
	fmt.Println("----Результат-----")

	appartmentServiceAmount := totalAmount / (entranceCount * appartmentPerEntranceCount)

	fmt.Printf("Каждая квартира должна заплатить по %d руб.\n", appartmentServiceAmount)

}
