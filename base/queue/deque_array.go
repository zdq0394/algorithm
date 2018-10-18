package queue

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
