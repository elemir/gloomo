package gloomo

import (
	"image"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2/text"
)

type Text struct {
	str      string
	fnt      font.Face
	bounds   image.Rectangle
	position image.Point
}

func NewText(fnt font.Face) *Text {
	return &Text{
		fnt: fnt,
	}
}

func (txt *Text) SetString(str string) {
	txt.str = str
	txt.bounds = text.BoundString(txt.fnt, txt.str).Add(txt.position)
}

func (txt *Text) SetPosition(pos image.Point) {
	txt.position = pos
	txt.bounds = text.BoundString(txt.fnt, txt.str).Add(txt.position)
}

func (txt Text) Bounds() image.Rectangle {
	return txt.bounds
}

func (txt Text) Draw(screen ViewPort) {
	// TODO(elemir): add color
	screen.DrawText(txt.str, txt.fnt, txt.bounds.Min, color.RGBA{R: 0xEF, G: 0xFF, B: 0xFF, A: 0xFF})
}
