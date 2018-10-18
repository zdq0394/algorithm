package queue

type MyCircularQueue struct {
	HeadIndex int
	TailIndex int
	Size      int
	Capacity  int
	Data      []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor2(k int) MyCircularQueue {
	return MyCircularQueue{
		HeadIndex: 0,
		TailIndex: 0,
		Size:      0,
		Capacity:  k,
		Data:      make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Data[this.TailIndex] = value
	this.TailIndex = (this.TailIndex + 1) % this.Capacity
	this.Size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.HeadIndex = (this.HeadIndex + 1) % this.Capacity
	this.Size--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[this.HeadIndex]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[this.TailIndex-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size == 0

}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Size == this.Capacity

}
