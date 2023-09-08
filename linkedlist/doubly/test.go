package doubly

import "fmt"

func DoublyLinkedListTest()  {
	dll := &DoublyLinkedList{}
	var key, val int
	val = 1
	fmt.Println("*********** AddHead", val)
	dll.AddHead(val)
	val = 2
	fmt.Println("*********** AddHead", val)
	dll.AddHead(val)
	val = 3
	fmt.Println("*********** AddHead", val)
	dll.AddHead(val)
	val = 4
	fmt.Println("*********** AddHead", val)
	dll.AddHead(val)
	dll.Print()
	val = 5
	fmt.Println("*********** AddTail", val)
	dll.AddTail(val)
	val = 6
	fmt.Println("*********** AddTail", val)
	dll.AddTail(val)
	val = 7
	fmt.Println("*********** AddTail", val)
	dll.AddTail(val)
	val = 8
	fmt.Println("*********** AddTail", val)
	dll.AddTail(val)
	dll.Print()
	key = 4
	val = 9
	fmt.Println("*********** Add", val, "before", key)
	dll.AddBefore(key, val)
	dll.Print()
	key = 5
	val = 10
	fmt.Println("*********** Add", val, "before", key)
	dll.AddBefore(key, val)
	dll.Print()
	key = 8
	val = 11
	fmt.Println("*********** Add", val, "before", key)
	dll.AddBefore(key, val)
	dll.Print()
	key = 9
	val = 12
	fmt.Println("*********** Add", val, "after", key)
	dll.AddAfter(key, val)
	dll.Print()
	key = 5
	val = 13
	fmt.Println("*********** Add", val, "after", key)
	dll.AddAfter(key, val)
	dll.Print()
	key = 8
	val = 14
	fmt.Println("*********** Add", val, "after", key)
	dll.AddAfter(key, val)
	dll.Print()
	fmt.Println("*********** DeleteHead")
	dll.DeleteHead()
	dll.Print()
	fmt.Println("*********** DeleteTail")
	dll.DeleteTail()
	dll.Print()
	key = 12
	fmt.Println("*********** DeleteKey", key)
	dll.DeleteKey(key)
	dll.Print()
	key = 5
	fmt.Println("*********** DeleteKey", key)
	dll.DeleteKey(key)
	dll.Print()
	key = 8
	fmt.Println("*********** DeleteKey", key)
	dll.DeleteKey(key)
	dll.Print()
	key = 4
	fmt.Println("*********** Find", key)
	fmt.Println(key, "found:",  dll.Find(key))
	key = 10
	fmt.Println("*********** Find", key)
	fmt.Println(key, "found:",  dll.Find(key))
	key = 11
	fmt.Println("*********** Find", key)
	fmt.Println(key, "found:",  dll.Find(key))
	key = 100
	fmt.Println("*********** Find", key)
	fmt.Println(key, "found:",  dll.Find(key))
}
