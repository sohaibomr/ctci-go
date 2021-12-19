package ch01

import "fmt"

func printMatrix(matrix [][]int) {
	for _, val := range matrix {
		fmt.Println(val)
	}
	fmt.Println("--------------------------")

}
func zeroMatrix(totalRows int, totalColumns int, matrix [][]int) {
	// find index of zero value
	// set that row and all values in that column to zero
	zeroFound := false
	var zeroRows []int
	var zeroCols []int
	fmt.Println("----------------- Original matrix ---------------------")
	printMatrix(matrix)
	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalColumns; j++ {
			if matrix[i][j] == 0 {
				zeroCols = append(zeroCols, j)
				zeroFound = true

			}
		}
		if zeroFound == true {
			zeroRows = append(zeroRows, i)
		}
		zeroFound = false
	}
	//	set rows to zero
	for _, val := range zeroRows {
		for j := 0; j < totalColumns; j++ {
			matrix[val][j] = 0
		}
	}
	for _, val := range zeroCols {
		for i := 0; i < totalRows; i++ {
			matrix[i][val] = 0
		}
	}
	fmt.Println("------------------ ZERO matrix ---------------")
	printMatrix(matrix)
}
