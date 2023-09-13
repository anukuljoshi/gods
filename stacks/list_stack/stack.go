package list_stack

import (
	"errors"

	"github.com/anukuljoshi/gods/linkedlist/singly"
)

// Stack is implemented using a SinglyListList
// with the Head as Top of Stack
type Stack struct {
	list *singly.SinglyLinkedList
}

// instantiate a new empty stack
func New() *Stack {
	stack := &Stack{
		list: singly.New(),
	}
	return stack
}

// get size of stack
func (s *Stack) Size() int {
	return s.list.Size()
}

// check if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.list.Size()==0
}

// get element from top of stack
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	return s.list.Head.Val, nil
}

// push an element to top of stack
func (s *Stack) Push(val int) {
	s.list.AddHead(val)
}

// pop element from top of stack
func (s *Stack) Pop() (int, error) {
	val, err := s.Top()
	if err!=nil {
		return val, err
	}
	s.list.DeleteHead()
	return val, nil
}

// print elements of stack
func (s *Stack) Print() {
	s.list.Print()
}
