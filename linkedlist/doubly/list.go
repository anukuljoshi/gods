package doubly

import (
	"fmt"
)

type DoublyListNode struct {
	Val int
	Next *DoublyListNode
	Prev *DoublyListNode
}

type DoublyLinkedList struct {
	Head *DoublyListNode
}

// add to head
func (dll *DoublyLinkedList) AddHead(val int) {
	node := &DoublyListNode{
		Val: val,
	}
	if dll.Head==nil {
		dll.Head = node
		return
	}
	// add node before head
	dll.Head.Prev = node
	node.Next = dll.Head
	// update head to new node
	dll.Head = node
}

// add to tail
func (dll *DoublyLinkedList) AddTail(val int) {
	node := &DoublyListNode{
		Val: val,
	}
	// list is empty
	if dll.Head==nil {
		dll.Head = node
		return
	}
	// find tail ptr
	ptr := dll.Head
	for ptr.Next!=nil {
		ptr = ptr.Next
	}
	// add node after ptr
	node.Prev = ptr
	ptr.Next = node
}

// add before key
func (dll *DoublyLinkedList) AddBefore(key, val int) {
	if dll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dummy := &DoublyListNode{
		Next: dll.Head,
	}
	// create new node with value
	node := &DoublyListNode{
		Val: val,
	}
	// find node to insert before
	prev := dummy
	ptr := dll.Head
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
	node.Prev = prev
	node.Next = ptr
	prev.Next = node
	ptr.Prev = node
	dll.Head = dummy.Next
}

// add after key
func (dll *DoublyLinkedList) AddAfter(key, val int) {
	// list empty
	if dll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	// create new node with value
	node := &DoublyListNode{
		Val: val,
	}
	// find node to insert after
	ptr := dll.Head
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
	// ptr is last node
	if ptr.Next==nil {
		dll.AddTail(val)
		return
	}
	// insert new node after ptr
	node.Prev = ptr
	node.Next = ptr.Next
	ptr.Next.Prev = node
	ptr.Next = node
}

// delete head
func (dll *DoublyLinkedList) DeleteHead() {
	if dll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dll.Head = dll.Head.Next
	dll.Head.Prev = nil
}

// delete tail
func (dll *DoublyLinkedList) DeleteTail() {
	if dll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dummy := &DoublyListNode{
		Next: dll.Head,
	}
	// find tail node
	prev := dummy
	ptr := dll.Head
	for ptr!=nil {
		if ptr.Next==nil {
			break
		}
		prev = ptr
		ptr = ptr.Next
	}
	prev.Next = nil
	dll.Head = dummy.Next
}

// delete key
func (dll *DoublyLinkedList) DeleteKey(key int) {
	if dll.Head==nil {
		fmt.Println("list is empty")
		return
	}
	dummy := &DoublyListNode{
		Next: dll.Head,
	}
	// find node with key
	prev := dummy
	ptr := dll.Head
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
	// key at last node
	if ptr.Next==nil {
		dll.DeleteTail()
		return
	}
	prev.Next = ptr
	ptr.Next.Prev = prev
	dll.Head = dummy.Next
}

// find a key in list
func (dll *DoublyLinkedList) Find(key int) bool {
	ptr := dll.Head
	for ptr!=nil {
		if ptr.Val==key {
			return true
		}
		ptr = ptr.Next
	}
	return false
}

// find length of list
func (dll *DoublyLinkedList) Length() int {
	ptr := dll.Head
	length := 0
	for ptr!=nil {
		length += 1
		ptr = ptr.Next
	}
	return length
}

// print elements of list
func (dll *DoublyLinkedList) Print() {
	fmt.Print("L2R: ")
	dll.PrintL2R()
	fmt.Print("R2L: ")
	dll.PrintR2L()
}

// print list elements from left to right
func (dll *DoublyLinkedList) PrintL2R()  {
	ptr := dll.Head
	for ptr!=nil {
		fmt.Printf("%d <-> ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println()
}

// print list elements from right to left
func (dll *DoublyLinkedList) PrintR2L()  {
	ptr := dll.Head
	for ptr!=nil && ptr.Next!=nil {
		ptr = ptr.Next
	}
	for ptr!=nil {
		fmt.Printf("%d <-> ", ptr.Val)
		ptr = ptr.Prev
	}
	fmt.Println()
}
