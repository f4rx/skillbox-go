package main

import "fmt"

func main()  {
	var a[5]int
	fmt.Println(a)
	b := c1(&a)
	fmt.Println(b)
	fmt.Println(a)
}

func c1(ia *[5]int) ([5]int) {
	ia[0] = 10
	return *ia
}
