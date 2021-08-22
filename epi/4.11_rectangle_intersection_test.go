package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/common"
	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestRectangleIntersection(t *testing.T) {
	tests := []struct {
		a        common.Rectangle
		b        common.Rectangle
		expected common.Rectangle
	}{
		// A encompasses B
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: 1, Y: 0}, TopRight: common.Point2D{X: 8, Y: 4}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 2, Y: 1}, TopRight: common.Point2D{X: 4, Y: 3}},
			expected: common.Rectangle{BottomLeft: common.Point2D{X: 2, Y: 1}, TopRight: common.Point2D{X: 4, Y: 3}},
		},
		// B encompasses A
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: 1, Y: 0}, TopRight: common.Point2D{X: 8, Y: 4}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 2, Y: 1}, TopRight: common.Point2D{X: 4, Y: 3}},
			expected: common.Rectangle{BottomLeft: common.Point2D{X: 2, Y: 1}, TopRight: common.Point2D{X: 4, Y: 3}},
		},
		// A intersecting B from the left
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: 0, Y: 0}, TopRight: common.Point2D{X: 8, Y: 4}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 7, Y: 1}, TopRight: common.Point2D{X: 12, Y: 3}},
			expected: common.Rectangle{BottomLeft: common.Point2D{X: 7, Y: 1}, TopRight: common.Point2D{X: 8, Y: 3}},
		},
		// A intersecting B from the right
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: 3, Y: 0}, TopRight: common.Point2D{X: 10, Y: 5}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 0, Y: 1}, TopRight: common.Point2D{X: 4, Y: 3}},
			expected: common.Rectangle{BottomLeft: common.Point2D{X: 3, Y: 1}, TopRight: common.Point2D{X: 4, Y: 3}},
		},
		// A intersecting B vertically (a cross)
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: 4, Y: -2}, TopRight: common.Point2D{X: 8, Y: 4}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 0, Y: 0}, TopRight: common.Point2D{X: 10, Y: 2}},
			expected: common.Rectangle{BottomLeft: common.Point2D{X: 4, Y: 0}, TopRight: common.Point2D{X: 8, Y: 2}},
		},
		//empty intersection - A touching left side of B
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: -4, Y: 1}, TopRight: common.Point2D{X: 0, Y: 3}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 0, Y: 0}, TopRight: common.Point2D{X: 8, Y: 4}},
			expected: common.Rectangle{},
		},
		// empty intersection - A touching bottom left corner of B
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: 0, Y: 0}, TopRight: common.Point2D{X: 8, Y: 4}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: -4, Y: -2}, TopRight: common.Point2D{X: 0, Y: 0}},
			expected: common.Rectangle{},
		},
		// no intersection
		{
			a:        common.Rectangle{BottomLeft: common.Point2D{X: -4, Y: 1}, TopRight: common.Point2D{X: -1, Y: 3}},
			b:        common.Rectangle{BottomLeft: common.Point2D{X: 0, Y: 0}, TopRight: common.Point2D{X: 8, Y: 4}},
			expected: common.Rectangle{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.RectangleIntersection(test.a, test.b))
	}
}
