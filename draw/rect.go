package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	gid "github.com/elemir/gloomo/id"
)

type Rect struct{}

func (r Rect) Draw(id gid.ID, img *ebiten.Image) {
	vector.DrawFilledRect(img, 0, 0, 64, 64, color.White, false)
}
