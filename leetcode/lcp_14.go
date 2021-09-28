package main

/**
 14. 最长公共前缀
	https://leetcode-cn.com/problems/longest-common-prefix/
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix = []byte(strs[0]) // 默认最长公共前缀就是数组下标0

	for i := 1; i < len(strs); i++ {
		// 寻找prefix和当前字符串的最长公共前缀
		if strs[i] == "" {
			// 当前为空
			return ""
		}
		current := []byte(strs[i])

		// 获取current和prefix中短的那个数组长度
		var min int
		if len(prefix) <= len(current) {
			min = len(prefix)
		} else {
			min = len(current)
		}

		var currentPrefix []byte
		for j := 0; j < min; j++ {
			if prefix[j] == current[j] {
				currentPrefix = append(currentPrefix, prefix[j])
			} else {
				break
			}
		}

		if len(currentPrefix) == 0 {
			return ""
		} else {
			prefix = currentPrefix
		}
	}

	return string(prefix)
}

//func main() {
//	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
//}
