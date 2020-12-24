/*
Задание 5*. Задача на постепенное наполнение корзин разной ёмкости (if и continue и break)
Представьте, что у вас есть три корзины разной ёмкости. Пользователю предлагается ввести, какое количество яблок помещается в каждую корзину. После этого программа должна заполнить все корзины по очереди, учитывая, какие корзины уже заполнены, строго соблюдая очерёдность заполнения и добавляя по одному яблоку в каждой итерации.
*/

package main

import (
	"fmt"
)

const basketsCount = 3

func main() {

	var basketsSize [basketsCount]int
	var baskets [basketsCount]int // Это переменные a, b, c

	for i := 0; i < basketsCount; i++ {
		fmt.Printf("Введите размер корзины %d: ", i+1)
		fmt.Scan(&basketsSize[i])
	}

	fulledBasketsCount := 0

	for i := 0; ; i++ {
		// fmt.Printf("Будет заполнена корзина %d\t", i)
		if fulledBasketsCount == basketsCount {
			break
		}
		fmt.Printf("%v\n", baskets)
		if baskets[i] < basketsSize[i] {
			baskets[i]++
			if baskets[i] == basketsSize[i] {
				fulledBasketsCount++
			}
		}

		if i == basketsCount-1 {
			i = -1
		}
	}

	fmt.Printf("Козины в конце заполнения: %v\n", baskets)

}
