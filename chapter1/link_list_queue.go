package chapter1

/**
  使用链表实现队列
*/

type LinkListQueue struct {
	first *node // 队头 保存最早进队列的数据
	last  *node // 队尾 保存最后进队列的数据
	n     int   // 队列长度
}

func NewLinkListQueue() *LinkListQueue {
	llq := new(LinkListQueue)
	llq.first = nil
	llq.last = nil
	llq.n = 0

	return llq
}

// 入队
func (p *LinkListQueue) Enqueue(s string) {
	node := new(node)
	node.item = s
	if p.isEmpty() {
		// 队列为空,首尾节点都指向该节点
		p.first = node
		p.last = node
	} else {
		// 插入队尾
		p.last.next = node
		p.last = node
	}
	p.n++
}

// 出队
func (p *LinkListQueue) Dequeue() string {
	if p.isEmpty() {
		return "nil"
	}

	// 队首出列
	item := p.first.item
	p.first = p.first.next
	p.n--
	if p.isEmpty() {
		// 出队之后 队列为空的话，last节点也要置为空
		p.last = nil
	}
	return item
}

// 队列长度
func (p *LinkListQueue) Size() int {
	return p.n
}

func (p *LinkListQueue) isEmpty() bool {
	return p.n == 0
}
