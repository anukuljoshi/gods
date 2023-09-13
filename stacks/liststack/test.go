package liststack

import "fmt"

func ListStackTest() {
	var s = New()
	var val int
	fmt.Println("******* IsEmpty", s.IsEmpty())
	for i:=1;i<=5;i++ {
		val = i
		fmt.Println("******* Push", val)
		s.Push(val)
	}
	s.Print()
	fmt.Println("******* IsEmpty", s.IsEmpty())
	fmt.Println("******* Size", s.Size())
	top, err := s.Top()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("******* Top", top)
	pop, err := s.Pop()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("******* Pop", pop)
	fmt.Println("******* Size", s.Size())
	for !s.IsEmpty() {
		pop, err := s.Pop()
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println("******* Pop", pop)
	}
	pop, err = s.Pop()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("******* Pop", pop)
	fmt.Println("******* Size", s.Size())
}
