package stack

import (
	"errors"
)

var (
	StackEmptyError = errors.New("Stack is empty.")
)

type Stack interface {
	Push(e int)
	Pop() (int, error)
	Peek() (int, error)
	Length() int
}
