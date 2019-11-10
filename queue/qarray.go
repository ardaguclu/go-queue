package queue

// QArray is slice based standard queue implementation without thread safety for concurrent access.
type QArray struct {
	Items []interface{}
}

// NewQArray initializes and returns new QArray.
func NewQArray() *QArray {
	return &QArray{
		Items: make([]interface{}, 0, 1024),
	}
}

// Push adds new item tail of the queue.
func (q *QArray) Push(n interface{}) {
	q.Items = append(q.Items, n)
}

// Pop returns first item from the queue.
func (q *QArray) Pop() (n interface{}) {
	if q.Len() > 0 {
		n = q.Items[0]
		q.Items[0] = nil
		q.Items = q.Items[1:]
	}

	return
}

// Len returns the item length of queue.
func (q *QArray) Len() int {
	return len(q.Items)
}
