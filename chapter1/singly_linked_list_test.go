package chapter1

import (
	"fmt"
	"testing"
)

// 测试删除单链的尾节点
func TestSinglyLinkedList_DelLastNode(t *testing.T) {
	l := NewSinglyLinkedList()
	if l.DelLastNode() != nil {
		t.Error("测试删除单链的尾节点 删除空链表返回不是nil")
	}
	l.Insert("1")
	if l.Size() != 1 {
		t.Error("测试删除单链的尾节点 链表长度不是1")
	}
	l.Insert("2")
	l.Insert("3")
	fmt.Println(l.String())
	if l.Size() != 3 {
		t.Error("测试删除单链的尾节点 链表长度不是3")
	}
	if l.DelLastNode().item != "3" {
		t.Error("测试删除单链的尾节点 删除尾节点不是3")
	}
	fmt.Println(l.String())
	if l.DelLastNode().item != "2" {
		t.Error("测试删除单链的尾节点 删除尾节点不是2")
	}
	fmt.Println(l.String())
	if l.DelLastNode().item != "1" {
		t.Error("测试删除单链的尾节点 删除尾节点不是1")
	}
	fmt.Println(l.String())
	l.DelLastNode()
	fmt.Println(l.String())
}

// 测试删除单链的首节点
func TestSinglyLinkedList_DelFirstNode(t *testing.T) {
	l := NewSinglyLinkedList()
	if l.DelFirstNode() != nil {
		t.Error("测试删除单链的首节点 删除空链表返回不是nil")
	}
	l.Insert("1")
	if l.Size() != 1 {
		t.Error("测试删除单链的首节点 链表长度不是1")
	}
	l.Insert("2")
	l.Insert("3")
	fmt.Println(l.String())
	if l.Size() != 3 {
		t.Error("测试删除单链的首节点 链表长度不是3")
	}
	if l.DelFirstNode().item != "1" {
		t.Error("测试删除单链的首节点 删除尾节点不是1")
	}
	fmt.Println(l.String())
	if l.DelFirstNode().item != "2" {
		t.Error("测试删除单链的首节点 删除尾节点不是2")
	}
	fmt.Println(l.String())
	if l.DelFirstNode().item != "3" {
		t.Error("测试删除单链的首节点 删除尾节点不是3")
	}
	fmt.Println(l.String())
	l.DelFirstNode()
	fmt.Println(l.String())
}

// 测试删除单链的某一个节点
func TestSinglyLinkedList_Del(t *testing.T) {
	l := NewSinglyLinkedList()
	if l.Del(-1) != nil {
		t.Error("测试删除单链的首节点 删除空链表返回不是nil")
	}
	if l.Del(3) != nil {
		t.Error("测试删除单链的首节点 删除空链表返回不是nil")
	}
	l.Insert("1")
	if l.Size() != 1 {
		t.Error("测试删除单链的首节点 链表长度不是1")
	}
	if l.Del(1).item != "1" {
		t.Error("测试删除单链的首节点 删除第1个节点返回不是1")
	}
	if l.Size() != 0 {
		t.Error("测试删除单链的首节点 链表长度不是0")
	}
	l.Insert("1")
	l.Insert("2")
	l.Insert("3")
	fmt.Println(l.String())
	if l.Del(2).item != "2" {
		t.Error("测试删除单链的首节点 删除第2个节点返回不是2")
	}
	fmt.Println(l.String())
	if l.Del(2).item != "3" {
		t.Error("测试删除单链的首节点 删除第2个节点返回不是3")
	}
	fmt.Println("------" + l.String())
	l.Insert("2")
	l.Insert("3")
	fmt.Println("======" + l.String())
	if l.Del(1).item != "1" {
		t.Error("测试删除单链的首节点 删除第1个节点返回不是1")
	}
	fmt.Println(l.String())
}

// 测试单链的find方法
func TestSinglyLinkedList_Find(t *testing.T) {
	l := NewSinglyLinkedList()
	if l.Find("1") != false {
		t.Error("测试单链的find方法 在空的链表内寻找1 失败")
	}
	l.Insert("1")
	if l.Find("1") != true {
		t.Error("测试单链的find方法 在空的链表内寻找1 失败")
	}
	l.Insert("2")
	l.Insert("3")
	if l.Find("2") != true {
		t.Error("测试单链的find方法 在空的链表内寻找2 失败")
	}
}
