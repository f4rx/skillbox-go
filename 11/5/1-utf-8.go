package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
		runeValue, _ := utf8.DecodeRuneInString(s[:i])
		// fmt.Println(s[0:1], s[0:i])
		fmt.Println(string(runeValue), runeValue)
		if strings.ToUpper(string(runeValue)) == string(runeValue) {
			c++
		}
		s = s[i+1:]
	}
	fmt.Println(c)
}
