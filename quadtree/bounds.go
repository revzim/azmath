package quadtree

import "github.com/revzim/azmath/vector2D"

type (
	Bounds struct {
		vector2D.Vector2D
		Radius float64 `json:"radius"`
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	}
	max struct {
		x float64
		y float64
	}
)

func NewVecBounds(vec *vector2D.Vector2D, width, height float64) *Bounds {
	return &Bounds{
		Vector2D: *vec,
		Width:    width,
		Height:   height,
	}
}

func NewBounds(x, y, width, height float64) *Bounds {
	return &Bounds{
		Vector2D: *vector2D.New(x, y),
		Width:    width,
		Height:   height,
	}
}

func (b *Bounds) Collides(oB *Bounds) bool {
	bMax := max{
		x: b.X + b.Width,
		y: b.Y + b.Height,
	}

	oBMax := max{
		x: oB.X + oB.Width,
		y: oB.Y + oB.Height,
	}

	if bMax.x < oB.X || b.X > oBMax.x || bMax.y < oB.Y || b.Y > oBMax.y {
		return false
	}

	return true
}
