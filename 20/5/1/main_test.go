package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeterminant(t *testing.T) {
	m1 := [size][size]int{
		{1, -2, 3},
		{4, 0, 6},
		{-7, 8, 9},
	}
	currentResult1 := determinant(m1)
	expectedResult1 := 204

	m2 := [size][size]int{
		{-1, 2, 5},
		{7, -4, 3},
		{-5, 0, 10},
	}
	currentResult2 := determinant(m2)
	expectedResult2 := -230

	assert.Equal(t, expectedResult1, currentResult1, "The two numbers should be the same.")
	assert.Equal(t, expectedResult2, currentResult2, "The two numbers should be the same.")
}
