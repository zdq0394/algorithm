package queue

import (
	"errors"
)

var (
	QueueEmptyError = errors.New("Queue is empty.")
)

type Queue interface {
	Add(e int)
	Remove() error
	Length() int
	Peek() (int, error)
	Empty() bool
}
