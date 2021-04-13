package chapter2

/**
  选择排序
	不断的选择剩余元素的最小值
*/

type selectionSort struct {
}

func NewSelectionSort() *selectionSort {
	s := new(selectionSort)
	return s
}

// 比较
func (p *selectionSort) less(v1 int, v2 int) bool {
	return v1 < v2
}

// 交换
func (p *selectionSort) exch(nums []int, i int, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

// 数组是否是已排序
func (p *selectionSort) isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return false
		}
	}
	return true
}

// Sort 排序实现
func (p *selectionSort) Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i // 最小的元素下标
		// 从剩余的元素中寻找最小值
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		p.exch(nums, i, min)
	}
}
