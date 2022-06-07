package chapter4

import (
	"os"
	"testing"
)

func TestDepthFirstPaths_PathTo(t *testing.T) {
	f, err := os.Open("./data/tinyG.txt")
	if err != nil {
		t.Errorf("TestNewGraphFromFile open file 失败: %s", err.Error())
	}
	g := NewGraphFromFile(f)

	path := NewDepthFirstPaths(g, 0)

	if !path.HasPathTo(1) {
		t.Errorf("TestDepthFirstPaths_PathTo 到节点1有路径")
	}
	if path.HasPathTo(7) {
		t.Errorf("TestDepthFirstPaths_PathTo 到节点7无路径")
	}

	t.Logf("TestDepthFirstPaths_PathTom 节点0到节点4路径: %v", path.PathTo(4))
	t.Logf("TestDepthFirstPaths_PathTom 节点0到节点1路径: %v", path.PathTo(1))

	path = NewDepthFirstPaths(g, 9)
	if !path.HasPathTo(12) {
		t.Errorf("TestDepthFirstPaths_PathTo 到节点12有路径")
	}

	if path.HasPathTo(6) {
		t.Errorf("TestDepthFirstPaths_PathTo 到节点6无路径")
	}

	t.Logf("TestDepthFirstPaths_PathTom 节点9到节点12路径: %v", path.PathTo(12))
	t.Logf("TestDepthFirstPaths_PathTom 节点9到节点11路径: %v", path.PathTo(11))
}
