package slice_stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

// instantiate a new empty stack
func New() *Stack {
	stack := &Stack{
		items: []int{},
	}
	return stack
}

// get size of stack
func (s *Stack) Size() int {
	return len(s.items)
}

// check if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size()==0
}

// get element from top of stack
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	return s.items[s.Size()-1], nil
}

// push element to top of stack
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// pop element from top of stack
func (s *Stack) Pop() (int, error) {
	val, err := s.Top()
	if err!=nil {
		return val, err
	}
	s.items = s.items[:s.Size()-1]
	return val, nil
}

// print elements in stack
func (s *Stack) Print() {
	fmt.Println(s.items)
}
