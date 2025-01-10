package container

import gid "github.com/elemir/gloomo/id"

type Set map[gid.ID]struct{}

func (s *Set) Add(id gid.ID) {
	if *s == nil {
		*s = make(map[gid.ID]struct{})
	}

	(*s)[id] = struct{}{}
}

func (s *Set) Get(id gid.ID) bool {
	if *s == nil {
		return false
	}

	_, ok := (*s)[id]

	return ok
}

func (s *Set) Delete(id gid.ID) {
	if *s == nil {
		return
	}

	delete(*s, id)
}
