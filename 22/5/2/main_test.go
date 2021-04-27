package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	tests := []struct {
		title       string
		arrayData   []int
		searchNumer int
		result      int
	}{
		{"Тестирование числа из массива", []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1},
		{"Тестирование в начале массива", []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 11},
		{"Тестирование в конце массива", []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{"Тестирование числа за пределами массива", []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
		{"Тестирование числа за пределами массива", []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1},
		{"Тестирование числа не содержащегося в массиве", []int{1, 2, 2, 2, 3, 4, 6, 6, 7, 8, 9, 10}, 5, -1},
	}

	for _, tt := range tests {
		//nolint:scopelint
		t.Run(tt.title, func(t *testing.T) {
			res := findIndex(tt.arrayData, tt.searchNumer)
			assert.Equal(t, tt.result, res, "The two numbers should be the same.")
		})
	}
}
