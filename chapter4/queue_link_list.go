package chapter4

type QueueLinkList struct {
	count int   // 队列长度
	root  *Node // 队首
	tail  *Node // 队尾
}

func NewQueueLinkList() *QueueLinkList {
	q := new(QueueLinkList)

	return q
}

func (q *QueueLinkList) Enqueue(v interface{}) {
	node := NewNode()
	node.val = v
	if q.IsEmpty() {
		// 队列为空
		q.root = node
		q.tail = node
	} else {
		node.prev = q.tail
		q.tail.next = node
		q.tail = node
	}
	q.count++
}

func (q *QueueLinkList) Dequeue() interface{} {
	var v interface{}
	if q.IsEmpty() {
		v = nil
	} else if q.Size() == 1 {
		// 需要处理root和tail
		v = q.tail.val
		q.tail = nil
		q.root = nil
	} else {
		// 只需要处理tail
		v = q.tail.val
		q.tail = q.tail.prev
		q.tail.next = nil
	}
	q.count--
	return v
}

func (q *QueueLinkList) IsEmpty() bool {
	return q.count == 0
}

func (q *QueueLinkList) Size() int {
	return q.count
}

func (q *QueueLinkList) Iteration() []interface{} {
	var res = make([]interface{}, 0, q.Size())
	for v := q.root; v != nil; v = v.next {
		res = append(res, v.val)
	}
	return res
}
