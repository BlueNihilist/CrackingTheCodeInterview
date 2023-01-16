package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []string{"Mr Johnson  ", " Dr Watson    ", "Hee hee haa haa      ", "^*%&&GYF^&^UBIU@W&YDGBW   "}
	testCasesLen := []int{10, 10, 15, 24}
	for i := range testCases {
		str := testCases[i][:testCasesLen[i]]
		str = strings.ReplaceAll(str, " ", "%20")
		fmt.Printf("\"%s\" -> \"%s\"\n", testCases[i], str)
	}
}
