package chapter3

/**
  二分查找
	默认nums是从小到大排序过的,如果查找到返回1,不存在返回-1
*/

func BinarySearch(nums []int, a int) int {
	// 这种边界写法没有下面那种简洁,但是下面那个边界不太好理解
	if (len(nums) == 1 && nums[0] != a) || len(nums) == 0 {
		return -1
	}

	mid := len(nums) / 2
	if nums[mid] == a {
		return 1
	} else if a < nums[mid] {
		// a在左半侧
		return BinarySearch(nums[:mid], a)
	} else {
		// a在右半侧
		return BinarySearch(nums[mid:], a)
	}
}

// 二分查找返回index
func BinarySearchReturnIndex(nums []int, target int, lo int, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		// 目标值在lo-mid
		return BinarySearchReturnIndex(nums, target, lo, mid-1)
	} else {
		// 目标值在mid-hi
		return BinarySearchReturnIndex(nums, target, mid+1, hi)
	}
}
