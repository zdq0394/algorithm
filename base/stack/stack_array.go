package stack

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
