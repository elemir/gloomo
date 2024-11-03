package render

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/elemir/gloomo/geom"
)

type Shader struct {
	shader   *ebiten.Shader
	uniforms map[string]any

	bounds geom.Rectangle
	img    *ebiten.Image
}

func NewShader(shader *ebiten.Shader, bounds geom.Rectangle, uniforms map[string]any) *Shader {
	return &Shader{
		shader:   shader,
		uniforms: uniforms,

		bounds: bounds,
	}
}

func (s Shader) Bounds() geom.Rectangle {
	return s.bounds
}

func (s *Shader) SetPosition(pos geom.Vec2) {
	s.bounds = s.bounds.Sub(s.bounds.Min).Add(pos)
}

func (s *Shader) SetUniform(key string, value any) {
	if s.uniforms == nil {
		s.uniforms = make(map[string]any)
	}

	s.uniforms[key] = value
}

func (s *Shader) SetImage(img *ebiten.Image) {
	s.img = img
}

func (s Shader) Draw(screen ViewPort) {
	var op ebiten.DrawRectShaderOptions

	op.GeoM.Translate(math.Round(s.bounds.Min[0]), math.Round(s.bounds.Min[1]))
	op.GeoM.Scale(1, 1)

	op.Uniforms = s.uniforms
	op.Blend = ebiten.Blend{
		BlendFactorSourceRGB:        ebiten.BlendFactorSourceAlpha,
		BlendFactorDestinationRGB:   ebiten.BlendFactorOneMinusSourceAlpha,
		BlendFactorSourceAlpha:      ebiten.BlendFactorOne,
		BlendFactorDestinationAlpha: ebiten.BlendFactorOneMinusSourceAlpha,
		BlendOperationRGB:           ebiten.BlendOperationAdd,
		BlendOperationAlpha:         ebiten.BlendOperationAdd,
	}
	op.Images[0] = s.img
	size := s.bounds.Max.Sub(s.bounds.Min).Round()

	screen.DrawRectShader(size.X, size.Y, s.shader, &op)
}
