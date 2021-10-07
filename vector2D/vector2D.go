package vector2D

import "math"

type (
	Vector2D struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}

	XF struct {
		*Vector2D
		Angle float64 `json:"angle"`
	}
)

const (
	fRounding = 10000
)

// New
func New(n ...float64) *Vector2D {
	if len(n) == 2 {
		return &Vector2D{n[0], n[1]}
	} else if len(n) == 1 {
		return &Vector2D{n[0], n[0]}
	}
	return nil
}

func NewXF(x, y, angle float64) *XF {
	xf := &XF{Angle: math.Round((math.Mod(angle*(180/math.Pi), 360.0) * fRounding)) / fRounding}
	xf.Vector2D = New(math.Round(x*fRounding)/fRounding, math.Round(y*fRounding)/fRounding)
	return xf
}

func (xf *XF) Compare(xf2 *XF, hasUpdated chan bool) {
	if xf.X != xf2.X || xf.Y != xf2.Y || xf.Angle != xf2.Angle {
		hasUpdated <- true
	} else {
		hasUpdated <- false
	}
}

var (
	ZeroVector = New(0)
	UnitVector = New(1)
	Unit45Deg  = New(.77, .707) // sqrt(0.5)
)

func (v *Vector2D) Add(v2 Vector2D) *Vector2D {
	tmpVec := New(v.X+v2.X, v.Y+v2.Y)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) Subtract(v2 Vector2D) *Vector2D {
	tmpVec := New(v.X-v2.X, v.Y-v2.Y)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) ScaleVec(v2 Vector2D) *Vector2D {
	tmpVec := New(v.X*v2.X, v.Y*v2.Y)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) RScaleVec(v2 Vector2D) *Vector2D {
	tmpVec := New(v.X/v2.X, v.Y/v2.Y)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) Scale(n float64) *Vector2D {
	tmpVec := New(v.X*n, v.Y*n)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) RScale(n float64) *Vector2D {
	if n == 0 {
		return ZeroVector
	}
	rScale := 1.0 / n
	tmpVec := New(v.X*rScale, v.Y*rScale)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) AddConst(n float64) *Vector2D {
	tmpVec := New(v.X+n, v.Y+n)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) SubtractConst(n float64) *Vector2D {
	tmpVec := New(v.X-n, v.Y-n)
	// v = tmpVec
	return tmpVec
}

func (v *Vector2D) DotProduct(v2 Vector2D) float64 {
	ans := v.X*v2.X + v.Y*v2.Y
	// v = tmpVec
	return ans
}

func (v *Vector2D) CrossProduct(v2 Vector2D) float64 {
	ans := v.X*v2.X - v.Y*v2.Y
	// v = tmpVec
	return ans
}

func (v *Vector2D) DistSquared(v2 Vector2D) float64 {
	return math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2)
}

func (v *Vector2D) Dist(v2 Vector2D) float64 {
	return math.Sqrt(v.DistSquared(v2))
}

func (v *Vector2D) Max(v2 Vector2D) *Vector2D {
	return New(math.Max(v.X, v2.X), math.Max(v.Y, v2.Y))
}

func (v *Vector2D) Min(v2 Vector2D) *Vector2D {
	return New(math.Min(v.X, v2.X), math.Min(v.Y, v2.Y))
}

func (v *Vector2D) Equal(v2 Vector2D) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v *Vector2D) Equals(v2 *Vector2D, tolerance float64) bool {
	return math.Abs(v.X-v2.X) <= tolerance && math.Abs(v.Y-v2.Y) <= tolerance
}

func (v *Vector2D) NotEqual(v2 Vector2D) bool {
	return v.X != v2.X || v.Y != v2.Y
}

func (v *Vector2D) Size() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2D) SizeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector2D) IsZero() bool {
	return v.X == 0.0 && v.Y == 0.0
}

func (v *Vector2D) IsNearlyZero(tolerance float64) bool {
	return math.Abs(v.X) <= tolerance && math.Abs(v.Y) <= tolerance
}

func (v *Vector2D) Lerp(v2 *Vector2D, tolerance float64) *Vector2D {
	x := v.X + tolerance*(v2.X-v.X)
	y := v.Y + tolerance*(v2.Y-v.Y)
	return New(x, y)
}
