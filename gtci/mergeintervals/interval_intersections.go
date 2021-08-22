package mergeintervals

import "github.com/HassanElsherbini/algo-prep/common"

// MEDIUM

//SOLUTION 1: TIME O(N + M) | SPACE O(N + M)
// func intervalIntersections2(intervalsA, intervalsB []common.Interval) []common.Interval {
// 	mergedSortedIntervals := mergeSortedIntervals(intervalsA, intervalsB)
// 	result := []common.Interval{}

// 	end := mergedSortedIntervals[0].End
// 	for i := 1; i < len(mergedSortedIntervals); i++ {
// 		interval := mergedSortedIntervals[i]
// 		if interval.Start <= end {
// 			intersection := common.Interval{
// 				interval.Start,
// 				common.Min(interval.End, end),
// 			}
// 			result = append(result, intersection)
// 			end = common.Max(interval.End, end)
// 		} else {
// 			end = interval.End
// 		}
// 	}

// 	return result
// }

// func mergeSortedIntervals(intervalsA, intervalsB []common.Interval) []common.Interval {
// 	result := []common.Interval{}
// 	var i, j int

// 	for i < len(intervalsA) || j < len(intervalsB) {
// 		var interval common.Interval
// 		if i >= len(intervalsA) {
// 			interval = intervalsB[j]
// 			j++
// 		} else if j >= len(intervalsB) {
// 			interval = intervalsA[i]
// 			i++
// 		} else if intervalsA[i].Start <= intervalsB[j].Start {
// 			interval = intervalsA[i]
// 			i++
// 		} else {
// 			interval = intervalsB[j]
// 			j++
// 		}

// 		result = append(result, interval)
// 	}

// 	return result
// }

//SOLUTION 2: TIME O(N + M) | SPACE O(1)  - NO NEED FOR MERGING
func IntervalIntersections(intervalsA, intervalsB []common.Interval) []common.Interval {
	var i, j int
	result := []common.Interval{}

	for i < len(intervalsA) && j < len(intervalsB) {
		intervalA := intervalsA[i]
		intervalB := intervalsB[j]

		isAOverlappingB := intervalA.Start <= intervalB.Start && intervalB.Start <= intervalA.End
		isBOverlappingA := intervalB.Start <= intervalA.Start && intervalA.Start <= intervalB.End

		if isAOverlappingB || isBOverlappingA {
			result = append(result, common.Interval{
				common.Max(intervalA.Start, intervalB.Start),
				common.Min(intervalA.End, intervalB.End),
			})
		}

		if intervalA.End <= intervalB.End {
			i++
		} else {
			j++
		}
	}

	return result
}
