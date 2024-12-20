package gloomo

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Camera is a special object for rendering.
type Camera struct {
	sceneSize image.Point
	bounds    image.Rectangle
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c Camera) Bounds() image.Rectangle {
	return c.bounds
}

func (c Camera) ViewPort(screen *ebiten.Image, mul float64) ViewPort {
	return ViewPort{
		img:    screen,
		bounds: c.bounds,
		mul:    mul,
	}
}

func (c *Camera) Move(shift image.Point) {
	pos := c.GetPosition()

	c.SetPosition(pos.Add(shift))
}

func (c *Camera) GetPosition() image.Point {
	return c.bounds.Max.Add(c.bounds.Min).Div(2)
}

func (c *Camera) SetPosition(pos image.Point) {
	x, y := pos.X, pos.Y
	w, h := c.bounds.Dx(), c.bounds.Dy()

	if x < w/2 {
		x = w / 2
	}

	if y < h/2 {
		y = h / 2
	}

	if x+w/2 > c.sceneSize.X {
		x = c.sceneSize.X - w/2
	}

	if y+h/2 > c.sceneSize.Y {
		y = c.sceneSize.Y - h/2
	}

	c.bounds = image.Rect(x-w/2, y-h/2, x+w/2, y+h/2)
}

func (c *Camera) SetSize(w, h int) {
	x, y := c.bounds.Min.X, c.bounds.Min.Y

	c.bounds = image.Rect(x, y, x+w, x+h)
}

func (c *Camera) RealMousePosition() image.Point {
	x, y := ebiten.CursorPosition()

	return c.bounds.Min.Add(image.Pt(x, y))
}
