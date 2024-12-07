package gloomo

import (
	"image"
	"image/color"
	"math/rand/v2"
)

type Tracer struct {
	trace []image.Point
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

func (s *Tracer) SetPosition(pos image.Point) {
	/*	if len(s.trace) > 100 {
			s.trace = s.trace[len(s.trace)-100:]
		}
	*/
	s.trace = append(s.trace, pos)
}

func (s Tracer) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1000, 1000)
}

func (s Tracer) Draw(screen ViewPort) {
	for _, pos := range s.trace {
		screen.Set(pos, s.clr)
	}
}
