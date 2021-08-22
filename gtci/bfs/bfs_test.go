package bfs_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/common"
	"github.com/HassanElsherbini/algo-prep/gtci/bfs"
	"github.com/stretchr/testify/assert"
)

func TestRightViewOfBinaryTree(t *testing.T) {
	tree1 := &common.BinaryTree{Value: 12}
	tree1.Left = &common.BinaryTree{Value: 7}
	tree1.Right = &common.BinaryTree{Value: 1}
	tree1.Left.Left = &common.BinaryTree{Value: 9}
	tree1.Right.Left = &common.BinaryTree{Value: 10}
	tree1.Right.Right = &common.BinaryTree{Value: 5}
	tree1.Left.Left.Left = &common.BinaryTree{Value: 3}

	tree2 := &common.BinaryTree{Value: 1}
	tree2.Left = &common.BinaryTree{Value: 2}
	tree2.Right = &common.BinaryTree{Value: 3}
	tree2.Left.Left = &common.BinaryTree{Value: 4}
	tree2.Left.Right = &common.BinaryTree{Value: 5}
	tree2.Right.Left = &common.BinaryTree{Value: 6}
	tree2.Right.Right = &common.BinaryTree{Value: 7}

	tests := []struct {
		tree     *common.BinaryTree
		expected []int
	}{
		{tree1, []int{12, 1, 5, 3}},
		{tree2, []int{1, 3, 7}},
	}

	for _, tt := range tests {
		actual := bfs.RightViewOfBinaryTree(tt.tree)
		assert.Equal(t, tt.expected, actual)
	}
}
