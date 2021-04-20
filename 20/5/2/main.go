package main

import (
	"fmt"
)

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

	s1 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	s2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	r2, err := multipleCustomMatrix(s1, s2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат: ", r2)
	}
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

func multipleCustomMatrix(matrixA [][]int, matrixB [][]int) ([][]int, error) {
	if len(matrixA) == 0 || len(matrixB) == 0 || len(matrixA[0]) != len(matrixB) {
		return nil, fmt.Errorf("неправильные размерности матрицы") //nolint:goerr113
	}

	rows := len(matrixA)
	columns := len(matrixB[0])

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for k := 0; k < len(matrixB); k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	return result, nil
}
