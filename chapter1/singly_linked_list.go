package chapter1

import "fmt"

// SinglyLinkedList 单向链表
type SinglyLinkedList struct {
	first *Node
	last  *Node
	n     int // 链表长度
}

func NewSinglyLinkedList() *SinglyLinkedList {
	linklist := new(SinglyLinkedList)
	return linklist
}

// Insert 尾插入链表
func (p *SinglyLinkedList) Insert(item string) {
	node := Node{item: item}
	if p.isEmpty() {
		p.first = &node
		p.last = &node
		p.n++
		return
	}
	oldLast := p.last
	oldLast.next = &node
	p.last = &node
	p.n++
}

// DelLastNode 删除尾节点
// 找到倒数第二个节点,然后item.next=nil即可
func (p *SinglyLinkedList) DelLastNode() *Node {
	if p.isEmpty() {
		return nil
	}
	delNode := p.last
	if p.Size() == 1 {
		p.first = nil
		p.last = nil
		p.n--
		return delNode
	}

	// 寻找倒数第二个节点
	item := p.first
	for ; item.next.next != nil; item = item.next {
	}
	item.next = nil
	p.last = item
	p.n--
	return delNode
}

// DelFirstNode 删除首节点
func (p *SinglyLinkedList) DelFirstNode() *Node {
	if p.isEmpty() {
		return nil
	}
	delNode := p.first
	if p.Size() == 1 {
		p.first = nil
		p.last = nil
		p.n--
		return delNode
	}

	p.first = p.first.next
	p.n--
	return delNode
}

func (p *SinglyLinkedList) Del(k int) *Node {
	if k <= 0 {
		// k必须大于0
		return nil
	}
	// 删除的节点超过链表长度
	if k > p.n {
		return nil
	}
	// 如果删除的节点是首尾节点的话,需要处理下first和last.
	// 其他的场景只需要删除特定的节点(k-1).next=k.next
	var delNode *Node
	if k == p.n {
		// 删除的是尾节点
		return p.DelLastNode()
	}
	if k == 1 {
		// 删除的是首节点
		return p.DelFirstNode()
	}
	// 删除的是中间节点
	t := p.first
	for i := 1; i < k-1; i++ {
		t = t.next
	}
	delNode = t.next
	t.next = t.next.next
	p.n--
	return delNode
}

// Find 查询item是否在链表内
func (p *SinglyLinkedList) Find(item string) bool {
	for t := p.first; t != nil; t = t.next {
		if t.item == item {
			return true
		}
	}
	return false
}

// 链表是否为空
func (p *SinglyLinkedList) isEmpty() bool {
	return p.n == 0
}

// Size 查询链表长度
func (p *SinglyLinkedList) Size() int {
	return p.n
}

// String 输出可视化链表
func (p *SinglyLinkedList) String() string {
	s := ""
	for item := p.first; item != nil; item = item.next {
		s += fmt.Sprintf("%s\t", item.item)
	}
	return s
}
