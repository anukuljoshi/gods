package list_queue

import (
	"errors"

	"github.com/anukuljoshi/gods/linkedlist/singly"
)

type Queue struct {
	items *singly.SinglyLinkedList
}

// instantiate an empty queue
func New() *Queue {
	queue := &Queue{
		items: singly.New(),
	}
	return queue
}

// get size of stack
func (s *Queue) Size() int {
	return s.items.Size()
}

// check if stack is empty
func (s *Queue) IsEmpty() bool {
	return s.items.Size()==0
}

// get element from start of queue
func (s *Queue) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	return s.items.Head.Val, nil
}

// enqueue an element to end of queue
func (s *Queue) Enqueue(val int) {
	s.items.AddTail(val)
}

// queue element from start of queue
func (s *Queue) Dequeue() (int, error) {
	val, err := s.Peek()
	if err!=nil {
		return val, err
	}
	s.items.DeleteHead()
	return val, nil
}

// print elements of stack
func (s *Queue) Print() {
	s.items.Print()
}
