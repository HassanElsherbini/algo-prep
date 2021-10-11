package topk

import "github.com/HassanElsherbini/algo-prep/common/binaryheap"

/* MEDIUM

SOLUTION: Time O(N * LOG(N)) | Space O(N)
- store all numbers with frequency > 1 in hashmap
- put all numbers in hash map into a min frequency heap
- set a unique counter = number of existing unique numbers
- till k = 0 or heap is empty:
	- extract least frequent number from heap
	- subtract freq - 1 from k
	- if k >= 0
		- increment unique counter

- if k > 0 : remove remaining k numbers from unique counter
- return counter
*/

func FindMaxDistinctNumbers(nums []int, k int) int {
	numFrequency := make(map[int]int)
	repeatedNums := []int{}

	for _, num := range nums {
		numFrequency[num]++
		if numFrequency[num] == 2 {
			repeatedNums = append(repeatedNums, num)
		}
	}

	comparator := func(a, b interface{}) int {
		assertedA := a.(int)
		assertedB := b.(int)
		if assertedA < assertedB {
			return 1
		}
		if assertedA > assertedB {
			return -1
		}
		return 0
	}

	minHeap := binaryheap.New(comparator)
	for _, num := range repeatedNums {
		minHeap.Insert(numFrequency[num])
	}

	result := len(numFrequency) - len(repeatedNums)
	for k > 0 && minHeap.Size() > 0 {
		frequency := minHeap.Extract().(int)
		k = k - (frequency - 1)
		if k >= 0 {
			result++
		}
	}

	if k > 0 {
		result -= k
	}

	return result
}
