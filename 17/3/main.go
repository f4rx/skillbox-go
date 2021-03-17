package main

import "fmt"



func main()  {
	c := 1 + 5
	fmt.Println(c)
	a := 2
	b := [5]int{1,2,3}
	fmt.Println(a, c, b)
	doPanic(10)

}

func doPanic(i int) {
	if i > 1 {
		doPanic(i-1)
	}

	// panic("ololo")
}