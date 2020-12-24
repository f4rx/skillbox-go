/*
Задание 4. Предыдущий урок на if
Есть три переменные со значениями 0. Первую нужно наполнить до десяти, вторую до 100 и третью до 1 000. Напишите цикл, в котором эти переменные будут увеличиваться на один. Используйте условия для пропуска переменных, которые уже достигли своих лимитов.
*/

package main

import (
	"fmt"
)

// Т.к. тут упор на if, то я делают вывод, что все 3 переменные заполняются за один проход цикла
func main() {

	a, b, c := 0, 0, 0

	for i := 1; i <= 1_000; i++ {
		// fmt.Printf("%d ", i)
		if a < 10 {
			a++
		}
		if b < 100 {
			b++
		}
		if c < 1_000 {
			c++
		}
	}
	// fmt.Printf("\n")

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

}
