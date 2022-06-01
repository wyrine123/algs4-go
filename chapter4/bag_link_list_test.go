package chapter4

import "testing"

func TestBagLinkList(t *testing.T) {
	bag := NewBagLinkList()
	if !bag.IsEmpty() || bag.Size() != 0 {
		t.Errorf("TestBagLinkList 背包不为空")
	}

	// 添加
	bag.Add(1)
	if bag.IsEmpty() {
		t.Errorf("TestBagLinkList 背包为空")
	}
	if bag.Size() != 1 {
		t.Errorf("TestBagLinkList 背包容量不为1")
	}

	t.Logf("背包第一次添加详情: %v", bag.Iteration())

	bag.Add(2)

	if bag.IsEmpty() {
		t.Errorf("TestBagLinkList 背包为空")
	}
	if bag.Size() != 2 {
		t.Errorf("TestBagLinkList 背包容量不为2")
	}

	t.Logf("背包第二次添加详情: %v", bag.Iteration())
}
