package main

import (
	"fmt"
)

func main() {

	fmt.Println("Расчёт роста бамбука")

	startHeight := 100
	growth := 50
	ate := 20
	growthPerDay := growth - ate

	height := growthPerDay*2 + growth/2 + startHeight
	fmt.Printf("Через 2.5 дня бамбук будет высотой: %d\n", height)

	/*
	День 1 - 50 см (на конец дня) (и плюс стартовый метр)
	День 2 - 80 50 - 20 + 50 (на конец дня)
	День 3 - 110 50 - 20 + 50 - 20 + 50
	День 4 - 140
	День 5 - 170
	День 6 - 200
	День 7 - 230

	Т.е. за первый день у нас прирост 50 см, и на второй и последующие дни - 30 см при расчётах на конец дня.
	*/
	daysFor3M := (300-startHeight-50)/growthPerDay + 1
	fmt.Printf("Полных дней, чтобы достичь высоты 3м: %d\n", daysFor3M)
}
