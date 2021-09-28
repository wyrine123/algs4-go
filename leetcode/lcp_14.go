package main

import "fmt"

/**
 14. 最长公共前缀
	https://leetcode-cn.com/problems/longest-common-prefix/
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix = []byte(strs[0]) // 默认最长公共前缀就是数组下标0
	// 为了节约内存空间 每次不再重新生成一个prefix,而是共用一个数组,通过一个下标来标识当前
	var prefixIndex = len(prefix)

	for i := 1; i < len(strs); i++ {
		// 寻找prefix和当前字符串的最长公共前缀
		if strs[i] == "" {
			// 当前为空
			return ""
		}
		current := []byte(strs[i])

		// 获取current和prefix中短的那个数组长度
		var min int
		if prefixIndex <= len(current) {
			min = prefixIndex
		} else {
			min = len(current)
		}

		// prefixIndex 重置为0
		prefixIndex = 0
		for j := 0; j < min; j++ {
			if prefix[j] == current[j] {
				prefixIndex++
			} else {
				break
			}
		}

		if prefixIndex == 0 {
			return ""
		}
	}

	return string(prefix[:prefixIndex])
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "f"}))

}
