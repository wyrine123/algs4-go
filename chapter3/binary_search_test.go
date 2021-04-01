package chapter3

import "testing"

type BinarySearchTestExamples struct {
	Nums   []int
	A      int
	Expect int
}

func TestBinarySearch(t *testing.T) {
	examples := []BinarySearchTestExamples{
		{
			[]int{1}, 1, 1,
		},
		{
			[]int{1}, 0, -1,
		},
		{
			[]int{2}, 1, -1,
		},
		{
			[]int{}, 1, -1,
		},
		{
			[]int{1, 2}, 1, 1,
		},
		{
			[]int{1, 2}, 2, 1,
		},
		{
			[]int{1, 2}, 3, -1,
		},
		{
			[]int{1, 2, 3}, 3, 1,
		},
		{
			[]int{1, 2, 3}, 4, -1,
		},
	}

	for _, example := range examples {
		result := BinarySearch(example.Nums, example.A)
		if result != example.Expect {
			t.Errorf("BinarySearch 结果和预期不一致, nums:%v, a:%d, expect:%d, result:%d",
				example.Nums, example.A, example.Expect, result)
		}
	}
}

func TestBinarySearchReturnIndex(t *testing.T) {
	examples := []BinarySearchTestExamples{
		{
			[]int{1}, 1, 0,
		},
		{
			[]int{1}, 0, -1,
		},
		{
			[]int{2}, 1, -1,
		},
		{
			[]int{}, 1, -1,
		},
		{
			[]int{1, 2}, 1, 0,
		},
		{
			[]int{1, 2}, 2, 1,
		},
		{
			[]int{1, 2}, 3, -1,
		},
		{
			[]int{1, 2, 3}, 3, 2,
		},
		{
			[]int{1, 2, 3}, 4, -1,
		},
	}

	for _, example := range examples {
		result := BinarySearchReturnIndex(example.Nums, example.A, 0, len(example.Nums)-1)
		if result != example.Expect {
			t.Errorf("TestBinarySearchReturnIndex 结果和预期不一致, nums:%v, a:%d, expect:%d, result:%d",
				example.Nums, example.A, example.Expect, result)
		}
	}
}
