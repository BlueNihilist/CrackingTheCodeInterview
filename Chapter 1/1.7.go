package main

import "fmt"

func rotateElem(matrix [][]int, i int, j int) [][]int {
	t := matrix[i][j]
	n := len(matrix[0]) - 1
	matrix[i][j] = matrix[n-j][i]
	matrix[n-j][i] = matrix[n-i][n-j]
	matrix[n-i][n-j] = matrix[j][n-i]
	matrix[j][n-i] = t
	return matrix
}

func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix[0])
	for i := 0; i < int((n+1)/2); i++ {
		for j := 0; j < int(n/2); j++ {
			matrix = rotateElem(matrix, i, j)
		}
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	testCases := [][][]int{{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}}
	for _, matrix := range testCases {
		printMatrix(matrix)
		fmt.Printf("formated to:\n")
		printMatrix(rotateMatrix(matrix))
		fmt.Printf("-----------------------------\n")
	}
}
