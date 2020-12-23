/*
Задача 5. Ресторан

В ресторане действуют следующие правила:

по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый!
по пятницам, если сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%;
если число гостей в одной компании превышает 5 человек, автоматически включается надбавка на обслуживание 10%.
Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
*/

package main

import (
	"fmt"
)

func main() {
	var dayNumber, visitors, check int

	fmt.Println("Введите день недели 1 - понедельник, .., 7 - воскресенье: ")
	fmt.Scan(&dayNumber)
	if dayNumber < 1 || dayNumber > 7 {
		fmt.Println("Невернй день недели")
		return
	}

	fmt.Println("Введите количество гостей: ")
	fmt.Scan(&visitors)
	if visitors == 0 {
		fmt.Println("Неверое количество гостей")
		return
	}

	fmt.Println("Введите сумму чека: ")
	fmt.Scan(&check)
	if check == 0 {
		fmt.Println("Неверая сумма")
		return
	}

	serviceTax := 100 // Будем дальше делить на 100, т.к. пока у нас нет флоатов
	if visitors > 5 {
		serviceTax = 110
		fmt.Printf("Добавлен сервисный сбор: %d%%\n", serviceTax - 100)
	}

	discount := 100
	if dayNumber == 1 {
		discount = 90 // -10%
		fmt.Printf("Скидка за понедельник: %d%%\n", 100 - discount)
	} else if dayNumber == 7 {
		if check > 10_000 {
			discount = 95 // -5%
			fmt.Printf("Скидка за понедельник: %d%%\n", 100 - discount)
		}
	}

	total := check * serviceTax * discount / 100 / 100
	fmt.Printf("Итого: %d руб.\n", total)

}
