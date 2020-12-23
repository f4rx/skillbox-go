package main

import (
	"fmt"
)

/*
Задача 1. Калькулятор скидки

Вы покупаете 3 товара в магазине. Если сумма вашего чека превышает 10 000 руб., вам будет сделана скидка 10%.

Напишите программу, которая запрашивает 3 стоимости товара и вычисляет сумму чека.
*/

func main() {
	var itemPrice1, itemPrice2, itemPrice3 int

	fmt.Print("Введите цену товара 1 ")
	fmt.Scan(&itemPrice1)

	fmt.Print("Введите цену товара 2 ")
	fmt.Scan(&itemPrice2)

	fmt.Print("Введите цену товара 3 ")
	fmt.Scan(&itemPrice3)

	sum := itemPrice1 + itemPrice2 + itemPrice3
	total := sum
	if sum > 10_000 {
		fmt.Println("Вам предоставлена скидки 10%")
		total = int(float64(sum) * 0.9)
	}
	println("Цена покупки: ", total)
}
