package main

import (
	"fmt"

	"github.com/anukuljoshi/gods/linkedlist/doubly"
	"github.com/anukuljoshi/gods/linkedlist/singly"
	"github.com/anukuljoshi/gods/queue/list_queue"
	"github.com/anukuljoshi/gods/queue/slice_queue"
	"github.com/anukuljoshi/gods/stacks/list_stack"
	"github.com/anukuljoshi/gods/stacks/slice_stack"
)


func main() {
	fmt.Println("1: Singly Linked List")
	fmt.Println("2: Doubly Linked List")
	fmt.Println("3: Stack using Slice")
	fmt.Println("4: Stack using LinkedList")
	fmt.Println("5: Queue using Slice")
	fmt.Println("6: Queue using LinkedList")
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
		slice_stack.SliceStackTest()
	case 4:
		// test function for linked list stack
		list_stack.ListStackTest()
	case 5:
		// test function for slice queue
		slice_queue.SliceQueueTest()
	case 6:
		// test function for list queue
		list_queue.ListQueueTest()
	default:
		fmt.Println("Incorrect Choice")
	}
}
