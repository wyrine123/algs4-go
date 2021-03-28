package chapter1

import "testing"

func TestNewLinkListQueue(t *testing.T) {
	queue := NewLinkListQueue()
	if queue.Size() != 0 {
		t.Error("TestNewLinkListQueue queue初始化size不是0")
	}
	if queue.first != nil {
		t.Error("TestNewLinkListQueue queue first不是nil")
	}
	if queue.last != nil {
		t.Error("TestNewLinkListQueue queue last不是nil")
	}

	queue.Enqueue("1")
	if queue.Size() != 1 {
		t.Error("TestNewLinkListQueue queue size不是1")
	}
	if queue.first != queue.last {
		t.Error("TestNewLinkListQueue queue first和last不相等")
	}
	queue.Enqueue("2")
	queue.Enqueue("3")
	if queue.n != 3 {
		t.Error("TestNewLinkListQueue queue size不是3")
	}
	if queue.Dequeue() != "1" {
		t.Error("TestNewLinkListQueue queue Dequeue不是1")
	}
	if queue.Dequeue() != "2" {
		t.Error("TestNewLinkListQueue queue Dequeue不是2")
	}
	if queue.Dequeue() != "3" {
		t.Error("TestNewLinkListQueue queue Dequeue不是3")
	}
	if queue.first != nil {
		t.Error("TestNewLinkListQueue queue first不是nil")
	}
	if queue.last != nil {
		t.Error("TestNewLinkListQueue queue last不是nil")
	}
}
