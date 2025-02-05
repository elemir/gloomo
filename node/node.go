package node

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	gid "github.com/elemir/gloomo/id"
)

// Node is a basic render-time component. All render-time repositories should create a node inside.
type Node struct {
	Draw     DrawFunc
	Position image.Point
	Size     image.Point
	ZIndex   int
}

// DrawFunc is a callback for drawing specific node type.
type DrawFunc func(gid.ID, *ebiten.Image)
