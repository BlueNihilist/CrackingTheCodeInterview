package main

import (
	"fmt"
	"strings"
)

type TestCase struct {
	s1 string
	s2 string
}

func isSubstring(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func stringRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	str := s1 + s1
	return isSubstring(str, s2)
}

func main() {
	testCases := []TestCase{{"waterbottle", "aterbottlew"}, {"aba", "ba"}, {"abc", "cda"}, {"abc", "bac"}, {"abac", "baca"}}
	for _, i := range testCases {
		fmt.Printf("\"%s\" -> \"%s\" | %t\n", i.s1, i.s2, stringRotation(i.s1, i.s2))
	}
}
