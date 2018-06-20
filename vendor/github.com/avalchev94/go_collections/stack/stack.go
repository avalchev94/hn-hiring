package stack

// Stack is a LIFO data structure, i.e. Last In First Out
type Stack struct {
	top    *node
	length int
}

type node struct {
	prev  *node
	value interface{}
}

// New creates empty stack
func New() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
	}
}

// Top peeks the value of the last added value in the stack
func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}

	return s.top.value
}

// Len returns the length/size of the stack
func (s *Stack) Len() int {
	return s.length
}

// Push inserts a new element at the top of the stack
func (s *Stack) Push(value interface{}) {
	s.top = &node{s.top, value}
	s.length++
}

// Pop removes the top element of the stack and returns it's value
func (s *Stack) Pop() interface{} {
	var popped interface{}
	if s.length > 0 {
		popped = s.top.value
		s.top = s.top.prev
		s.length--
	}

	return popped
}

// Empty returns true if stack has no elements.
func (s *Stack) Empty() bool {
	return s.length <= 0
}
