package chapter4

type BagLinkList struct {
	v    int   // 背包容量
	root *Node // 根结点
	tail *Node // 尾节点
}

func NewBagLinkList() *BagLinkList {
	return new(BagLinkList)
}

func (bag *BagLinkList) Add(item interface{}) {
	node := NewNode()
	node.val = item
	if bag.IsEmpty() {
		bag.root = node
		bag.tail = node
	} else {
		bag.tail.next = node
		node.prev = bag.tail
		bag.tail = node
	}
	bag.v++
}

func (bag *BagLinkList) IsEmpty() bool {
	return bag.v == 0
}

func (bag *BagLinkList) Size() int {
	return bag.v
}

func (bag *BagLinkList) Iteration() []interface{} {
	data := make([]interface{}, 0, bag.Size())
	for node := bag.root; node != nil; node = node.next {
		data = append(data, node.val)
	}
	return data
}
