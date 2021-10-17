package linkedlist

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	len  int
}

func New() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) Add(value interface{}) *Node {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}

	l.len++
	return newNode
}

func (l *LinkedList) Remove(node *Node) {
	if l.Head == nil {
		return
	}

	if l.Head == node {
		l.Head = l.Head.Next
	} else {
		cur := l.Head
		for cur.Next != nil && cur.Next != node {
			cur = cur.Next
		}
		// node not found
		if cur.Next == nil {
			return
		}

		cur.Next = cur.Next.Next
	}

	node.Next = nil
	l.len--
}

func (l *LinkedList) Len() int {
	return l.len
}
