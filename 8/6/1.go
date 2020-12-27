package main

import (
	"fmt"
)

func main() {
	var month string
	fmt.Print("Введите месяц года: ")
	

	switch fmt.Scan(&month); month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Зима")
	case "март", "апрель", "май" :
		fmt.Println("Весна")
	case "июнь", "июль", "август":
		fmt.Println("Лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Зима")
	default:
		fmt.Println("Ошибка ввода")
	}
}
