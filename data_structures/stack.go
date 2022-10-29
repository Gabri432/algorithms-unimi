package algorithmsunimi

type stack struct {
	list *linkedList
}

// Verifies if the stack is empty
func (s *stack) isEmpty() bool {
	if s.list.head == nil {
		return true
	}
	return false
}

// Add a new element to the stack
func (s *stack) push(elem int) {
	s.list.push(elem)
}

// Removes an element from the stack
func (s *stack) pop() *listNode {
	p := s.list.head
	var previous *listNode
	if p != nil {
		previous = s.list.head
		s.list.head = p.next
	}
	return previous
}

// Returns the element at the top of the stack
func (s *stack) top() *listNode {
	return s.list.head
}
