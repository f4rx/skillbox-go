package main

import (
	"fmt"
	"time"
)

const RIGHT = true
const LEFT = false

func main(){

	l := -10
	r := 10
	direction := RIGHT
	startPostion := 0
	currentPostion := 0

	for  {
		if currentPostion == startPostion {
			fmt.Println("КПП")
		} else {
			fmt.Println("позиция: ", currentPostion)
		}

		if direction == RIGHT {
			currentPostion ++
		} else {
			currentPostion --
		}
		if currentPostion == r {
			direction = LEFT
		} else if currentPostion == l {
			direction = RIGHT
		}

		time.Sleep(300 * time.Millisecond)
	}

}
