package graph

type SearchInterface interface {
	Marked(v int) bool
	Count() int
}

type DepthFirstSearch struct {
	graph  Graph
	marked []bool
	count  int
}

func NewDepthFirstSearch(g Graph, s int) *DepthFirstSearch {
	d := new(DepthFirstSearch)
	d.graph = g
	d.marked = make([]bool, g.GetV())
	d.dfs(g, s)
	return d
}

func (s *DepthFirstSearch) dfs(g Graph, v int) {
	s.marked[v] = true
	s.count++
	for node := g.GetAdj(v); node != nil; node = node.Next {
		if !s.marked[node.Val] {
			s.dfs(g, node.Val)
		}
	}
}

// Marked v和s是连通的吗
func (s *DepthFirstSearch) Marked(v int) bool {
	return s.marked[v]
}

// Count 与s连通的顶点总数
func (s *DepthFirstSearch) Count() int {
	return s.count
}
