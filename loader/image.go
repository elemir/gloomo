package loader

import (
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Image struct {
	assetsPath string
	cache      map[string]*ebiten.Image
}

func NewImage(assetsPath string) *Image {
	return &Image{
		assetsPath: assetsPath,
		cache:      make(map[string]*ebiten.Image),
	}
}

func (loader *Image) Load(path string) (*ebiten.Image, error) {
	if img, ok := loader.cache[path]; ok {
		return img, nil
	}

	img, _, err := ebitenutil.NewImageFromFile(filepath.Join(loader.assetsPath, path))
	if err != nil {
		return nil, err
	}

	loader.cache[path] = img

	return img, nil
}
