package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software" // 5
	s := "Проверка с Русскими СиМвоЛами"
	// s := "FOO Bar CoM foo" // 3
	// s := "FOO Bar CoM" // 3
	// s := "" // 0

	s = s + " "
	c := 0
	for {
		i := strings.Index(s, " ")
		if i < 0 {
			break
		}
		tmpS := s[0:i]
		for j, runeValue := range tmpS {
			if j > 0 {
				break
			}
			if strings.ToUpper(string(runeValue)) == string(runeValue) {
				c++
				fmt.Println(string(runeValue), runeValue)
			}
		}
		s = s[i+1:]
	}
	fmt.Println(c)
}
