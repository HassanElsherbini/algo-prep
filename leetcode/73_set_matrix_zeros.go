package leetcode

/*Medium

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's,
and return the matrix.

You must do it in place.

Example 1:

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]



SOLUTION 1: Time O(M * N) | Space O(M + N)
- creat a zeroRow and zeroCol arrays to store the cols and rows that need to be set to zero
- for each cell:
	- if it is zero:
		- add its row and column to zeroRow and zeroCol arrays

- for each cell:
	- if its row or col are in zeroRow or zeroCol:
		- set the cell's value to zero

*/

// func SetMatrixZeros(matrix [][]int) {
// 	isZeroRow := make([]bool, len(matrix))
// 	isZeroCol := make([]bool, len(matrix[0]))

// 	for i, row := range matrix {
// 		for j, val := range row {
// 			if val == 0 {
// 				isZeroRow[i] = true
// 				isZeroCol[j] = true
// 			}
// 		}
// 	}

// 	for i, row := range matrix {
// 		for j, _ := range row {
// 			if isZeroRow[i] || isZeroCol[j] {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// }

/*------------------------------------------------------------------------------------

SOLUTION 2: Time O(M * N) | Space O(1)
- using first cell of each row/col to indicate whether they need to be set to zero

- if first cell is zero mark the first col as a zero col
- for each cell in matrix :
	- if cell is zero set the first cell of its row and col to zero

- for each cell in matrix starting from row 2 and col 2 :
	- if first cell of its row or col is zero:
		- set the cell to zero

- if first cell matrix[0][0] =  zero:
	- set first row to zero
- if first col is marked:
	- set first col to zero
*/

func SetMatrixZeros(matrix [][]int) {
	isFirstColZero := matrix[0][0] == 0

	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if isFirstColZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
