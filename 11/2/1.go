package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int64
	var base int
	fmt.Scan(&i)
	fmt.Scan(&base)
	a := strconv.FormatInt(i, base)
	fmt.Println(a)
}
