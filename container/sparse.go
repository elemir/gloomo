package container

import (
	"iter"

	gid "github.com/elemir/gloomo/id"
)

type element[T any] struct {
	id  gid.ID
	val T
}

func elem[T any](id gid.ID, val T) element[T] {
	return element[T]{
		id:  id,
		val: val,
	}
}

type SparseArray[T any] struct {
	indexes map[gid.ID]int
	elems   []element[T]
}

func (sa *SparseArray[T]) Set(id gid.ID, val T) {
	if sa.indexes == nil {
		sa.indexes = make(map[gid.ID]int)
	}

	idx, ok := sa.indexes[id]
	if ok {
		sa.elems[idx] = elem(id, val)

		return
	}

	idx = len(sa.elems)
	sa.elems = append(sa.elems, elem(id, val))
	sa.indexes[id] = idx
}

func (sa *SparseArray[T]) Get(id gid.ID) (T, bool) {
	idx, ok := sa.indexes[id]
	if !ok {
		return Zero[T](), false
	}

	return sa.elems[idx].val, true
}

func (sa *SparseArray[T]) Delete(id gid.ID) {
	idx, ok := sa.indexes[id]
	if !ok {
		return
	}

	lastElem, elems := popSlice(sa.elems)
	sa.elems = elems
	sa.elems[idx] = lastElem
	sa.indexes[lastElem.id] = idx

	delete(sa.indexes, id)
}

func popSlice[T any](slice []T) (T, []T) {
	idx := len(slice)

	return slice[idx-1], slice[:idx-1]
}

func (sa *SparseArray[T]) Items() iter.Seq2[gid.ID, T] {
	return func(yield func(gid.ID, T) bool) {
		for _, elem := range sa.elems {
			if !yield(elem.id, elem.val) {
				return
			}
		}
	}
}
