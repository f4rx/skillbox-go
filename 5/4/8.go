/*
8. Игра “Угадывание числа” (дополнительно)

Что нужно сделать
Ну, и какой же компьютер без игр? Давайте научим его играть в “угадывание чисел”. Пользователь загадывает число от 1 до 10 (включительно). Программа пытается это число угадать, для этого она выводит число, а пользователь должен ответить угадала программа, загаданное число больше или меньше.

Советы и рекомендации
Программа не должна делать больше 4 попыток в процессе угадывания.
*/

package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Println("Загадайте число от 1 до 10 (включительно). Дальше программа будет пробовать угадать.")
	fmt.Println("Если угадала - напишите !, если меньше <, если больше >")

	startHightNumber := 10
	startLowNumber := 1
	hightNumber := startHightNumber
	lowNumber := startLowNumber
	guessNumber := hightNumber / 2

	for {
		fmt.Printf("Это число %d ? ", guessNumber)
		fmt.Scan(&input)
		if guessNumber < startLowNumber || guessNumber > startHightNumber {
			fmt.Println("Вы мухлюете =)")
			return
		}

		// fmt.Printf("DEBUG, ! guessNumber=%d, lowNumber=%d, hightNumber=%d\n", guessNumber, lowNumber, hightNumber)

		if input == "<" {
			hightNumber = guessNumber
			if (guessNumber+lowNumber+1)/2 >= hightNumber {
				guessNumber--
			} else {
				guessNumber = (guessNumber + lowNumber + 1) / 2
			}
		} else if input == ">" {
			lowNumber = guessNumber
			if (hightNumber+guessNumber+1)/2 <= lowNumber {
				guessNumber++
			} else {
				guessNumber = (hightNumber + guessNumber + 1) / 2
			}
		} else if input == "!" {
			fmt.Println("Ура, угадали")
			return
		} else {
			fmt.Println("Такая команда не известна")
			return
		}

		// fmt.Printf("DEBUG, guessNumber=%d, lowNumber=%d, hightNumber=%d\n", guessNumber, lowNumber, hightNumber)

	}

}
