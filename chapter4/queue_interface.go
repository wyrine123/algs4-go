package chapter4

// QueueInterface 队列interface
type QueueInterface interface {
	Enqueue(v interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
	Iteration() []interface{}
}
