package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/geom"
)

type Node interface {
	Draw(vp ViewPort)
	Bounds() geom.Rectangle
}

type parNode struct {
	Node
	parallax float64
}

type Scene struct {
	nodes  []parNode
	camera *Camera

	w, h float64

	background color.Color
}

func NewScene(w, h float64) *Scene {
	return &Scene{
		w: w,
		h: w,
	}
}

func (s *Scene) Size() (float64, float64) {
	return s.w, s.h
}

func (s *Scene) SetCamera(camera *Camera) {
	s.camera = camera
	s.camera.sceneSize = geom.Vec2{s.w, s.h}
}

func (s *Scene) SetBackground(clr color.Color) {
	s.background = clr
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.camera.Update()

	var nodes []parNode

	for _, node := range s.nodes {
		min := s.camera.Bounds().Min.Mul(node.parallax)
		max := s.camera.Bounds().Size().Add(min)
		camera := geom.Rectangle{
			Min: min,
			Max: max,
		}

		if node.Bounds().Overlaps(camera) {
			nodes = append(nodes, node)
		}
	}

	/*
		sort.Slice(objs, func(i, j int) bool {
			return objs[i].ZIndex() < objs[j].ZIndex()
		})
	*/

	if s.background != nil {
		screen.Fill(s.background)
	}

	for _, obj := range nodes {
		vp := s.camera.ViewPort(screen, obj.parallax)
		obj.Draw(vp)
	}
}

/*
	func (s *Scene) AddParallaxedObject(obj Node, mul float64) {
		s.objs = append(s.objs, node{
			Node:     obj,
			parallax: mul,
		})
	}
*/

func (s *Scene) AddNode(node Node, opts ...AddNodeOpt) {
	s.nodes = append(s.nodes, parNode{
		Node:     node,
		parallax: 1.0,
	})
}

// TODO(evgenii.omelchenko): should be functional options with support of parallax and z index
type AddNodeOpt struct{}

func (s *Scene) Layout(w, h int) {
	s.camera.SetSize(float64(w), float64(h))
}
