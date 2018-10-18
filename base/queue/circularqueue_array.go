package queue

type CircularQueueArrayImpl struct {
	HeadIndex int
	TailIndex int
	Size      int
	Capacity  int
	Data      []int
}

func NewCircularQueue(capacity int) CircularQueue {
	return &CircularQueueArrayImpl{
		HeadIndex: 0,
		TailIndex: 0,
		Size:      0,
		Capacity:  capacity,
		Data:      make([]int, capacity),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (q *CircularQueueArrayImpl) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.Data[q.TailIndex] = value
	q.TailIndex = (q.TailIndex + 1) % q.Capacity
	q.Size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (q *CircularQueueArrayImpl) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.HeadIndex = (q.HeadIndex + 1) % q.Capacity
	q.Size--
	return true

}

/** Get the front item from the queue. */
func (q *CircularQueueArrayImpl) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Data[q.HeadIndex]
}

/** Get the last item from the queue. */
func (q *CircularQueueArrayImpl) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Data[q.TailIndex-1]
}

/** Checks whether the circular queue is empty or not. */
func (q *CircularQueueArrayImpl) IsEmpty() bool {
	return q.Size == 0
}

/** Checks whether the circular queue is full or not. */
func (q *CircularQueueArrayImpl) IsFull() bool {
	return q.Size == q.Capacity
}
