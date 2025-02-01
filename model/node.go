package model

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	gid "github.com/elemir/gloomo/id"
)

type Node struct {
	Draw     DrawFunc
	Position image.Point
	Size     image.Point
	ZIndex   int
}

type DrawFunc func(gid.ID, *ebiten.Image)
