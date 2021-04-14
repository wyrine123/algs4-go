package main

import "fmt"

/**
  两数相加
	https://leetcode-cn.com/problems/add-two-numbers/
	非空 的链表
*/

// todo 结果直接使用l1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	if l1 == nil || l2 == nil {
		return result
	}

	carry := 0 // 进位

	result.Val = -1
	resultNode := result

	for l1Node, l2Node := l1, l2; l1Node != nil || l2Node != nil || carry != 0; {
		sum := 0
		node := new(ListNode)
		if l1Node != nil && l2Node != nil {
			sum = l1Node.Val + l2Node.Val + carry
			if sum < 10 {
				carry = 0
				node.Val = sum
			} else {
				// 需要进位
				carry = 1
				node.Val = sum - 10
			}
			resultNode.Next = node
			resultNode = node

			l1Node = l1Node.Next
			l2Node = l2Node.Next
			continue
		}
		if l1Node != nil && l2Node == nil {
			// l1不为空，l2为空
			sum = l1Node.Val + carry
			if sum < 10 {
				carry = 0
				node.Val = sum
			} else {
				// 需要进位
				carry = 1
				node.Val = sum - 10
			}
			resultNode.Next = node
			resultNode = node

			l1Node = l1Node.Next
			continue
		}
		if l2Node != nil && l1Node == nil {
			// l2不为空,l1为空
			sum = l2Node.Val + carry
			if sum < 10 {
				carry = 0
				node.Val = sum
			} else {
				// 需要进位
				carry = 1
				node.Val = sum - 10
			}
			resultNode.Next = node
			resultNode = node

			l2Node = l2Node.Next
			continue
		}
		// l1,l2都为空,但carry!=0
		// 新增一个node
		node.Val = carry
		resultNode.Next = node
		carry = 0
	}

	// 删除第一个节点
	result = result.Next
	return result
}

func main() {
	l1 := new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 7

	result := addTwoNumbers(l1, l2)
	for node := result; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Println()
}
