package binaryheap

type BinaryHeap struct {
	Items      []interface{}
	comparator func(a, b interface{}) int
	size       int
}

func New(c func(a, b interface{}) int) *BinaryHeap {
	return &BinaryHeap{comparator: c}
}

func (bh *BinaryHeap) Insert(values ...interface{}) {
	for _, value := range values {
		bh.Items = append(bh.Items, value)
		bh.size++
		bh.heapifyUp(bh.size - 1)
	}
}

func (bh *BinaryHeap) Extract() interface{} {
	if bh.size == 0 {
		return nil
	}

	extracted := bh.Items[0]
	bh.Items[0] = bh.Items[bh.size-1]
	bh.Items = bh.Items[:bh.size-1]
	bh.size--

	bh.heapifyDown(0)

	return extracted
}

func (bh *BinaryHeap) Peek() interface{} {
	if bh.size == 0 {
		return nil
	}

	return bh.Items[0]
}

func (bh *BinaryHeap) Size() int {
	return bh.size
}

func (bh *BinaryHeap) heapifyUp(index int) {
	for bh.comparator(bh.Items[index], bh.Items[parent(index)]) > 0 {
		bh.swap(index, parent(index))
		index = parent(index)
	}
}

func (bh *BinaryHeap) heapifyDown(index int) {
	// till index is less than size of heap
	// choose higher priority child
	// if that child's priority is higher than parent
	// then swap with parent and set index to the child's index
	// else return
	for index < bh.size {
		leftIndex := left(index)
		rightIndex := right(index)
		var childToCompare int

		if leftIndex < bh.size && rightIndex < bh.size {
			if bh.comparator(bh.Items[leftIndex], bh.Items[rightIndex]) >= 0 {
				childToCompare = leftIndex
			} else {
				childToCompare = rightIndex
			}
		} else if leftIndex < bh.size {
			childToCompare = leftIndex
		} else {
			return
		}

		if bh.comparator(bh.Items[childToCompare], bh.Items[index]) > 0 {
			bh.swap(index, childToCompare)
			index = childToCompare
		} else {
			return
		}
	}
}

func (bh *BinaryHeap) swap(a, b int) {
	bh.Items[a], bh.Items[b] = bh.Items[b], bh.Items[a]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (2 * index) + 1
}

func right(index int) int {
	return (2 * index) + 2
}
