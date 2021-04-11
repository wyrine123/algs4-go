package chapter1

/**
  动态连通性问题 QuickFind
	id数组里保存对象所在的分量
*/

type UFQuickFind struct {
	id    []int
	count int
}

func NewUFQuickFind(n int) *UFQuickFind {
	uf := new(UFQuickFind)
	uf.id = make([]int, n, n)
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	uf.count = n

	return uf
}

// Union 在p,q之间添加一条链接
func (uf *UFQuickFind) Union(p int, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)
	if pId == qId {
		// 已经在同一个分量上
		return
	}
	for i, v := range uf.id {
		// 把所在分量是p的对象的分量都调整为q
		if v == pId {
			uf.id[i] = qId
		}
	}
	uf.count--
}

// Find 查找p所在分量的标识符
func (uf *UFQuickFind) Find(p int) int {
	return uf.id[p]
}

// 检测p,q是否在一个分量上
func (uf *UFQuickFind) connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count 连通分量的数量
func (uf *UFQuickFind) Count() int {
	return uf.count
}
