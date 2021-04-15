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

// ReverseInitSingleLinkList 逆序构造单链
func ReverseInitSingleLinkList(nums []int) *Node {
	// stack 头插方式
	var header *Node
	for _, i := range nums {
		oldHeader := header
		node := new(Node)
		node.Val = i
		node.Next = oldHeader
		header = node
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
	if i <= 0 {
		return header
	}
	if i == 1 {
		// 头插
		return SingleLinkListHeaderInsert(header, val)
	}
	// 其他位置插入 需要找到该位置的前置节点
	prev := new(Node)
	node := header
	for j := 1; j < i; j++ {
		prev = node
		if node != nil {
			node = node.Next
		}
	}
	if prev == nil {
		// 没到i位置,链表就结束了
		return header
	}
	insertNode := new(Node)
	insertNode.Val = val
	insertNode.Next = prev.Next
	prev.Next = insertNode
	return header
}

// SingleLinkListReverse 单链逆序
func SingleLinkListReverse(header *Node) *Node {
	// 头插 实现逆序
	var reverseLinkListHeader *Node
	for node := header; node != nil; node = node.Next {
		oldHeader := reverseLinkListHeader
		newNode := new(Node)
		newNode.Val = node.Val
		newNode.Next = oldHeader
		reverseLinkListHeader = newNode
	}
	return reverseLinkListHeader
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
