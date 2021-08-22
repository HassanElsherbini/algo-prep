package mergeintervals_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/common"
	"github.com/HassanElsherbini/algo-prep/gtci/mergeintervals"
	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		intervals []common.Interval
		expected  []common.Interval
	}{
		{
			[]common.Interval{{1, 4}, {2, 5}, {7, 9}},
			[]common.Interval{{1, 5}, {7, 9}},
		},
		{
			[]common.Interval{{6, 7}, {2, 4}, {5, 9}},
			[]common.Interval{{2, 4}, {5, 9}},
		},
		{
			[]common.Interval{{1, 4}, {2, 6}, {3, 5}},
			[]common.Interval{{1, 6}},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, mergeintervals.MergeIntervals(tt.intervals))
	}
}

func TestIntervalIntersections(t *testing.T) {
	tests := []struct {
		intervalA []common.Interval
		intervalB []common.Interval
		expected  []common.Interval
	}{
		{
			[]common.Interval{{1, 3}, {5, 6}, {7, 9}},
			[]common.Interval{{2, 3}, {5, 7}},
			[]common.Interval{{2, 3}, {5, 6}, {7, 7}},
		},
		{
			[]common.Interval{{1, 3}, {5, 7}, {9, 12}},
			[]common.Interval{{5, 10}},
			[]common.Interval{{5, 7}, {9, 10}},
		},
	}

	for _, tt := range tests {
		assert.Equal(t,
			tt.expected,
			mergeintervals.IntervalIntersections(tt.intervalA, tt.intervalB),
		)
	}
}
