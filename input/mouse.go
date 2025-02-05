package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
)

func toEbitenMouseButton(button MouseButton) (ebiten.MouseButton, bool) {
	switch button {
	case MouseButtonLeft:
		return ebiten.MouseButtonLeft, true
	case MouseButtonRight:
		return ebiten.MouseButtonRight, true
	}

	return 0, false
}

type Mouse struct{}

func (m *Mouse) IsPressed(button MouseButton) bool {
	emb, ok := toEbitenMouseButton(button)
	if !ok {
		return false
	}

	return inpututil.IsMouseButtonJustPressed(emb)
}

func (m *Mouse) Position() image.Point {
	return image.Pt(ebiten.CursorPosition())
}
