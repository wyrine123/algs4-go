package chapter4

type SearchInterface interface {
	Marked(v int) bool // v和s是连通的吗
	Count() int        // 与s连通的顶点总数
}
