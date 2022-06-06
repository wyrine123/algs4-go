package chapter4

import (
	"fmt"
	"io"
	"strconv"
)

type Graph struct {
	v   int            // 顶点数目(vertex)
	e   int            // 边的数据(edge)
	adj []BagInterface // 临接表
}

// NewGraph 初始化顶点数目为v的图
func NewGraph(v int) *Graph {
	g := new(Graph)
	g.v = v
	g.adj = make([]BagInterface, v, v)

	// 初始化
	for i := 0; i < v; i++ {
		g.adj[i] = new(BagLinkList)
	}

	return g
}

func NewGraphFromFile(in io.Reader) *Graph {
	var (
		v, e   int
		v1, v2 int
		//line string
		g   *Graph
		err error
	)
	// 第一行是顶点的数量
	_, err = fmt.Fscan(in, &v)
	if err != nil {
		panic("NewGraphFromFile 读取顶点数量错误: " + err.Error())
	}
	g = NewGraph(v)
	// 第二行是边的数量
	_, err = fmt.Fscan(in, &e)
	if err != nil {
		panic("NewGraphFromFile 读取边数量错误: " + err.Error())
	}

	for i := 0; i < e; i++ {
		_, err = fmt.Fscan(in, &v1)
		if err != nil {
			panic("NewGraphFromFile 第" + strconv.Itoa(i+1) + "行第一个数错误: " + err.Error())
		}
		_, err = fmt.Fscan(in, &v2)
		if err != nil {
			panic("NewGraphFromFile 第" + strconv.Itoa(i+1) + "行第二个数错误: " + err.Error())
		}
		g.AddEdge(v1, v2)
	}

	return g
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddEdge(v, w int) {
	g.e++

	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

func (g *Graph) Adj(v int) []int {
	data := make([]int, 0, g.adj[v].Size())
	for _, v := range g.adj[v].Iteration() {
		value, ok := v.(int)
		if !ok {
			panic("Graph 临接表里node val不是int类型")
		}
		data = append(data, value)
	}
	return data
}
func (g *Graph) ToString() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", g.V(), g.E())

	for v := 0; v < g.V(); v++ {
		s += fmt.Sprintf("%d: ", v)
		for _, w := range g.Adj(v) {
			s += fmt.Sprintf("%d ", w)
		}
		s += "\n"
	}
	return s
}
