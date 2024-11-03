package gloomo

import (
	"image/color"

	"golang.org/x/image/font"

	"github.com/elemir/gloomo/geom"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Text struct {
	str      string
	fnt      font.Face
	bounds   geom.Rectangle
	position geom.Vec2
}

func NewText(fnt font.Face) *Text {
	return &Text{
		fnt: fnt,
	}
}

func (txt *Text) SetString(str string) {
	txt.str = str
	txt.bounds = geom.FromRectangle(text.BoundString(txt.fnt, txt.str)).Add(txt.position)
}

func (txt *Text) SetPosition(pos geom.Vec2) {
	txt.position = pos
	txt.bounds = geom.FromRectangle(text.BoundString(txt.fnt, txt.str)).Add(txt.position)
}

func (txt Text) Bounds() geom.Rectangle {
	return txt.bounds
}

func (txt Text) Draw(screen ViewPort) {
	// TODO(elemir): add color
	screen.DrawText(txt.str, txt.fnt, txt.bounds.Min, color.RGBA{R: 0xEF, G: 0xFF, B: 0xFF, A: 0xFF})
}
