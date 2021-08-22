package common

import "math"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type Interval struct {
	Start int
	End   int
}

type Point2D struct {
	X int
	Y int
}

func (p Point2D) DistanceFromOrigin() float64 {
	xDiff := math.Pow(float64(p.X), 2)
	yDiff := math.Pow(float64(p.Y), 2)

	return math.Sqrt(xDiff + yDiff)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

type Rectangle struct {
	BottomLeft Point2D
	TopRight   Point2D
}
