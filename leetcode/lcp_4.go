package main

/**
  4. 寻找两个正序数组的中位数
	https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// nums1, nums2 有序
	totalCount := len(nums1) + len(nums2)
	var (
		isOdd  bool // 是否是奇数
		median int  // 中位数
	)
	if totalCount%2 != 0 {
		isOdd = true
	}
	median = totalCount / 2

	// 归并排序
	var (
		i1, i2    int   // 下标
		mergeNums []int // 归并后的数组
	)
	for i1 < len(nums1) || i2 < len(nums2) {
		if i1 == len(nums1) {
			// i1到头,i2没到头
			mergeNums = append(mergeNums, nums2[i2])
			i2++
		} else if i2 == len(nums2) {
			// i1没超,i2已到头
			mergeNums = append(mergeNums, nums1[i1])
			i1++
		} else {
			// i1, i2都没超过
			if nums1[i1] <= nums2[i2] {
				mergeNums = append(mergeNums, nums1[i1])
				i1++
			} else {
				mergeNums = append(mergeNums, nums2[i2])
				i2++
			}
		}
	}

	if len(mergeNums) == 0 {
		return 0
	}
	// 寻找第median个数
	if isOdd {
		return float64(mergeNums[median])
	} else {
		return (float64(mergeNums[median]) + float64(mergeNums[median-1])) / 2
	}
}
