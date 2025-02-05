package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Size  image.Point
	Steps map[string][]*ebiten.Image
}

type AnimatedSprite struct {
	Animation *Animation
	Position  image.Point
	ZIndex    int

	Current string
	Counter int
}
