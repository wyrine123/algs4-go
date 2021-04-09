package chapter2

import (
	"fmt"
	"testing"
)

func TestMaxPQ(t *testing.T) {
	maxPQ := NewMaxPQ()
	fmt.Println(maxPQ.String())

	maxPQ.Insert(1)
	maxPQ.Insert(2)
	maxPQ.Insert(3)
	maxPQ.Insert(4)
	fmt.Println(maxPQ.String())
	if maxPQ.Size() != 4 {
		t.Error("maxPQ 长度不是4")
	}
	if maxPQ.DelMax() != 4 {
		t.Error("maxPQ  DelMax不是4")
	}
	if maxPQ.Size() != 3 {
		t.Error("maxPQ 长度不是3")
	}
	if maxPQ.DelMax() != 3 {
		t.Error("maxPQ  DelMax不是3")
	}
	if maxPQ.DelMax() != 2 {
		t.Error("maxPQ  DelMax不是2")
	}
	if maxPQ.DelMax() != 1 {
		t.Error("maxPQ  DelMax不是1")
	}
}
