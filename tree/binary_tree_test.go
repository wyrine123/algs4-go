package tree

import (
	"testing"
)

// 完全二叉树测试struct
type CompleteBinaryTreeExample struct {
	id             int   // 用来定位
	construct      []int // 构造入参
	preOrderResult []int // 前序输出
}

// 二叉树测试struct
type BinaryTreeExample struct {
	id             int      // 用来定位
	construct      []string // 构造入参
	preOrderResult []int    // 前序输出
}

// 测试创建二叉树
func TestCreateCompleteBinaryTree(t *testing.T) {
	var examples []CompleteBinaryTreeExample

	example1 := CompleteBinaryTreeExample{}
	example1.id = 1
	example1.construct = []int{1, 2, 3, 4, 5, 6, 7}
	example1.preOrderResult = []int{1, 2, 4, 5, 3, 6, 7}
	examples = append(examples, example1)

	example2 := CompleteBinaryTreeExample{}
	example2.id = 2
	example2.construct = []int{1, 2, 3, 4, 5, 6, 7, 8}
	example2.preOrderResult = []int{1, 2, 4, 8, 5, 3, 6, 7}
	examples = append(examples, example2)

	example3 := CompleteBinaryTreeExample{}
	example3.id = 3
	example3.construct = []int{}
	example3.preOrderResult = []int{}
	examples = append(examples, example3)

	example4 := CompleteBinaryTreeExample{}
	example4.id = 4
	example4.construct = []int{1}
	example4.preOrderResult = []int{1}
	examples = append(examples, example4)

	for _, example := range examples {
		tree := CreateCompleteBinaryTree(example.construct)
		actual := PreOrderTraverse(tree)
		if !isTwoIntSliceSame(actual, example.preOrderResult) {
			// 实际和预想不相同
			t.Errorf("TestCreateCompleteBinaryTree 测试失败, id:%d, 预想: %v, 实际: %v", example.id, example.preOrderResult, actual)
		}
	}
}

func TestCreateBinaryTree(t *testing.T) {
	var examples []BinaryTreeExample

	example1 := BinaryTreeExample{}
	example1.id = 1
	example1.construct = []string{}
	example1.preOrderResult = []int{}
	examples = append(examples, example1)

	example2 := BinaryTreeExample{}
	example2.id = 2
	example2.construct = []string{"1"}
	example2.preOrderResult = []int{1}
	examples = append(examples, example2)

	example3 := BinaryTreeExample{}
	example3.id = 3
	example3.construct = []string{"null"}
	example3.preOrderResult = []int{}
	examples = append(examples, example3)

	example4 := BinaryTreeExample{}
	example4.id = 4
	example4.construct = []string{"1", "2", "null"}
	example4.preOrderResult = []int{1, 2}
	examples = append(examples, example4)

	example5 := BinaryTreeExample{}
	example5.id = 5
	example5.construct = []string{"1", "2", "null", "null", "5"}
	example5.preOrderResult = []int{1, 2, 5}
	examples = append(examples, example5)

	example6 := BinaryTreeExample{}
	example6.id = 6
	example6.construct = []string{"1", "2", "3", "null", "5", "null", "7"}
	example6.preOrderResult = []int{1, 2, 5, 3, 7}
	examples = append(examples, example6)

	for _, example := range examples {
		tree, err := CreateBinaryTree(example.construct)
		if err != nil {
			t.Errorf("TestCreateCompleteBinaryTree 测试初始化树失败, id:%d", example.id)
			continue
		}
		actual := PreOrderTraverse(tree)
		if !isTwoIntSliceSame(actual, example.preOrderResult) {
			// 实际和预想不相同
			t.Errorf("TestCreateCompleteBinaryTree 测试失败, id:%d, 预想: %v, 实际: %v", example.id, example.preOrderResult, actual)
		}
	}
}

// 判断2个int slice是否相同
// 每个index上的值相同,则认为是相同
func isTwoIntSliceSame(p []int, q []int) bool {
	if len(p) != len(q) {
		return false
	}

	for i := 0; i < len(p); i++ {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}
