package main

import (
	"fmt"
)

func main() {
	sum(1, 2)
	diff(3, 5)
	fmt.Println()
	changeOrder(sum, diff, 1, 2, 3, 5)
}

func sum(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
}

func diff(a, b int) {
	fmt.Printf("%d - %d = %d\n", a, b, a - b)
}

func changeOrder(f1 func(int, int), f2 func(int, int), f1Arg1, f1Arg2, f2Arg1, f2Arg2 int) {
	f2(f2Arg1, f2Arg2)
	f1(f1Arg1, f1Arg2)
}
