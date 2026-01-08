package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) Append(val int) {
	newNode := &Node{Value: val}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (list *LinkedList) Display() {
	current := list.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
}
