package bfs

import (
	"container/list"

	"github.com/HassanElsherbini/algo-prep/common"
)

/* EASY

  SOLUTION:
		- viewing a tree from the right means that only the right most node of each tree level is visible.
		- Traverse the tree in level order using a queue and take the right most node of each level
*/

func RightViewOfBinaryTree(root *common.BinaryTree) []int {
	result := []int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelNodeCount := queue.Len()
		if rightMostNode, ok := queue.Back().Value.(*common.BinaryTree); ok {
			result = append(result, rightMostNode.Value)
		}

		for i := 0; i < levelNodeCount; i++ {
			val := queue.Remove(queue.Front())
			if node, ok := val.(*common.BinaryTree); ok {
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
		}
	}
	return result
}
