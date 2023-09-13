package slice_queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

// instantiate a new empty queue
func New() *Queue {
	queue := &Queue{
		items: []int{},
	}
	return queue
}

// get size of queue
func (q *Queue) Size() int {
	return len(q.items)
}

// check if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size()==0
}

// get element from start of queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// push element to end of queue
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// dequeue element from start of queue
func (q *Queue) Dequeue() (int, error) {
	val, err := q.Peek()
	if err!=nil {
		return val, err
	}
	q.items = q.items[1:]
	return val, nil
}

// print elements in queue
func (q *Queue) Print() {
	fmt.Println(q.items)
}
