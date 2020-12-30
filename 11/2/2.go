package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Golang is programming language"
	s2 := strings.ReplaceAll(s1, "a", "")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1)-len(s2))

	c := 0
	for _, char := range s1 {
		if char == 'a' {
			c++
		}
	}
	fmt.Println(c)

	c = 0
	for {
		i := strings.Index(s1, "a");
		if i < 0 {
			break
		}
		c++
		s1 = s1[i+1:]
	}
	fmt.Println(c)

}
