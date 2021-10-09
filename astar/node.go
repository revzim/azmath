package astar

type (

	// node --
	// STORES PATHER DATA
	node struct {
		pather   Pather
		cost     float64
		priority float64
		parent   *node
		open     bool
		closed   bool
		index    int
	}

	// patherNodes --
	// MAPS PATHERS TO NODES
	patherNodes map[Pather]*node
)

// get --
// GETS PATHER NODE
func (nm patherNodes) get(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p] = n
	}
	return n
}
