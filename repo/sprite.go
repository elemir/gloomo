package repo

import (
	"iter"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/draw"
	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/node"
)

type Sprite struct {
	Nodes  Collection[node.Node]
	Images Collection[*ebiten.Image]

	drawFunc node.DrawFunc
}

func (s *Sprite) List() iter.Seq2[gid.ID, node.Sprite] {
	return func(yield func(gid.ID, node.Sprite) bool) {
		for id, nd := range s.Nodes.Items() {
			img, ok := s.Images.Get(id)
			if !ok {
				continue
			}

			sprite := node.Sprite{
				Image:    img,
				Position: nd.Position,
			}

			if !yield(id, sprite) {
				return
			}
		}
	}
}

func (s *Sprite) Upsert(id gid.ID, sprite node.Sprite) {
	if s.drawFunc == nil {
		s.drawFunc = draw.Sprite(s)
	}

	s.Nodes.Set(id, node.Node{
		Draw:     s.drawFunc,
		Position: sprite.Position,
		ZIndex:   sprite.ZIndex,
		Size:     sprite.Image.Bounds().Size(),
	})

	s.Images.Set(id, sprite.Image)
}

func (s *Sprite) Get(id gid.ID) (node.Sprite, bool) {
	nd, ok := s.Nodes.Get(id)
	if !ok {
		return node.Sprite{}, false
	}

	img, ok := s.Images.Get(id)
	if !ok {
		return node.Sprite{}, false
	}

	return node.Sprite{
		Image:    img,
		Position: nd.Position,
	}, true
}
