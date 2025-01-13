package repo

import (
	"iter"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
)

type Collection[T any] interface {
	Set(id gid.ID, val T)
	Get(id gid.ID) (T, bool)
	Items() iter.Seq2[gid.ID, T]
}

type Node struct {
	Nodes Collection[model.Node]
}

func (n *Node) List() iter.Seq2[gid.ID, model.Node] {
	return func(yield func(gid.ID, model.Node) bool) {
		for i, node := range n.Nodes.Items() {
			if !yield(i, node) {
				return
			}
		}
	}
}

func (n *Node) Upsert(id gid.ID, node model.Node) {
	n.Nodes.Set(id, node)
}

func (n *Node) Get(id gid.ID) (model.Node, bool) {
	return n.Nodes.Get(id)
}
