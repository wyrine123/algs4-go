package main

import "fmt"

/**
  两数相加
	https://leetcode-cn.com/problems/add-two-numbers/
	非空 的链表
*/

// todo 结果直接使用l1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	carry := 0 // 进位

	// 使用l1保存结果
	// 这儿循环内保证 l1一定比l2长,所以就不判断 l2 != nil 这个条件了
	for l1Node, l2Node := l1, l2; l1Node != nil || carry != 0; {
		sum := 0
		if l1Node != nil && l2Node != nil {
			sum = l1Node.Val + l2Node.Val + carry
			if sum >= 10 {
				// 需要进位
				carry = 1
				sum -= 10
			} else {
				carry = 0
			}
			if (l1Node.Next == nil && l2Node.Next != nil) || (l1Node.Next == nil && carry != 0) {
				// l2比l1链表长,或者链表结束,但是carry为不为0 需要给l1新增一个进位节点
				zeroNode := new(ListNode)
				zeroNode.Val = carry
				l1Node.Next = zeroNode
				carry = 0 // 清除进位
			}
			l1Node.Val = sum // 原地保存

			l1Node = l1Node.Next
			l2Node = l2Node.Next
			continue
		}
		if l1Node != nil && l2Node == nil {
			// l1不为空，l2为空
			sum = l1Node.Val + carry
			if sum >= 10 {
				// 需要进位
				carry = 1
				sum -= 10
			} else {
				carry = 0
			}
			if l1Node.Next == nil && carry != 0 {
				// l1结束,但是还有进位
				zeroNode := new(ListNode)
				zeroNode.Val = carry
				l1Node.Next = zeroNode
				carry = 0 // 清除进位
			}
			l1Node.Val = sum // 原地保存
			l1Node = l1Node.Next
			continue
		}
	}

	return l1
}

func main() {
	l1 := new(ListNode)
	l1.Val = 9
	l1.Next = new(ListNode)
	l1.Next.Val = 9
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 9

	l2 := new(ListNode)
	l2.Val = 9
	//l2.Next = new(ListNode)
	//l2.Next.Val = 6
	//l2.Next.Next = new(ListNode)
	//l2.Next.Next.Val = 7

	result := addTwoNumbers(l2, l1)
	for node := result; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Println()
}
