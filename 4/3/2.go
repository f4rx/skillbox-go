/*
Задача 2. Складываем в уме

Напишите программу, которая проверяет то, как вы умеете складывать два числа в уме. Программа должна выводить два разных сообщения на верный и неверный ответ пользователя. В последнем случае, надо показывать правильный результат. Пользователь должен ввести с клавиатуры два целых числа, а потом результат их сложения.
*/

package main

import (
	"fmt"
)

func main() {
	var a, b, inputResult int
	fmt.Print("Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)
	fmt.Print("Введите произведение чисел (a*b): ")
	fmt.Scan(&inputResult)

	result := a * b
	if inputResult == result {
		fmt.Println("Всё верно!")
	} else {
		fmt.Println("Вы ошиблись, правильный результат:", result)
	}

}
