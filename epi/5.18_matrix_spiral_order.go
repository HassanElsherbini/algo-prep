package epi

// 5.18  Compute the spiral ordering of a 2D array (N x N)

/* SOLUTION 1: Time O(N) | Space O(1) excluding result array
- while left boundary < right boundary || top boundary < bottom boundary:
	- add top row: all numbers from left to right boundary
	- increment top boundary
	- add left row: add all numbers from top to bottom boundary
	- decrement right boundary
	- add bottom row: add all numbers from right to left boundary
	- decrement bottom boundary
	- add left row: add all numbers from right to left boundary
	- incement left boundary
*/

func MatrixSpiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}

	rows := len(matrix)
	cols := len(matrix[0])

	top, right, bottom, left := 0, cols-1, rows-1, 0

	for left <= right && top <= bottom {
		//top
		for l := left; l <= right; l++ {
			result = append(result, matrix[top][l])
		}
		top++

		//right
		for t := top; t <= bottom; t++ {
			result = append(result, matrix[t][right])
		}
		right--

		//bottom
		for r := right; r >= left; r-- {
			result = append(result, matrix[bottom][r])
		}
		bottom--

		//left
		for b := bottom; b >= top; b-- {
			result = append(result, matrix[b][left])
		}
		left++
	}

	return result
}
