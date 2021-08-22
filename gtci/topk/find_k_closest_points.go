package topk

import (
	"github.com/HassanElsherbini/algo-prep/common"
)

/* EASY

SOLUTION: Time O(K * LOG(K) + N-K * LOG(K)) | Space O(K)
- add first K points to max heap
- for each point in remaining points:
	- if point is closer to origin than the current K closest point (the root):
		- remove K closest point
		- add point to heap
- return all K points in the heap
*/

func FindKClosestPoints(points []common.Point2D, k int) []common.Point2D {
	comparator := func(a, b interface{}) int {
		pointA := a.(common.Point2D)
		pointB := b.(common.Point2D)

		aDist := pointA.DistanceFromOrigin()
		bDist := pointB.DistanceFromOrigin()

		if aDist > bDist {
			return 1
		}

		if aDist < bDist {
			return -1
		}
		return 0
	}

	maxHeap := common.NewBinaryHeap(comparator)
	for i := 0; i < k; i++ {
		maxHeap.Insert(points[i])
	}

	for i := k; i < len(points); i++ {
		point := points[i]
		kClosest := maxHeap.Peek().(common.Point2D)
		if point.DistanceFromOrigin() < kClosest.DistanceFromOrigin() {
			maxHeap.Extract()
			maxHeap.Insert(point)
		}
	}

	result := []common.Point2D{}
	for _, item := range maxHeap.Items {
		result = append(result, item.(common.Point2D))
	}

	return result
}
