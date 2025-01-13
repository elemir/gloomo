package draw

import (
	"github.com/hajimehoshi/ebiten/v2"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
)

type SpriteRepo interface {
	Get(id gid.ID) (model.Sprite, bool)
}

func Sprite(repo SpriteRepo) model.DrawFunc {
	return func(id gid.ID, img *ebiten.Image) {
		sprite, ok := repo.Get(id)
		if !ok {
			return
		}

		var opts ebiten.DrawImageOptions
		opts.GeoM.Translate(float64(sprite.Position.X), float64(sprite.Position.Y))

		img.DrawImage(sprite.Image, &opts)
	}
}
