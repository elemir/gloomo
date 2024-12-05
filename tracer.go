package gloomo

import (
	"image/color"
	"math/rand/v2"

	"github.com/elemir/gloomo/geom"
)

type Tracer struct {
	trace []geom.Vec2
	clr   color.Color
}

func NewTracer() *Tracer {
	clr := color.RGBA{
		uint8(rand.Uint() % 255),
		uint8(rand.Uint() % 255),
		uint8(rand.Uint() % 255),
		255,
	}

	return &Tracer{
		clr: clr,
	}
}

func (s *Tracer) SetPosition(pos geom.Vec2) {
	/*	if len(s.trace) > 100 {
			s.trace = s.trace[len(s.trace)-100:]
		}
	*/
	s.trace = append(s.trace, pos)
}

func (s Tracer) Bounds() geom.Rectangle {
	return geom.Rect(0, 0, 1000, 1000)
}

func (s Tracer) Draw(screen ViewPort) {
	for _, pos := range s.trace {
		screen.Set(pos, s.clr)
	}
}
