package chapter4

import "testing"

func newGraph() *Graph {
	g := NewGraph(3)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 0)

	return g
}

func TestGraphDegree(t *testing.T) {
	g := newGraph()

	if GraphDegree(g, 0) != 4 {
		t.Errorf("TestGraphDegree 节点0的度数不是4, 结果是:%d", GraphDegree(g, 0))
	}
	if GraphDegree(g, 1) != 1 {
		t.Errorf("TestGraphDegree 节点1的度数不是1, 结果是:%d", GraphDegree(g, 1))
	}
	if GraphDegree(g, 2) != 1 {
		t.Errorf("TestGraphDegree 节点2的度数不是1, 结果是:%d", GraphDegree(g, 2))
	}
}

func TestGraphMaxDegree(t *testing.T) {
	g := newGraph()

	if GraphMaxDegree(g) != 4 {
		t.Errorf("TestGraphMaxDegree 最大度数不是4, 结果是:%d", GraphMaxDegree(g))
	}
}

func TestGraphAvgDegree(t *testing.T) {
	g := newGraph()

	if GraphAvgDegree(g) != 2 {
		t.Errorf("TestGraphAvgDegree 平均度数不是2, 结果是:%.2f", GraphAvgDegree(g))
	}
}

func TestGraphNumberOfSelfLoops(t *testing.T) {
	g := newGraph()

	if GraphNumberOfSelfLoops(g) != 1 {
		t.Errorf("TestGraphNumberOfSelfLoops 自环个数不是1, 结果是:%d", GraphNumberOfSelfLoops(g))
	}
}

func TestGraph_ToString(t *testing.T) {
	g := newGraph()

	t.Log(g.ToString())
}
