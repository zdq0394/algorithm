package queue

import (
	"errors"
)

var (
	CircularQueueEmptyError = errors.New("Circular queue is empty.")
	CircularQueueFullError  = errors.New("Circular queue is full.")
)

type CircularQueue interface {
	/** Insert an element into the circular queue. Return true if the operation is successful. */
	EnQueue(value int) bool
	/** Delete an element from the circular queue. Return true if the operation is successful. */
	DeQueue() bool
	/** Get the front item from the queue. */
	Front() int
	/** Get the last item from the queue. */
	Rear() int
	/** Checks whether the circular queue is empty or not. */
	IsEmpty() bool
	/** Checks whether the circular queue is full or not. */
	IsFull() bool
}
