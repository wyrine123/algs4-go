package chapter4

// GraphDegree 计算节点v的度数
func GraphDegree(g *Graph, v int) int {
	return len(g.Adj(v))
}

// GraphMaxDegree 图的最大度数
func GraphMaxDegree(g *Graph) int {
	maxDegree := 0
	for v := 0; v < g.V(); v++ {
		if GraphDegree(g, v) > maxDegree {
			maxDegree = GraphDegree(g, v)
		}
	}
	return maxDegree
}

// GraphAvgDegree 计算所有顶点的平均度数
func GraphAvgDegree(g *Graph) float64 {
	return 2 * float64(g.E()) / float64(g.V())
}

// GraphNumberOfSelfLoops 计算自环的个数
func GraphNumberOfSelfLoops(g *Graph) int {
	num := 0
	for v := 0; v < g.V(); v++ {
		// 寻找每个节点临接表里是否有指向自己的
		for _, item := range g.Adj(v) {
			if item == v {
				num++
			}
		}
	}
	// 注意这儿要除以2,addEdge时,指向自己的会被添加2次
	return num / 2
}
