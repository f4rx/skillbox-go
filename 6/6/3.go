/*
Задание 2. Сумма двух чисел по единице
Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле следующим образом: берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.
*/

package main

import (
	"fmt"
)

func main() {

	var purchasePrice, discount float64
	fmt.Print("Введите цену товара: ")
	fmt.Scan(&purchasePrice)
	// Тут не понятно скидка вводится в рублях или в процентах.
	fmt.Print("Введите размер скидки (руб.): ")
	fmt.Scan(&discount)

	if discount > 2_000 {
		discount = 2_000
	}

	if discount/purchasePrice > 0.3 {
		discount = purchasePrice * 0.3
	}

	purchasePriceTotal := purchasePrice - discount

	fmt.Println("Цена товара: ", purchasePrice)
	fmt.Println("Скидка: ", discount)
	fmt.Println("Итоговая цена: ", purchasePriceTotal)

}
