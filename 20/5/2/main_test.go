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

func TestMultipleCustomMatrix(t *testing.T) {
	// ##############
	// Test1
	// ###############
	t1m1 := [][]int{
		{1, 2},
		{1, 2},
	}

	expectedResult1 := [][]int{
		{3, 6},
		{3, 6},
	}

	currentResult1, err := multipleCustomMatrix(t1m1, t1m1)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedResult1, currentResult1, "The two matrix should be the same.")
	}

	// ##############
	// Test2
	// ###############
	t2m1 := [][]int{
		{1, 2},
		{3, 4},
		{2, 2},
	}

	t2m2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	expectedResult2 := [][]int{
		{9, 12, 15},
		{19, 26, 33},
		{10, 14, 18},
	}

	currentResult2, err := multipleCustomMatrix(t2m1, t2m2)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedResult2, currentResult2, "The two matrix should be the same.")
	}

	// ##############
	// Test3
	// ###############
	t3m1 := [][]int{
		{1, 2},
		{3, 4},
		{2, 2},
	}

	t3m2 := [][]int{
		{1, 2, 3},
	}

	currentResult3, err := multipleCustomMatrix(t3m1, t3m2)
	if assert.Error(t, err) {
		assert.Nil(t, currentResult3)
	}
}
