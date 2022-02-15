package graph

import "testing"

func genGraph() *Graph {
	graph := NewGraph()
	graph.AddEdge(0, 5)
	graph.AddEdge(4, 3)
	graph.AddEdge(0, 1)
	graph.AddEdge(9, 12)
	graph.AddEdge(6, 4)
	graph.AddEdge(5, 4)
	graph.AddEdge(0, 2)
	graph.AddEdge(11, 12)
	graph.AddEdge(9, 10)
	graph.AddEdge(0, 6)
	graph.AddEdge(7, 8)
	graph.AddEdge(9, 11)
	graph.AddEdge(5, 3)

	return graph
}

func Test_Graph(t *testing.T) {
	graph := genGraph()
	if graph.GetV() != 13 {
		t.Errorf("Test_Graph graph.GetV() 应该是13，实际是：%d", graph.GetV())
	}
	if graph.GetE() != 13 {
		t.Errorf("Test_Graph graph.GetE() 应该是13，实际是：%d", graph.GetE())
	}

	t.Logf("Test_Graph字符串表示：%s", graph.ToString())
}
