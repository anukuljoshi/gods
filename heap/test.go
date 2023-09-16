package heap

import "fmt"

func MinHeapTest() {
	items := []int{5,6,7,2,8,9,3,1}
	heap := NewMinHeap(items...)
	fmt.Println("*********** MinHeap", heap.Items())
	fmt.Println("*********** MinHeap Size", heap.Size())
	fmt.Println("*********** MinHeap Peek", heap.Peek())
	fmt.Println("*********** MinHeap Pop", heap.Pop())
	fmt.Println("*********** MinHeap Size", heap.Size())
	fmt.Println("*********** MinHeap", heap.Items())
	heap.Push(4)
	fmt.Println("*********** MinHeap", heap.Items())
}

func MaxHeapTest() {
	items := []int{5,6,7,2,8,9,3,1}
	heap := NewMaxHeap(items...)
	fmt.Println("*********** MaxHeap", heap.Items())
	fmt.Println("*********** MaxHeap Size", heap.Size())
	fmt.Println("*********** MaxHeap Peek", heap.Peek())
	fmt.Println("*********** MaxHeap Pop", heap.Pop())
	fmt.Println("*********** MaxHeap Size", heap.Size())
	fmt.Println("*********** MaxHeap", heap.Items())
	heap.Push(4)
	fmt.Println("*********** MaxHeap", heap.Items())
}
