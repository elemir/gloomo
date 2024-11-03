package gloomo

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/geom"
)

type Sprite struct {
	image *ebiten.Image

	angle    geom.Angle
	position geom.Vec2
}

func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{
		image: img,
	}
}

func (s *Sprite) SetAngle(angle geom.Angle) {
	s.angle = angle
}

func (s *Sprite) SetPosition(pos geom.Vec2) {
	s.position = pos
}

func (s Sprite) Bounds() geom.Rectangle {
	return geom.FromRectangle(s.image.Bounds()).Add(s.position)
}

func (s Sprite) Draw(screen ViewPort) {
	var op ebiten.DrawImageOptions

	op.GeoM.Translate(-float64(s.image.Bounds().Dx())/2, -float64(s.image.Bounds().Dx())/2)
	op.GeoM.Rotate(float64(s.angle))
	op.GeoM.Translate(math.Round(s.position[0]), math.Round(s.position[1]))
	op.GeoM.Scale(1, 1)

	screen.DrawImage(s.image, &op)
}
