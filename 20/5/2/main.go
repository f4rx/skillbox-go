package main

import "fmt"

func main() {
	m1 := [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	m2 := [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	r := multipleMatrix(m1, m2)
	fmt.Println("Результат: ", r)
}

func multipleMatrix(matrixA [3][5]int, matrixB [5][4]int) [3][4]int {
	var r [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			r[i][j] = matrixA[i][0]*matrixB[0][j] +
				matrixA[i][1]*matrixB[1][j] +
				matrixA[i][2]*matrixB[2][j] +
				matrixA[i][3]*matrixB[3][j] +
				matrixA[i][4]*matrixB[4][j]
		}
	}
	return r
}
