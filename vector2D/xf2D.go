package vector2D

import "math"

type (
	XF struct {
		*Vector2D
		Angle float64 `json:"angle"`
	}
)

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
