package chapter1

/**
  动态连通性问题 QuickUnion
*/

type UFQuickUnion struct {
	id    []int
	count int
}

func NewUFQuickUnion(n int) *UFQuickUnion {
	uf := new(UFQuickUnion)
	uf.id = make([]int, n, n)
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	uf.count = n

	return uf
}

// Union 在p,q之间添加一条链接
func (uf *UFQuickUnion) Union(p int, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)
	if pId == qId {
		// 已经在同一个分量上
		return
	}
	// 不在同一个分量上,让他们同源
	uf.id[pId] = qId
	uf.count--
}

// Find 查找p所在分量的标识符
func (uf *UFQuickUnion) Find(p int) int {
	value := uf.id[p]
	for value != uf.id[value] {
		value = uf.id[value]
	}
	return value
}

// 检测p,q是否在一个分量上
func (uf *UFQuickUnion) connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count 连通分量的数量
func (uf *UFQuickUnion) Count() int {
	return uf.count
}
