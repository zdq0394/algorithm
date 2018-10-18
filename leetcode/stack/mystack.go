package stack

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
type dequeArrayImpl struct {
	Array []int
}

func NewDequeArray() Deque {
	return &dequeArrayImpl{
		Array: []int{},
	}
}

var _ Deque = &dequeArrayImpl{}

func (q *dequeArrayImpl) AddFront(e int) {
	q.Array = append([]int{e}, q.Array...)
}
func (q *dequeArrayImpl) RemoveFront() error {
	if q.Length() == 0 {
		return DequeEmptyError
	}
	q.Array = q.Array[1:]
	return nil
}
func (q *dequeArrayImpl) PeekFront() (int, error) {
	if q.Length() == 0 {
		return 0, DequeEmptyError
	}
	return q.Array[0], nil
}
func (q *dequeArrayImpl) AddBack(e int) {
	q.Array = append(q.Array, e)
}
func (q *dequeArrayImpl) RemoveBack() error {
	if q.Length() == 0 {
		return DequeEmptyError
	}
	q.Array = q.Array[:q.Length()-1]
	return nil
}
func (q *dequeArrayImpl) PeekBack() (int, error) {
	if q.Length() == 0 {
		return 0, DequeEmptyError
	}
	return q.Array[q.Length()-1], nil
}
func (q *dequeArrayImpl) Length() int {
	return len(q.Array)
}
func (q *dequeArrayImpl) Empty() bool {
	return q.Length() == 0
}

type MyStack struct {
	dq Deque
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		dq: NewDequeArray(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.dq.AddBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	e, _ := this.dq.PeekBack()
	this.dq.RemoveBack()
	return e
}

/** Get the top element. */
func (this *MyStack) Top() int {
	e, _ := this.dq.PeekBack()
	return e
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.dq.Empty()
}
