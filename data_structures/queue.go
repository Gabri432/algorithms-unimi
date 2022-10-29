package algorithmsunimi

type queue struct {
	list *linkedList
	head *listNode
	end  *listNode
}

// Verifies if a queue is empty
func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

// It adds a new element to the queue
func (q *queue) enqueue(elem int) {
	n := newNode(elem)
	if q.head == nil {
		q.head = n
	}
	q.end = n
}

// It removes an element from the queue
func (q *queue) dequeue() *listNode {
	first := q.head
	next := q.head.next
	if next == nil {
		q.end = nil
	}
	return first
}

// Returns the first element of a queue
func (q *queue) first() *listNode {
	return q.head
}
