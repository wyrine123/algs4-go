package link_list

import "fmt"

/**
  单链
*/

// InitSingleLinkList 初始化单链
func InitSingleLinkList(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	// 初始化头节点
	header := new(Node)
	header.Val = nums[0]

	// 记录最后一个节点
	lastNode := header
	for _, i := range nums[1:] {
		node := new(Node)
		node.Val = i
		lastNode.Next = node
		lastNode = node
	}
	return header
}

// SingleLinkListHeaderInsert 单链头插
func SingleLinkListHeaderInsert(header *Node, val int) *Node {
	// 保存原先的头节点
	oldHeader := header
	// 构造新节点
	node := new(Node)
	node.Val = val
	// 把新节点的next指向旧头节点
	node.Next = oldHeader
	return node
}

// SingleLinkListInsert 单链在第i个位置插入
func SingleLinkListInsert(header *Node, i int, val int) *Node {
	if i == 1 {
		// 头插
		return SingleLinkListHeaderInsert(header, val)
	}
	// 其他位置插入
	// todo
	return nil
}

// PrintLinkList 打印输出链表
func PrintLinkList(header *Node) {
	for node := header; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Println()
}

// GetLinkListLength 获取单链长度
func GetLinkListLength(header *Node) int {
	length := 0
	for node := header; node != nil; node = node.Next {
		length++
	}
	return length
}
