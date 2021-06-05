package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortBubbleReverse(t *testing.T) {
	tests := []struct {
		title     string
		arrayData []int
		result    []int
	}{
		{"Сортировка 1", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Сортировка 2", []int{3, 2, 1}, []int{3, 2, 1}},
		{"Сортировка 2", []int{4, 1, 2, 3}, []int{4, 3, 2, 1}},
		{"Сортировка 2", []int{4, 1, 7, 4}, []int{7, 4, 4, 1}},
		{"Сортировка 3", []int{}, []int{}},
	}

	for _, tt := range tests {
		//nolint:scopelint
		t.Run(tt.title, func(t *testing.T) {
			res := sortBubbleReverse(tt.arrayData...)
			assert.Equal(t, tt.result, res, "The two arrays should be the same.")
		})
	}
}
