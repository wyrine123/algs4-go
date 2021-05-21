package main

/**
  1. 两数之和
	https://leetcode-cn.com/problems/two-sum/
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // value=>index

	for i, item := range nums {
		leftValue := target - item
		if leftIndex, ok := m[leftValue]; ok {
			return []int{i, leftIndex}
		} else {
			// 不存在
			m[item] = i
		}
	}
	return nil
}
