package singly

import (
	"fmt"
)

func SinglyLinkedListTest() {
	sll := &SinglyLinkedList{}
	var key, val int
	val = 1
	fmt.Println("*********** AddHead", val)
	sll.AddHead(val)
	val = 2
	fmt.Println("*********** AddHead", val)
	sll.AddHead(val)
	val = 3
	fmt.Println("*********** AddHead", val)
	sll.AddHead(val)
	val = 4
	fmt.Println("*********** AddHead", val)
	sll.AddHead(val)
	sll.Print()
	val = 5
	fmt.Println("*********** AddTail", val)
	sll.AddTail(val)
	val = 6
	fmt.Println("*********** AddTail", val)
	sll.AddTail(val)
	val = 7
	fmt.Println("*********** AddTail", val)
	sll.AddTail(val)
	val = 8
	fmt.Println("*********** AddTail", val)
	sll.AddTail(val)
	sll.Print()
	key = 4
	val = 9
	fmt.Println("*********** Add", val, "before", key)
	sll.AddBefore(key, val)
	sll.Print()
	key = 5
	val = 10
	fmt.Println("*********** Add", val, "before", key)
	sll.AddBefore(key, val)
	sll.Print()
	key = 8
	val = 11
	fmt.Println("*********** Add", val, "before", key)
	sll.AddBefore(key, val)
	sll.Print()
	key = 9
	val = 12
	fmt.Println("*********** Add", val, "after", key)
	sll.AddAfter(key, val)
	sll.Print()
	key = 5
	val = 13
	fmt.Println("*********** Add", val, "after", key)
	sll.AddAfter(key, val)
	sll.Print()
	key = 8
	val = 14
	fmt.Println("*********** Add", val, "after", key)
	sll.AddAfter(key, val)
	sll.Print()
	fmt.Println("*********** DeleteHead")
	sll.DeleteHead()
	sll.Print()
	fmt.Println("*********** DeleteTail")
	sll.DeleteTail()
	sll.Print()
	key = 12
	fmt.Println("*********** DeleteKey", key)
	sll.DeleteKey(key)
	sll.Print()
	key = 5
	fmt.Println("*********** DeleteKey", key)
	sll.DeleteKey(key)
	sll.Print()
	key = 8
	fmt.Println("*********** DeleteKey", key)
	sll.DeleteKey(key)
	sll.Print()
	key = 4
	fmt.Println("*********** Contains", key, sll.Contains(key))
	key = 10
	fmt.Println("*********** Contains", key, sll.Contains(key))
	key = 11
	fmt.Println("*********** Contains", key, sll.Contains(key))
	key = 100
	fmt.Println("*********** Contains", key, sll.Contains(key))
}
