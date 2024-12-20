package gloomo

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

// ViewPort is a wrapper around ebiten.Image that allows to set bounds.
type ViewPort struct {
	img    *ebiten.Image
	bounds image.Rectangle
	mul    float64
}

func (i ViewPort) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {
	x, y := i.pos(image.Point{})
	options.GeoM.Translate(x, y)

	i.img.DrawImage(img, options)
}

func (i ViewPort) DrawRectShader(width, height int, shader *ebiten.Shader, options *ebiten.DrawRectShaderOptions) {
	x, y := i.pos(image.Point{})
	options.GeoM.Translate(x, y)

	i.img.DrawRectShader(width, height, shader, options)
}

func (i ViewPort) DrawText(str string, fnt font.Face, pos image.Point, color color.Color) {
	x, y := i.pos(pos)

	text.Draw(i.img, str, fnt, int(x), int(y), color)
}

func (i ViewPort) DrawRect(x, y, w, h int, clr color.Color) {
	dx, dy := i.pos(image.Pt(x, y))

	vector.DrawFilledRect(i.img, float32(dx), float32(dy), float32(w), float32(h), clr, false)
}

func (i ViewPort) Set(pos image.Point, clr color.Color) {
	x, y := i.pos(pos)

	i.img.Set(int(x), int(y), clr)
}

func (i ViewPort) pos(pt image.Point) (float64, float64) {
	x, y := float64(i.bounds.Min.X), float64(i.bounds.Min.Y)

	return math.Round(-i.mul*x + float64(pt.X)), math.Round(-i.mul*y + float64(pt.Y))
}
