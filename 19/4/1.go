package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Это просто чтобы с предыдущего урока попрактироваться
	log.SetLevel(log.WarnLevel)
	// log.SetLevel(log.InfoLevel)
}

func main() {

	_ = joinArrays([4]int{1, 3, 5, 7}, [5]int{1, 2, 4, 8, 10})
	_ = joinArrays([4]int{1, 3, 5, 7}, [5]int{2, 4, 8, 10, 12})
	_ = joinArrays([4]int{1, 1, 1, 1}, [5]int{2, 2, 2, 2, 2})
	_ = joinArrays([4]int{8, 8, 8, 8}, [5]int{1, 2, 3, 4, 6})
}

func joinArrays(a [4]int, b [5]int) (c [9]int) {
	ai, bi := 0, 0
	for i := 0; i < len(c); i++ {
		log.Info(ai, bi, i)
		if bi >= len(b) || ai < len(a) && a[ai] < b[bi] {
			c[i] = a[ai]
			ai++
		} else if bi < len(b) {
			c[i] = b[bi]
			bi++
		}
	}
	fmt.Println(c)
	return
}
