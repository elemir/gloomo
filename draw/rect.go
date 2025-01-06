package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	gid "github.com/elemir/gloomo/id"
)

func Rect(repo NodeRepo) Func {
	return func(id gid.ID, img *ebiten.Image) {
		node, ok := repo.Get(id)
		if !ok {
			return
		}

		vector.DrawFilledRect(img, float32(node.Position.X), float32(node.Position.Y),
			float32(node.Size.X), float32(node.Size.Y), color.White, false)
	}
}
