package gloomo

import (
	"cmp"
	"iter"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/node"
)

type NodeRepository interface {
	List() iter.Seq2[gid.ID, node.Node]
}

type Render struct {
	repo NodeRepository
}

func NewRender(repo NodeRepository) *Render {
	return &Render{
		repo: repo,
	}
}

type idNode struct {
	id   gid.ID
	node node.Node
}

func (r *Render) Draw(screen *ebiten.Image) {
	var idNodes []idNode

	for id, node := range r.repo.List() {
		idNodes = append(idNodes, idNode{
			id:   id,
			node: node,
		})
	}

	slices.SortFunc(idNodes, func(a, b idNode) int {
		compare := cmp.Compare(a.node.ZIndex, b.node.ZIndex)
		if compare != 0 {
			return compare
		}

		compare = cmp.Compare(a.node.Position.X, b.node.Position.X)
		if compare != 0 {
			return compare
		}

		return cmp.Compare(a.node.Position.Y, b.node.Position.Y)
	})

	for _, idNode := range idNodes {
		idNode.node.Draw(idNode.id, screen)
	}
}
