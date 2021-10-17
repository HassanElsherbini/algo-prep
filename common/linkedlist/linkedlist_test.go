package linkedlist_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/common/linkedlist"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyList(t *testing.T) {
	list := linkedlist.New()

	t.Log("Add head")
	head := list.Add(1)
	checkList(t, list, []*linkedlist.Node{head})

	t.Log("Remove head")
	list.Remove(head)
	checkList(t, list, []*linkedlist.Node{})

	t.Log("Remove when empty")
	list.Remove(head)
	checkList(t, list, []*linkedlist.Node{})

}

func TestMultiNodeList(t *testing.T) {
	list := linkedlist.New()
	head := list.Add(1)
	n2 := list.Add(2)
	n3 := list.Add(3)
	n4 := list.Add(4)
	n5 := list.Add(5)

	t.Log("Add nodes")
	checkList(t, list, []*linkedlist.Node{head, n2, n3, n4, n5})

	t.Log("Remove Head")
	list.Remove(head)
	checkList(t, list, []*linkedlist.Node{n2, n3, n4, n5})

	t.Log("Remove from mid")
	list.Remove(n3)
	checkList(t, list, []*linkedlist.Node{n2, n4, n5})

	t.Log("Remove tail")
	list.Remove(n5)
	checkList(t, list, []*linkedlist.Node{n2, n4})

}

func checkList(t *testing.T, l *linkedlist.LinkedList, nodes []*linkedlist.Node) {
	assert.Equalf(t, len(nodes), l.Len(), "List length not equal")
	for n, i := l.Head, 0; n != nil && i < len(nodes); n, i = n.Next, i+1 {
		require.Samef(t, nodes[i], n, "Node Address not the same")
		require.Equalf(t, nodes[i].Value, n.Value, "Node value not equal")
	}
}
