package astar

type (
	// prioQueue --
	// IMPL heap.Interface & HOLDS NODES
	prioQueue []*node
)

func (pq prioQueue) Len() int {
	return len(pq)
}

func (pq prioQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq prioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *prioQueue) Push(x interface{}) {
	n := len(*pq)
	no := x.(*node)
	no.index = n
	*pq = append(*pq, no)
}

func (pq *prioQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	no := old[n-1]
	no.index = -1
	*pq = old[0 : n-1]
	return no
}
