package queue

type Queue struct {
	elements []int32
}

// Add new element in the Queue
func (q *Queue) Enqueue(element int32) {
	q.elements = append(q.elements, element)
}

// Remove and return the first element from Queue
func (q *Queue) Dequeue() int32 {
	value := q.elements[0]
	q.elements = q.remove()
	return value
}
func (q Queue) remove() []int32 {
	new := q.elements[1:]
	return new
}

// Return the first element
func (q Queue) Peek() int32 {
	return q.elements[0]
}

// Create a queue
func New(length uint16) *Queue {
	return &Queue{make([]int32, 0, length)}
}
