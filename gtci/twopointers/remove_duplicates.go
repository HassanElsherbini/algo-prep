package twopointers

/* EASY

SOLUTION:
-if array size is less than 2 return 0
-while placement index is less than n-1
	-find next element that is not equal to the element at placement index
	-set found element at placementIndx+1
	-increment placementIndx

*/

func RemoveDuplicates(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	placementIndx := 0
	for searchIndx := 1; searchIndx < len(arr); searchIndx++ {
		if arr[searchIndx] != arr[placementIndx] {
			placementIndx++
			arr[placementIndx] = arr[searchIndx]
		}
	}

	return placementIndx + 1
}
