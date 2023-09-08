package main

import (
	"fmt"

	"github.com/anukuljoshi/gods/linkedlist/doubly"
	"github.com/anukuljoshi/gods/linkedlist/singly"
)


func main() {
	fmt.Println("1: Singly Linked List")
	fmt.Println("2: Doubly Linked List")
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
	default:
		fmt.Println("Incorrect Choice")
	}
}
