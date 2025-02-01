package model

import "github.com/hajimehoshi/ebiten/v2"

type Animation struct {
	Steps []*ebiten.Image
}

type AnimatedSprite struct {
	Animation Animation

	Current string
	Counter int
}
