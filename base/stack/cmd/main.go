package main

import (
	"fmt"

	"github.com/zdq0394/algorithm/base/stack"
)

func main() {
	s := stack.NewStack()
	s.Push(2)
	s.Push(3)
	s.Push(4)
	var v int
	var e error
	v, e = s.Pop()
	fmt.Println(v, e)
	v, e = s.Pop()
	fmt.Println(v, e)
	fmt.Println(s.Length())
	v, e = s.Pop()
	fmt.Println(v, e)
	fmt.Println(s.Length())
	v, e = s.Pop()
	fmt.Println(v, e)
}
