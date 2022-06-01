package chapter4

// GraphInterface 图 interface定义
type GraphInterface interface {
	V() int
	E() int
	AddEdge(v, w int)
	Adj(v int) []int
	ToString() string
}
