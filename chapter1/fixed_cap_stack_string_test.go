package chapter1

import "testing"

func TestNewFixedCapStackString(t *testing.T) {
	stack := NewFixedCapStackString(4)
	if stack.Size() != 0 {
		t.Error("TestNewFixedCapStackString 初始化stack size不是0")
	}
	stack.Push("1")
	if stack.Size() != 1 {
		t.Error("TestNewFixedCapStackString stack size不是1")
	}
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	stack.Push("5")
	if stack.Size() != 5 {
		t.Error("TestNewFixedCapStackString stack size不是5")
	}
	if stack.Pop() != "5" {
		t.Error("TestNewFixedCapStackString stack pop不是5")
	}
	stack.Pop() // "4"
	stack.Pop() // "3"
	stack.Pop() // "2"
	if stack.Size() != 1 {
		t.Error("TestNewFixedCapStackString stack size不是1")
	}
}
