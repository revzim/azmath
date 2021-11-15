package geom

import (
	"github.com/revzim/azmath"
	"github.com/revzim/azmath/vector2D"
)

type (
	Circle struct {
		GeomData `json:"data"`
	}
)

func NewCircle(opts GeomDataOpts) *Circle {
	return &Circle{
		GeomData: NewGeomData(CircleKey, opts.GeomXF),
	}
}

func (c *Circle) Intersects(g Geometry) (bool, bool) {
	if g.Type() == Types[CircleKey] {
		c2 := g.GetGeomXF()
		squaredDist := SquaredDist(c.Pos.X, c2.Pos.X, c.Pos.Y, c2.Pos.Y)
		return squaredDist < ((c.Radius * c.Radius) + (c2.Radius * c2.Radius)), true
		// if sqrdDist, ok := SquaredDist(c.Pos.X, c2.Pos.X, c.Pos.Y, c2.Pos.Y); ok {
		// 	c2 := g.GetGeomXF()
		// 	return sqrdDist < ((c.Radius * c.Radius) + (c2.Radius * c2.Radius)), true
		// }
	}
	return false, false
}

func (c *Circle) PointIntersection(x, y float64) bool {
	squaredDist := (x-c.Pos.X)*(x-c.Pos.X) + (y-c.Pos.Y)*(y-c.Pos.Y)
	return squaredDist < (c.Radius * c.Radius)
}

func (c *Circle) GetGeomXF() GeomXF {
	return c.GeomXF
}

func (c *Circle) GetClosestPoint(g Geometry) vector2D.Vector2D {
	b := g.GetGeomXF()
	return *vector2D.New(azmath.Clamp3D(b.Bounds.X.Min, c.Pos.X, b.Bounds.X.Max), azmath.Clamp3D(b.Bounds.Y.Min, c.Pos.Y, b.Bounds.Y.Max))
}

func (c *Circle) Type() geomtype {
	return c.geomType
}
