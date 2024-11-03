package geom

import (
	"math"
)

type Angle float64

func (a Angle) Normalize() Angle {
	rad := math.Remainder(float64(a), 2*math.Pi)
	if rad <= -math.Pi {
		rad = math.Pi
	}

	return Angle(rad)
}

func (a Angle) Abs() Angle {
	return Angle(math.Abs(float64(a)))
}
