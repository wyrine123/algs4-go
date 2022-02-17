package graph

import "testing"

func TestDepthFirstSearch_Marked(t *testing.T) {
	g := genGraph()
	d := NewDepthFirstSearch(*g, 0)

	if d.Marked(7) == true {
		t.Errorf("TestDepthFirstSearch_Marked d.Marked(7) should: %t, actual: %t", false, true)
	}
	if d.Marked(1) == false {
		t.Errorf("TestDepthFirstSearch_Marked d.Marked(7) should: %t, actual: %t", true, false)
	}
	if d.Count() != 7 {
		t.Errorf("TestDepthFirstSearch_Marked d.Count() should: %d, actual: %d", 7, d.Count())
	}
}
