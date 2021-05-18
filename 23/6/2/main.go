package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{
		"В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!",
		"Друг мой эльф! Яшке б свёз птиц южных чащ!",
		"Любя, съешь щипцы, — вздохнёт мэр, — кайф жгуч.",
		"Шеф взъярён тчк щипцы с эхом гудбай Жюль.",
		"Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф.",
		"Экс-граф? Плюш изъят. Бьём чуждый цен хвощ!",
		"Эх, чужак! Общий съём цен шляп (юфть) — вдрызг!",
		"Эх, чужд кайф, сплющь объём вши, грызя цент.",
		"Чушь: гид вёз кэб цапф, юный жмот съел хрящ.",
		"Съешь [же] ещё этих мягких французских булок да выпей чаю",
	}

	chars := []rune{'а', 'р', 'ш'}
	a, _ := parseTest(sentences, chars)
	print2DArray(a)
}

func parseTest(sentences []string, chars []rune) ([][]int, error) {
	// Вроде это не нужно тут, и так верну пустой слайс
	if len(sentences) == 0 || len(chars) == 0 {
		return nil, fmt.Errorf("входные данные - пустые массив") //nolint:goerr113
	}

	result := make([][]int, len(sentences))

	for i, s := range sentences {
		words := strings.Split(s, " ")
		lastWord := words[len(words)-1]

		r2 := make([]int, len(chars))
		for j, c := range chars {
			r2[j] = getPosition(lastWord, c)
		}
		result[i] = r2
	}
	return result, nil
}

func getPosition(s string, r rune) int {
	// https://blog.golang.org/strings#TOC_6.
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

func print2DArray(a [][]int) {
	for _, l := range a {
		fmt.Println(l)
	}
}
