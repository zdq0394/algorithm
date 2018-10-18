package queue

import "errors"

var (
	DequeEmptyError = errors.New("Deque is empty.")
)

type Deque interface {
	AddFront(e int)
	RemoveFront() error
	PeekFront() (int, error)
	AddBack(e int)
	RemoveBack() error
	PeekBack() (int, error)
	Length() int
	Empty() bool
}
