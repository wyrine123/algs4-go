package chapter4

/**
  深度优先搜索
*/

type DepthFirstSearch struct {
	markedV []bool // 节点是否被标记过
	countV  int    // 访问过的节点数
}

func NewDepthFirstSearch(g *Graph, s int) *DepthFirstSearch {
	d := new(DepthFirstSearch)
	d.markedV = make([]bool, g.V(), g.V())

	d.dfs(g, s)
	return d
}

func (s *DepthFirstSearch) dfs(g *Graph, v int) {
	s.markedV[v] = true // 该节点被访问过
	s.countV++
	for _, w := range g.Adj(v) {
		if !s.Marked(w) {
			// 没有被访问过,继续递归
			s.dfs(g, w)
		}
	}
}

func (s *DepthFirstSearch) Marked(w int) bool {
	return s.markedV[w]
}

func (s *DepthFirstSearch) Count() int {
	return s.countV
}
