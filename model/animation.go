package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/geom"
)

type Animation struct {
	Steps  []*ebiten.Image
	Mirror bool
}

type AnimationSheet struct {
	Size       image.Point
	Animations map[string]Animation
}

type AnimatedSprite struct {
	AnimationSheet *AnimationSheet
	Position       geom.Vec2
	ZIndex         int

	Current string
	Counter int
}
