package chapter1

/**
  动态连通性问题 加权的QuickUnion
*/

type UFQuickUnionWeighted struct {
	id    []int
	sz    []int // 权重大小
	count int
}

func NewUFQuickUnionWeighted(n int) *UFQuickUnionWeighted {
	uf := new(UFQuickUnionWeighted)
	uf.id = make([]int, n, n)
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.sz[i] = 1 // 初始化权重都是1
	}
	uf.count = n

	return uf
}

// Union 在p,q之间添加一条链接
func (uf *UFQuickUnionWeighted) Union(p int, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)
	if pId == qId {
		// 已经在同一个分量上
		return
	}
	// 不在同一个分量上,让他们同源
	// 比较权重,小的树要加到大的树上
	if uf.sz[pId] < uf.sz[qId] {
		uf.sz[qId] += uf.sz[pId]
		uf.id[pId] = qId
	} else {
		uf.sz[pId] += uf.sz[qId]
		uf.id[qId] = pId
	}

	uf.count--
}

// Find 查找p所在分量的标识符
func (uf *UFQuickUnionWeighted) Find(p int) int {
	value := uf.id[p]
	for value != uf.id[value] {
		value = uf.id[value]
	}
	return value
}

// 检测p,q是否在一个分量上
func (uf *UFQuickUnionWeighted) connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count 连通分量的数量
func (uf *UFQuickUnionWeighted) Count() int {
	return uf.count
}
