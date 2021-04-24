package main

/**
  省份数量
  https://leetcode-cn.com/problems/number-of-provinces/
*/

type unionFind struct {
	nums   []int // 数据
	weight []int // 权重
}

func NewUnionFind(count int) *unionFind {
	u := new(unionFind)
	u.nums = make([]int, count, count)
	u.weight = make([]int, count, count)

	for i := 0; i < count; i++ {
		u.nums[i] = i
		u.weight[i] = 1
	}

	return u
}

func (p *unionFind) union(i, j int) {
	// 已经是联通的
	iRoot := p.find(i)
	jRoot := p.find(j)
	if iRoot == jRoot {
		return
	}

	// 这儿判断下 是否已经联通
	// 防止出现死循环
	// 判断权重,小的权重往大的上合并
	if p.weight[iRoot] > p.weight[jRoot] {
		p.nums[jRoot] = iRoot
		p.weight[iRoot] += p.weight[jRoot]
	} else {
		p.nums[iRoot] = jRoot
		p.weight[jRoot] += p.weight[iRoot]
	}
}

func (p *unionFind) find(i int) int {
	root := i
	for p.nums[root] != root {
		root = p.nums[root]
	}
	return root
}

// 找到分割后的节点数量
func (p *unionFind) findPods() int {
	// 只要找到 p.num[i]==i 的数量就行
	count := 0
	for i, v := range p.nums {
		if i == v {
			count++
		}
	}
	return count
}

func (p *unionFind) isConnect(i, j int) bool {
	return p.find(i) == p.find(j)
}

func findCircleNum(isConnected [][]int) int {
	cityNum := len(isConnected)
	if cityNum == 0 {
		return 0
	}
	u := NewUnionFind(cityNum)

	for i1, v := range isConnected {
		for i2, j := range v {
			if j == 1 {
				u.union(i1, i2)
			}
		}
	}

	return u.findPods()
}

//[
//[1,0,0,0,0,0,0,0,0,1,0,0,0,0,0],
//[0,1,0,1,0,0,0,0,0,0,0,0,0,1,0],
//[0,0,1,0,0,0,0,0,0,0,0,0,0,0,0],
//[0,1,0,1,0,0,0,1,0,0,0,1,0,0,0],
//[0,0,0,0,1,0,0,0,0,0,0,0,1,0,0],
//[0,0,0,0,0,1,0,0,0,0,0,0,0,0,0],
//[0,0,0,0,0,0,1,0,0,0,0,0,0,0,0],
//[0,0,0,1,0,0,0,1,1,0,0,0,0,0,0],
//[0,0,0,0,0,0,0,1,1,0,0,0,0,0,0],
//[1,0,0,0,0,0,0,0,0,1,0,0,0,0,0],
//[0,0,0,0,0,0,0,0,0,0,1,0,0,0,0],
//[0,0,0,1,0,0,0,0,0,0,0,1,0,0,0],
//[0,0,0,0,1,0,0,0,0,0,0,0,1,0,0],
//[0,1,0,0,0,0,0,0,0,0,0,0,0,1,0],
//[0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]]
//func main() {
//	var isConnected [][]int
//	isConnected = append(isConnected, []int{1,0,0,0,0,0,0,0,0,1,0,0,0,0,0})
//	isConnected = append(isConnected, []int{0,1,0,1,0,0,0,0,0,0,0,0,0,1,0})
//	isConnected = append(isConnected, []int{0,0,1,0,0,0,0,0,0,0,0,0,0,0,0})
//	isConnected = append(isConnected, []int{0,1,0,1,0,0,0,1,0,0,0,1,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,0,1,0,0,0,0,0,0,0,1,0,0})
//	isConnected = append(isConnected, []int{0,0,0,0,0,1,0,0,0,0,0,0,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,0,0,0,1,0,0,0,0,0,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,0,0,0,0,1,1,0,0,0,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,1,0,0,0,1,1,0,0,0,0,0,0})
//	isConnected = append(isConnected, []int{1,0,0,0,0,0,0,0,0,1,0,0,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,0,0,0,0,0,0,0,1,0,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,1,0,0,0,0,0,0,0,1,0,0,0})
//	isConnected = append(isConnected, []int{0,0,0,0,1,0,0,0,0,0,0,0,1,0,0})
//	isConnected = append(isConnected, []int{0,1,0,0,0,0,0,0,0,0,0,0,0,1,0})
//	isConnected = append(isConnected, []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,1})
//
//	fmt.Println(findCircleNum(isConnected))
//}
