package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software" // 5
	// s := "FOO Bar CoM foo" // 3
	// s := "FOO Bar CoM" // 3
	// s := "" // 0

	s = s + " "
	c := 0
	for {
		i := strings.Index(s, " ");
		if i < 0 {
			break
		}
		fmt.Println(s[0:1], s[0:i])
		if strings.ToUpper(s[0:1]) == s[0:1] {
			c++
		}
		s = s[i+1:]
	}
	fmt.Println(c)
}
