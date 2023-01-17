package main

import (
	"fmt"
	"strconv"
)

func stringCompression(str string) string {
	res := []rune{}
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		curRune := s[i]
		noRune := 1
		for i < len(s)-1 {
			if s[i+1] == curRune {
				i += 1
				noRune += 1
			} else {
				break
			}
		}
		res = append(res, curRune)
		res = append(res, []rune(strconv.Itoa(noRune))...)
	}
	return string(res)
}

func main() {
	testCases := []string{"aabcccccaaa", "a", "qwertyuiopasdfghjklzxcvbnm", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbcccdddddddddddddddddd"}
	for _, i := range testCases {
		fmt.Printf("\"%s\" -> \"%s\"\n", i, stringCompression(i))
	}
}
