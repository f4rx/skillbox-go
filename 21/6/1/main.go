package main

import "fmt"

func main() {
	result := calcalute(2, 240, 1.0)
	fmt.Println(result)
}

func calcalute(x int16, y uint8, z float32) float32 {
	// S = 2 × x + y ^ 2 − 3/z
	return 2.0*float32(x) + float32(y)*float32(y) - 3.0/z
}
