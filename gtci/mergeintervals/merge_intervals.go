package mergeintervals

import (
	"sort"

	"github.com/HassanElsherbini/algo-prep/common"
)

/* MEDIUM

SOLUTION:
 - sort intervals by end time
 - for each intervals
     if interval start overlaps the running merged interval
       merge interval with the running merged interval
     else
       add running merged interval to result
       set interval as the new running merged interval

 - add remaining running merged interval to result, since the loop will end without adding the last running merged interval
*/

func MergeIntervals(intervals []common.Interval) []common.Interval {
	var result []common.Interval

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].Start < intervals[b].Start
	})

	start := intervals[0].Start
	end := intervals[0].End

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval.Start <= end {
			end = common.Max(end, interval.End)
		} else {
			result = append(result, common.Interval{Start: start, End: end})
			start = interval.Start
			end = interval.End
		}
	}

	result = append(result, common.Interval{Start: start, End: end})

	return result
}
