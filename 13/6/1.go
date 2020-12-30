package main

import (
	"fmt"
)

func main() {
	f1()
	f2()
	fmt.Println()
	changeOrder(f1, f2)
}

func f1() {
	fmt.Printf("Func f1\n")
}

func f2() {
	fmt.Printf("Func f2\n")
}

func changeOrder(f1 func(), f2 func()) {
	f2()
	f1()
}
