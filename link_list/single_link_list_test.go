package link_list

import (
	"fmt"
	"testing"
)

// 测试初始化单链
func TestInitSingleLinkListInit(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	PrintLinkList(InitSingleLinkList(nums))
}

// 测试获取单链长度
func TestGetLinkListLength(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	header := InitSingleLinkList(nums)
	fmt.Println(GetLinkListLength(header))
}

// 测试单链头插
func TestSingleLinkListHeaderInsert(t *testing.T) {
	nums := []int{1}
	header := InitSingleLinkList(nums)
	PrintLinkList(SingleLinkListHeaderInsert(header, -1))
}
