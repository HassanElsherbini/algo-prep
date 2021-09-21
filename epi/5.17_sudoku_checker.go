package epi

import (
	"math"
)

// 5.17 Sudoku Checker

/* SOLUTION: Time O(N^2) | Space O(N)
- check that each row has no duplicate values
- check that each col has no duplicate values
- check that each √n x √n sub-grids has no duplicate values
*/

func IsValidSudoku(grid [][]int) bool {
	dim := len(grid)

	// check rows
	for i := 0; i < dim; i++ {
		if hasDuplicate(grid, i, i+1, 0, dim) {
			return false
		}
	}

	// check cols
	for i := 0; i < dim; i++ {
		if hasDuplicate(grid, 0, dim, i, i+1) {
			return false
		}
	}

	// check √n x √n sub-grids
	subGridDim := int(math.Sqrt(float64(dim)))
	for i := 0; i < subGridDim; i++ {
		for j := 0; j < subGridDim; j++ {
			if hasDuplicate(grid, i*subGridDim, (i+1)*subGridDim, j*subGridDim, (j+1)*subGridDim) {
				return false
			}
		}
	}

	return true
}

func hasDuplicate(grid [][]int, rowStart, rowEnd, colStart, colEnd int) bool {
	visited := make(map[int]bool)
	for r := rowStart; r < rowEnd; r++ {
		for c := colStart; c < colEnd; c++ {
			if grid[r][c] != 0 {
				if visited[grid[r][c]] {
					return true
				}
				visited[grid[r][c]] = true
			}
		}
	}

	return false
}
