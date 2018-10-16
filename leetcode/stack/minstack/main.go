package main

import (
	"sync"
)

type MinStack struct {
	value    []int
	minValue []int
	lock     sync.Mutex
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		value:    make([]int, 0),
		minValue: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = append(this.value, x)
	if len(this.minValue) == 0 {
		this.minValue = append(this.minValue, x)
	} else {
		currentValue := this.minValue[len(this.minValue)-1]
		if x < currentValue {
			this.minValue = append(this.minValue, x)
		} else {
			this.minValue = append(this.minValue, currentValue)
		}
	}
}

func (this *MinStack) Pop() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if len(this.value) == 0 {
		panic("Empty stack.")
	}
	this.value = this.value[:len(this.value)-1]
	this.minValue = this.minValue[:len(this.minValue)-1]

}

func (this *MinStack) Top() int {
	return this.value[len(this.value)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValue[len(this.minValue)-1]
}

func main() {

}
