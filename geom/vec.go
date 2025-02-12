package geom

import (
	"fmt"
	"image"
	"math"
)

type Vec2 [2]float64

func FromPoint(pt image.Point) Vec2 {
	return Vec2{
		float64(pt.X),
		float64(pt.Y),
	}
}

func (v Vec2) Mul(k float64) Vec2 {
	return Vec2{
		v[0] * k,
		v[1] * k,
	}
}

func (v Vec2) Div(k float64) Vec2 {
	return Vec2{
		v[0] / k,
		v[1] / k,
	}
}

func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{
		v[0] + u[0],
		v[1] + u[1],
	}
}

func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{
		v[0] - u[0],
		v[1] - u[1],
	}
}

func (v Vec2) Normalize() Vec2 {
	if v.Length() == 0 {
		return v
	}

	return v.Mul(1 / v.Length())
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec2) Distance(u Vec2) float64 {
	return v.Sub(u).Length()
}

func (v Vec2) Round() image.Point {
	return image.Pt(int(math.Round(v[0])), int(math.Round(v[1])))
}

func (v Vec2) Dot(u Vec2) float64 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec2) Cross(u Vec2) float64 {
	return v[0]*u[1] - v[1]*u[0]
}

func (v Vec2) Rotate(angle Angle) Vec2 {
	return Vec2{
		angle.Cos()*v[0] - angle.Sin()*v[1],
		angle.Sin()*v[0] + angle.Cos()*v[1],
	}
}

func (v Vec2) Angle() Angle {
	if v.Length() == 0 {
		return 0
	}

	return Angle(math.Atan2(v[1], v[0]))
}

func (v Vec2) Unpack() (float64, float64) {
	return v[0], v[1]
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", v[0], v[1])
}
