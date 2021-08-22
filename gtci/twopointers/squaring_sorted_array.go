package twopointers

// EASY

func SquareSortedArray(arr []int) []int {
	squares := make([]int, len(arr))
	i := 0
	j, resultIndx := len(arr)-1, len(arr)-1

	for i <= j {
		if arr[j]*arr[j] > arr[i]*arr[i] {
			squares[resultIndx] = arr[j] * arr[j]
			j--
		} else {
			squares[resultIndx] = arr[i] * arr[i]
			i++
		}

		resultIndx--
	}

	return squares
}
