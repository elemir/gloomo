package repo

import (
	"iter"

	"github.com/elemir/gloomo/draw"
	gid "github.com/elemir/gloomo/id"
)

type Collection[T any] interface {
	Set(id gid.ID, val T)
	Get(id gid.ID) (T, bool)
	Items() iter.Seq2[gid.ID, T]
}

type Node struct {
	Nodes Collection[draw.Node]
}

func (n *Node) List() iter.Seq2[gid.ID, draw.Node] {
	return func(yield func(gid.ID, draw.Node) bool) {
		for i, node := range n.Nodes.Items() {
			if !yield(i, node) {
				return
			}
		}
	}
}

func (n *Node) Upsert(id gid.ID, node draw.Node) {
	n.Nodes.Set(id, node)
}

func (n *Node) Get(id gid.ID) (draw.Node, bool) {
	return n.Nodes.Get(id)
}
