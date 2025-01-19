package repo

import (
	"iter"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/draw"
	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
)

type Sprite struct {
	Nodes  Collection[model.Node]
	Images Collection[*ebiten.Image]

	drawFunc model.DrawFunc
}

func (s *Sprite) List() iter.Seq2[gid.ID, model.Sprite] {
	return func(yield func(gid.ID, model.Sprite) bool) {
		for id, node := range s.Nodes.Items() {
			img, ok := s.Images.Get(id)
			if !ok {
				continue
			}

			sprite := model.Sprite{
				Image:    img,
				Position: node.Position,
			}

			if !yield(id, sprite) {
				return
			}
		}
	}
}

func (s *Sprite) Upsert(id gid.ID, sprite model.Sprite) {
	if s.drawFunc == nil {
		s.drawFunc = draw.Sprite(s)
	}

	s.Nodes.Set(id, model.Node{
		Draw:     s.drawFunc,
		Position: sprite.Position,
		Size:     sprite.Image.Bounds().Size(),
	})

	s.Images.Set(id, sprite.Image)
}

func (s *Sprite) Get(id gid.ID) (model.Sprite, bool) {
	node, ok := s.Nodes.Get(id)
	if !ok {
		return model.Sprite{}, false
	}

	img, ok := s.Images.Get(id)
	if !ok {
		return model.Sprite{}, false
	}

	return model.Sprite{
		Image:    img,
		Position: node.Position,
	}, true
}
