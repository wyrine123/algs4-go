package main

import "fmt"

/**
  61. 旋转链表
  	https://leetcode-cn.com/problems/rotate-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	/**
	1. 先获取链表的长度length
	2. 然后k%length取余 actualCount
	3. 找到倒数第 actualCount+1 的节点,保存后续节点,然后把该节点的next置为nil
	4. 把截断的链表和原先的拼接到一起
	*/
	if k == 0 {
		// 不需要翻转
		return head
	}
	if head == nil {
		// 链表长度为0
		return head
	}
	length := getListLength(head)
	actualCount := k % length // 不重复翻转的次数
	if actualCount == 0 {
		// 不需要翻转
		return head
	}

	var newHead *ListNode
	for node, i := head, 1; node != nil; node = node.Next {
		if i == length-actualCount {
			newHead = node.Next
			node.Next = nil
		}
		i++
	}

	// 找到newHead最后的节点
	node := newHead
	for ; node != nil && node.Next != nil; node = node.Next {
	}
	if node != nil {
		node.Next = head
	}
	return newHead
}

// 	获取链表长度
func getListLength(head *ListNode) int {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	return length
}

func main() {
	headNode := new(ListNode)
	headNode.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	headNode.Next = node2

	//node3 := new(ListNode)
	//node3.Val = 3
	//node2.Next = node3
	//
	//node4 := new(ListNode)
	//node4.Val = 4
	//node3.Next = node4
	//
	//node5 := new(ListNode)
	//node5.Val = 5
	//node4.Next = node5

	newNode := rotateRight(headNode, 2)
	for node := newNode; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
}
