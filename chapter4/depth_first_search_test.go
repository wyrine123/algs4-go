package chapter4

import (
	"os"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
	f, err := os.Open("./data/tinyG.txt")
	if err != nil {
		t.Fatalf("TestDepthFirstSearch open file err : %s", err.Error())
	}
	g := NewGraphFromFile(f)

	dfs := NewDepthFirstSearch(g, 0)
	if dfs.Count() != 7 {
		t.Errorf("TestDepthFirstSearch Count !=7, actual = %d", dfs.Count())
	}

	dfs = NewDepthFirstSearch(g, 7)
	if dfs.Count() != 2 {
		t.Errorf("TestDepthFirstSearch Count !=2, actual = %d", dfs.Count())
	}

	dfs = NewDepthFirstSearch(g, 9)
	if dfs.Count() != 4 {
		t.Errorf("TestDepthFirstSearch Count !=4, actual = %d", dfs.Count())
	}
}
