package render

import (
	"iter"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/draw"
	gid "github.com/elemir/gloomo/id"
)

type NodeRepository interface {
	List() iter.Seq2[gid.ID, draw.Node]
}

type Render struct {
	repo NodeRepository
}

func NewRender(repo NodeRepository) *Render {
	return &Render{
		repo: repo,
	}
}

func (r *Render) Draw(screen *ebiten.Image) {
	for id, node := range r.repo.List() {
		/*		halfSize := node.Size.Div(2)
				rect := image.Rectangle{
					Min: node.Position.Sub(halfSize),
					Max: node.Position.Add(halfSize),
				}
		*/
		node.Draw(id, screen)
	}
}
