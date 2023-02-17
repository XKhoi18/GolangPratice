package main

import (
	"fmt"
)

func main() {
	A1 := [][]int{
		{1, 1, 2},
		{2, 0, 0},
	}
	A2 := [][]int{
		{5, -1},
		{-3, 2},
		{0, 4},
	}
	A3 := [][]int{
		{1, 1, -2},
		{3, 2, 4},
		{-1, -2, -2},
	}
	A4 := [][]int{
		{-1, 0},
		{1, 0},
	}
	fmt.Println(Solution(A1)) //3
	fmt.Println(Solution(A2)) //0
	fmt.Println(Solution(A3)) //7
	fmt.Println(Solution(A4)) //5
}

func Solution(A [][]int) int {
	n := len(A)
	m := len(A[0])
	var matrixSum float64 = 0
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrixSum += float64(A[i][j])
		}
	}

	for startRow := 0; startRow < n; startRow++ {
		for startCol := 0; startCol < m; startCol++ {
			for endRow := 0; endRow < n; endRow++ {
				for endCol := 0; endCol < m; endCol++ {
					var currSum float64 = 0
					count := 0
					for row := startRow; row <= endRow; row++ {
						for col := startCol; col <= endCol; col++ {
							currSum += float64(A[row][col])
							count += 1
						}
					}
					if count > 0 && matrixSum-currSum == currSum {
						result += 1
					}
				}
			}
		}
	}
	return result
}
