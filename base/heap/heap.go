package heap

import (
	"errors"
)

var (
	HeapEmptyError = errors.New("Heap is empty.")
)

type Heap interface {
	AddElement(e int)
	RemoveElement() error
	TopElement() (int, error)
	Print()
}
