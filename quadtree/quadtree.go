package quadtree

/**
 * JAVA IMPL OF QUADTREE
 * https://gamedevelopment.tutsplus.com/tutorials/quick-tip-use-quadtrees-to-detect-likely-collisions-in-2d-space--gamedev-374
 * 'A quadtree is a data structure used to divide a 2D region into more manageable parts.
 * It's an extended binary tree, but instead of two child nodes it has four'
 *
 * DEBATED ON NIL/NON-NIL SLICE WENT WITH NON-NIL DUE TO JSON ENCODING
 */

type (
	Object interface {
		GetObject() interface{}
		GetBounds() *Bounds
	}

	Quadtree struct {
		MaxObjects int         `json:"max_objects"`
		MaxLevels  int         `json:"max_levels"`
		Level      int         `json:"level"`
		Count      int         `json:"count"`
		Bounds     *Bounds     `json:"bounds"`
		Objects    []*Object   `json:"objects"`
		Tree       []*Quadtree `json:"tree"`
	}
)

const (
	RightTop = iota
	LeftTop
	LeftBot
	RightBot
)

func New(level int, bounds *Bounds) *Quadtree {
	return &Quadtree{
		Level:   level,
		Bounds:  bounds,
		Objects: []*Object{},
		Tree:    []*Quadtree{}, // NewDefaultTree(),
	}
}

// NewDefaultTree --
// CREATES 4 TREE SUBNODES
func NewDefaultTree(args ...int) []*Quadtree {
	cnt := 4
	if len(args) == 1 {
		cnt = args[0]
	}
	return make([]*Quadtree, cnt)
}

// Insert --
// INSERT
func (qt *Quadtree) Insert(object Object) {
	qt.Count++
	objectBounds := object.GetBounds()
	if len(qt.Tree) > 0 {
		index := qt.getIndex(objectBounds)
		if index != -1 {
			qt.Tree[index].Insert(object)
			return
		}
	}

	qt.Objects = append(qt.Objects, &object)

	shouldHandleSubNodes := len(qt.Objects) > qt.MaxObjects && qt.Level < qt.MaxLevels
	if shouldHandleSubNodes {
		if len(qt.Tree) <= 0 {
			qt.split()
		}
		i := 0
		for range qt.Objects {
			index := qt.getIndex((*qt.Objects[i]).(Object).GetBounds())
			if index != -1 {
				tmpObject := *qt.Objects[i]                              // TMP OBJ
				qt.Objects = append(qt.Objects[:i], qt.Objects[i+1:]...) // REMOVE OBJ
				qt.Tree[index].Insert(tmpObject)
			} else {
				i++
			}
		}
	}
}

// Get --
// GET ALL OBJECTS THAT *COULD* COLLIDE
func (qt *Quadtree) getPossibleObjects(object Object) []*Object {
	objects := qt.Objects
	index := qt.getIndex(object.GetBounds())
	if len(qt.Tree) > 0 {
		if index != -1 {
			objects = append(objects, qt.Tree[index].getPossibleObjects(object)...)
		} else {
			for i := range qt.Tree {
				objects = append(objects, qt.Tree[i].getPossibleObjects(object)...)
			}
		}
	}

	return objects
}

// Intersects --
// TESTS FOR SIMPLE BOUNDS INTERSECTION
func (qt *Quadtree) Intersects(object Object) []*Object {
	var intersections []*Object
	possibleCollisions := qt.getPossibleObjects(object)
	for i := range possibleCollisions {
		if (*possibleCollisions[i]).GetBounds().Collides(object.GetBounds()) {
			intersections = append(intersections, possibleCollisions[i])
		}
	}
	return intersections
}

// GetAllIntersectingObjects --
// RETURNS ALL INTERSECTING OBJECTS
func (qt *Quadtree) GetAllIntersectingObjects(object Object) []*Object {
	var objects []*Object
	arr := qt.getPossibleObjects(object)
	for i := range arr {
		currObj := (*arr[i])
		currBounds := currObj.GetBounds()
		if currBounds.Collides(object.GetBounds()) {
			// log.Println(currObj.GetObject())
			objects = append(objects, arr[i])
		}

	}
	return objects
}

// GetAllIntersectingObjects --
// CALLBACK FOR EACH INTERSECTING OBJECT
func (qt *Quadtree) GetAllIntersectingObjectsCB(object Object, cb func(object *Object)) {
	arr := qt.getPossibleObjects(object)
	for i := range arr {
		currObj := (*arr[i])
		currBounds := currObj.GetBounds()
		if currBounds.Collides(object.GetBounds()) {
			cb(arr[i])
		}
	}
}

// Clear --
// CLEARS OBJECTS, TREE, & RESETS COUNT
func (qt *Quadtree) Clear() {
	qt.Objects = nil
	for i := range qt.Tree {
		if qt.Tree[i] != nil {
			qt.Tree[i].Clear()
			qt.Tree[i] = nil
		}
	}
	qt.Tree = NewDefaultTree()
	qt.Count = 0
}

func (qt *Quadtree) split() {
	if len(qt.Tree) != 4 {
		width := qt.Bounds.Width / 2
		height := qt.Bounds.Height / 2
		x := qt.Bounds.X
		y := qt.Bounds.Y
		qt.populateSplit(x, y, width, height)
	}
}

func (qt *Quadtree) populateSplit(x, y, width, height float64) {
	lvl := qt.Level + 1
	qt.Tree[0] = New(lvl, NewBounds(x+width, y, width, height))
	qt.Tree[1] = New(lvl, NewBounds(x, y, width, height))
	qt.Tree[2] = New(lvl, NewBounds(x, y+height, width, height))
	qt.Tree[3] = New(lvl, NewBounds(x+width, y+height, width, height))
}

func (qt *Quadtree) getIndex(b *Bounds) int {
	verticalMidpt := b.X + (b.Width / 2)
	horizontalMidpt := b.Y + (b.Height / 2)

	topQuadrant := b.Y < horizontalMidpt && b.Y+b.Height < horizontalMidpt
	botQuadrant := b.Y > horizontalMidpt

	objectCanFitInLeftQuad := b.X < verticalMidpt && b.X+b.Width < verticalMidpt
	objectCanFitInRightQuad := b.X > verticalMidpt
	if objectCanFitInLeftQuad {
		if topQuadrant {
			return LeftTop
		} else if botQuadrant {
			return LeftBot
		}
	} else if objectCanFitInRightQuad {
		if topQuadrant {
			return RightTop
		} else if botQuadrant {
			return RightBot
		}
	}

	return -1
}
