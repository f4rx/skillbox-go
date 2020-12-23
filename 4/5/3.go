/*
Задача 3. Банкомат

Пользователи банкомата хотят снимать деньги. Но банкомат умеет выдавать только купюры по 100 рублей, а максимальная сумма снятия — 100 000 рублей.

Напишите программу, которая проверяет допустимость суммы средств, которую ввёл пользователь. Сумма не должна быть меньше 100 и больше 100 000 руб.
*/

package main

import (
	"fmt"
)

func main() {

	var cashRequest int
	minAmount := 100
	maxAmount := 100_000

	fmt.Println("*Имитация банкомата*")
	fmt.Print("Какая сумма вам необходима для снятия: ")
	fmt.Scan(&cashRequest)

	if cashRequest < minAmount {
		fmt.Printf("Сумма не может быть меньше %d\n", minAmount)
	} else if cashRequest > maxAmount {
		fmt.Printf("Сумма не может быть больше %d\n", maxAmount)
	} else if cashRequest % minAmount != 0 {
		fmt.Printf("Сумма должна быть кратна %d\n", minAmount)
	} else {
		fmt.Printf("Вам доступно %d для снятия\n", cashRequest)
	}
}
