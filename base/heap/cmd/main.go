package main

import (
	"github.com/zdq0394/algorithm/base/heap"
)

func tmin() {
	h := heap.NewEmptyMinHeap()
	h.AddElement(4)
	h.AddElement(5)
	h.AddElement(6)
	h.AddElement(1)
	h.AddElement(2)
	h.AddElement(3)
	h.Print()
	h.RemoveElement()
	h.Print()
	h.RemoveElement()
	h.Print()
}

func tmax() {
	h := heap.NewEmptyMaxHeap()
	h.AddElement(4)
	h.AddElement(5)
	h.AddElement(6)
	h.AddElement(1)
	h.AddElement(2)
	h.AddElement(3)
	h.Print()
	h.RemoveElement()
	h.Print()
	h.RemoveElement()
	h.Print()
}

func main() {
	tmin()
	tmax()
}
