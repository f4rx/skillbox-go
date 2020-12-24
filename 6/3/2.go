package main

import (
	"fmt"
)

const answer = "яйцо"
const stopWord = "надоело"


func main(){

	var input string
	for  {
		fmt.Println("Что было раньше?")
		fmt.Scan(&input)
		if input == stopWord {
			break
		}
		fmt.Println(answer)
	}

	fmt.Println("Пока неудачник")
}
