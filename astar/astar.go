package astar

import "container/heap"

// A* PATHFINDNIG IMPL
// INSPIRED BY github.com/beefstack/go-astar

type (

	// Pather --
	// A* PATHER INTERFACE
	Pather interface {
		// Neighbors --
		// RETURNS DIRECT NEIGHBORS FOR PATH
		Neighbors() []Pather
		// NeighborCost --
		// RETURNS MOVE COST
		NeighborCost(destination Pather) float64
		// EstimatedCost --
		// ESTIMATE MOVE COSTS BTWN NON-ADJACENT NODES
		EstimatedCost(destination Pather) float64
	}
)

// Path --
// RETURNS SHORT PATH, DIST, TRUE BTWN PATHER NODES
// RESULTING IN FIXED: XI -> XF INSTEAD OF XF -> XI
func Path(init, destination Pather) (path []Pather, distance float64, found bool) {
	nm := patherNodes{}
	nq := &prioQueue{}
	heap.Init(nq)
	destinationNode := nm.get(destination)
	destinationNode.open = true
	heap.Push(nq, destinationNode)
	for {
		if nq.Len() == 0 {
			// NO PATH
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(init) {
			// FOUND PATH LOOP & APPEND
			p := []Pather{}
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		// LOOP & WEIGH AGAINST COST
		for _, neighbor := range current.pather.Neighbors() {
			cost := current.cost + current.pather.NeighborCost(neighbor)
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.priority = cost + neighbor.EstimatedCost(init)
				neighborNode.parent = current
				heap.Push(nq, neighborNode)
			}
		}
	}
}
