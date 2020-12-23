/*
Задача 6. Студенты* (сложная задача)

Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов, на K групп. Напишите программу, которая поможет старосте сделать это: он вводит N, K и порядковый номер студента, а программа определяет, в какую группу он попадает.

Подсказка: в одну группу могут попадать студенты из разных частей списка.
*/

package main

import (
	"fmt"
)

func main() {
	var studentNumber, groupCount int
	var group int

	fmt.Print("Введите номер студента: ")
	fmt.Scan(&studentNumber)

	fmt.Print("Введите количество групп: ")
	fmt.Scan(&groupCount)

	/*
	Ожидаемый результат на 3 группы
	Номер студента  1 2 3 4 5 6 7 8 9 10
	Группа студента 1 2 3 1 2 3 1 2 3 1
	*/

	if studentNumber <= groupCount {
		group = studentNumber
	} else {
		groupTmp := studentNumber % groupCount
		if groupTmp == 0 {
			group = groupCount
		} else {
			group = groupTmp
		}
	}

	fmt.Printf("Группа студента: %d\n", group)

}
