package main

import (
	"fmt"
	"strings"
)

var letters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func palindromePermutation(str string) bool {
	m := make(map[rune]int)
	s := strings.ToLower(str)
	noEvens := 0
	for _, i := range s {
		if strings.ContainsRune(letters, i) {
			m[i] += 1
		}
	}
	for _, val := range m {
		if val%2 == 1 {
			noEvens += 1
		}
		if noEvens >= 2 {
			return false
		}
	}
	return true
}

func main() {
	testCases := []string{"Tact Coa", "ab ba", "ABBBBBB", "ABBBBB", "!!!!!@@@@@", "!!H!!!"}
	for _, i := range testCases {
		fmt.Printf("palindromePermutation(\"%s\") = %t\n", i, palindromePermutation(i))
	}
}
