package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image    *ebiten.Image
	Position image.Point
	ZIndex   int
}
