package main

import "fmt"

// https://tproger.ru/problems/algorithm-that-displays-all-valid-combinations-of-pairs-of-brackets/

// https://coderoad.ru/29321063/%D0%B2%D1%81%D0%B5-%D0%B4%D0%BE%D0%BF%D1%83%D1%81%D1%82%D0%B8%D0%BC%D1%8B%D0%B5-%D0%BA%D0%BE%D0%BC%D0%B1%D0%B8%D0%BD%D0%B0%D1%86%D0%B8%D0%B8-n-%D0%BF%D0%B0%D1%80-%D1%81%D0%BA%D0%BE%D0%B1%D0%BE%D0%BA
// тут в каментах пример

//https://overcoder.net/q/152180/%D0%BF%D0%BE%D0%B8%D1%81%D0%BA-%D0%B2%D1%81%D0%B5%D1%85-%D0%BA%D0%BE%D0%BC%D0%B1%D0%B8%D0%BD%D0%B0%D1%86%D0%B8%D0%B9-%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D0%BB%D1%8C%D0%BD%D1%8B%D1%85-%D1%81%D0%BA%D0%BE%D0%B1%D0%BE%D0%BA

func main() {
	n := 3
	// result := make([]string, 0)
	// for i := 1; i <= n; i++ {
	// 	a := nPair(i)
	// 	result = append(result, a...)
	// }

	result := nPair(n)
	fmt.Printf("%v\n", result)
}

func nPair(n int) []string {
	if n == 0 {
		return []string{""}
	}

	result := make([]string, 0)

	for i := 0; i < n; i++ {
		lefts := nPair(i)
		rights := nPair(n - i - 1)

		for l := 0; l < len(lefts); l++ {
			for r := 0; r < len(rights); r++ {
				result = append(result, "("+lefts[l]+")"+rights[r])
			}
		}
	}

	return result

}
