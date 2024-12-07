package gloomo

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image *ebiten.Image

	angle    float64
	position image.Point
}

func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{
		image: img,
	}
}

func (s *Sprite) SetAngle(angle float64) {
	s.angle = angle
}

func (s *Sprite) SetPosition(pos image.Point) {
	s.position = pos
}

func (s Sprite) Bounds() image.Rectangle {
	return s.image.Bounds().Add(s.position)
}

func (s Sprite) Draw(screen ViewPort) {
	var op ebiten.DrawImageOptions

	op.GeoM.Translate(-float64(s.image.Bounds().Dx())/2, -float64(s.image.Bounds().Dx())/2)
	op.GeoM.Rotate(float64(s.angle))
	op.GeoM.Translate(math.Round(float64(s.position.X)), math.Round(float64(s.position.Y)))
	op.GeoM.Scale(1, 1)

	screen.DrawImage(s.image, &op)
}
