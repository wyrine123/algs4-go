package chapter4

import (
	"os"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	f, err := os.Open("./data/tinyG.txt")
	if err != nil {
		t.Errorf("TestBreadthFirstSearch open file 失败: %s", err.Error())
	}
	g := NewGraphFromFile(f)

	s := NewBreadthFirstSearch(g, 0)
	if s.Count() != 7 {
		t.Errorf("TestBreadthFirstSearch 与节点0相连的节点数不是7, actual: %d", s.Count())
	}
	if !s.Marked(3) {
		t.Error("TestBreadthFirstSearch 节点0与节点3不相连")
	}
	if s.Marked(11) {
		t.Error("TestBreadthFirstSearch 节点0与节点11相连")
	}
}
