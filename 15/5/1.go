package main

import "fmt"

const arSize = 10

func main() {
	var a [arSize]int
	for i := 0; i < arSize; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&a[i])
	}

	evenCount := 0 // Чётные 2, 4, 6...
	oddCount := 0  // Нечётные 1, 3, 5...

	for _, v := range a {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Printf("Количество чётных чисел в массиве: %d\n", evenCount)
	fmt.Printf("Количество нечётных чисел в массиве: %d\n", oddCount)
}
