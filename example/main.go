package main

import (
	"image"
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo"
	"github.com/elemir/gloomo/container"
	"github.com/elemir/gloomo/draw"
	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
	"github.com/elemir/gloomo/repo"
)

type Render interface {
	Draw(img *ebiten.Image)
}

type Game struct {
	w, h   int
	render Render
}

func (g *Game) Draw(img *ebiten.Image) {
	g.render.Draw(img)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(w int, h int) (int, int) {
	g.w, g.h = w, h

	return w, h
}

func main() {
	var nodes container.SparseArray[model.Node]
	var idGen gid.Generator

	nodeRepo := &repo.Node{
		Nodes: &nodes,
	}

	rend := gloomo.NewRender(nodeRepo)
	id := idGen.New()
	drawRect := draw.Rect(nodeRepo)

	nodeRepo.Upsert(id, model.Node{
		Draw:     drawRect,
		Position: image.Pt(100, 100),
		Size:     image.Pt(64, 64),
	})

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(&Game{
		render: rend,
	}); err != nil {
		slog.Error("Unable to run game", slog.Any("err", err))
		os.Exit(-1)
	}
}
