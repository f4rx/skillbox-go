package main

import (
	"fmt"
)

func main(){

	total := 0

	for i:=0; i<=40; i++ {
		if i % 4 != 0 {
			continue
		}
		total += i
	}

		fmt.Println("total: ", total)
}
