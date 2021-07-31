package tree

import (
	"errors"
	"math"
	"strconv"
)

// CreateCompleteBinaryTree 根据数组创建完全二叉树
func CreateCompleteBinaryTree(values []int) *BinaryTreeNode {
	if len(values) == 0 {
		return nil
	}

	var indexTreeNode = make([]*BinaryTreeNode, len(values)+1) // 保存每个位置的node

	for i := 0; i < len(values); i++ {
		index := i + 1
		node := new(BinaryTreeNode)
		node.Val = values[i]
		indexTreeNode[index] = node

		// 向下取整,找到当前节点的父节点,增加到父节点中
		parentIndex := int(math.Floor(float64(index) / 2))
		if parentIndex != 0 {
			// 不处理根节点
			if index == parentIndex*2 {
				// 左节点
				indexTreeNode[parentIndex].Left = node
			} else {
				// 右节点
				indexTreeNode[parentIndex].Right = node
			}
		}
	}

	return indexTreeNode[1] // 根节点在数组index=1上
}

// CreateBinaryTree 根据数组创建二叉树
// null代表该节点为空, 其他的解析成int类型来处理
func CreateBinaryTree(values []string) (*BinaryTreeNode, error) {
	if len(values) == 0 {
		return nil, nil
	}

	var indexTreeNode = make([]*BinaryTreeNode, len(values)+1) // 保存每个位置的node

	for i := 0; i < len(values); i++ {
		index := i + 1
		var node *BinaryTreeNode
		if values[i] == "null" {
			indexTreeNode[index] = node
		} else {
			// 解析成int类型
			intV, err := strconv.Atoi(values[i])
			if err != nil {
				return nil, errors.New("error param")
			}
			node = new(BinaryTreeNode)
			node.Val = intV
			indexTreeNode[index] = node
		}

		// 向下取整,找到当前节点的父节点,增加到父节点中
		partnerIndex := int(math.Floor(float64(index) / 2))
		if partnerIndex != 0 {
			// 不处理根节点

			// 先判断父节点是否为空
			if indexTreeNode[partnerIndex] == nil {
				return nil, errors.New("error param")
			}

			if index == partnerIndex*2 {
				// 左节点
				indexTreeNode[partnerIndex].Left = node
			} else {
				// 右节点
				indexTreeNode[partnerIndex].Right = node
			}
		}
	}

	return indexTreeNode[1], nil // 根节点在数组index=1上
}

// PreOrderTraverse 二叉树前序遍历
// 中间节点->左节点->右节点
func PreOrderTraverse(root *BinaryTreeNode) []int {
	var result []int
	var preOrder func(node *BinaryTreeNode)

	preOrder = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)

	return result
}

// InOrderTraverse 二叉树中序遍历
// 左节点->中间节点->右节点
func InOrderTraverse(root *BinaryTreeNode) []int {
	var result []int
	var inOrder func(node *BinaryTreeNode)

	inOrder = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		result = append(result, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)

	return result
}

// PostOrderTraverse 二叉树后续遍历
// 左节点->右节点->中间节点
func PostOrderTraverse(root *BinaryTreeNode) []int {
	var result []int
	var postOrder func(node *BinaryTreeNode)

	postOrder = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		result = append(result, node.Val)
	}
	postOrder(root)

	return result
}
