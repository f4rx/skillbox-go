package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		x, y := getRandomPoint()
		x1, y1 := newCoordinate(x, y)
		fmt.Printf("(x,y) -> (x1, y1): (%d,%d) -> (%d,%d) \n", x, y, x1, y1)
	}

}

func newCoordinate(x, y int) (int, int) {
	return 2*x + 10, -3*y + 5
}

func getRandomPoint() (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(100)
	y = rand.Intn(100)
	return
}
