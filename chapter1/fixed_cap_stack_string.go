package chapter1

import "fmt"

/**
  定长的字符串栈(slice实现)
	 这儿用slice实现吧,array定义时必须要声明数组长度
*/

type FixedCapStackString struct {
	s []string
	n int
}

// construct
func NewFixedCapStackString(cap int) *FixedCapStackString {
	fc := new(FixedCapStackString)
	fc.s = make([]string, cap, cap)
	fc.n = 0

	return fc
}

// push
func (p *FixedCapStackString) Push(s string) {
	if p.n+1 > cap(p.s) {
		// push后容量已经超出原数组,需要扩容
		fmt.Printf("push触发了resize,resize后的数组大小: %d\n", cap(p.s)*2)
		p.reSize(cap(p.s) * 2)
	}
	p.s[p.n] = s
	p.n++
}

// 重新调整数组大小
func (p *FixedCapStackString) reSize(n int) {
	newS := make([]string, n, n)
	// 复制
	for i := 0; i < p.n; i++ {
		newS[i] = p.s[i]
	}
	p.s = newS
}

// pop
func (p *FixedCapStackString) Pop() string {
	if p.isEmpty() {
		return "nil"
	}
	p.n--
	// 如果pop后的数据量只有原数组的1/4,resize
	if p.n < cap(p.s)/4 {
		fmt.Printf("pop触发了resize,n大小:%d, resize后的数组大小: %d\n", p.n, cap(p.s)/2)
		p.reSize(cap(p.s) / 2)
	}
	item := p.s[p.n]
	p.s[p.n] = ""
	return item
}

// 是否为空
func (p *FixedCapStackString) isEmpty() bool {
	return p.n == 0
}

// 栈容量
func (p *FixedCapStackString) Size() int {
	return p.n
}
