package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// s := "a10 10 20b 20 30c30 30 dd"
	s := " a10 10 20b 20 30c30 30 dd "
	// s := ""

	sArr := strings.Split(s, " ")
	for _, v  := range sArr {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		fmt.Println(i)
	}
}
