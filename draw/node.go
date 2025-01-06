package draw

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	gid "github.com/elemir/gloomo/id"
)

type NodeRepo interface {
	Get(id gid.ID) (Node, bool)
}

type Node struct {
	Draw     DrawFunc
	Position image.Point
	Size     image.Point
}

type DrawFunc func(gid.ID, *ebiten.Image)
