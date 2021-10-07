package vector

import (
	"math"

	"github.com/revzim/azmath"
)

type (
	Vector struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}
)

func New(n ...float64) *Vector {
	if len(n) == 1 {
		return &Vector{n[0], n[0], n[0]}
	} else if len(n) == 3 {
		return &Vector{n[0], n[1], n[2]}
	}
	return nil
}

var (
	ZeroVector     = New(0)
	OneVector      = New(1)
	UpVector       = New(0, 0, 1)
	DownVector     = New(0, 0, -1)
	ForwardVector  = New(1, 0, 0)
	BackwardVector = New(-1, 0, 0)
	RightVector    = New(0, 1, 0)
	LeftVector     = New(0, -1, 0)
	XAxisVector    = New(1, 0, 0)
	YAxisVector    = New(0, 1, 0)
	ZAxisVector    = New(0, 0, 1)
)

func (v *Vector) Add(v2 Vector) *Vector {
	tmpVec := New(
		v.X+v2.X,
		v.Y+v2.Y,
		v.Z+v2.Z,
	)
	return tmpVec
}

func (v *Vector) Subtract(v2 Vector) *Vector {
	tmpVec := New(
		v.X-v2.X,
		v.Y-v2.Y,
		v.Z-v2.Z,
	)
	return tmpVec
}

func (v *Vector) Scale(v2 Vector) *Vector {
	tmpVec := New(
		v.X*v2.X,
		v.Y*v2.Y,
		v.Z*v2.Z,
	)
	return tmpVec
}

func (v *Vector) Divide(v2 Vector) *Vector {
	tmpVec := New(
		v.X/v2.X,
		v.Y/v2.Y,
		v.Z/v2.Z,
	)
	return tmpVec
}

func (v *Vector) AddConst(n float64) *Vector {
	tmpVec := New(
		v.X+n,
		v.Y+n,
		v.Z+n,
	)
	return tmpVec
}

func (v *Vector) SubtractConst(n float64) *Vector {
	tmpVec := New(
		v.X-n,
		v.Y-n,
		v.Z-n,
	)
	return tmpVec
}

func (v *Vector) ScaleConst(n float64) *Vector {
	tmpVec := New(
		v.X*n,
		v.Y*n,
		v.Z*n,
	)
	return tmpVec
}

func (v *Vector) DivideConst(n float64) *Vector {
	if n == 0 {
		return ZeroVector
	}
	rScale := 1.0 / n
	tmpVec := New(
		v.X*rScale,
		v.Y*rScale,
		v.Z*rScale,
	)
	return tmpVec
}

func (v *Vector) CrossProduct(v2 Vector) *Vector {
	tmpVec := New(
		v.Y*v2.Z-v.Z*v2.Y,
		v.Z*v2.X-v.X*v2.Z,
		v.X*v2.Y-v.Y*v2.X,
	)
	return tmpVec
}

func (v *Vector) DotProduct(v2 Vector) float64 {
	ans := v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
	return ans
}

func (v *Vector) Equals(v2 Vector) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

func (v *Vector) NearlyEquals(v2 Vector, tolerance float64) bool {
	return math.Abs(v.X-v2.X) <= tolerance && math.Abs(v.Y-v2.Y) <= tolerance && math.Abs(v.Z-v2.Z) <= tolerance
}

func (v *Vector) NotEqual(v2 Vector) bool {
	return v.X != v2.X || v.Y != v2.Y || v.Z != v2.Z
}

func (v *Vector) Negative() *Vector {
	tmpVec := New(-v.X, -v.Y, -v.Z)
	return tmpVec
}

func (v *Vector) Size() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) SizeSquare() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector) Size2D() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

func (v *Vector) NearlyZero(tolerance float64) bool {
	return math.Abs(v.X) <= tolerance && math.Abs(v.Y) <= tolerance && math.Abs(v.Z) <= tolerance
}

func (v *Vector) Normalize(tolerance float64) bool {
	squareSum := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if squareSum > tolerance {
		scale := azmath.FastInvSqrt64(squareSum)
		v.X *= scale
		v.Y *= scale
		v.Z *= scale
		return true
	}
	return false
}

func (v *Vector) Dist(v2 Vector) float64 {
	return math.Sqrt(v.DistSquared(v2))
}

func (v *Vector) DistSquared(v2 Vector) float64 {
	return math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2) + math.Pow(v2.Z-v.Z, 2)
}

func (v *Vector) DistXY(v2 Vector) float64 {
	return math.Sqrt(v.DistSquaredXY(v2))
}

func (v *Vector) DistSquaredXY(v2 Vector) float64 {
	return math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2)
}
