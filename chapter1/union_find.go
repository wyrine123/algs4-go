package chapter1

/**
  动态连通性问题
	QuickFind
	QuickUnion
	加权QuickUnion
*/

type UF interface {
	Union(p int, q int)
	Find(p int) int
	Count() int
}
