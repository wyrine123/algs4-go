package chapter4

type PathInterface interface {
	HasPathTo(v int)    // 是否存在s到v的路径
	PathTo(v int) []int // s到v的路径,不存在返回nil
}
