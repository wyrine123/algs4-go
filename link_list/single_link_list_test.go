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

// 测试逆序初始化单链
func TestReverseInitSingleLinkList(t *testing.T) {
	nums := []int{}
	PrintLinkList(ReverseInitSingleLinkList(nums))
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

// 测试单链位置插入
func TestSingleLinkListInsert(t *testing.T) {
	nums := []int{2, 3}
	header := InitSingleLinkList(nums)
	PrintLinkList(SingleLinkListInsert(header, 5, 1))
}

// 测试单链逆序
func TestSingleLinkListReverse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	header := InitSingleLinkList(nums)
	PrintLinkList(SingleLinkListReverse(header))
}
