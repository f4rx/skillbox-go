package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSubstring(t *testing.T) {
	tests := []struct {
		title        string
		sourceString string
		subString    string
		result       bool
	}{
		{"Поиск 1", "привет мир!", "вет", true},
		{"Поиск 2", "привет мир!", "мир", true},
		{"Поиск 3", "привет мир", "мирный", false},
		{"Поиск 4", "Программирование - это просто", "вание", true},
		{"Поиск 6", "Программирование - это просто", "корабль", false},
	}

	for _, tt := range tests {
		//nolint:scopelint
		t.Run(tt.title, func(t *testing.T) {
			res := searchSubstring(tt.sourceString, tt.subString)
			assert.Equal(t, tt.result, res, "The two arrays should be the same.")
		})
	}
}
