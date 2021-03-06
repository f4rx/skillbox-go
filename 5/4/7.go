/*
7. Счастливый билет

Что нужно сделать

Все мы в детстве, да и не только в детстве, получив бумажный билет, проверяли, а не является ли он “счастливым”. Давайте напишем программу, в которую пользователь будет вводить четырехзначный номер билета, а программа будет выводить, является ли он счастливым (сумма старших двух разрядов равна сумме двух младших разрядов) или даже зеркальным.
*/

package main

import (
	"fmt"
)

func main() {
	var ticket int

	fmt.Println("Программа: счастливый билет")
	fmt.Print("введите номер билета (4х значный): ")
	fmt.Scan(&ticket)

	if ticket > 9999 || ticket < 1000 {
		fmt.Println("Неправильный номер билета")
		return
	}

	c1 := ticket / 1000
	c2 := ticket % 1000 / 100
	c3 := ticket % 100 / 10
	c4 := ticket % 10

	// Для дебага
	// fmt.Printf("%d %d %d %d\n", c1, c2, c3, c4)

	if c1 == c4 && c2 == c3 {
		fmt.Println("Зеркальнй номер")
	} else if c1+c2 == c3+c4 {
		fmt.Println("Счастливый номер")
	} else {
		fmt.Println("Обычный номер")
	}

}
