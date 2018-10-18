package queue

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
	Empty() bool
}
type stackArrayImpl struct {
	Array            []int
	LastElementIndex int //-1 means no elements in stack
}

func NewStack() Stack {
	return &stackArrayImpl{
		Array:            []int{},
		LastElementIndex: -1,
	}
}

func (s *stackArrayImpl) Push(e int) {
	s.Array = append(s.Array, e)
	s.LastElementIndex++
}

func (s *stackArrayImpl) Pop() (int, error) {
	if s.LastElementIndex == -1 {
		return 0, StackEmptyError
	}
	v := s.Array[s.LastElementIndex]
	s.Array = s.Array[:s.LastElementIndex]
	s.LastElementIndex--
	return v, nil
}

func (s *stackArrayImpl) Peek() (int, error) {
	if s.LastElementIndex == -1 {
		return 0, StackEmptyError
	}
	return s.Array[s.LastElementIndex], nil
}

func (s *stackArrayImpl) Length() int {
	return s.LastElementIndex + 1
}

func (s *stackArrayImpl) Empty() bool {
	return s.Length() == 0
}

type MyQueue struct {
	inStack  Stack
	outStack Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack:  NewStack(),
		outStack: NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	if this.outStack.Empty() {
		for !this.inStack.Empty() {
			e, _ := this.inStack.Pop()
			this.outStack.Push(e)
		}
	}
	e, _ := this.outStack.Pop()
	return e
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.outStack.Empty() {
		for !this.inStack.Empty() {
			e, _ := this.inStack.Pop()
			this.outStack.Push(e)
		}
	}
	e, _ := this.outStack.Peek()
	return e
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.inStack.Empty() && this.outStack.Empty()
}
