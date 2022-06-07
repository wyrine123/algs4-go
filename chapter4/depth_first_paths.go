package chapter4

type DepthFirstPaths struct {
	markedV []bool // 节点是否被标记过
	edgeTo  []int  // index的节点是由哪个节点指向的
	s       int    // 起点
}

func NewDepthFirstPaths(g *Graph, s int) *DepthFirstPaths {
	path := new(DepthFirstPaths)
	path.s = s
	path.markedV = make([]bool, g.V())
	path.edgeTo = make([]int, g.V())

	path.dfs(g, s)

	return path
}

func (path *DepthFirstPaths) dfs(g *Graph, v int) {
	path.markedV[v] = true
	for _, w := range g.Adj(v) {
		if !path.Marked(w) {
			// 没有被标记过
			path.edgeTo[w] = v
			path.dfs(g, w)
		}
	}
}

func (path *DepthFirstPaths) Marked(w int) bool {
	return path.markedV[w]
}

func (path *DepthFirstPaths) HasPathTo(v int) bool {
	return path.Marked(v)
}

func (path *DepthFirstPaths) PathTo(v int) []int {
	if !path.HasPathTo(v) {
		return nil
	}

	var tmp []int
	for w := v; w != path.s; w = path.edgeTo[w] {
		tmp = append(tmp, w)
	}
	tmp = append(tmp, path.s)

	// 反转数组
	var res = make([]int, 0, len(tmp))
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}

	return res
}
