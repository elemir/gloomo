package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/node"
)

type RectRepo interface {
	Get(id gid.ID) (node.Node, bool)
}

func Rect(repo RectRepo) node.DrawFunc {
	return func(id gid.ID, screen *ebiten.Image) {
		node, ok := repo.Get(id)
		if !ok {
			return
		}

		vector.DrawFilledRect(screen, float32(node.Position.X), float32(node.Position.Y),
			float32(node.Size.X), float32(node.Size.Y), color.White, false)
	}
}
