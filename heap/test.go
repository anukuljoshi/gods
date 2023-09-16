package heap

import "fmt"

func MinHeapTest() {
	items := []int{5,4,7,2,8,9,3,6}
	heap := NewMinHeap(items...)
	fmt.Println("*********** MinHeap", heap.Items())
	fmt.Println("*********** MinHeap Size", heap.Size())
	fmt.Println("*********** MinHeap Peek", heap.Peek())
	fmt.Println("*********** MinHeap Pop", heap.Pop())
	fmt.Println("*********** MinHeap Size", heap.Size())
	fmt.Println("*********** MinHeap", heap.Items())
	heap.Push(1)
	fmt.Println("*********** MinHeap", heap.Items())
}
