package slicestack

import (
	"errors"
	"fmt"
)

type Stack []int

// get size of stack
func (s *Stack) Size() int {
	return len(*s)
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
	return (*s)[s.Size()-1], nil
}

// push element to top of stack
func (s *Stack) Push(val int) {
	(*s) = append((*s), val)
}

// pop element from top of stack
func (s *Stack) Pop() (int, error) {
	val, err := s.Top()
	if err!=nil {
		return val, err
	}
	(*s) = (*s)[:s.Size()-1]
	return val, nil
}

// print elements in stack
func (s *Stack) Print() {
	fmt.Println(*s)
}
