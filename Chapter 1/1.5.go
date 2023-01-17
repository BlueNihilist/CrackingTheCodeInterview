package main

import "fmt"

type stringPair struct {
	str1 string
	str2 string
}

func replaceCharCheck(x stringPair) bool {
	noDiffrences := 0
	for i := range x.str1 {
		if x.str1[i] != x.str2[i] {
			noDiffrences += 1
			if noDiffrences > 1 {
				return false
			}
		}
	}
	return true
}

// could make the add check and remove check the same function

func removeCharCheck(x stringPair) bool {
	noDiffrences := 0
	for i := range x.str2 {
		if x.str1[i+noDiffrences] != x.str2[i] {
			noDiffrences += 1
			if x.str1[i+noDiffrences] != x.str2[i] {
				return false
			}
			if noDiffrences > 1 {
				return false
			}
		}
	}
	return true
}

func addCharCheck(x stringPair) bool {
	noDiffrences := 0
	for i := range x.str2 {
		if x.str1[i] != x.str2[i+noDiffrences] {
			noDiffrences += 1
			if x.str1[i] != x.str2[i+noDiffrences] {
				return false
			}
			if noDiffrences > 1 {
				return false
			}
		}
	}
	return true
}

func oneAway(x stringPair) bool {
	if len(x.str1) == len(x.str2) {
		return replaceCharCheck(x)
	} else if (len(x.str1) - len(x.str2)) == 1 {
		return removeCharCheck(x)
	} else if (len(x.str1) - len(x.str2)) == -1 {
		return addCharCheck(x)
	} else {
		return false
	}
}

func main() {
	testCases := []stringPair{{"pale", "ple"}, {"pale", "pe"}, {"pales", "pale"}, {"palemen", "pale"}, {"pale", "bale"}, {"pale", "bake"}}
	for _, i := range testCases {
		fmt.Printf("oneAway(\"%s\", \"%s\") = %t\n", i.str1, i.str2, oneAway(i))
	}
}
