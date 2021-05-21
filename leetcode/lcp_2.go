package main

/**
  2. 两数相加
	https://leetcode-cn.com/problems/add-two-numbers/
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// l1, l2不为空
	var carry int // 进位
	// 结果保存在l1中
	for l1Node, l2Node := l1, l2; l1Node != nil; l1Node, l2Node = l1Node.Next, l2Node.Next {
		result := l1Node.Val + l2Node.Val + carry
		if result < 10 {
			carry = 0
			l1Node.Val = result
		} else {
			carry = 1 // 进位
			l1Node.Val = result - 10
		}

		if l1Node.Next == nil {
			if l2Node.Next == nil {
				// l1, l2 next都不存在了
				// 只需要注意进位
				if carry != 0 {
					// 需要在l1,l2上增加一个空节点
					node1, node2 := ListNode{}, ListNode{}
					l2Node.Next = &node2
					l1Node.Next = &node1
				}
			} else {
				// l1 next不存在,l2 next存在
				// 补充l1 next
				node1 := ListNode{}
				l1Node.Next = &node1
			}
		} else {
			// l1 next 不为空
			if l2Node.Next == nil {
				// l2 next为空
				// 补充l2 next
				node2 := ListNode{}
				l2Node.Next = &node2
			}
		}
	}
	return l1
}
