package gloomo

import (
	"image/color"

	"github.com/elemir/gloomo/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// ViewPort is a wrapper around ebiten.Image that allows to set bounds.
type ViewPort struct {
	img    *ebiten.Image
	bounds geom.Rectangle
	mul    float64
}

func (i ViewPort) DrawImage(image *ebiten.Image, options *ebiten.DrawImageOptions) {
	x, y := i.bounds.Min.Mul(-i.mul).Unpack()
	options.GeoM.Translate(x, y)
	i.img.DrawImage(image, options)
}

func (i ViewPort) DrawRectShader(width, height int, shader *ebiten.Shader, options *ebiten.DrawRectShaderOptions) {
	x, y := i.bounds.Min.Mul(-i.mul).Unpack()
	options.GeoM.Translate(x, y)
	i.img.DrawRectShader(width, height, shader, options)
}

func (i ViewPort) DrawText(str string, fnt font.Face, pos geom.Vec2, color color.Color) {
	rpos := i.bounds.Min.Mul(-i.mul).Add(pos).Round()

	text.Draw(i.img, str, fnt, rpos.X, rpos.Y, color)
}

func (i ViewPort) Set(pos geom.Vec2, clr color.Color) {
	rpos := i.bounds.Min.Mul(-i.mul).Add(pos).Round()

	i.img.Set(rpos.X, rpos.Y, clr)
}
