package stack

// LIFO Last In First Out
type Stack struct {
	elements []int32
	rep      int32
}

func (s *Stack) Add(new int32) {
	s.elements = append(s.elements, new)
	s.rep++
}

func (s *Stack) Remove() int32 {
	element := s.elements[s.rep]
	s.elements = s.remove()
	s.rep--
	return element
}

func (s Stack) remove() []int32 {
	return s.elements[:s.rep]
}

func (s Stack) Peek() int32 {
	return s.elements[s.rep]
}

func New(length uint32) *Stack {
	return &Stack{make([]int32, 0, length), -1}
}
