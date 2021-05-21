package main

/**
  3. 无重复字符的最长子串
	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	var (
		slideWindow []int32
		maxLength   int // 最大长度
	)
	for _, item := range s {
		if exist, index := isInSlice(slideWindow, item); exist {
			// 存在
			if len(slideWindow) > maxLength {
				maxLength = len(slideWindow)
			}
			// 从左侧删除
			slideWindow = append(slideWindow, item)
			slideWindow = slideWindow[index+1:]
		} else {
			// 不存在
			slideWindow = append(slideWindow, item)
		}
	}
	if len(slideWindow) > maxLength {
		maxLength = len(slideWindow)
	}
	return maxLength
}

// 判断target是否在slice中,并返回对应的index
func isInSlice(nums []int32, target int32) (bool, int) {
	for i, item := range nums {
		if target == item {
			return true, i
		}
	}
	return false, 0
}
