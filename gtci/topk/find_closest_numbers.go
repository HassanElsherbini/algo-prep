package topk

import (
	"math"
	"sort"

	"github.com/HassanElsherbini/algo-prep/common"
	"github.com/HassanElsherbini/algo-prep/common/binaryheap"
)

/* MEDIUM

SOLUTION (min heap): Time O(lOG(n) + K * LOG(K)) | Space O(K)
- find number y that is the closest to x using binary first search
- place all k-1 numbers before y and all k numbers after y in a min heap where closeness to x is the priority
- extract top k numbers from heap and return them sorted

SOLUTION (two pointers): Time O(lOG(N) + K) | Space O(1)
- find index y of number closest to x using binary first search
- use two pointers a and b, both starting at index at y
- untill k-1 numbers are found:
	- if number at a is closer to x than number at b:
		- decrement a
	- else increment b

	return all numbers between a and b inclusive
*/

func FindClosestNumbersUsingTwoPointers(nums []int, k, x int) []int {
	if k > len(nums) {
		return nil
	}

	y := findPositionOfClosestNumber(nums, x)
	start := y
	end := y

	for i := 1; i < k; i++ {
		if start-1 >= 0 && end+1 < len(nums) {
			distToA := math.Abs(float64(x - nums[start-1]))
			distToB := math.Abs(float64(x - nums[end+1]))
			if distToA < distToB {
				start--
			} else {
				end++
			}
		} else if start-1 >= 0 {
			start--
		} else {
			end++
		}
	}

	return nums[start : end+1]
}

func FindClosestNumbers(nums []int, k, x int) []int {
	comparator := func(a, b interface{}) int {
		assertedA := a.(int)
		assertedb := b.(int)
		distToA := math.Abs(float64(x - assertedA))
		distToB := math.Abs(float64(x - assertedb))

		if distToA < distToB {
			return 1
		}
		if distToA > distToB {
			return -1
		}
		return 0
	}

	minHeap := binaryheap.New(comparator)
	if k > len(nums) {
		return nil
	}

	y := findPositionOfClosestNumber(nums, x)
	start := common.Max(y-k, 0)
	end := common.Min(y+k, len(nums)-1)

	for i := start; i <= end; i++ {
		minHeap.Insert(nums[i])
	}

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, minHeap.Extract().(int))
	}

	sort.Ints(result)
	return result
}

func findPositionOfClosestNumber(nums []int, key int) int {
	if key < nums[0] {
		return 0
	}

	n := len(nums)
	if key > nums[n-1] {
		return n - 1
	}

	var start int
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2
		if key == nums[mid] {
			return mid
		}

		if key > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if (nums[start] - key) < (key - nums[end]) {
		return start
	}
	return end
}
