package chapter2

import (
	"fmt"
	"strings"
)

/**
  优先队列-使用二叉堆实现
	k节点的父节点k/2,k节点的2个子节点2k,2k+1
*/

type maxPQ struct {
	pq []int // 数组保存的二叉堆
	n  int   // 堆大小pq[1:], 0不使用
}

func NewMaxPQ() *maxPQ {
	p := new(maxPQ)
	p.pq = make([]int, 1, 1)
	p.pq[0] = 0 // 直接填充,不使用

	return p
}

// Insert 插入数据
func (p *maxPQ) Insert(num int) {
	p.n++
	//p.pq[p.n] = num
	p.pq = append(p.pq, num)
	p.swim(p.n)
}

// DelMax 删除最大的数
func (p *maxPQ) DelMax() int {
	if p.isEmpty() {
		return -1
	}
	max := p.pq[1]
	// 把最后数移到第一个
	p.pq[1] = p.pq[p.n]
	p.n--
	p.sink(1)
	p.pq = p.pq[:p.n]
	return max
}

func (p *maxPQ) String() string {
	s := ""
	for i := 1; i <= p.n; i++ {
		s += fmt.Sprintf("%d\t", p.pq[i])
	}
	s = strings.TrimRight(s, "\t")
	return s
}

// 比较2个节点上数的大小
func (p *maxPQ) less(i int, j int) bool {
	return p.pq[i] < p.pq[j]
}

// 交换2个节点上的数
func (p *maxPQ) exch(i int, j int) {
	t := p.pq[i]
	p.pq[i] = p.pq[j]
	p.pq[j] = t
}

// 是否为空
func (p *maxPQ) isEmpty() bool {
	return p.n == 0
}

// Size 优先队列的大小
func (p *maxPQ) Size() int {
	return p.n
}

// 上浮
func (p *maxPQ) swim(k int) {
	// 比较该节点与它的父节点,如果该节点比它的父节点大,交换
	// 然后继续比较,直到比它的父节点小或者已经上浮到根节点
	for ; k > 1 && p.less(k/2, k); k = k / 2 {
		p.exch(k, k/2)
	}
}

// 下沉
func (p *maxPQ) sink(k int) {
	// 比较该节点是否比2个子节点小,如果有比2个节点小的,和2个节点中比较大的那个节点交换
	// 然后继续比较,直到都比2个子节点大或者没有子节点为止
	j := 2 * k                  // 默认左子节点是最小的
	for ; j <= p.n; j = 2 * j { // 有左子节点
		if j < p.n && p.less(j, j+1) {
			// 左右节点都在,并且右节点比左节点大
			j++
		}
		// j保存的就是子节点中最大的节点
		if !p.less(k, j) {
			// 父节点比子节点大
			break
		}
		p.exch(k, j)
	}
}
