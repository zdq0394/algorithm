package queue

import (
	"testing"
)

func TestCircularQueue(t *testing.T) {
	obj := Constructor2(10)
	param_1 := obj.EnQueue(100)
	if !param_1 {
		t.Fail()
	}
	param_2 := obj.DeQueue()
	if !param_2 {
		t.Fail()
	}
	param_3 := obj.Front()
	if param_3 != -1 {
		t.Fail()
	}
	param_4 := obj.Rear()
	if param_4 != -1 {
		t.Fail()
	}
	param_5 := obj.IsEmpty()
	if !param_5 {
		t.Fail()
	}
	param_6 := obj.IsFull()
	if param_6 {
		t.Fail()
	}
}
