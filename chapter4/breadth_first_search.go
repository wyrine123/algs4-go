package chapter4

type BreadthFirstSearch struct {
	countV  int
	markedV []bool
	queue   *QueueLinkList
}

func NewBreadthFirstSearch(g *Graph, v int) *BreadthFirstSearch {
	s := new(BreadthFirstSearch)
	s.markedV = make([]bool, g.V())
	s.queue = NewQueueLinkList()
	s.bfs(g, v)

	return s
}

func (s *BreadthFirstSearch) bfs(g *Graph, v int) {
	s.countV++
	s.queue.Enqueue(v)
	s.markedV[v] = true
	for w := s.queue.Dequeue(); w != nil; w = s.queue.Dequeue() {
		for _, item := range g.Adj(*w) {
			if !s.Marked(item) {
				s.markedV[item] = true
				s.countV++
				s.queue.Enqueue(item)
			}

		}
	}
}

func (s *BreadthFirstSearch) Count() int {
	return s.countV
}

func (s *BreadthFirstSearch) Marked(v int) bool {
	return s.markedV[v]
}
