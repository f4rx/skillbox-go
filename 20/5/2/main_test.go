package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleMatrix1(t *testing.T) {
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

	currentResult := multipleMatrix(m1, m2)

	expectedResult := [3][4]int{
		{175, 190, 205, 220},
		{400, 440, 480, 520},
		{625, 690, 755, 820},
	}

	assert.Equal(t, expectedResult, currentResult, "The two numbers should be the same.")
}

func TestMultipleMatrix2(t *testing.T) {
	m1 := [3][5]int{
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3},
	}

	m2 := [5][4]int{
		{4, 4, 4, 4},
		{5, 5, 5, 5},
		{6, 6, 6, 6},
		{7, 7, 7, 7},
		{8, 8, 8, 8},
	}

	currentResult := multipleMatrix(m1, m2)

	expectedResult := [3][4]int{
		{30, 30, 30, 30},
		{60, 60, 60, 60},
		{90, 90, 90, 90},
	}

	assert.Equal(t, expectedResult, currentResult, "The two numbers should be the same.")
}
