package main

import "fmt"

func main() {
	myFunction(1, 2, a)

	myFunction(1, 2, func(a int, b int) int { fmt.Printf("%d - %d = %d\n", a, b, a-b); return a + b })

	f := func(a int, b int) int {
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
		return a * b
	}
	myFunction(1, 2, f)
}

func myFunction(a, b int, f func(int, int) int) {
	defer f(a, b)
	fmt.Println("my function")
}

func a(a int, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	return a + b
}
