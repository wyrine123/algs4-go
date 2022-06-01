package chapter4

type SearchInterface interface {
	Search(g *Graph, s int) // 找到和起点s连通的所有顶点
	Marked(v int) bool      // v和s是连通的吗
	Count() int             // 与s连通的顶点总数
}
