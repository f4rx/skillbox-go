package main

import (
	"fmt"
)

func main() {

	fmt.Println("Расчёт роста бамбука (со звёздочкой)")

	var startHeight, growth, ate, days int

	fmt.Print("Введите высоту санженца (см): ")
	fmt.Scan(&startHeight)

	fmt.Print("Введите величину роста за день (см): ")
	fmt.Scan(&growth)

	fmt.Print("Введите величину агрессивности гусениц (см): ")
	fmt.Scan(&ate)

	fmt.Print("Введите количество дней для расчёта роста (будет рассчитан рост на конец дня): ")
	fmt.Scan(&days)

	growthPerDay := growth - ate
	height := growth + growthPerDay*(days-1) + startHeight
	fmt.Printf("К вечеру %d-го дня бамбук будет высотой: %d\n", days, height)

	var requiredHeight int
	fmt.Print("Для расчёта дней для достижения необходимой высоты (к вечеру), укажите требуемую высоту бабмука (см): ")
	fmt.Scan(&requiredHeight)
	daysForRequiredHeight := (requiredHeight-startHeight-growth)/growthPerDay + 1
	fmt.Printf("Полных дней, чтобы достичь высоты %dсм: %d\n", requiredHeight, daysForRequiredHeight)

}
