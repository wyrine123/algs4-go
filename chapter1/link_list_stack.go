package chapter1

/**
  使用链表实现stack
	链表是一种递归的数据结构
      它为空(null),或者是指向一个节点(node)的引用,该节点包括一个元素和一个指向另一条链表的引用
*/

type LinkListStack struct {
	first *node // 首节点 这儿要使用指针,这样就可以判断==nil
	n     int   // 栈长度
}

// 节点
type node struct {
	item string // 节点值
	next *node  // next node
}

func NewLinkListStack() *LinkListStack {
	lls := new(LinkListStack)
	lls.first = nil
	lls.n = 0

	return lls
}

// push
func (p *LinkListStack) Push(s string) {
	oldFirst := p.first
	node := new(node)
	node.item = s
	node.next = oldFirst
	p.first = node
	p.n++
}

// pop
func (p *LinkListStack) Pop() string {
	if p.isEmpty() {
		return "nil"
	}
	item := p.first.item
	p.first = p.first.next
	p.n--
	return item
}

// size
func (p *LinkListStack) Size() int {
	return p.n
}

// 是否为空
func (p *LinkListStack) isEmpty() bool {
	return p.first == nil
}
