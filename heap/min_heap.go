package heap

// min heap implementation
type MinHeap struct {
	items []int
}

// instantiate a heap 
func NewMinHeap(items ...int) *MinHeap {
	heap := &MinHeap{
		items: items,
	}
	// heapify
	for i:=(len(heap.items)/2)-1;i>=0;i-- {
		heap.Heapify(i)
	}
	return heap
}

// get size of heap
func (h *MinHeap) Size() int {
	return len(h.items)
}

// get if heap is empty
func (h *MinHeap) IsEmpty() bool {
	return h.Size()==0
}

// heapify
func (h *MinHeap) Heapify(index int) {
	// find smallest from index and its children
	smallest := index
	left := 2*index + 1
	right := 2*index + 2
	if left<len(h.items) && h.items[left]<h.items[smallest] {
		smallest = left
	}
	if right<len(h.items) && h.items[right]<h.items[smallest] {
		smallest = right
	}
	if smallest!=index {
		// swap smallest and index
		h.items[smallest], h.items[index] = h.items[index], h.items[smallest]
		h.Heapify(smallest)
	}
}

func (h *MinHeap) BubbleUp(index int) {
	for index >= 0 {
		parent := (index-1)/2
		if h.items[parent]>h.items[index] {
			h.items[parent], h.items[index] = h.items[index], h.items[parent]
		}else {
			break
		}
		index = parent
	}
}

// get min element from heap
func (h *MinHeap) Peek() int {
	return h.items[0]
}

// insert an element into heap
func (h *MinHeap) Push(val int) {
	h.items = append(h.items, val)
	size := h.Size()
	h.BubbleUp(size-1)
}

// get and delete min element from heap
func (h *MinHeap) Pop() int {
	size := h.Size()
	// store first element to return
	min_value := h.items[0]

	// swap first and last element
	h.items[0], h.items[size-1] = h.items[size-1], h.items[0]

	// remove last element
	h.items = h.items[:size-1]

	// heapify the first element
	h.Heapify(0)
	return min_value
}

// returns items of heap as list
func (h *MinHeap) Items() []int {
	return h.items
}
