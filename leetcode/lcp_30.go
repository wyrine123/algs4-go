package main

import (
	"fmt"
	"math"
)

/**
  魔塔游戏
	https://leetcode-cn.com/problems/p0NxJO/

	迭代数组
	数组前n个值的和必须>=0,如果<0,需要把前n个值里最小的数值移到最后,保证前n个值的和>=0即可.
	这个前n个值不需要保证顺序,只需要记录最小值即可.
*/

func magicTower(nums []int) int {
	nSum := 0           // 前n个值的和
	var shiftNums []int // 移位数组
	// 排序后的数据
	minHeap := NewMinHeap()
	for i := 0; i < len(nums); i++ {
		minHeap.Insert(nums[i])
		nSum += nums[i]
		if nSum < 0 {
			// 需要把最小的值弹出 移到 移位数组
			minValue, _ := minHeap.DeleteMin()
			shiftNums = append(shiftNums, minValue)
			nSum += -minValue // 加回去
		}
	}

	for _, v := range shiftNums {
		nSum += v
	}

	if nSum < 0 {
		return -1
	} else {
		return len(shiftNums)
	}
}

type MinHeap struct {
	Element []int
}

// MinHeap构造方法
func NewMinHeap() *MinHeap {
	// 第一个元素仅用于结束insert中的 for 循环
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

// 插入数字,插入数字需要保证堆的性质
func (H *MinHeap) Insert(v int) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	// 上浮
	for ; H.Element[i/2] > v; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// 删除并返回最小值
func (H *MinHeap) DeleteMin() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1] < H.Element[child] {
			child++
		}
		// 下滤一层
		if lastElement > H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil
}

// 堆的大小
func (H *MinHeap) Length() int {
	return len(H.Element) - 1
}

// 获取最小堆的最小值
func (H *MinHeap) Min() (int, error) {
	if len(H.Element) > 1 {
		return H.Element[1], nil
	}
	return 0, fmt.Errorf("heap is empty")
}

// MinHeap格式化输出
func (H *MinHeap) String() string {
	return fmt.Sprintf("%v", H.Element[1:])
}
