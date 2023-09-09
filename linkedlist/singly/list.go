package singly

import (
	"fmt"
)

type SinglyListNode struct {
	Val int
	Next *SinglyListNode
}

type SinglyLinkedList struct {
	Head *SinglyListNode
}

// insert a node at head
func (sll *SinglyLinkedList) AddHead(val int)  {
	node := &SinglyListNode{
		Val: val,
		Next: sll.Head,
	}
	sll.Head = node
}

// insert a node at tail
func (sll *SinglyLinkedList) AddTail(val int)  {
	node := &SinglyListNode{
		Val: val,		
	}
	// list is empty
	if sll.Head==nil {
		sll.Head = node
		return
	}
	// find tail ptr
	ptr := sll.Head
	for ptr.Next!=nil {
		ptr = ptr.Next
	}
	ptr.Next = node
}

// insert a Node with val before the first occurrence of key
func (sll *SinglyLinkedList) AddBefore(key, val int) {
	// list empty
	if sll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dummy := &SinglyListNode{
		Next: sll.Head,
	}
	// create new node with value
	node := &SinglyListNode{
		Val: val,
	}
	// find node to insert before
	prev := dummy
	ptr := sll.Head
	for ptr!=nil {
		if ptr.Val==key {
			break
		}
		prev = ptr
		ptr = ptr.Next
	}
	// key not found
	if ptr==nil {
		fmt.Printf("%d not found in list\n", key)
		return
	}
	// insert node between prev and ptr
	prev.Next = node
	node.Next = ptr
	sll.Head = dummy.Next
}

// add after key
func (sll *SinglyLinkedList) AddAfter(key, val int) {
	// list empty
	if sll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	// create new node with new value
	node := &SinglyListNode{
		Val: val,
	}
	// find node to insert after
	ptr := sll.Head
	for ptr!=nil {
		if ptr.Val==key {
			break
		}
		ptr = ptr.Next
	}
	// key not found
	if ptr==nil {
		fmt.Printf("%d not found in list\n", key)
		return
	}
	// insert new node after ptr
	node.Next = ptr.Next
	ptr.Next = node
}

// deletion at head
func (sll *SinglyLinkedList) DeleteHead()  {
	if sll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	sll.Head = sll.Head.Next
}

// deletion at tail
func (sll *SinglyLinkedList) DeleteTail() {
	if sll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dummy := &SinglyListNode{
		Next: sll.Head,
	}
	// find tail node
	prev := dummy
	ptr := sll.Head
	for ptr!=nil {
		if ptr.Next==nil {
			break
		}
		prev = ptr
		ptr = ptr.Next
	}
	prev.Next = nil
	sll.Head = dummy.Next
}

// delete key
func (sll *SinglyLinkedList) DeleteKey(key int) {
	if sll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dummy := &SinglyListNode{
		Next: sll.Head,
	}
	// find node with key
	prev := dummy
	ptr := sll.Head
	for ptr!=nil {
		if ptr.Val==key {
			break
		}
		prev = ptr
		ptr = ptr.Next
	}
	// key not found
	if ptr==nil {
		fmt.Printf("%d not found in list\n", key)
		return
	}
	prev.Next = ptr.Next
	sll.Head = dummy.Next
}

// find key in list
func (sll *SinglyLinkedList) Contains(key int) bool {
	ptr := sll.Head
	for ptr!=nil {
		if ptr.Val==key {
			return true
		}
		ptr = ptr.Next
	}
	return false
}

// find length of LinkedList
func (sll *SinglyLinkedList) Length() int {
	ptr := sll.Head
	length := 0
	for ptr!=nil {
		length += 1
		ptr = ptr.Next
	}
	return length
}

// print elements of list
func (sll *SinglyLinkedList) Print()  {
	ptr := sll.Head
	for ptr!=nil {
		fmt.Printf("%d -> ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println()
}
