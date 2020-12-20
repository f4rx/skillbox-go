package main

import (
	"fmt"
)

func main() {

	fmt.Println("Давай копить!")
	fmt.Println("А на что? Введи цель:")

	var goal string
	fmt.Scan(&goal)

	total := 0

	var money int

	fmt.Println("Сумма 1:")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 2:")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 3:")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 4:")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 5:")
	fmt.Scan(&money)
	total += money

	fmt.Println("==============")
	fmt.Print("Ты копил на ", goal, ". Получилось собрать ", total, " руб. \n")

}
