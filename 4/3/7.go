/*
Задача 7. Меню ресторана

Напишите программу, которая выводит меню бизнес-ланча ресторана, в зависимости от дня недели.

В меню есть общая часть, а есть уникальная часть, которая зависит от дня недели. Пользователь должен ввести номер дня недели, от 1 (понедельник) до 7 (воскресенье), а программа должна вывести на экран день недели и меню этого дня.
*/

package main

import (
	"fmt"
)

func main() {
	var dayNumber int
	fmt.Print("Введите номер длня недели 1 - понедельник,.. 7 - воскресенье: ")
	fmt.Scan(&dayNumber)

	if dayNumber == 1 {
		fmt.Println("меню понедельника...")
	} else if dayNumber == 2 {
		fmt.Println("меню вторника...")
	} else if dayNumber == 3 {
		fmt.Println("меню среды...")
	} else if dayNumber == 4 {
		fmt.Println("меню четверг...")
	} else if dayNumber == 5 {
		fmt.Println("меню пятницы...")
	} else if dayNumber == 6 {
		fmt.Println("меню суббота...")
	} else if dayNumber == 7 {
		fmt.Println("меню воскресенье...")
	} else {
		fmt.Println("Неправильный день недели...")
	}
}
