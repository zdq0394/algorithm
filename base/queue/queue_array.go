package queue

type QueueArrayImpl struct {
	Array []int
}

var _ Queue = &QueueArrayImpl{}

func NewQueue() Queue {
	return &QueueArrayImpl{
		Array: []int{},
	}
}

func (q *QueueArrayImpl) Add(e int) {
	q.Array = append(q.Array, e)
}

func (q *QueueArrayImpl) Remove() error {
	if len(q.Array) == 0 {
		return QueueEmptyError
	}
	if len(q.Array) == 1 {
		q.Array = []int{}
		return nil
	}
	q.Array = q.Array[1:]
	return nil
}

func (q *QueueArrayImpl) Length() int {
	return len(q.Array)
}

func (q *QueueArrayImpl) Peek() (int, error) {
	if len(q.Array) == 0 {
		return 0, QueueEmptyError
	}
	return q.Array[0], nil
}

func (q *QueueArrayImpl) Empty() bool {
	return q.Length() == 0
}
