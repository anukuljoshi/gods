package array

import (
	"errors"
	"fmt"
)

type Stack []int

// size
func (s *Stack) Size() int {
	return len(*s)
}

// isEmpty
func (s *Stack) IsEmpty() bool {
	return s.Size()==0
}

// top
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	return (*s)[s.Size()-1], nil
}

// push
func (s *Stack) Push(val int) {
	(*s) = append((*s), val)
}

// pop
func (s *Stack) Pop() (int, error) {
	val, err := s.Top()
	if err!=nil {
		return val, err
	}
	(*s) = (*s)[:s.Size()-1]
	return val, nil
}

// print
func (s *Stack) Print() {
	fmt.Println(*s)
}
