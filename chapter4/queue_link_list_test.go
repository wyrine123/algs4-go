package chapter4

import "testing"

func TestQueueLinkList(t *testing.T) {
	q := NewQueueLinkList()

	if !q.IsEmpty() {
		t.Error("TestQueueLinkList 队列初始不为空")
	}
	t.Logf("TestQueueLinkList 队列初始: %v", q.Iteration())

	q.Enqueue(1)

	if q.Size() != 1 {
		t.Error("TestQueueLinkList 队列长度不为1")
	}
	t.Logf("TestQueueLinkList Enqueue 1后: %v", q.Iteration())

	v := q.Dequeue()
	if !q.IsEmpty() {
		t.Error("TestQueueLinkList 队列DeQueue后不为空")
	}
	if v.(int) != 1 {
		t.Error("TestQueueLinkList 队列DeQueue值不为1")
	}
	t.Logf("TestQueueLinkList Enqueue 1后 Dequeue: %v", q.Iteration())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Error("TestQueueLinkList 队列长度不为3")
	}
	t.Logf("TestQueueLinkList Enqueue 1, 2, 3后: %v", q.Iteration())

	q.Dequeue()

	if q.Size() != 2 {
		t.Error("TestQueueLinkList 队列长度不为2")
	}
	t.Logf("TestQueueLinkList Enqueue 1, 2, 3, Dequeue后: %v", q.Iteration())
}
