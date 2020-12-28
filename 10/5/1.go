package main

import (
	"fmt"
	"math"
)

func main() {
	// формула
	// (x^n)/n!
	// e^x = 1 + x + x^2/2! + x^3/3

	var x, n float64
	fmt.Printf("Введите x: ")
	fmt.Scan(&x)
	fmt.Printf("Введите n (точность после запятой): ")
	fmt.Scan(&n)
	epsilon := 1.0 / math.Pow(10.0, n)


	var result, prevResult float64
	factorial := 1.0
	for i:=0.0; ; i++ {
		if i != 0 {
			factorial *= i
		}
		result += math.Pow(x, i) / factorial
		if result - prevResult < epsilon {
			break
		}
		prevResult = result
	}

	fmt.Printf("e^%.0f = %.16f\n", x, result)

}
