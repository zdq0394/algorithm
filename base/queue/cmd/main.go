package main

import (
	"fmt"

	"github.com/zdq0394/algorithm/base/queue"
)

func main() {
	q := queue.NewQueue()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	var v int
	var e error
	v, e = q.Peek()
	fmt.Println(v, e)
	q.Remove()
	v, e = q.Peek()
	fmt.Println(v, e)
	fmt.Println("Length:", q.Length())
}
