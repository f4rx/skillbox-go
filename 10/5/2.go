package main

import (
	"fmt"
	"math"
)

func main() {
	var amount, percentPerMonth float64
	var years int

	fmt.Printf("Введите начальный депозите: ")
	fmt.Scan(&amount)
	fmt.Printf("Введите процентную ставку в месяц: ")
	fmt.Scan(&percentPerMonth)
	fmt.Printf("Введите количество лет: ")
	fmt.Scan(&years)

	bankPart := 0.0

	for i:=0; i < years * 12; i++ {
		amount += amount * (percentPerMonth / 100)
		customersPart := math.Trunc(amount * 100) / 100
		bankPart += amount - customersPart
		amount = customersPart
	}

	fmt.Printf("%.2f %.16f\n", amount, bankPart)

}
