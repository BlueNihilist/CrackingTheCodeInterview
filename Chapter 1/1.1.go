package main

import (
	"fmt"
	"strings"
)

var letters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func isUnique(str string) bool {
	s := strings.ToLower(str)
	m := make(map[rune]int)
	for _, i := range s {
		_, err := m[i]
		if (err == false) && (strings.ContainsRune(letters, i)) {
			m[i] = 1
		} else if err == true {
			return false
		}

	}
	return true
}

func main() {
	testCases := []string{"Hello!", "Helo!", "Helo!!!!", "HelLo!", "I don't ev kw ha  bu"}
	for _, k := range testCases {
		fmt.Printf("isUnique(\"%s\") = %t\n", k, isUnique(k))
	}
}
