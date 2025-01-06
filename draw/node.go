package draw

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/id"
)

type NodeRepo interface{}

type Node struct {
	Draw     DrawFunc
	Position image.Point
	Size     image.Point
}

type DrawFunc func(id.ID, *ebiten.Image)
