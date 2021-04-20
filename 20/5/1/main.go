package main

import "fmt"

const size = 3

func main() {
	m := [size][size]int{
		{1, -2, 3},
		{4, 0, 6},
		{-7, 8, 9},
	}

	d := determinant(m)
	fmt.Println("Определитель: ", d)
}

func determinant(m [size][size]int) int {
	return m[0][0]*m[1][1]*m[2][2] -
		m[0][0]*m[1][2]*m[2][1] -
		m[0][1]*m[1][0]*m[2][2] +
		m[0][1]*m[1][2]*m[2][0] +
		m[0][2]*m[1][0]*m[2][1] -
		m[0][2]*m[1][1]*m[2][0]
}
