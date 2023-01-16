package main

import "fmt"

func checkPermutation(str1 string, str2 string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, i := range str1 {
		m1[i] += 1
	}
	for _, i := range str2 {
		m2[i] += 1
	}

	for key, val := range m1 {
		if val != m2[key] {
			return false
		}
	}
	for key, val := range m2 {
		if val != m1[key] {
			return false
		}
	}

	return true
}

func main() {
	testCases1 := []string{"ABBA", "ABBA", "ABBA", "COcoa", "taco cat", "timbER", "timbER "}
	testCases2 := []string{"ABBA", "ABBACS", "BABA", "coCOa", "tactaco ", "timbRE", "timbER"}
	for i := range testCases1 {
		fmt.Printf("checkPermutation(\"%s\", \"%s\") = %t\n", testCases1[i], testCases2[i], checkPermutation(testCases1[i], testCases2[i]))
	}
}
