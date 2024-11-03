package gloomo

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/geom"
)

// Camera is a special object for rendering.
type Camera struct {
	sceneSize geom.Vec2
	bounds    geom.Rectangle
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c Camera) Bounds() geom.Rectangle {
	return c.bounds
}

func (c Camera) ViewPort(screen *ebiten.Image, mul float64) ViewPort {
	return ViewPort{
		img:    screen,
		bounds: c.bounds.Round(),
		mul:    mul,
	}
}

func (c *Camera) Move(shift geom.Vec2) {
	c.bounds = c.bounds.Add(shift)
}

func (c *Camera) SetPosition(pos geom.Vec2) {
	x, y := pos.Unpack()
	w, h := c.bounds.Size().Unpack()

	if x < w/2 {
		x = w / 2
	}

	if y < h/2 {
		y = h / 2
	}

	if x+w/2 > c.sceneSize[0] {
		x = c.sceneSize[0] - w/2
	}

	if y+h/2 > c.sceneSize[1] {
		y = c.sceneSize[1] - h/2
	}

	c.bounds = geom.Rect(x-w/2, y-h/2, w, h)
}

func (c *Camera) SetSize(w, h float64) {
	x, y := c.bounds.Min[0], c.bounds.Min[1]

	c.bounds = geom.Rect(x, y, w, h)
}

func (c *Camera) RealMousePosition() geom.Vec2 {
	x, y := ebiten.CursorPosition()

	return c.bounds.Min.Add(geom.FromPoint(image.Pt(x, y)))
}
