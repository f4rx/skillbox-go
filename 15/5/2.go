package main

import "fmt"

const arSize = 10

func main() {
	var a [arSize]int
	for i := 0; i < arSize; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&a[i])
	}

	fmt.Printf("Исходный массив %v\n", a)
	b := revertArray(a)
	fmt.Printf("Реверсированный массив %v\n", b)

}

func revertArray(sourceArray [arSize]int) (newArray [arSize]int) {
	for i, v := range sourceArray {
		newArray[arSize-i-1] = v
	}
	return
}
