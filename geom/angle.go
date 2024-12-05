package geom

import (
	"fmt"
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

func (a Angle) String() string {
	return fmt.Sprintf("%.f°", a*180/math.Pi)
}

func (a Angle) Abs() Angle {
	return Angle(math.Abs(float64(a)))
}

func (a Angle) Sin() float64 {
	return math.Sin(float64(a))
}

func (a Angle) Cos() float64 {
	return math.Cos(float64(a))
}
