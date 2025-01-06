package id

type Generator struct {
	lastIdx ID
}

func (g *Generator) New() ID {
	g.lastIdx++

	return g.lastIdx - 1
}
