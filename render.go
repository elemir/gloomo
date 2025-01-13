package gloomo

import (
	"iter"

	"github.com/hajimehoshi/ebiten/v2"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
)

type NodeRepository interface {
	List() iter.Seq2[gid.ID, model.Node]
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
		node.Draw(id, screen)
	}
}
