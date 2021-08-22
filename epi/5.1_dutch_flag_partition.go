package epi

/* 5.1 The Dutch national Flag Problem

SOLUTION: Time O(N) | Space O(1)
- while left ptr < right ptr:
	- if [leftPtr] < [pivot]:
		- increment leftptr
	- else if [rightPtr] >= [pivot]:
		- decrement rightptr
	- else:
		- swap [leftPtr] with [rightPtr]

- repeat the same process to place numbers equal to the pivot infront of the numbers less than the pivot.
	(leftPtr will be pointing to the position infront of the last placed number that is less than the pivot)
*/

func DutchFlagPartition(nums []int, pivot int) {
	leftPtr := 0
	rightPtr := len(nums) - 1
	pivotVal := nums[pivot]

	for leftPtr < rightPtr {
		if nums[leftPtr] < pivotVal {
			leftPtr++
		} else if nums[rightPtr] >= pivotVal {
			rightPtr--
		} else {
			nums[leftPtr], nums[rightPtr] = nums[rightPtr], nums[leftPtr]
		}
	}

	rightPtr = len(nums) - 1
	for leftPtr < rightPtr {
		if nums[leftPtr] == pivotVal {
			leftPtr++
		} else if nums[rightPtr] > pivotVal {
			rightPtr--
		} else {
			nums[leftPtr], nums[rightPtr] = nums[rightPtr], nums[leftPtr]
		}
	}
}
