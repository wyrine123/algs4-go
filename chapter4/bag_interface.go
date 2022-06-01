package chapter4

// BagInterface 背包interface
type BagInterface interface {
	Add(item interface{})
	IsEmpty() bool
	Size() int
	Iteration() []interface{}
}
