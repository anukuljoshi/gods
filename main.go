package main

import (
	"fmt"

	"github.com/anukuljoshi/gods/linkedlist/doubly"
	"github.com/anukuljoshi/gods/linkedlist/singly"
	"github.com/anukuljoshi/gods/stacks/slicestack"
)


func main() {
	fmt.Println("1: Singly Linked List")
	fmt.Println("2: Doubly Linked List")
	fmt.Println("3: Stack using Slice")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		// test function for singly linked list
		singly.SinglyLinkedListTest()
	case 2:
		// test function for doubly linked list
		doubly.DoublyLinkedListTest()
	case 3:
		// test function for slice stack
		slicestack.SliceStackTest()
	default:
		fmt.Println("Incorrect Choice")
	}
}
