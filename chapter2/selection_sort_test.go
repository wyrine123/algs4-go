package chapter2

import (
	"math/rand"
	"testing"
	"time"
)

// 测试选择排序
func TestSelectionSort_Sort(t *testing.T) {
	examples := [][]int{
		[]int{},
		[]int{1},
		[]int{3, 2, 1, 5, -1, 10},
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		example := make([]int, 1000, 1000)
		for j := 0; j < 1000; j++ {
			example[j] = rand.Intn(100000)
		}
		examples = append(examples, example)
	}
	s := NewSelectionSort()

	for i, example := range examples {
		s.Sort(example)
		if !s.isSorted(example) {
			t.Errorf("测试选择排序 index:%d 排序后不对", i)
		}
	}
}
