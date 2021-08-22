package epi

import "github.com/HassanElsherbini/algo-prep/common"

/* 4.11 Rectangle Intersection

SOLUTION: Time O(1) | Space O(1)
	- if the two rectangles don't intersecting rectangle (bottom left point of one rectangle falling within the x and y
		ranges of the other rectangle):
			- return empty rectangle

	- calculate the bottom-left and top-right points of the intersection rectangle and return it
*/
func RectangleIntersection(a, b common.Rectangle) common.Rectangle {
	if !formsIntersectingRectangle(a, b) {
		return common.Rectangle{}
	}

	return common.Rectangle{
		BottomLeft: common.Point2D{
			X: common.Max(a.BottomLeft.X, b.BottomLeft.X),
			Y: common.Max(a.BottomLeft.Y, b.BottomLeft.Y),
		},
		TopRight: common.Point2D{
			X: common.Min(a.TopRight.X, b.TopRight.X),
			Y: common.Min(a.TopRight.Y, b.TopRight.Y),
		},
	}
}

func formsIntersectingRectangle(a, b common.Rectangle) bool {
	intersectingX := (a.BottomLeft.X >= b.BottomLeft.X && a.BottomLeft.X < b.TopRight.X) || (b.BottomLeft.X >= a.BottomLeft.X && b.BottomLeft.X < a.TopRight.X)
	intersectingY := (a.BottomLeft.Y >= b.BottomLeft.Y && a.BottomLeft.Y < b.TopRight.Y) || (b.BottomLeft.Y >= a.BottomLeft.Y && b.BottomLeft.Y < a.TopRight.Y)

	return intersectingX && intersectingY
}
