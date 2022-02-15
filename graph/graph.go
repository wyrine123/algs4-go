package graph

import (
	"algs4-go/link_list"
	"fmt"
)

type Graph struct {
	v   int                     // 顶点数目
	e   int                     // 边的数目
	adj map[int]*link_list.Node // 临接表
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.adj = make(map[int]*link_list.Node)

	return graph
}

func (g *Graph) AddEdge(v int, w int) {
	g.e++ // 边数量+1
	wNode := new(link_list.Node)
	wNode.Val = w
	if g.adj[v] == nil {
		g.adj[v] = wNode
		g.v++
	} else {
		// 这儿要使用头插
		wNode.Next = g.adj[v]
		g.adj[v] = wNode
	}

	vNode := new(link_list.Node)
	vNode.Val = v
	if g.adj[w] == nil {
		g.adj[w] = vNode
		g.v++
	} else {
		// 这儿要使用头插
		vNode.Next = g.adj[w]
		g.adj[w] = vNode
	}
}

func (g *Graph) GetV() int {
	return g.v
}

func (g *Graph) GetE() int {
	return g.e
}

func (g *Graph) GetAdj(v int) *link_list.Node {
	return g.adj[v]
}

func (g *Graph) ToString() string {
	var s string
	s += fmt.Sprintf("%d vertices, %d edges\n", g.GetV(), g.GetE())
	// 注意这儿的map是没有顺序的
	for k, v := range g.adj {
		s += fmt.Sprintf("%d: ", k)
		for node := v; node != nil; node = node.Next {
			s += fmt.Sprintf("%d ", node.Val)
		}
		s += "\n"
	}
	return s
}

// GetGraphDegree 获取图节点v的度数
func GetGraphDegree(graph Graph, v int) int {
	var degree int
	root := graph.GetAdj(v)
	for node := root; node != nil; node = node.Next {
		degree++
	}
	return degree
}

// GetGraphMaxDegree 计算所有顶点的最大度数
func GetGraphMaxDegree(graph Graph) int {
	var max int
	for k := range graph.adj {
		kMax := GetGraphDegree(graph, k)
		if kMax > max {
			max = kMax
		}
	}
	return max
}

// GetGraphAvgDegree 获取所有顶点的平均度数
func GetGraphAvgDegree(graph Graph) float64 {
	var sum int
	for k := range graph.adj {
		sum += GetGraphDegree(graph, k)
	}
	return float64(sum) / float64(graph.GetV())

	// todo 更简单方法
	return float64(graph.GetE() * 2.0 / graph.GetV())
}

// GetGraphNumOfSelfLoops 计算自环的个数
// 自环是自己指向自己
func GetGraphNumOfSelfLoops(graph Graph) int {
	var count int
	for k := range graph.adj {
		root := graph.GetAdj(k)
		for node := root; node != nil; node = node.Next {
			if node.Val == k {
				count++
			}
		}
	}
	return count
}
