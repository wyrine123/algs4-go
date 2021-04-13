package chapter2

/**
  排序算法
	选择排序
	插入排序
	希尔排序
	归并排序
	快速排序
	堆排序
*/

type SortInterface interface {
	less(v1 int, v2 int) bool
	exch(nums []int, i int, j int)
	isSorted(nums []int) bool
	Sort(nums []int)
}
