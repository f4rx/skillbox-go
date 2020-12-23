/*
Задача 8. Зарплата* (сложная задача)

В отделе маркетинга работают 3 сотрудника. У каждого — разные зарплаты. Напишите программу, которая вычисляет разницу между самой высокой и низкой зарплатой сотрудника, а также среднюю зарплату отдела.*/

package main

import (
	"fmt"
)

func main() {
	var salary1, salary2, salary3 int
	fmt.Print("Введите зарплату первого сотрудника: ")
	fmt.Scan(&salary1)
	fmt.Print("Введите зарплату второго сотрудника: ")
	fmt.Scan(&salary2)
	fmt.Print("Введите зарплату третьего сотрудника: ")
	fmt.Scan(&salary3)

	avgSalary := (salary1 + salary2 + salary3) / 3

	highSalary := salary1
	lowSalary := salary1

	if salary2 > highSalary {
		highSalary = salary2
	}
	if salary3 > highSalary {
		highSalary = salary3
	}

	if salary2 < lowSalary {
		lowSalary = salary2
	}
	if salary3 < lowSalary {
		lowSalary = salary3
	}

	fmt.Println("Разница между самой большой и маленькой зарплатой: ", highSalary - lowSalary)
	fmt.Println("Средняя зарплата: ", avgSalary)
}
