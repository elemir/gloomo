package gloomo

import (
	"image"
	"image/color"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Node interface {
	Draw(vp ViewPort)
	Bounds() image.Rectangle
}

type node struct {
	Node
	parallax float64
	zIndex   int
}

type Scene struct {
	nodes  []node
	camera *Camera

	w, h int

	background color.Color
}

func NewScene(w, h int) *Scene {
	return &Scene{
		w: w,
		h: w,
	}
}

func (s *Scene) Size() (int, int) {
	return s.w, s.h
}

func (s *Scene) SetCamera(camera *Camera) {
	s.camera = camera
	s.camera.sceneSize = image.Point{s.w, s.h}
}

func (s *Scene) SetBackground(clr color.Color) {
	s.background = clr
}

func (s *Scene) Draw(screen *ebiten.Image) {
	var nodes []node

	for _, node := range s.nodes {
		min := s.camera.Bounds().Min

		fMinX, fMinY := float64(min.X), float64(min.Y)
		fMinX, fMinY = fMinX*node.parallax, fMinY*node.parallax
		fMinX, fMinY = math.Round(fMinX), math.Round(fMinY)

		min = image.Pt(int(fMinX), int(fMinY))
		max := s.camera.Bounds().Size().Add(min)

		camera := image.Rectangle{
			Min: min,
			Max: max,
		}

		if node.Bounds().Overlaps(camera) {
			nodes = append(nodes, node)
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].zIndex < nodes[j].zIndex
	})

	if s.background != nil {
		screen.Fill(s.background)
	}

	for _, node := range nodes {
		vp := s.camera.ViewPort(screen, node.parallax)
		node.Draw(vp)
	}
}

// AddNode is used for adding node to scene with specified options.
func (s *Scene) AddNode(baseNode Node, opts ...AddNodeOpt) {
	node := node{
		Node:     baseNode,
		parallax: 1,
	}

	for _, opt := range opts {
		opt(&node)
	}

	s.nodes = append(s.nodes, node)
}

// AddNodeOpt is used for configure node that added on scene.
type AddNodeOpt func(*node)

func WithParallax(parallax float64) AddNodeOpt {
	return func(node *node) {
		node.parallax = parallax
	}
}

func WithZIndex(zIndex int) AddNodeOpt {
	return func(node *node) {
		node.zIndex = zIndex
	}
}

func (s *Scene) Layout(w, h int) {
	s.camera.SetSize(w, h)
}
