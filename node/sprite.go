package node

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is a basic useful render-time aggregate that incapsulate an image rendering.
type Sprite struct {
	Image    *ebiten.Image
	Position image.Point
	ZIndex   int
}
