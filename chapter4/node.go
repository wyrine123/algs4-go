package chapter4

type Node struct {
	val  interface{} // 值
	prev *Node       // 前一个节点
	next *Node       // 后一个节点
}

func NewNode() *Node {
	return new(Node)
}
