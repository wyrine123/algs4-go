package chapter1

import "testing"

func TestNewLinkListStack(t *testing.T) {
	stack := NewLinkListStack()
	if stack.Size() != 0 {
		t.Error("TestNewLinkListStack 初始化stack size不是0")
	}

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	if stack.Size() != 3 {
		t.Error("TestNewLinkListStack stack size不是3")
	}
	if stack.Pop() != "3" {
		t.Error("TestNewLinkListStack stack pop不是3")
	}
	if stack.Pop() != "2" {
		t.Error("TestNewLinkListStack stack pop不是2")
	}
	if stack.Pop() != "1" {
		t.Error("TestNewLinkListStack stack pop不是1")
	}
	if stack.Size() != 0 {
		t.Error("TestNewLinkListStack stack size不是0")
	}
}
