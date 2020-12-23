/*

4. Сумма без сдачи

Что нужно сделать

Программное обеспечение банкоматов постоянно решает задачу, как имеющимися купюрами сформировать сумму, введенную пользователем. Давайте попробуем решить похожую задачу, и определить, сможет ли пользователь заплатить за товар без сдачи или нет. Для этого он будет вводить стоимость товара и номиналы трех монет. 


*/

package main

import (
	"fmt"
)

func main() {
	var coin1, coin2, coin3 int
	var total int

	fmt.Println("Расчёт сдачи в магазите")

	fmt.Print("Введите итоговую сумму: ")
	fmt.Scan(&total)

	fmt.Print("Введите монету 1: ")
	fmt.Scan(&coin1)
	fmt.Print("Введите монету 2: ")
	fmt.Scan(&coin2)
	fmt.Print("Введите монету 3: ")
	fmt.Scan(&coin3)

	withoutChange := false

	if coin1 + coin2 + coin3 < total {
		fmt.Println("Ваших средств не хватает на покупку")
		return
	}

	if coin1 + coin2 + coin3 == total {
		withoutChange = true
	} else if coin1 + coin2 == total {
		withoutChange = true
	} else if coin1 + coin3 == total {
		withoutChange = true
	} else if coin2 + coin3 == total {
		withoutChange = true
	} else if coin1 == total || coin2 == total || coin3 == total {
		withoutChange = true
	}

	if withoutChange {
		fmt.Println("Можно раслпатиться без сдачи")
	} else {
		fmt.Println("Нелья раслпатиться без сдачи")
	}

}
