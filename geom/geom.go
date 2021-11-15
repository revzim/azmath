package geom

import (
	"github.com/revzim/azmath"
	"github.com/revzim/azmath/vector2D"
)

type (
	geomtype byte

	Geometry interface {
		Type() geomtype
		// Area() float64
		// Perimeter() float64
		PointIntersection(x, y float64) bool
		Intersects(Geometry) (bool, bool)
		GetClosestPoint(Geometry) vector2D.Vector2D
		GetGeomXF() GeomXF
		// SquaredDist(Geometry) (float64, bool)
	}

	// GeomData --
	// ACCOMPANYING DATA FOR GEOMETRY TYPE
	GeomData struct {
		ID       string `json:"id"`
		GeomXF   `json:"xf"`
		key      string
		geomType geomtype
		Data     map[string]interface{} `json:"data"`
	}

	// GeomOpts --
	// HELPER STRUCT FOR GEOMDATA OPTIONS
	GeomDataOpts struct {
		ID string
		GeomXF
		Bounds
	}

	GeomXF struct {
		Origin vector2D.Vector2D `json:"origin"`
		Pos    vector2D.Vector2D `json:"pos"`
		Bounds Bounds            `json:"bounds,omitempty"`
		Radius float64           `json:"radius,omitempty"`
		Width  float64           `json:"width,omitempty"`
		Length float64           `json:"length,omitempty"`
		// Height float64 `json:"height,omitempty"`
	}

	Bound struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	}

	Bounds struct {
		X Bound `json:"x"`
		Y Bound `json:"y"`
		// Z bound `json:"z"`
	}
)

const (
	SquareKey = "square"
	CircleKey = "circle"
)

var (
	Types = map[string]geomtype{
		CircleKey: 0x00,
		SquareKey: 0x01,
	}
)

func NewGeomData(key string, xf GeomXF) GeomData {
	return GeomData{
		key:      key,
		geomType: Types[key],
		GeomXF:   xf, // Size{},
	}
}

func NewBounds(x, y Bound) Bounds {
	return Bounds{
		X: x,
		Y: y,
	}
}

func Intersects(g, g2 Geometry) (bool, bool) {
	if g.Type() == g2.Type() {
		return g.Intersects(g2)
		// g2Geom := g2.GetGeomXF()
		// return g.PointIntersection(g2Geom.Pos.X, g2Geom.Pos.Y), true
	} else {
		return geomIntersect(g, g2)
	}
}

// GeomIntersect --
// DIFF GEOM INTERSECTION
func geomIntersect(g, g2 Geometry) (bool, bool) {
	var closestPtToCircleCenter vector2D.Vector2D
	var circleXF GeomXF
	switch g.Type() {
	case Types[SquareKey]:
		closestPtToCircleCenter = g.GetClosestPoint(g2)
		circleXF = g2.GetGeomXF()
	case Types[CircleKey]:
		closestPtToCircleCenter = g2.GetClosestPoint(g)
		circleXF = g.GetGeomXF()
	default:
		return false, false
	}

	squaredDist := azmath.SquaredDist(closestPtToCircleCenter.X, circleXF.Pos.X, closestPtToCircleCenter.Y, circleXF.Pos.Y)
	return squaredDist < circleXF.Radius*circleXF.Radius, true
}
