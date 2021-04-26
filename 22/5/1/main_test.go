package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	t.Parallel()
	arrayData := [...]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	tests := []struct {
		title       string
		arrayData   [10]int
		searchNumer int
		result      int
	}{
		{"Тестирование числа из массива", arrayData, 8, 4},
		{"Тестирование числа за пределами массива", arrayData, 10, 0},
	}

	for _, tt := range tests {
		//nolint:scopelint
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			res := findIndex(tt.arrayData, tt.searchNumer)
			assert.Equal(t, tt.result, res, "The two numbers should be the same.")
		})
	}
}
