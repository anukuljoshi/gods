package list_stack

import (
	"errors"

	"github.com/anukuljoshi/gods/linkedlist/singly"
)

// Stack is implemented using a SinglyListList
// with the Head as Top of Stack
type Stack struct {
	items *singly.SinglyLinkedList
}

// instantiate a new empty stack
func New() *Stack {
	stack := &Stack{
		items: singly.New(),
	}
	return stack
}

// get size of stack
func (s *Stack) Size() int {
	return s.items.Size()
}

// check if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.items.Size()==0
}

// get element from top of stack
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	return s.items.Head.Val, nil
}

// push an element to top of stack
func (s *Stack) Push(val int) {
	s.items.AddHead(val)
}

// pop element from top of stack
func (s *Stack) Pop() (int, error) {
	val, err := s.Top()
	if err!=nil {
		return val, err
	}
	s.items.DeleteHead()
	return val, nil
}

// print elements of stack
func (s *Stack) Print() {
	s.items.Print()
}
