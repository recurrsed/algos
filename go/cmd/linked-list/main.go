package main

import (
	linkedlist "github.com/recurrsed/algos/internal/linked-list"
)

func main() {

	list := linkedlist.LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Display()
}
