package exercism

import (
	"fmt"
	"strings"
)

// Letter                           Value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10

// type letterValues struct {
// 	letters [6]string
// 	values  [6]int
// }

func scrabbleScore(s string) int32 {
	var result int32
	letters := [7]string{"aeioulnrst", "dg", "bcmp", "fhvwy", "k", "jx", "qz"}
	// exam string
	countLetters := [7]int32{1, 2, 3, 4, 5, 8, 10}
	for index, value := range letters {
		if !strings.ContainsAny(s, value) {
			countLetters[index] = 0
		}
	}

	// iter s
	fmt.Println(countLetters)
	for _, value := range s {

		count := strings.Count(s, string(value))
		s = strings.ReplaceAll(s, string(value), "")
		fmt.Println(count)
		for index, letter := range letters {
			if strings.ContainsAny(letter, string(value)) {
				countLetters[index] *= int32(count)
			}
			fmt.Println(countLetters)

		}
		fmt.Println(s)
	}
	fmt.Println(countLetters)
	for _, value := range countLetters {
		result += value
	}
	return result
}

func ScrabbleScore(s string) int32 {
	var result int32
	s = strings.ToLower(s)
	letters := map[int32]string{1: "aeioulnrst", 2: "dg", 3: "bcmp", 4: "fhvwy", 5: "k", 8: "jx", 10: "qz"}
	for key, value := range letters {
		for _, letter := range s {
			if strings.Contains(value, string(letter)) {
				result += key
			}
		}
	}
	return result
}
