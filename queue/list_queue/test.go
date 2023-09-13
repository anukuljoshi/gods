package list_queue

import "fmt"

func ListQueueTest()  {
	var s = New()
	var val int
	fmt.Println("******* IsEmpty", s.IsEmpty())
	for i:=1;i<=5;i++ {
		val = i
		fmt.Println("******* Enqueue", val)
		s.Enqueue(val)
	}
	s.Print()
	fmt.Println("******* IsEmpty", s.IsEmpty())
	fmt.Println("******* Size", s.Size())
	top, err := s.Peek()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("******* Peek", top)
	pop, err := s.Dequeue()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("******* Dequeue", pop)
	fmt.Println("******* Size", s.Size())
	s.Print()
	for !s.IsEmpty() {
		pop, err := s.Dequeue()
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println("******* Dequeue", pop)
	}
	pop, err = s.Dequeue()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("******* Dequeue", pop)
	fmt.Println("******* Size", s.Size())
}
