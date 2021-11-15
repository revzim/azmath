package geom

import (
	"github.com/revzim/azmath"
	"github.com/revzim/azmath/vector2D"
)

type (
	Square struct {
		GeomData `json:"data"`
	}
)

func NewSquare(opts GeomDataOpts) *Square {
	return &Square{
		GeomData: NewGeomData(SquareKey, opts.GeomXF),
	}
}

// func (s *Square) Intersects(s2 *Square) (bool, bool) {
// 	return (s.Bounds.X.Min <= s2.Bounds.X.Max && s.Bounds.X.Max >= s2.Bounds.X.Min &&
// 		s.Bounds.Y.Min <= s2.Bounds.Y.Max && s.Bounds.Y.Max >= s2.Bounds.Y.Min), true
// }

func (s *Square) Intersects(g Geometry) (bool, bool) {
	if g.Type() == Types[SquareKey] {
		s2 := g.GetGeomXF()
		sPosX := s.Pos.X
		sPosY := s.Pos.Y
		s2PosX := s2.Pos.X
		s2PosY := s2.Pos.Y
		return ((s.Bounds.X.Min+sPosX) <= (s2.Bounds.X.Max+s2PosX) && (s.Bounds.X.Max+sPosX) >= (s2.Bounds.X.Min+s2PosX) &&
			(s.Bounds.Y.Min+sPosY) <= (s2.Bounds.Y.Max+s2PosY) && (s.Bounds.Y.Max+sPosY) >= (s2.Bounds.Y.Min+s2PosY)), true
	}
	return false, false

}

func (s *Square) PointIntersection(x, y float64) bool {
	return (x >= (s.Bounds.X.Min+s.Pos.X) && x <= (s.Bounds.X.Max+s.Pos.X) &&
		y >= (s.Bounds.Y.Min+s.Pos.Y) && y <= (s.Bounds.Y.Max+s.Pos.Y))
}

func (s *Square) GetClosestPoint(g Geometry) vector2D.Vector2D {
	c := g.GetGeomXF()
	return *vector2D.New(azmath.Clamp3D(s.Bounds.X.Min, c.Pos.X, s.Bounds.X.Max), azmath.Clamp3D(s.Bounds.Y.Min, c.Pos.Y, s.Bounds.Y.Max))
}

func (s *Square) GetGeomXF() GeomXF {
	return s.GeomXF
}

func (s *Square) Type() geomtype {
	return s.geomType
}
